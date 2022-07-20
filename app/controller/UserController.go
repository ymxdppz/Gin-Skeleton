package controller

import (
	"github.com/gin-gonic/gin"
	"practice/lib"
)

func Info(c *gin.Context) {
	m := map[string]any{
		"str": "this is userinfo",
	}
	c.JSON(200, lib.NewResult().SetData(m).GetResult())
}
