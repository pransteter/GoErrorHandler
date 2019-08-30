package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type parameters struct {
	Soma soma `json:"soma"`
}

type soma struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
	Result  int `json:"result"`
}

func main() {
	defer ErrorHandler()

	jsonFile, fileErr := os.Open("parameters.json")

	if fileErr != nil {
		log.Println(fileErr)
	}

	byteJsonFile, byteJsonErr := ioutil.ReadAll(jsonFile)

	if byteJsonErr != nil {
		log.Println(byteJsonErr)
	}

	parameters := parameters{}

	jsonErr := json.Unmarshal(byteJsonFile, &parameters)

	if jsonErr != nil {
		log.Println(jsonErr)
	}

	result := parameters.Soma.Number1 + parameters.Soma.Number2

	if result != parameters.Soma.Result {
		panic(CreateException(GENERIC_ERROR, "A soma está errada."))
	} else {
		log.Println("Está certo.")
	}

	log.Printf("O resultado é: %v\n", result)

	defer jsonFile.Close()
}
