package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type AuthResponse struct {
	Access_token string
	Token_type   string
	Expires_in   int
}

func GetAuthToken() string {
	endpoint := "https://accounts.spotify.com/api/token"

	authOptions := url.Values{}
	authOptions.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(authOptions.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID")+":"+os.Getenv("CLIENT_SECRET"))))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var authResponse AuthResponse
	json.Unmarshal(responseData, &authResponse)

	return authResponse.Access_token
}
