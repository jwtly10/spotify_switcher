# spotify_switcher
GO Application to make switching music services easier

I wanted to try Apple Music but could never use it properly as most of my music playlist was missing. This CLI tool fixes that. 

It uses Spotify dev API to pull playlist data, and scrapes the Apple Music site to save all of the songs from Spotify. Can also be set to headless where an instance of Chrome is opened, but hidden.

Required .env vars
CLIENT_ID - Spotify Developer Client ID
CLIENT_SECRET - Spotify Developer Client Secret
PLAYLIST_ID - Playlist ID

Demo:

https://github.com/jwtly10/spotify_switcher/assets/39057715/4a802c27-e8f9-4c5d-8b27-2c703d59a12c


<img width="1664" alt="spotify-switcher" src="https://github.com/jwtly10/spotify_switcher/assets/39057715/cf18af2c-a08c-47b5-94dc-daf33d40a35a">
