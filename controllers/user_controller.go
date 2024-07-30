package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	user, err := c.userService.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.DATA_CREATED_SUCCESSFULLY, user.DocId))
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	users, err := c.userService.GetAllUsers(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, users))
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	user, err := c.userService.GetUserByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, user))
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	user.DocId = ctx.Param("id")
	if err := c.userService.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_UPDATED_SUCCESSFULLY, nil))
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if err := c.userService.DeleteUser(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_DELETED_SUCCESSFULLY, nil))
}
