# Spotify client

Go package that obtaint metadata about the current playing song and status on your desktop app Spotify app, also manipulate player options ( play/pause, volume , etc)

Works on Linux and MacOS .


## Instalation
```
$ go get github.com/ernesto27/spotifyclient
```

## Example
Get metadata about current song track playing.
```go
package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ernesto27/spotifyclient"
)
func main() {

	// Example GetMetadata from current song playing
	meta, err := spotifyclient.GetCurrentTrack()
	if err != nil {
		fmt.Println("Seems that you don't have the spotify app desktop installed  or is not open :(")
		log.Fatalf("failed getting metadata, err: %s", err.Error())
	}

    // Print artist name
    fmt.Println(meta.ArtistName[0])
    
    // RESPONSE
    /*
    Field: ArtistName	Value: [Europe]
    Field: AlbumName	Value: The Final Countdown (Expanded Edition)
    Field: TrackName	Value: The Final Countdown
    Field: DiscNumber	Value:
    Field: Duration	Value: 310333
    Field: PlayedCount	Value: 0
    Field: TrackNumber	Value: 1
    Field: Popularity	Value: 77
    Field: ID	Value: spotify:track:3MrRksHupTVEQ7YbA0FsZK
    Field: ArtworkURL	Value: https://i.scdn.co/image/ab67616d0000b2732d925cec3072ed1b74e5188f
    Field: URL	Value: spotify:track:3MrRksHupTVEQ7YbA0FsZK
    */

```
#
## Methods

### Get current track
```go
meta, err := spotifyclient.GetCurrentTrack()
if err != nil {
    fmt.Println("Seems that you don't have the spotify app desktop installed  or is not open :(")
    log.Fatalf("failed getting metadata, err: %s", err.Error())
}
fmt.Println(meta)
```


### Get state
```go
state, err := spotifyclient.GetState()
if err != nil {
	log.Fatalf("err: %s", err.Error())
}
fmt.Println(state)
```

### Play song
```go
spotifyclient.Play()
```

### Pause song
```go
spotifyclient.Pause()
```

### Volumen up
```go
spotifyclient.VolumeUp()
```

### Volumen down
```go
spotifyclient.VolumeDown()
```

### Play pause
```go
spotifyclient.PlayPause()
```

### Set volumen
```go
spotifyclient.SetVolume(50)
```

### Next track
```go
spotifyclient.Next()
```

### Prev track
```go
spotifyclient.Prev()
```

### Play url track
```go
spotifyclient.PlayTrack("spotify:track:5ponLS88v37duDQHewRDaX")
```















