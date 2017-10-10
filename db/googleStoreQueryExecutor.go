package db

import (
	"context"
	"errors"

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
	return nil, errors.New("NOT IMPLEMENTED")
}
