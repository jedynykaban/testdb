package main

import (
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	serviceConfigSectionName   = "app"
	reactorConfigSectionName   = "reactor"
	datastoreConfigSectionName = "datastore"
	hairaseteConfigSectionName = "hairasete"
)

const (
	hostEntry               = "host"
	portEntry               = "port"
	cacheMaxAgeEntry        = "cacheMaxAge"
	logLevelEntry           = "loglevel"
	logOutputEntry          = "logoutput"
	logFormatEntry          = "logformat"
	timeoutEntry            = "timeout"
	projectNameEntry        = "projectName"
	mediaPrefixEntry        = "mediaPrefix"
	resourcePathFormatEntry = "resourcePathFormat"
	pathEntry               = "path"
	shisanAesKeyEntry       = "shisanAesKey"
)

// ServiceConfig is a base config for the service.
type ServiceConfig struct {
	Host               string
	Port               string
	CacheMaxAge        string
	LogLevel           log.Level
	LogOutput          io.Writer
	LogFormat          string
	ResourcePathFormat string
	RepositoryType     string
	ShisanAesKey       string
}

func (sc *ServiceConfig) log() {
	log.Infoln("Service host:", sc.Host)
	log.Infoln("Service port:", sc.Port)
	log.Infoln("Service cache max age:", sc.CacheMaxAge)
	log.Infoln("Service log level:", sc.LogLevel)
	var output string
	if sc.LogOutput == os.Stderr {
		output = "stderr"
	} else {
		output = "stdout"
	}
	log.Infoln("Service log output:", output)
	log.Infoln("Service log format:", sc.LogFormat)
	log.Infoln("Service resource path format:", sc.ResourcePathFormat)
	log.Infoln("Service shisan AES key: [REDACTED]")
}

const (
	serviceHostDefault               = ""
	servicePortDefault               = "8000"
	serviceCacheMaxAgeDefault        = "300" // in seconds so 5 minutes
	serviceLogLevelDefault           = "info"
	serviceLogOutputDefault          = "stdout"
	serviceLogFormatDefault          = "json"
	serviceResourcePathFormatDefault = "/v2/resources/%s"
)

// ReactorConfig is a base config for the reactor connection.
type ReactorConfig struct {
	Host        string
	Port        string
	Timeout     time.Duration
	MediaPrefix string
}

func (rc *ReactorConfig) log() {
	log.Infoln("Reactor host:", rc.Host)
	log.Infoln("Reactor port:", rc.Port)
	log.Infoln("Reactor timeout:", rc.Timeout)
	log.Infoln("Reactor media prefix:", rc.MediaPrefix)
}

const (
	reactorHostDefault        = "se-02.adtomafusion.com"
	reactorPortDefault        = "80"
	reactorTimeoutDefault     = 1000 * time.Millisecond
	reactorMediaPrefixDefault = "mosaiqio.dev."
)

// DatastoreConfig is a base config for the datastore connection.
type DatastoreConfig struct {
	ProjectName string
}

func (dc *DatastoreConfig) log() {
	log.Infoln("Datastore project name:", dc.ProjectName)
}

const (
	datastoreProjectNameDefault = "mosaiqio-dev"
)

// HairaseteConfig is a base config for the hairasete connection.
type HairaseteConfig struct {
	Host string
	Port string
	Path string
}

func (hc *HairaseteConfig) log() {
	log.Infoln("Hairasete host:", hc.Host)
	log.Infoln("Hairasete port:", hc.Port)
	log.Infoln("Hairasete path:", hc.Path)
}

const (
	hairaseteHostDefault = "???"
	hairasetePortDefault = "???"
	hairasetePathDefault = "/tokens/validate"
)

// Config is a full omakase config.
type Config struct {
	Service   ServiceConfig
	Reactor   ReactorConfig
	Datastore DatastoreConfig
	Hairasete HairaseteConfig
}

// Log logs the settings stored in config.
func (c *Config) Log() {
	c.Service.log()
	c.Reactor.log()
	c.Datastore.log()
	c.Hairasete.log()
}

func setDefaults() {
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, hostEntry), serviceHostDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, portEntry), servicePortDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, cacheMaxAgeEntry), serviceCacheMaxAgeDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, logLevelEntry), serviceLogLevelDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, logOutputEntry), serviceLogOutputDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, logFormatEntry), serviceLogFormatDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", serviceConfigSectionName, resourcePathFormatEntry), serviceResourcePathFormatDefault)

	viper.SetDefault(fmt.Sprintf("%s.%s", reactorConfigSectionName, hostEntry), reactorHostDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", reactorConfigSectionName, portEntry), reactorPortDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", reactorConfigSectionName, timeoutEntry), reactorTimeoutDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", reactorConfigSectionName, mediaPrefixEntry), reactorMediaPrefixDefault)

	viper.SetDefault(fmt.Sprintf("%s.%s", datastoreConfigSectionName, projectNameEntry), datastoreProjectNameDefault)

	viper.SetDefault(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, hostEntry), hairaseteHostDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, portEntry), hairasetePortDefault)
	viper.SetDefault(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, pathEntry), hairasetePathDefault)
}

func translateLogLevel(level string) log.Level {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Warn("Uknown log level set in config. Setting up log level to DEBUG.")
		return log.DebugLevel
	}
	return lvl
}

func translateLogOutput(out string) io.Writer {
	if out == "stderr" {
		return os.Stderr
	}
	return os.Stdout
}

func buildConfig() Config {
	logLevel := translateLogLevel(viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, logLevelEntry)))
	logOutput := translateLogOutput(viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, logOutputEntry)))
	return Config{
		Service: ServiceConfig{
			Host:               viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, hostEntry)),
			Port:               viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, portEntry)),
			CacheMaxAge:        viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, cacheMaxAgeEntry)),
			LogLevel:           logLevel,
			LogOutput:          logOutput,
			LogFormat:          viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, logFormatEntry)),
			ResourcePathFormat: viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, resourcePathFormatEntry)),
			ShisanAesKey:       viper.GetString(fmt.Sprintf("%s.%s", serviceConfigSectionName, shisanAesKeyEntry)),
		},
		Reactor: ReactorConfig{
			Host:        viper.GetString(fmt.Sprintf("%s.%s", reactorConfigSectionName, hostEntry)),
			Port:        viper.GetString(fmt.Sprintf("%s.%s", reactorConfigSectionName, portEntry)),
			Timeout:     viper.GetDuration(fmt.Sprintf("%s.%s", reactorConfigSectionName, timeoutEntry)),
			MediaPrefix: viper.GetString(fmt.Sprintf("%s.%s", reactorConfigSectionName, mediaPrefixEntry)),
		},
		Datastore: DatastoreConfig{
			ProjectName: viper.GetString(fmt.Sprintf("%s.%s", datastoreConfigSectionName, projectNameEntry)),
		},
		Hairasete: HairaseteConfig{
			Host: viper.GetString(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, hostEntry)),
			Port: viper.GetString(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, portEntry)),
			Path: viper.GetString(fmt.Sprintf("%s.%s", hairaseteConfigSectionName, pathEntry)),
		},
	}
}

func getConfig() Config {
	// set defaults first
	setDefaults()

	// env vars setup
	// viper.SetEnvPrefix("omakase")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.AutomaticEnv()

	// // config file setup
	// viper.SetConfigType("toml")
	// viper.SetConfigName("omakase")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"config_file_name": "omakase.toml",
	// 		"error":            err,
	// 	}).Warn("Problem reading config file, using defaults or environment variables (if set)")
	// }
	config := buildConfig()
	return config
}
