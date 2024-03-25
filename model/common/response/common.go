package response

type MsgArrayItem struct {
	Tv    string `json:"tv"`
	Ps    string `json:"ps"`
	Pz    string `json:"pz"`
	Bp    string `json:"bp"`
	Fv    string `json:"fv"`
	Oa    string `json:"oa"`
	Ob    string `json:"ob"`
	A     string `json:"a"`
	B     string `json:"b"`
	C     string `json:"c"`
	D     string `json:"d"`
	Ch    string `json:"ch"`
	Ot    string `json:"ot"`
	Tlong string `json:"tlong"`
	F     string `json:"f"`
	Ip    string `json:"ip"`
	G     string `json:"g"`
	Mt    string `json:"mt"`
	Ov    string `json:"ov"`
	H     string `json:"h"`
	I     string `json:"i"`
	It    string `json:"it"`
	Oz    string `json:"oz"`
	L     string `json:"l"`
	N     string `json:"n"`
	O     string `json:"o"`
	P     string `json:"p"`
	Ex    string `json:"ex"`
	S     string `json:"s"`
	T     string `json:"t"`
	U     string `json:"u"`
	V     string `json:"v"`
	W     string `json:"w"`
	Nf    string `json:"nf"`
	Y     string `json:"y"`
	Z     string `json:"z"`
	Ts    string `json:"ts"`
}

type StockMsg struct {
	MsgArray []MsgArrayItem `json:"msgArray"`
}

type FinancialEntry struct {
	Date       string  `json:"date"`
	StockID    string  `json:"stock_id"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	OriginName string  `json:"origin_name"`
}

type Epsdata struct {
	Data []FinancialEntry `json:"data"`
}
