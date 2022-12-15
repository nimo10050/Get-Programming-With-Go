//package main
//
//import (
//	"fmt"
//	"log"
//	"os"
//	"regexp"
//	"strings"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type PatiTradeInfo struct {
//	InsureTradeNo  string `db:"insure_trade_no"`
//	PersonInsureId string `db:"person_insure_id"`
//	TradeOtherInfo string `db:"trade_other_info"`
//	DeleteFlag     int    `db:"delete_flag"`
//}
//
//type TradeLogs struct {
//	TradeTypeCode   string `db:"trade_type_code"`
//	TradeParamInfo  string `db:"trade_param_info"`
//	TradeResultInfo string `db:"trade_result_info"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//
//	// 日志初始化
//	logFile, err := os.OpenFile("./cleandata202106.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		fmt.Println("open log file failed, err:", err)
//		return
//	}
//	log.SetOutput(logFile)
//	log.SetFlags(log.Lmicroseconds | log.Ldate)
//
//	// 数据库初始化
//	log.Println("======数据库连接初始化======")
//	var username = "emr"
//	var password = "Wowjoy@2018"
//	var ip = "10.10.32.21"
//	var port = "31591"
//	var dbName = "sis"
//	log.Println("用户名: ", username)
//	log.Println("密码: ", password)
//	log.Println("ip: ", ip)
//	log.Println("端口: ", port)
//	log.Println("数据库名称: ", dbName)
//	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, ip, port, dbName)
//	log.Println(dbURL)
//	database, err := sqlx.Open("mysql", dbURL)
//	if err != nil {
//		log.Println("数据库连接失败!", err)
//		defer Db.Close()
//		return
//	}
//	Db = database
//	log.Println("======数据库初始化成功!======")
//	log.Println()
//}
//
//func main() {
//	var currentDate = "2021-06-07"
//	var endDate = "2021-06-30"
//	for {
//		cleanData(currentDate)
//		time.Sleep(10 * time.Second)
//		currentDate = getNextDay(currentDate)
//		if endDate == currentDate {
//			cleanData(currentDate)
//			log.Println("当前时间: ", currentDate, "等于截止日期:", endDate, " program exit! ")
//			return
//		}
//	}
//
//}
//
//func cleanData(date string) {
//	log.Println()
//	log.Println("开始清理数据 ===> ", date)
//	queryTradeLog(date)
//	log.Println("数据清理完成 <===", date)
//}
//
//func getNextDay(date string) string {
//	loc, _ := time.LoadLocation("Local")
//	t, _ := time.ParseInLocation("2006-01-02", date, loc)
//	d, _ := time.ParseDuration("24h")
//	var nextDay = t.Add(d)
//	return nextDay.Format("2006-01-02")
//}
//
//func queryTradeInfo(tradeOtherInfo string, date string) []PatiTradeInfo {
//	var tradeInfo []PatiTradeInfo
//	err := Db.Select(&tradeInfo, "select insure_trade_no, person_insure_id, trade_other_info, delete_flag from 			si_pati_trade where trade_other_info=? and to_days(insure_trade_time)=to_days(?)", tradeOtherInfo, date)
//	if err != nil {
//		log.Println("执行查询 si_pati_trade sql 失败", err)
//		return nil
//	}
//	return tradeInfo
//
//}
//
//func queryTradeLog(date string) {
//	var tradeLogs []TradeLogs
//	err := Db.Select(&tradeLogs, "select trade_type_code, trade_param_info, trade_result_info from si_trade_logs_2 where to_days(create_time) = to_days(?)", date)
//
//	if err != nil {
//		log.Println("执行查询交易日志 sql 失败, ", err)
//		return
//	}
//	var tradeLogsLen = len(tradeLogs)
//	log.Println("查询到 ", len(tradeLogs), " 条交易日志")
//
//	for i := 0; i < tradeLogsLen; i++ {
//		tradeLog := tradeLogs[i]
//		var isRefund = strings.Index(tradeLog.TradeTypeCode, "42") != -1
//		var tradeResultIsRefundSettleId = len(tradeLog.TradeResultInfo) == 33
//		// 退费交易 && 退费结算id存在
//
//		if isRefund && tradeResultIsRefundSettleId {
//			log.Println()
//			log.Println("=======", i, "======")
//			// 从退费交易日志中拿到退费结算id
//			var oriSettleRecordId = getSettleRecordIdFromLog(tradeLog.TradeParamInfo)
//			// 从退费交易日志中拿到原始结算ID
//			var refundSettleRecordId = getRefundSettleRecordIdFromLog(tradeLog.TradeResultInfo)
//			log.Println("原始结算id: ", oriSettleRecordId, "退费结算id: ", refundSettleRecordId)
//			if oriSettleRecordId == "" || refundSettleRecordId == "" {
//				continue
//			}
//			// 根据原始结算ID, 查询结算记录(理论上一正一负)
//			var tradeInfos = queryTradeInfo(oriSettleRecordId, date)
//			if len(tradeInfos) == 2 {
//				log.Println("根据退费 id 查询到 2 条结算记录 => ", tradeInfos)
//				first := tradeInfos[0]
//				second := tradeInfos[1]
//				var refundTradeInfo PatiTradeInfo
//				if first.DeleteFlag == 1 || second.DeleteFlag == 2 {
//					refundTradeInfo = second
//				} else if first.DeleteFlag == 2 || second.DeleteFlag == 1 {
//					refundTradeInfo = first
//				}
//				if refundTradeInfo != (PatiTradeInfo{}) {
//					log.Println("更新对应的退费数据: ", refundTradeInfo)
//					updateTradeInfo(oriSettleRecordId, refundSettleRecordId)
//				}
//			} else {
//				log.Println("根据原始结算 id 期望查询到 2 条结算记录, 但是实际查询出 ", len(tradeInfos), " 条数据.")
//			}
//			log.Println("=======", i, "======")
//		}
//
//	}
//
//}
//
//func updateTradeInfo(oriSettleRecordId string, refundSettleRecordId string) {
//	res, exeErr := Db.Exec("update si_pati_trade set trade_other_info = ? where trade_other_info = ? and delete_flag=2 ", refundSettleRecordId, oriSettleRecordId)
//
//	if exeErr != nil {
//		log.Println("执行 update si_pati_trade sql 失败! ", exeErr)
//		return
//	}
//
//	row, err := res.RowsAffected()
//	if err != nil {
//		log.Println("获取更新行数失败, ", err)
//		return
//	}
//	log.Println("更新 si_pati_trade 表成功, success rows: ", row)
//}
//
//func getSettleRecordIdFromLog(tradeParamInfo string) string {
//	reg := regexp.MustCompile(`<prm_yka103>([^</]+)</prm_yka103>`)
//	if reg == nil {
//		log.Println("正则表达式编译失败")
//		return ""
//	}
//	// 根据规则提取关键信息
//	regResult := reg.FindAllStringSubmatch(tradeParamInfo, -1)
//	return regResult[0][1]
//}
//
//func getRefundSettleRecordIdFromLog(tradeResultInfo string) string {
//	tradeResultInfo = strings.Replace(tradeResultInfo, "{", "", -1)
//	tradeResultInfo = strings.Replace(tradeResultInfo, "}", "", -1)
//	return strings.Replace(tradeResultInfo, "prm_yka198=", "", -1)
//}
