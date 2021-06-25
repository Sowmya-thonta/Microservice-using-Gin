package controllers

import (
	"microservices/domain/httperrors"
	"microservices/domain/users"
	"microservices/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func respond(c *gin.Context, isXml bool, httpCode int, body interface{}) {
	if isXml {
		c.XML(httpCode, body)
		return
	}
	c.JSON(httpCode, body)
}

func (controller usersController) Create(ctx *gin.Context) {

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//Return bad request
		httpErr := httperrors.NewBadRequestError("Invalid JSON ")
		ctx.JSON(httpErr.Code, httpErr)
		return
	}

	createdUser, err := services.UsersService.Create(user)
	if err != nil {
		ctx.JSON(err.Code, err)
		return

	}
	ctx.JSON(http.StatusCreated, createdUser)

}

func (controller usersController) Get(ctx *gin.Context) {
	isXml := ctx.GetHeader("Accept") == "aplication/xml"
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("Invalid UserId")
		respond(ctx, isXml, httpErr.Code, httpErr)
		return

	}

	user, getErr := services.UsersService.Get(userId)
	if getErr != nil {
		respond(ctx, isXml, getErr.Code, getErr)
		return
	}

	respond(ctx, isXml, http.StatusOK, user)
}
