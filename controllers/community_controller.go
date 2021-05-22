// Package controllers 控制器调用
package controllers

import (
	"Dream/dto"
	"Dream/models"
	"Dream/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var chat models.Chat
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		log.Printf("AddComment: %v", err)
		c.JSON(http.StatusOK, dto.ResultInfo{
			Status:  false,
			Message: "please input right format!",
		})
		return
	}
	ok := services.AddComment(chat)
	if !ok {
		c.JSON(http.StatusOK, dto.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, dto.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func DeleteComment(c *gin.Context) {
	ok := services.DeleteComment(c.Param("id"))
	if !ok {
		c.JSON(http.StatusOK, dto.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, dto.ResultInfo{
		Status:  true,
		Message: "success",
	})
}

func GetCommentsByDream(c *gin.Context) {
	dreams, ok := services.GetCommentsByDream(c.Param("id"))
	if !ok {
		c.JSON(http.StatusOK, dto.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, dto.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    dreams,
	})
}
