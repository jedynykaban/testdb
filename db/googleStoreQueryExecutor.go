package db

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

type datastoreStorage struct {
	client *datastore.Client
	ctx    context.Context
}

//New : new instance of datastoreStorage, data access object
func New(ctx context.Context, client *datastore.Client) *datastoreStorage {
	return &datastoreStorage{
		client: client,
		ctx:    ctx,
	}
}

func (d *datastoreStorage) GetLicenses() ([]License, error) {
	query := datastore.NewQuery("License")
	var result []License
	_, err := d.client.GetAll(d.ctx, query, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}
