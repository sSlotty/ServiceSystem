package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type RegisterService interface {
	RegisterUser(username string, password string, email string, fullName string) bool
}

type RegisterInformations struct {
	UserID   string
	Username string
	Password string
	Email    string
	FullName string
	CreateAt string
	UpdateAt string
}
type data struct {
	data string `bson:"data"`
}

func RegisterUser(info *RegisterInformations) bool {
	client := MongoDBConnection()
	collection := client.Database("ServiceSystem").Collection("Users")
	status := SaveData(collection, info)
	if status {
		return true
	}
	return false
}

//check user exist by username or email in mongo db
func CheckUserExist(username string, email string) bool {
	client := MongoDBConnection()
	collection := client.Database("ServiceSystem").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursorUser, err := collection.Find(ctx, bson.M{"username": username})
	cursorEmail, err := collection.Find(ctx, bson.M{"email": email})

	if err != nil {
		log.Fatal(err)
	}

	var user []bson.M
	if err := cursorUser.All(ctx, &user); err != nil {
		log.Fatal(err)
	}

	var mail []bson.M
	if err := cursorEmail.All(ctx, &mail); err != nil {
		log.Fatal(err)
	}
	if len(user) > 0 || len(mail) > 0 {
		return true
	} else {
		return false
	}

}

//check get username form mondb
