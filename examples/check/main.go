package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamboto2000/ccc"
)

func main() {
	card, err := ccc.Check("537941")
	if err != nil {
		panic(err.Error())
	}

	js, err := json.Marshal(card)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(js))
}
