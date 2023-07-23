# spotify_switcher
Application to make switching music services easier

Calls the Spotify to pull tracks from a playlist and outputs a CSV list that can be used as input for a webscraper to collate a new playlist in another streaming platform. 

Utilizes Go Web Scraping Packages to add songs in Apple Music

Required .env vars
CLIENT_ID - Spotify Developer Client ID
CLIENT_SECRET - Spotify Developer Client Secret
PLAYLIST_ID - Playlist ID
