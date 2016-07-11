package data

import (
	"encoding/json"
	"fmt"
	//	"reflect"
)

type QueryDrawRes struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	DrawId     int    `json:"drawId"`
	DrawStatus int    `json:"drawStatus"`
	EndTime    string `json:"endTime"`
	GameId     int    `json:"gameId"`
	PrizeTime  string `json:"prizeTime"`
	StartTime  string `json:"startTime"`
}

var DrawRespStr string = "{\"drawId\":0,\"drawStatus\":1,\"endTime\":\"\",\"gameId\":10002,\"prizeTime\":\"\",\"startTime\":\"\",\"code\":\"000\",\"message\":\"\"}"

type WTicket struct {
	BetAmount        float64
	BetTime          string
	DrawId           int
	GameType         int
	PoolAmount       float64
	Tax              float64
	AccountCode      string
	AccumulatePoints float64
	TicketId         string
	PrizeAmount      float64
	PrizeLevelId     int
	BetNo            int
	GameId           int
}

type PrL struct {
	BetNo        int
	PrizeAmount  float64
	PrizeLevelId int
	TicketId     string
}

type BetResp struct {
	Message    string
	SerialNo   string
	TicketId   string
	PrizeLevel []PrL
	WinTicket  []WTicket
	Code       string
}

type SumBase struct {
	BetAmount   float64
	PrizeLevel  int
	PrizeAmount float64
}

func ParseDrawId(res string) int {
	drawId := 0
	//var tmp map[string]interface{}
	var r QueryDrawRes
	if err := json.Unmarshal([]byte(res), &r); err == nil && r.DrawStatus == 1 {
		drawId = r.DrawId
	} else {
		fmt.Println("Error in ParseDrawId:", err)
		fmt.Println("resp: ", res, "\n")
	}
	return drawId
}
