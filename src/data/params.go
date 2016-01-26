package data

import (
	//	"fmt"
	//	"encoding/json"
	"time"
)

var ServerIp = "192.169.23.24"
var ServerPort = "8080"

var QueryDrawUrl = "/incubator/bc/bet/saleDraw.do"
var BetUrl = "/incubator/bc/bet/bet.do"

var BetAmount1 float64 = 1.0
var BetAmount3 float64 = 3.0
var BetAmount5 float64 = 5.0
var BetAmount10 float64 = 10.0

var DrawId int = 0

type Bet struct {
	AccountCode string  `json:"accountcode"`
	BetAmount   float32 `json:"betamount"`
	DrawId      int     `json:"drawid"`
	GameId      int     `json:"gameid"`
	OldSerialNo int     `json:"oldserialno"`
	SerialNo    int     `json:"serialno"`
}

var BetReq10001 string = "{\"accountCode\":\"testaccount\",\"betAmount\":0.0,\"drawId\":0,\"gameId\":10001,\"oldSerialNo\":0,\"serialNo\":0}"
var BetReq10002 string = "{\"accountCode\":\"testaccount\",\"betAmount\":3.0,\"drawId\":0,\"gameId\":10002,\"oldSerialNo\":0,\"serialNo\":0}"
var BetReq10003 string = "{\"accountCode\":\"testaccount\",\"betAmount\":0.0,\"drawId\":0,\"gameId\":10003,\"oldSerialNo\":0,\"serialNo\":0}"

type Draw struct {
	GameId string
}

var QueryDrawStr string = "{\"GameId\":\"10002\"}"

func CreateSerialNo() string {
	SerialNo := ""
	SerialNo = time.Now().Format("20060102150405")[6:] + RandNum(4) + RandNum(6)
	return SerialNo
}
