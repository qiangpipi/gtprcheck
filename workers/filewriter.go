package workers

import (
	"fmt"
	"gtprcheck/data"
	"os"
	"strconv"
	"time"
)

var id10001 []int = []int{100, 0, 11, 12, 13, 14, 15}
var id10002 []int = []int{100, 0, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
var id10003 []int = []int{100, 0, 31, 32, 33, 34, 35, 36, 37, 38, 39}

func FileWorker(f *os.File, q <-chan map[int]data.Sum, gi int) {
	var tmp string
	if gi == 10001 {
		tmp = "BetCount,BetAmount,0Count,0Amount,11Count,11Amount,12Count,12Amount,13Count,13Amount,14Count,14Amount,15Count,15Amount\n"
	} else if gi == 10002 {
		tmp = "BetCount,BetAmount,0Count,0Amount,21Count,21Amount,22Count,22Amount,23Count,23Amount,24Count,24Amount,25Count,25Amount,26Count,26Amount,27Count,27Amount,28Count,28Amount,29Count,29Amount,30Count,30Amount\n"
	} else if gi == 10003 {
		tmp = "BetCount,BetAmount,0Count,0Amount,31Count,31Amount,32Count,32Amount,33Count,33Amount,34Count,34Amount,35Count,35Amount,36Count,36Amount,37Count,37Amount,38Count,38Amount,39Count,39Amount\n"
	} else if gi == 20004 {
	}
	if _, err := f.WriteString(tmp); err != nil {
		fmt.Println("Write file header error: ", err)
	}
	for {
		tmp = ""
		sd := <-q
		if gi == 10001 {
			for _, v := range id10001 {
				tmp += strconv.Itoa(sd[v].Count)
				tmp += ", "
				tmp += strconv.FormatFloat(sd[v].Amount, 'G', -1, 64)
				tmp += ", "
			}
		} else if gi == 10002 {
			for _, v := range id10002 {
				tmp += strconv.Itoa(sd[v].Count)
				tmp += ", "
				tmp += strconv.FormatFloat(sd[v].Amount, 'G', -1, 64)
				tmp += ", "
			}
		} else if gi == 10003 {
			for _, v := range id10003 {
				tmp += strconv.Itoa(sd[v].Count)
				tmp += ", "
				tmp += strconv.FormatFloat(sd[v].Amount, 'G', -1, 64)
				tmp += ", "
			}
		} else if gi == 20004 {
		}
		tmp += time.Now().Format("2006-01-02T15:04:05.999")
		tmp += "\n"
		if _, err := f.WriteString(tmp); err != nil {
			fmt.Println("Write file data error: ", err)
		}
	}
}
