package controllers

import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"cryptocur/utils"
	"cryptocur/model"
)

func FetchData() {
	endpoint := "blockfacts/price?asset=BTC&denominator=USD"

	cli := &http.Client{}
	request, err := http.NewRequest("GET", utils.BASE_URL + endpoint, nil)
	request.Header.Set("X-API-KEY", "XdWagWWwMVh5HQ76t5KYRU0gWvQxQf")
	request.Header.Set("X-API-SECRET", "QHmemdNJ9bp1mX6bI4F6KTV5IMUPAlqLGUPidhXVACrHw")
	if err != nil {
		panic(err)
	}
	response, err := cli.Do(request)

	if err != nil {
		fmt.Println("Error in api call")
		log.Fatal(err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	var bitcoinObject model.Bitcoin
    json.Unmarshal(responseBytes, &bitcoinObject)

    fmt.Println("Struct is:", bitcoinObject)
}
