package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	_mongoURL      = os.Getenv("MONGO_URL")
	_mongoDataBase = os.Getenv("MONGO_DATABASE")
)

// IMongoDriver is interface
type IMongoDriver interface {
	Database() *mongo.Database
}

// Driver is struct
type Driver struct {
	DB *mongo.Database
}

// NewMongoDriver is func
func NewMongoDriver() IMongoDriver {
	result := &Driver{}
	if err := result.EstablishConnection(); err != nil {
		log.Fatal(err)
	}
	return result
}

// Database is func
func (d *Driver) Database() *mongo.Database {
	return d.DB
}

// EstablishConnection is func
func (d *Driver) EstablishConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		return fmt.Errorf("driver.EstablishConnection() on WithTimeout")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(_mongoURL))
	if err != nil {
		return fmt.Errorf("driver.EstablishConnection() on NewClient:%w", err)
	}
	if err := client.Connect(ctx); err != nil {
		return fmt.Errorf("driver.EstablishConnection() on Connect:%w", err)
	}
	d.DB = client.Database(_mongoDataBase)
	return nil
}
