package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// respose, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	// resposeData, err := ioutil.ReadAll(respose.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(resposeData))
}

func getClientID() string {
	return os.Getenv("CLIENT_ID")
}

func getClientSecret() string {
	return os.Getenv("CLIENT_SECRET")
}
