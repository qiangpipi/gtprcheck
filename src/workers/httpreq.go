package workers

import (
	"data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func ReqWorker(q chan<- data.SumBase, betGame string, betAmount float64) {
	//	var prList []data.PrLevel
	var res data.BetResp
	var sBase data.SumBase
	//Internal counter for 1 thread
	var count int64 = 0
	var totalMs int64 = 0
	for {
		count++
		a := time.Now().UnixNano()
		r := HttpReq(data.BetUrl, CreateParams(betAmount, betGame))
		b := time.Now().UnixNano()
		totalMs += ((b - a) / 1000000)
		if count%1000 == 0 {
			fmt.Println("1000 req took ", totalMs, " Milliseconds")
			count = 0
			totalMs = 0
		}
		//Check if "code\":\"000" existing in r, then put into queue resp
		if strings.Contains(r, "code\":\"000") {
			if err := json.Unmarshal([]byte(r), &res); err == nil {
				for _, v := range res.PrizeLevel {
					if v.PrizeLevelId >= 21 && v.PrizeLevelId <= 30 {
						sBase.BetAmount = 1.0
					} else {
						sBase.BetAmount = betAmount
					}
					sBase.PrizeLevel = v.PrizeLevelId
					sBase.PrizeAmount = v.PrizeAmount
					q <- sBase
				}
			} else {
				fmt.Println("Resp: ", r)
			}
		} else {
			fmt.Println("###########Server return error##########")
			fmt.Println("Resp: ", r)
		}
	}
}

func HttpReq(path, params string) (res string) {
	u := "http://" + data.ServerIp + ":" + data.ServerPort + path
	resp, err := http.Post(u,
		"application/x-www-form-urlencoded",
		strings.NewReader("req="+params))
	if err == nil {
		d, _ := ioutil.ReadAll(resp.Body)
		if str, err := url.QueryUnescape(string(d)); err == nil {
			res = str
		} else {
			fmt.Println("Error in HttpReq: ", err, "\n")
			fmt.Println("Resp: ", string(d))
		}
	} else {
		fmt.Println("Error in HttpReq: ", err, "\n")
		fmt.Println("Params: ", params, "\n")
	}
	defer resp.Body.Close()
	return res
}

func CreateParams(amount float64, reqStr string) string {
	var tmp map[string]interface{}
	params := ""
	if err := json.Unmarshal([]byte(reqStr), &tmp); err == nil {
		tmp["serialNo"] = data.CreateSerialNo()
		tmp["betAmount"] = amount
		if d, ok := tmp["gameId"].(float64); ok && int(d) == 10002 {
			tmp["drawId"] = data.DrawId
		}
	} else {
		fmt.Println("Error in CreateParams: ", err)
	}
	if b, err := json.Marshal(tmp); err == nil {
		//		params = "?req=" + string(b)
		params = string(b)
	} else {
		fmt.Println("Error: ", err)
	}
	return params
}

func GetDrawId() int {
	drawId := 0
	//	params := "?req={\"gameId\":\"10002\"}"
	params := "{\"gameId\":\"10002\"}"
	res := HttpReq(data.QueryDrawUrl, params)
	drawId = data.ParseDrawId(res)
	return drawId
}
