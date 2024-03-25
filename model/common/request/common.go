package request

type StockImformation struct {
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	PEratio       string `json:"PEratio"`
	DividendYield string `json:"DividendYield"`
	PBratio       string `json:"PBratio"`
	Type          string `json:"type"`
}

type StockImformation2 struct {
	Code          string `json:"SecuritiesCompanyCode"`
	Name          string `json:"CompanyName"`
	PEratio       string `json:"PriceEarningRatio"`
	DividendYield string `json:"YieldRatio"`
	PBratio       string `json:"PriceBookRatio"`
	Type          string `json:"type"`
}

type StockCode struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	PEratio string `json:"peratio"`
	Type    string `json:"type"`
}

type FavoritedStock struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	PEratio string `json:"peratio"`
	Type    int    `json:"type"`
}
