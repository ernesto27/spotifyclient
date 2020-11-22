# Spotify client

Go package that obtaint metadata about the current playing song and status on your desktop app Spotify app, also manipulate player options ( play/pause, volume , etc)

At the moment is support only the macOS operating system ,  this package uses the appleScript spotify commands to works.

https://en.wikipedia.org/wiki/AppleScript

In the future will be implemented on Linux also


In the future will be implemented on Linux also


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

	v := reflect.ValueOf(meta)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
    }

    // Print artist name
    fmt.Println(meta.ArtistName[])
    
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
    Field: AlbumArtist	Value: Europe
    Field: ArtworlURL	Value: https://i.scdn.co/image/ab67616d0000b2732d925cec3072ed1b74e5188f
    Field: SpotifyURL	Value: spotify:track:3MrRksHupTVEQ7YbA0FsZK
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















