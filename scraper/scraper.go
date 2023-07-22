package scraper

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Test() string {
	return "test"
}

func ScrapeAppleMusic(page *rod.Page, song string, artist string) int {
	var browserSong string

	// Search for song
	page.MustSearch(".search-input__text-field").MustClick().MustInput(song + ", " + artist).MustType(input.Enter).MustWaitLoad()
	// page.MustScreenshot("searchSong.png")

	fmt.Println("Searching for " + song + ", " + artist)
	time.Sleep(1 * time.Second)
	// Check if song is found
	potentialMatches := page.MustElements(".top-search-lockup__primary")

	for _, b := range potentialMatches {
		t, err := b.Text()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Check name: " + t)
		if t == song {
			browserSong = t
		}

	}

	// If not found skip
	if browserSong != song {
		fmt.Println(browserSong)
		fmt.Println("Song not found")
		return 0
	}
	// Else add song to library

	// clear search first
	// page.MustSearch(".search-input__cancel-button").MustClick()

	// page.MustScreenshot("songFound.png")

	page.MustSearch("/html[@class='js-focus-visible']/body/div[@class='app-container svelte-197510h']/div[@id='scrollable-page']/main[@class='svelte-11bt8wm']/div[@class='content-container svelte-11bt8wm']/div[@class='desktop-search-page svelte-e9u219']/div[@class='section svelte-ubaf1n with-top-spacing']/div[@class='section-content svelte-ubaf1n']/ul[@class='grid svelte-1ntv2c0 grid--flow-row grid--custom-columns grid--top-results']/li[@class='grid-item svelte-1ntv2c0'][1]/div[@class='top-search-lockup svelte-1an0vgx']/div[@class='top-search-lockup__icons svelte-1an0vgx']/div[@class='cloud-buttons svelte-199vsqz']/button[@class='add-to-library-button svelte-1wzng6v add-to-library-button--icon-only']").MustClick()

	// page.MustScreenshot("songAdded.png")

	return 1
	// browser.MustClose()
}
