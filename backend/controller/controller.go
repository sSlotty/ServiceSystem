package controller

import (
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sSlotty/authentication-service/dto"
	"sSlotty/authentication-service/service"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

// Login
func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		fmt.Println(err)
		return "no data found"
	}

	isUserAuthenticated := controller.loginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {

		data := service.GetUser(credential.Username)
		if !(data["data"] == "not found") {
			x := data["data"].([]bson.M)
			userID := x[0]["userid"].(string)
			return controller.jWtService.GenerateToken(userID, credential.Username, true)

		}

	}
	return ""
}

// RegisterController Register
func RegisterController(ctx *gin.Context) {
	var credential dto.RegisterCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		fmt.Println(err)
		return
	}
	if credential.Password == credential.ConfPassword {
		//controller.RegisterUser(credential.Username, credential.Password, credential.Email, credential.FullName)
		status := service.CheckUserExist(credential.Username, credential.Email)
		if !status {
			//haspassword
			hashedPassword, _ := HashPassword(credential.Password)
			info := &service.RegisterInformations{
				UserID:   xid.NewWithTime(time.Now()).String(),
				Username: credential.Username,
				Password: hashedPassword, Email: credential.Email,
				FullName: credential.FullName,
				CreateAt: time.Now().String(),
				UpdateAt: time.Now().String(),
			}
			statusRegister := service.RegisterUser(info)
			if statusRegister {
				ctx.JSON(http.StatusCreated, gin.H{"message": "success to created account ⚡", "data": bson.M{}})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "register failed please try again or contact admin", "data": bson.M{}})
			}
		} else {
			ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": "username , email is already exist", "data": bson.M{}})
		}
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "password and confirm password not match", "data": bson.M{}})
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
