package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"os"
	"sSlotty/authentication-service/controller"
	"sSlotty/authentication-service/middleware"
	"sSlotty/authentication-service/service"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.Default()

	server.Use(cors.Default())

	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	uri := os.Getenv("MONGO_URI")

	// check connection mongoDB
	DB_status := service.CheckConnection(uri)
	if DB_status {
		fmt.Println("MongoDB is connected")
	} else {
		fmt.Println("MongoDB is not connected")
	}

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"messsage": "Please use /login or /register ✨ ⚡️ ", "data": bson.M{}})
	})

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"message": "success to login account ⚡️", "data": bson.M{"token": token}})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "data": bson.M{}})

		}
	})

	server.POST("/signup", func(ctx *gin.Context) {
		controller.RegisterController(ctx)
	})

	v1 := server.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": bson.M{}})
		})
	}

	port := "8080"
	err = server.Run(":" + port)
	if err != nil {
		return
	}
}
