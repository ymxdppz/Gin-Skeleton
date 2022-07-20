package model

import "time"

type MerchantModel struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	MerchantId  string    `json:"merchant_id"`
	Name        string    `json:"name"`
	Account     string    `json:"account"`
	Password    string    `json:"password"`
	Pid         string    `json:"pid"`
	ProxyList   string    `json:"proxy_list"`
	Level       string    `json:"level"`
	MachineLock string    `json:"machine_lock"`
	Status      string    `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	ModifyTime  time.Time `json:"modify_time"`
}

func (m *MerchantModel) TableName() string {
	return "merchant"
}
