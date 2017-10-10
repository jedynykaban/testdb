package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/datastore"

	"github.com/jedynykaban/testdb/db"

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

	ctx := context.TODO()
	storeClient := createStoreClient(ctx)
	if storeClient != nil {
		log.Info("Store client successfully created")
	}

	repo := db.New(ctx, storeClient)
	testResult, err := repo.GetLicenses()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range testResult {
		log.Info(item.Name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	log.Info("testdb stoppped")
}

func createStoreClient(ctx context.Context) *datastore.Client {
	log.Info("config.Datastore.ProjectName: " + config.Datastore.ProjectName)
	client, err := datastore.NewClient(ctx, config.Datastore.ProjectName)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
