package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/protocol"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var restData []protocol.Person
var count int32

func runRestClient() {
	isDataAvailable := make(chan bool, 1)
	personChannel := make(chan string, 1)
	wg := &sync.WaitGroup{}
	i := 0

	funcDBLength := funcDB("getlength", 0)
	wg.Add(1)

	for i < funcDBLength.(int)+1 {
		go Thread1(i, wg, isDataAvailable, personChannel)
		out := <-isDataAvailable
		if out {
			fmt.Println("isDataAvailable: ", out)
			fmt.Println("Performing Http Post..., ", i)
			personData := <-personChannel
			responseBody := bytes.NewBuffer([]byte(personData))

			resp, err := http.Post("http://localhost:8082/vishal", "application/json; charset=utf-8", responseBody)
			if err != nil {
				log.Fatalln(err)
			}
			// fmt.Println("response log : ", resp)
			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)

			_ = string(bodyBytes)
		} else {
			fmt.Println("Data is not available")
		}
		time.Sleep(10 * time.Millisecond)
		i++
	}
	wg.Wait()
}

func runRestServer(ctx *gin.Context) {
	var input Person
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("rest server data : ", input)

	if input.Name != "" {
		go runGrpcClient(input)

	} else {
		fmt.Println("Empty data recieved")
	}
}
