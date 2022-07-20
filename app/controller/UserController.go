package controller

import (
	"gin-skeleton/app/lib"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	m := map[string]any{
		"str": "this is userinfo",
	}
	c.JSON(200, lib.NewResult().SetData(m).GetResult())
}
