package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// 铜仁结算数据清理

// 1. 把 si_pati_trade 表的 recil_center_code 的数据 修改为 520603

// 2. 查询 si_pati_trade 重复的结算数据， 并注释掉

// 3. 把虚假的工伤结算数据修改掉 修改 si_pati_trade 表 medi_insure_id=1024， fi_reckon 的 reckon 后缀改为 1024

// 4. 异地
