package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamboto2000/ccc"
)

func main() {
	card, err := ccc.CheckWithProx("537941", "http://185.198.188.55:8080")
	if err != nil {
		panic(err.Error())
	}

	js, err := json.Marshal(card)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(js))
}
