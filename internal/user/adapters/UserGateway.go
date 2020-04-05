package adapters

import (
	"context"
	"os"
	"time"

	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/internal/user/domains"
	"github.com/ShotaroOkada/Online_CA_Tech_Dojo_Golang/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	_userCollectionName = os.Getenv("USER_COLLECTION_NAME")
)

// UserGateway is struct
type UserGateway struct {
	MongoDriver db.IMongoDriver
}

// NewUserGateway is func
func NewUserGateway(driver db.IMongoDriver) domains.UserRepository {
	result := &UserGateway{}
	result.MongoDriver = driver
	return result
}

// Create is func
func (u UserGateway) Create(user *domains.User) (token string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := u.MongoDriver.Database().Collection(_userCollectionName).InsertOne(ctx, bson.M{
		"id":   user.ID,
		"name": user.Name,
	}); err != nil {
		return "", err
	}
	token = "hoge"
	return token, nil
}

// Get is func
func (u UserGateway) Get(token string) (*domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"token": token,
	}
	var result bson.M
	err := u.MongoDriver.Database().Collection(_userCollectionName).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	var user *domains.User
	user = &domains.User{
		ID:   result["id"].(primitive.ObjectID).String(),
		Name: result["name"].(string),
	}

	return user, nil
}

// Update is func
func (u UserGateway) Update(user domains.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"id": user.ID,
	}
	_, err := u.MongoDriver.Database().Collection(_userCollectionName).UpdateOne(ctx, filter, user)
	if err != nil {
		return err
	}

	return nil
}
