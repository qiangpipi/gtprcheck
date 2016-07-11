package main

import (
	"fmt"
	"gtprcheck/data"
	"gtprcheck/workers"
	"os"
	"os/signal"
	"runtime"
)

var NumT10001 int = 50
var NumT10002 int = 10
var NumT10003 int = 50

func main() {
	//Signal channel for stop
	runtime.GOMAXPROCS(runtime.NumCPU())
	cs := make(chan os.Signal, 1)
	//Queues for game 10001#################################################
	//resp10001 := make(chan data.SumBase, 1000)
	//sumfile10001 := make(chan map[int]data.Sum, 10)

	//Queues for game 10002#################################################
	//resp10002 := make(chan data.SumBase, 1000)
	//sumfile10002 := make(chan map[int]data.Sum, 10)

	//Queues for game 10003#################################################
	resp10003 := make(chan data.SumBase, 1000)
	sumfile10003 := make(chan map[int]data.Sum, 10)

	//Get draw Id and set global var
	data.DrawId = workers.GetDrawId()

	//Define behavior for signal Interrupt and Kill
	signal.Notify(cs, os.Interrupt, os.Kill)

	//File for sumdata #####################################################
	f10003, err := os.OpenFile("f10003.csv", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0666)
	if err == nil {
		go workers.FileWorker(f10003, sumfile10003, 10003)
	} else {
		fmt.Println("File open error: ", err)
	}
	//Start http req threads and put formated response to respone chan
	for i := 0; i < NumT10003; i++ {
		go workers.ReqWorker(resp10003, data.BetReq10003, data.BetAmount1)
	}

	//Start sumarizing threads
	//Get response from response chan
	//and put sum result to sumarized chan
	for i := 0; i < 5; i++ {
		go workers.BetSumWorker(resp10003, sumfile10003)
	}

	//Start file thread
	//Get sum result from sumarized chan
	//and write data into csv file

	s := <-cs
	fmt.Println("\nSignal received:", s)
	defer f10003.Close()
	//Stop http req threads

	//Stop sumarizing threads

	//Stop file thread
}
