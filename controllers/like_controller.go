package controllers

import (
	"Dream/common"
	"Dream/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HasLike(c *gin.Context) {
	userId := c.Param("userId")
	dreamId := c.Param("dreamId")
	hasLike, err := services.HasLike(userId, dreamId)
	if err {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  true,
		Message: "success",
		Data:    hasLike,
	})
}

func Like(c *gin.Context) {
	userId := c.Param("userId")
	dreamId := c.Param("dreamId")
	hasLike, err := services.Like(userId, dreamId)
	if err {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	if !hasLike {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  true,
			Message: "success",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  false,
		Message: "already liked",
	})
}

func Unlike(c *gin.Context) {
	userId := c.Param("userId")
	dreamId := c.Param("dreamId")
	hasLike, err := services.Like(userId, dreamId)
	if err {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  false,
			Message: "database error",
		})
		return
	}
	if hasLike {
		c.JSON(http.StatusOK, common.ResultInfo{
			Status:  true,
			Message: "success",
		})
		return
	}
	c.JSON(http.StatusOK, common.ResultInfo{
		Status:  false,
		Message: "hasn't liked",
	})
}
