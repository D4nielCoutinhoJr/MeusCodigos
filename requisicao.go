package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	User User `json: "user"`
}

type User struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Sectors    []Sectors  `json:"sectors"`
	Permission Permission `json:"permission"`
	Token      string     `json:"token"`
	Disabled   bool       `json:"disabled"`
}

type Sectors struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	const Url = "http://188.166.125.5:9090/v1/authenticate"

	body := []byte(`{"email":"admin@admin.com","password":"12345"}`)

	res, err := http.Post(Url, "Json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var response Response
	err = json.Unmarshal(resBody, &response)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
