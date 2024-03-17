package system

import (
	"gland_test/global"
)

type SysMyFavourite struct {
	global.SKW_MODEL
	StockID   string `json:"code" gorm:"index;comment:股票代號"`
	StockName string `json:"name" gorm:"index;comment:股票名稱"`
}

func (SysMyFavourite) TableName() string {
	return "sys_myfavourite"
}
