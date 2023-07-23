package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"

	"github.com/Jeffail/gabs"

	"github.com/jwtly10/spotify_switcher/auth"
	"github.com/jwtly10/spotify_switcher/scraper"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Track struct {
	song   string
	artist string
}

func main() {

	spotifyToken := auth.GetAuthToken()
	savedTracks := getPlaylistTracks(spotifyToken, os.Getenv("PLAYLIST_ID"))

	fmt.Println(strconv.Itoa(len(savedTracks)) + " Songs loaded from Spotify")

	// Load Browser
	url, _ := launcher.NewUserMode().Launch()
	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage("https://music.apple.com/gb/search").MustWaitLoad()

	loaded := 0
	for _, track := range savedTracks {
		loaded = loaded + scraper.ScrapeAppleMusic(page, track.song, track.artist)
	}

	fmt.Println(strconv.Itoa(loaded) + " songs loaded")
}

func getPlaylistTracks(apiToken string, playlistID string) []Track {
	fmt.Println("Getting playlist tracks...")
	endpoint := "https://api.spotify.com/v1/playlists/" + playlistID + "/tracks"

	results := []Track{}
	results = recursivelyGetResults(results, endpoint, apiToken)

	return results
}

func recursivelyGetResults(results []Track, url_call string, apiToken string) []Track {
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

		artist := strings.Join(artists, ", ")
		track := Track{name, artist}
		results = append(results, track)
		artists = nil
	}

	if jsonParsed.Search("next").Data() != nil {
		results = recursivelyGetResults(results, jsonParsed.Search("next").Data().(string), apiToken)
	}

	return results
}

func writeResultsToCSV(result []Track) {
	csvFile, err := os.Create("results.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	for _, row := range result {
		_, err := csvFile.WriteString(row.song + ", " + row.artist + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
