package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type LoginService interface {
	LoginUser(username string, password string) bool
}
type LoginInformations struct {
	username string
	password string
}

// LoginUser implements LoginService
func (info *LoginInformations) LoginUser(username string, password string) bool {
	status := false
	var data = bson.M{}
	data = GetUser(username)
	if !(data["data"] == "not found") {
		x := data["data"].([]bson.M)
		passwordRaw := CheckPasswordHash(password, x[0]["password"].(string))
		usernameRaw := username == x[0]["username"].(string)
		if passwordRaw && usernameRaw {
			status = true
		}
	} else {
		status = false

	}
	return status
}

func StaticLoginService() LoginService {
	return &LoginInformations{
		username: "testmail@gmail.com",
		password: "testing",
	}
}

//get data form mongodb by username
func GetUser(username string) bson.M {
	client := MongoDBConnection()
	collection := client.Database("ServiceSystem").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}
	var user []bson.M
	if err := cursor.All(ctx, &user); err != nil {
		log.Fatal(err)
	}
	if len(user) > 0 {
		return bson.M{"data": user}
	}
	return bson.M{"data": "not found"}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
