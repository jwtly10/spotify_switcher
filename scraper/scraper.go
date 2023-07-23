package scraper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Test() string {
	return "test"
}

func ScrapeAppleMusic(page *rod.Page, song string, artist string) int {
	var songFound bool
	var songIndex int

	// clear this just in case
	time.Sleep(1 * time.Second)
	page.MustSearch(".search-input__text-field").MustClick().MustType(input.Escape).MustWaitLoad()
	time.Sleep(1 * time.Second)

	// Search for song
	page.MustSearch(".search-input__text-field").MustClick().MustInput(song + ", " + artist).MustType(input.Enter).MustWaitLoad()
	// page.MustScreenshot("searchSong.png")

	fmt.Println("Searching for " + song + ", " + artist)
	time.Sleep(1 * time.Second)

	// Check if song is found
	potentialMatches := page.MustElements(".top-search-lockup__primary")

out:
	for index, b := range potentialMatches {
		t, err := b.Text()
		if err != nil {
			fmt.Println(err)
		}
		// Checking that the song of this element is similar to the song we are looking for
		if strings.Contains(t, song) {
			songFound = true
			songIndex = index + 1
			fmt.Println("Index: " + strconv.Itoa(songIndex))
			break out
		}

	}

	if !songFound {
		fmt.Println("Song not found")
		return 0
	}

	// page.MustScreenshot("songFound.png")

	page.MustSearch("/html[@class='js-focus-visible']/body/div[@class='app-container svelte-197510h']/div[@id='scrollable-page']/main[@class='svelte-11bt8wm']/div[@class='content-container svelte-11bt8wm']/div[@class='desktop-search-page svelte-e9u219']/div[@class='section svelte-ubaf1n with-top-spacing']/div[@class='section-content svelte-ubaf1n']/ul[@class='grid svelte-1ntv2c0 grid--flow-row grid--custom-columns grid--top-results']/li[@class='grid-item svelte-1ntv2c0'][" + strconv.Itoa(songIndex) + "]/div[@class='top-search-lockup svelte-1an0vgx']/div[@class='top-search-lockup__icons svelte-1an0vgx']/div[@class='cloud-buttons svelte-199vsqz']/button[@class='add-to-library-button svelte-1wzng6v add-to-library-button--icon-only']").MustClick()

	fmt.Println("Added " + song + ", " + artist + " to library")

	// page.MustScreenshot("songAdded.png")

	return 1
	// browser.MustClose()
}
