package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/services"
)

type ItemController struct {
	itemService services.ItemService
}

func NewItemController(itemService services.ItemService) *ItemController {
	return &ItemController{itemService}
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	item, err := c.itemService.CreateItem(item)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, ToSuccessResponse(global_constant.DATA_CREATED_SUCCESSFULLY, item.DocId))
}

func (c *ItemController) GetAllItems(ctx *gin.Context) {
	var requestFilterBody common.RequestFilterBodyType
	if err := ctx.ShouldBindJSON(&requestFilterBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	items, err := c.itemService.GetAllItems(requestFilterBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, items))
}

func (c *ItemController) GetItemByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	item, err := c.itemService.GetItemByID(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_FETCHED_SUCCESSFULLY, item))
}

func (c *ItemController) UpdateItem(ctx *gin.Context) {
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}

	item.DocId = ctx.Param("id")

	if err := c.itemService.UpdateItem(item); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_UPDATED_SUCCESSFULLY, nil))
}

func (c *ItemController) DeleteItem(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := c.itemService.DeleteItem(idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, ToErrorResponse(global_constant.ERROR_WHILE_PROCESSING, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, ToSuccessResponse(global_constant.DATA_DELETED_SUCCESSFULLY, nil))
}
