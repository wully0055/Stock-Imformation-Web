package system

import (
	"gland_test/global"
)

type SysStockTable struct {
	global.SKW_MODEL
	StockID   string `json:"stockID" gorm:"index;comment:股票代號"`
	StockName string `json:"stockName" gorm:"index;comment:股票名稱"`
	PEratio   string `json:"peratio" gorm:"index;comment:本益比"`
}

func (SysStockTable) TableName() string {
	return "sys_stocktable"
}