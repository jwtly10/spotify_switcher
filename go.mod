module github.com/jwtly10/spotify_switcher

go 1.20

require github.com/Jeffail/gabs v1.4.0

require (
	github.com/joho/godotenv v1.5.1
	github.com/jwtly10/spotify_switcher/scraper v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-rod/rod v0.114.0 // indirect
	github.com/ysmood/fetchup v0.2.3 // indirect
	github.com/ysmood/goob v0.4.0 // indirect
	github.com/ysmood/got v0.34.1 // indirect
	github.com/ysmood/gson v0.7.3 // indirect
	github.com/ysmood/leakless v0.8.0 // indirect
)

replace github.com/jwtly10/spotify_switcher/auth => ./auth

replace github.com/jwtly10/spotify_switcher/scraper => ./scraper
