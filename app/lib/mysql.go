package lib

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"practice/config"
)

var DbDriver map[string]*gorm.DB

func InitMysql(driver string) {
	dbConfig := config.GetConfigByMap("database." + driver)
	var c config.MySQLConfig
	err := mapstructure.Decode(dbConfig, &c)
	if err != nil {
		panic(err)
	}
	dns := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database +
		"?charset=" + c.Charset + "&parseTime=True&loc=Asia%2FShanghai"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("mysql conn err %s", err))
	}
	dbPool, _ := db.DB()
	dbPool.SetMaxIdleConns(int(c.MaxIdleConn))
	dbPool.SetMaxOpenConns(int(c.SetMaxOpenConn))
	err = dbPool.Ping()
	if err != nil {
		panic(fmt.Errorf("mysql conn err %s", err))
	}
	DbDriver = make(map[string]*gorm.DB)
	DbDriver[driver] = db
}

func GetDb(driver string) *gorm.DB {
	if db, err := DbDriver[driver]; err {
		return db
	}
	panic(fmt.Errorf("DB Pool is not init, please check it"))
}
