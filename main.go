package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/Jeffail/gabs"

	"github.com/jwtly10/spotify_switcher/auth"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	spotifyToken := auth.GetAuthToken()

	savedTracks := getPlaylistTracks(spotifyToken, os.Getenv("PLAYLIST_ID"))

	writeResultsToCSV(savedTracks)
	fmt.Println("Done!")

}

func getPlaylistTracks(apiToken string, playlistID string) []string {
	fmt.Println("Getting playlist tracks...")
	endpoint := "https://api.spotify.com/v1/playlists/" + playlistID + "/tracks"

	results := []string{}
	results = recursivelyGetResults(results, endpoint, apiToken)

	// fmt.Println(len(results))

	// for x := 0; x < 10; x++ {
	// 	fmt.Println(results[x])
	// }

	return results
}

func recursivelyGetResults(results []string, url_call string, apiToken string) []string {
	req, err := http.NewRequest("GET", url_call, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+apiToken)

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

	jsonParsed, err := gabs.ParseJSON(responseData)
	if err != nil {
		log.Fatal(err)
	}

	sizeOfResults, err := jsonParsed.Search("items", "track").ArrayCount()

	artists := []string{}
	for x := 0; x < sizeOfResults; x++ {
		name := strings.ReplaceAll(jsonParsed.Search("items", "track").Index(x).Path("name").Data().(string), ",", " ")
		numberOfArtists, err := jsonParsed.Search("items", "track").Index(x).Path("artists").ArrayCount()
		if err != nil {
			log.Fatal(err)
		}

		for y := 0; y < numberOfArtists; y++ {
			artists = append(artists, jsonParsed.Search("items", "track").Index(x).Path("artists").Index(y).Path("name").Data().(string))
		}
		results = append(results, name+", "+strings.Join(artists, ", "))
		artists = nil
	}

	if jsonParsed.Search("next").Data() != nil {
		results = recursivelyGetResults(results, jsonParsed.Search("next").Data().(string), apiToken)
	}

	return results
}

func writeResultsToCSV(result []string) {
	csvFile, err := os.Create("results.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	for _, row := range result {
		_, err := csvFile.WriteString(row + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
