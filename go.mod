module github.com/jwtly10/spotify_switcher

go 1.20

require (
	github.com/Jeffail/gabs v1.4.0
	github.com/jwtly10/spotify_switcher/auth v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.5.1 // indirect

replace github.com/jwtly10/spotify_switcher/auth => ./auth
