package workers

import (
	"gtprcheck/data"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	//	"strings"
)

func BetSumWorker(qsum <-chan data.SumBase, qfile chan<- map[int]data.Sum) {
	var ts map[int]data.Sum
	ts = make(map[int]data.Sum)
	var s data.Sum
	for {
		sBase := <-qsum
		if v, ok := ts[100]; ok {
			s = v
			s.Count += 1
			s.Amount += sBase.BetAmount
		} else {
			s.Count = 1
			s.Amount = sBase.BetAmount
		}
		ts[100] = s

		if v, ok := ts[sBase.PrizeLevel]; ok {
			s = v
			s.Count += 1
			s.Amount += sBase.PrizeAmount
		} else {
			s.Count = 1
			s.Amount = sBase.PrizeAmount
		}
		ts[sBase.PrizeLevel] = s

		if ts[100].Count == 10000 {
			fmt.Println("map 100 before clean: ", ts)
			tmp := make(map[int]data.Sum)
			for k, v := range ts {
				tmp[k] = v
			}
			qfile <- tmp
			for k, _ := range ts {
				ts[k] = data.Sum{Count: 0, Amount: 0.0}
			}
		}
	}
}
