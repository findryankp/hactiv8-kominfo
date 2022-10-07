package controllers

import (
	"net/http"
	"orders/configs"
	"orders/models"

	"github.com/gin-gonic/gin"
)

func byId(id, model interface{}) (bool, interface{}) {
	query := configs.DB.First(&model, id)
	if query.Error != nil {
		return false, query.Error
	}
	return true, nil
}

func GetOrder(c *gin.Context) {
	var items []models.Item
	configs.DB.Model(&models.Item{}).Preload("Order").Find(&items)
	configs.ResponseJson(c, http.StatusOK, true, "-", items)
}

func CreateOrder(c *gin.Context) {
	var formOrder models.FormOrder
	c.ShouldBindJSON(&formOrder)

	order := models.Order{
		CustomerName: formOrder.CustomerName,
		OrderedAt:    formOrder.OrderedAt,
	}

	configs.DB.Create(&order)

	items := models.Item{
		ItemCode:    formOrder.Item[len(formOrder.Item)-1].ItemCode,
		Description: formOrder.Item[len(formOrder.Item)-1].Description,
		Quantity:    formOrder.Item[len(formOrder.Item)-1].Quantity,
		OrderId:     order.ID,
	}

	configs.DB.Create(&items)

	configs.ResponseJson(c, http.StatusCreated, true, "-", order)
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	isExist, msg := byId(c.Param("id"), &order)
	if !isExist {
		configs.ResponseJson(c, http.StatusBadRequest, true, "Data not found", msg)
	}

	var formOrder models.FormOrder
	c.ShouldBindJSON(&formOrder)
	configs.DB.Model(&order).Updates(formOrder)

	updateItem := models.Item{
		ItemCode:    formOrder.Item[len(formOrder.Item)-1].ItemCode,
		Description: formOrder.Item[len(formOrder.Item)-1].Description,
		Quantity:    formOrder.Item[len(formOrder.Item)-1].Quantity,
		OrderId:     order.ID,
	}

	var item models.Item
	configs.DB.Model(&item).Where("order_id = ?", order.ID).Updates(updateItem)

	configs.DB.Model(&models.Item{}).Where("order_id = ?", order.ID).Preload("Order").Find(&item)
	configs.ResponseJson(c, http.StatusCreated, true, "-", item)
}

func DeleteOrder(c *gin.Context) {
	var order models.Order
	configs.DB.Where("id = ?", c.Param("id")).First(&order)

	var item models.Item
	configs.DB.Where("order_id = ?", c.Param("id")).First(&item)

	configs.DB.Delete(
		&order,
		&item,
	)

	configs.ResponseJson(c, http.StatusOK, true, "-", nil)
}
