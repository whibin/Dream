package controllers

import (
	"Dream/common"
	"Dream/models"
	"Dream/services"
	"Dream/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpenId(c *gin.Context) {
	openId := utils.GetOpenId(c.Query("appId"), c.Query("code"), c.Query("secret"))
	fmt.Println(openId)
	c.JSON(http.StatusOK, common.ResultInfo{
		Status: true,
		Data:   openId,
	})
}

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.SaveUser(user)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.UpdateUser(user)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func GetUser(c *gin.Context) {
	openId := c.Query("open_id")
	user, ok := services.GetUser(openId)
	if !ok {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    user,
	})
}
