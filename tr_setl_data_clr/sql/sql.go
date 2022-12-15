package sql

// 因为 xx 原因, 部分结算数据的本地对账中心编码存错了。 比如万山 的本地对账中心编码应该是 520603.
// 查询出本地对账中心编码
var selectPatiTradeInfoSql = "select insure_trade_no, person_insure_id, trade_other_info, delete_flag " +
	" from si_pati_trade " +
	" where delete_flag = 0" +
	" and medi_insure_id = '1024'" +
	" and settle_record_type in (2, 4)" +
	" and reconcil_center_code != ?" +
	" and to_days(insure_trade_time) between to_days(?) and to_days(?)"

// 查询 si_pati_trade 表
func getSelectPatiTradeInfoSql() string {
	return selectPatiTradeInfoSql
}
