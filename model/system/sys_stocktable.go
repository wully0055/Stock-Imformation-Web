package system

import (
	"gland_test/global"
)

type SysStockTable struct {
	global.SKW_MODEL
	StockID   string `json:"code" gorm:"index;comment:股票代號"`
	StockName string `json:"name" gorm:"index;comment:股票名稱"`
	Type      string `json:"type" gorm:"index;comment:上市/上櫃"`
}

func (SysStockTable) TableName() string {
	return "sys_stocktable"
}
