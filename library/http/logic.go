package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetId(c *gin.Context) {
	biz := c.Query("biz")
	id, err := s.GetId(biz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetSnowFlakeId(c *gin.Context) {
	id := s.SnowFlakeGetId()
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetRedisId(c *gin.Context) {
	biz := c.Query("biz")
	id, err := s.RedisGetId(biz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
