package main

import (
	"fmt"
	"strconv"
	"time"
)

// GOROUTINE RUNNER
func HandleCompany(c chan int) {
	c <- 1
	//------------HANDLE------------
	row := company[0]
	company = company[1:]
	//HANDLE COMPANY
	fmt.Println("HANDLE " + row.data)
	//
	time.Sleep(time.Second * 3)
	// CREATE GOROUTINE TO HANDLER JOB
	//
	//
	//	
	//	CONTINUE HANDLER JOB...
	//
	//
	//
	//JOBJOBJOBJOBJOBJOBJOBJOBJBOBJOB
	//---------------- ---------------
	<-c
}

func main() {
	buffer := 3
	channelCompany := make(chan int, buffer)

	for len(company) > 0 {
		for len(channelCompany) < buffer {
			go HandleCompany(channelCompany)
			fmt.Println("LEN BUFF " + strconv.Itoa(len(channelCompany)))
		}
	}
	time.Sleep(5 * time.Second)
}

type row struct {
	data string
}

// FAKE DATABASE COMPANY
var company = []row{
	row{
		data: "DacTech",
	},
	row{
		data: "Fsoft",
	},
	row{
		data: "Mgm",
	},
	row{
		data: "Sun*",
	},
	row{
		data: "Datahouse",
	},
	row{
		data: "Openweb",
	},
	row{
		data: "Kms",
	},
	row{
		data: "Google Inc",
	},
	row{
		data: "Facebook Inc",
	},
}
