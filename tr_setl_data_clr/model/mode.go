package model

type PatiTradeInfo struct {
	InsureTradeNo    string `db:"insure_trade_no"`
	PersonInsureId   string `db:"person_insure_id"`
	TradeOtherInfo   string `db:"trade_other_info"`
	InsureSettleInfo string `db:"insure_settle_info"`
	DeleteFlag       int    `db:"delete_flag"`
}

type TradeLogs struct {
	TradeTypeCode   string `db:"trade_type_code"`
	TradeParamInfo  string `db:"trade_param_info"`
	TradeResultInfo string `db:"trade_result_info"`
}

type SettleOutParam struct {
	// 清算中心
	CenterCode string `json:"prm_ykb139"`
}
