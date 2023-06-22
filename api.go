package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func funcDB(action string, index int) interface{} {
	personList := []Person{{"vishal", 24.5, 1996}, {"Amit", 25.5, 1992}, {"Manuwela", 26.5, 1991}, {"Anaya", 27.5, 1990}, {"Rahul", 28.5, 1990}}

	if index >= len(personList) {
		return -1
	}
	if action == "getlength" {
		return len(personList)
	} else {
		name := personList[index].Name.getInfo()
		age := personList[index].Age.getInfo()
		year := personList[index].Year.getInfo()
		return Person{Name: name, Age: age, Year: year}
	}
}

func Thread1(i int, wg *sync.WaitGroup, isDataAvailable chan bool, personChannel chan string) {
	funcData := funcDB("getData", i)

	val, exist := funcData.(int)
	// fmt.Println("data exist ? : ", exist)
	if exist && val == -1 {
		personChannel <- ""
		isDataAvailable <- false
	} else {
		data := Person{Name: funcData.(Person).Name, Age: funcData.(Person).Age, Year: funcData.(Person).Year}
		// fmt.Println("funcDB data : ", data)

		p1, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Marshaling error : ", err)
		}

		fmt.Println("marshalled data : ", string(p1))
		personChannel <- string(p1)
		isDataAvailable <- true
	}
}

func Thread2(data chan string, wg *sync.WaitGroup, isDataAvailable chan bool, ctx *gin.Context) {
	for {
		select {
		case out := <-data:
			fmt.Println("isDataAvailable: ", out)
			if len(data) != 0 {
				var unMarshalledPerson Person
				if len(data) == 0 {
					fmt.Println("Error: Empty JSON Received")
				}

				chanData := <-data
				err := json.Unmarshal([]byte(chanData), &unMarshalledPerson)
				if err != nil {
					fmt.Println("Unmarshaling error : ", err)
				}
				fmt.Println("unmarshalled data : ", unMarshalledPerson)
			}
		}
	}
}
