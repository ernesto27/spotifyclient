// +build darwin

package spotifyclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
)

// Metadata contains Spotify player metadata
type SpotifyMetadata struct {
	ArtistName  []string `json:"artist"`
	AlbumName   string   `json:"album"`
	TrackName   string   `json:"name"`
	DiscNumber  string   `json:"disc_number"`
	Duration    int      `json:"duration"`
	PlayedCount int      `json:"played_count"`
	TrackNumber int      `json:"track_number"`
	Popularity  int      `json:"popularity"`
	ID          string   `json:"id"`
	ArtworkURL  string   `json:"artwork_url"`
	URL         string   `json:"spotify_url"`
}

// State of Spotify app
type State struct {
	TrackID  string `json:"track_id"`
	Volume   int    `json:"volume"`
	Position int    `json:"position"`
	State    string `json:"state"`
}

func GetCurrentTrack() (SpotifyMetadata, error) {
	command := `
		on escape_quotes(string_to_escape)
			set AppleScript's text item delimiters to the "\""
			set the item_list to every text item of string_to_escape
			set AppleScript's text item delimiters to the "\\\""
			set string_to_escape to the item_list as string
			set AppleScript's text item delimiters to ""
			return string_to_escape
		end escape_quotes

		tell application "Spotify"
			set ctrack to "{"
			set ctrack to ctrack & "\"artist\": \"" & my escape_quotes(current track's artist) & "\""
			set ctrack to ctrack & ",\"album\": \"" & my escape_quotes(current track's album) & "\""
			set ctrack to ctrack & ",\"disc_number\": " & current track's disc number
			set ctrack to ctrack & ",\"duration\": " & current track's duration
			set ctrack to ctrack & ",\"played_count\": " & current track's played count
			set ctrack to ctrack & ",\"track_number\": " & current track's track number
			set ctrack to ctrack & ",\"popularity\": " & current track's popularity
			set ctrack to ctrack & ",\"id\": \"" & current track's id & "\""
			set ctrack to ctrack & ",\"name\": \"" & my escape_quotes(current track's name) & "\""
			set ctrack to ctrack & ",\"album_artist\": \"" & my escape_quotes(current track's album artist) & "\""
			set ctrack to ctrack & ",\"artwork_url\": \"" & current track's artwork url & "\""
			set ctrack to ctrack & ",\"spotify_url\": \"" & current track's spotify url & "\""
			set ctrack to ctrack & "}"
		end tell
		`
	info := runAppleScript(command)
	fmt.Println(info)

	r := regexp.MustCompile("\"artist\": \"([a-zA-Z 0-9 '/]+)\"")
	res := r.ReplaceAllString(info, "\"artist\": [\"$1\"]")

	data := SpotifyMetadata{}
	json.Unmarshal([]byte(res), &data)
	return data, nil
}

func Play() {
	command := `tell application "Spotify" to play`
	runAppleScript(command)
}

func Pause() {
	command := `tell application "Spotify" to pause`
	runAppleScript(command)
}

func VolumeUp() {
	command := `
		on min(x, y)
		if x ≤ y then
			return x
		else
			return y
		end if
		end min

		tell application "Spotify" to set sound volume to (my min(sound volume + 10, 100))
	`
	runAppleScript(command)
}

func VolumeDown() {
	command := `
		on max(x, y)
		if x ≤ y then
			return y
		else
			return x
		end if
		end max
	
		tell application "Spotify" to set sound volume to (my max(sound volume - 10, 0))
	`
	runAppleScript(command)
}

func PlayPause() {
	command := `tell application "Spotify" to playpause`
	runAppleScript(command)
}

func SetVolume(value int) {
	command := fmt.Sprintf(`tell application "Spotify" to set sound volume to %d`, value)
	runAppleScript(command)
}

func PlayTrack(value string) {
	command := fmt.Sprintf(`tell application "Spotify" to play track "%s"`, value)
	runAppleScript(command)
}

func Next() {
	command := `tell application "Spotify" to next track`
	runAppleScript(command)
}

func Prev() {
	command := `tell application "Spotify" to previous track`
	runAppleScript(command)
}

func GetState() (State, error) {
	command := `
		tell application "Spotify"
			set cstate to "{"
			set cstate to cstate & "\"track_id\": \"" & current track's id & "\""
			set cstate to cstate & ",\"volume\": " & sound volume
			set cstate to cstate & ",\"position\": " & (player position as integer)
			set cstate to cstate & ",\"state\": \"" & player state & "\""
			set cstate to cstate & "}"
	
		return cstate
		end tell
  
	`
	info := runAppleScript(command)
	fmt.Println(info)
	data := State{}
	json.Unmarshal([]byte(info), &data)
	return data, nil
}

func runAppleScript(command string) string {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	info := out.String()
	return info
}
