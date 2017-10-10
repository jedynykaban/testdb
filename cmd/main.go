package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/datastore"
	log "github.com/Sirupsen/logrus"
)

var config Config

func init() {
	config = getConfig()
	setupLogging(config.Service.LogOutput, config.Service.LogLevel, config.Service.LogFormat)
}

func setupLogging(output io.Writer, level log.Level, format string) {
	log.SetOutput(output)
	log.SetLevel(level)
	if strings.EqualFold(format, "json") {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	log.Info("testdb started")

	var storeClient = createStoreClient()
	if storeClient != nil {
		log.Info("Store client successfully created")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	log.Info("testdb stoppped")
}

func createStoreClient() *datastore.Client {
	ctx := context.TODO()
	log.Info("config.Datastore.ProjectName: " + config.Datastore.ProjectName)
	client, err := datastore.NewClient(ctx, config.Datastore.ProjectName)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
