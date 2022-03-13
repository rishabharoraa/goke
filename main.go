package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// generated go struct
type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

func main() {

	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any?blacklistFlags=nsfw,religious,political,racist,sexist,explicit&type=single")

	if err != nil {
		fmt.Println("Error: could not fetch joke :(")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: could not fetch joke :(")
		os.Exit(1)
	}

	debug := string(body)
	fmt.Println(debug)

	var jokeResponse JokeResponse
	if err := json.Unmarshal(body, &jokeResponse); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(jokeResponse.Joke)

}
