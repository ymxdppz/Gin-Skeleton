package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"practice/app/lib"
	"practice/app/model"
)

var _loginController *loginController

type loginController struct {
	controller *BaseController
}

func NewLoginController() *loginController {
	if _loginController == nil {
		_loginController = &loginController{
			controller: &BaseController{},
		}
	}
	return _loginController
}

func (l *loginController) Login(c *gin.Context) {
	db := lib.GetDb("defualt")
	merchantModel := &model.MerchantModel{}
	err := db.Select("merchant_id", "name", "account", "create_time").First(merchantModel).Error
	if err != nil {
		c.JSON(200, lib.NewResult().SetErrInfo(5000, "操作失败").GetResult())
	}
	fmt.Println(merchantModel)
	c.JSON(200, lib.NewResult().SetData(map[string]any{}).GetResult())
}

func (l *loginController) Logout(c *gin.Context) {
	m := map[string]any{
		"str": "this is logout",
	}
	c.JSON(200, lib.NewResult().SetData(m).GetResult())
}
