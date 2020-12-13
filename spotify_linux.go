// +build linux

package spotifyclient

import (
	"fmt"
	"log"
	"reflect"

	"github.com/godbus/dbus"
)

var conn *dbus.Conn

func init() {
	conn = getConn()
}

func getConn() *dbus.Conn {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

const (
	sender = "org.mpris.MediaPlayer2.spotify"
	path   = "/org/mpris/MediaPlayer2"
	member = "org.mpris.MediaPlayer2.Player"

	metadataMessage       = member + ".Metadata"
	playMessage           = member + ".Play"
	pauseMessage          = member + ".Pause"
	playPauseMessage      = member + ".PlayPause"
	playbackStatusMessage = member + ".PlaybackStatus"
	nextMessage           = member + ".Next"
	previousMessage       = member + ".Previous"
	openURI               = member + ".OpenUri"
)

// State of Spotify app
type State struct {
	TrackID  string
	Volume   int
	Position int
	State    string
}

// Metadata contains Spotify player metadata
type SpotifyMetadata struct {
	ArtistName  []string `spotify:"xesam:artist"`
	AlbumName   string   `spotify:"xesam:album"`
	TrackName   string   `spotify:"xesam:title"`
	DiscNumber  int32    `spotify:"xesam:discNumber"`
	Duration    uint64   `spotify:"mpris:length"`
	PlayedCount int32
	TrackNumber int32 `spotify:"xesam:trackNumber"`
	Popularity  int32
	ID          string `spotify:"mpris:trackid"`
	ArtworkURL  string
	URL         string `spotify:"xesam:url"`
}

// parseMetadata returns a parsed Metadata struct
func parseMetadata(variant dbus.Variant) *SpotifyMetadata {
	metadataMap := variant.Value().(map[string]dbus.Variant)
	metadataStruct := new(SpotifyMetadata)

	valueOf := reflect.ValueOf(metadataStruct).Elem()
	typeOf := reflect.TypeOf(metadataStruct).Elem()

	for key, val := range metadataMap {
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			if field.Tag.Get("spotify") == key {
				field := valueOf.Field(i)
				field.Set(reflect.ValueOf(val.Value()))
			}
		}
	}

	return metadataStruct
}

// GetMetadata returns the current metadata from the Spotify app
func GetCurrentTrack() (*SpotifyMetadata, error) {
	obj := conn.Object(sender, path)
	property, err := obj.GetProperty(metadataMessage)
	if err != nil {
		return nil, err
	}

	return parseMetadata(property), nil
}

func Play() {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)
		c := obj.Call(playMessage, 0)
		if c.Err != nil {
			fmt.Println("Error on play")
		}
	}
}

func Pause() {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)
		c := obj.Call(pauseMessage, 0)
		if c.Err != nil {
			fmt.Println("Error on pause")
		}
	}
}

func PlayPause() {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)
		c := obj.Call(playPauseMessage, 0)
		if c.Err != nil {
			fmt.Println("Error on Play Pause")
		}
	}
}

func VolumeUp() {
	fmt.Println("Not support on linux")
}

func VolumeDown() {
	fmt.Println("Not support on linux")
}

func SetVolume(value int) {
	fmt.Println("Not support on linux")
}

func PlayTrack(value string) {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)

		c := obj.Call(openURI, 0, value)
		if c.Err != nil {
			fmt.Println("Error on OpenURI")
		}
	}
}

func GetState() (State, error) {
	obj := conn.Object(sender, path)
	property, err := obj.GetProperty(playbackStatusMessage)

	data := State{}

	if err != nil {
		data.State = "Unknown"
		return data, err
	}

	fmt.Println(property)

	data.State = property.String()

	return data, nil
}

func Next() {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)
		c := obj.Call(nextMessage, 0)
		if c.Err != nil {
			fmt.Println("Error on Next")
		}
	}
}

func Prev() {
	started, err := IsServiceStarted(conn)
	if err != nil {
		fmt.Println("Spotify app is not open")
	} else if started {
		obj := conn.Object(sender, path)
		c := obj.Call(previousMessage, 0)
		if c.Err != nil {
			fmt.Println("Error on Prev")
		}
	}
}

// IsServiceStarted checks if the Spotify app is running
func IsServiceStarted(conn *dbus.Conn) (bool, error) {
	started := false

	err := conn.Object(
		"org.freedesktop.DBus",
		"/org/freedesktop/DBus",
	).Call(
		"org.freedesktop.DBus.NameHasOwner",
		0,
		sender,
	).Store(
		&started,
	)
	if err != nil {
		return false, err
	}

	return started, nil
}
