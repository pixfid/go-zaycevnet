package api

import (
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var (
	client *ZClient
)

func TestNewZClient(t *testing.T) {
	client = NewZClient("", "kmskoNkYHDnl3ol2")
	client.Auth()
	println("Token:", client.accessToken)
}

func TestZClient_TOP(t *testing.T) {
	result, err := client.TOP(1)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}
	t.Logf("Artist Name %s.", result.Tracks[0].ArtistName)
	t.Logf("Track  Name %s.", result.Tracks[0].Track)
	t.Logf("Track  Date %d.", result.Tracks[0].Date)
	t.Logf("Track  Bitrate %d.", result.Tracks[0].Bitrate)
	t.Logf("Track  Duration %s.", result.Tracks[0].Duration)

}

func TestZClient_MusicSetList(t *testing.T) {
	result, err := client.MusicSetList(0)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}
	t.Logf("Artist Name %s.", result.List[0].Name)
	t.Logf("Track  Count %d.", result.List[0].TracksCount)
	t.Logf("Track  ID %d.", result.List[0].ID)
	t.Logf("Track  URL %s.", result.List[0].URL)
}

func TestZClient_MusicSetDetile(t *testing.T) {
	result, err := client.MusicSetDetile(5735)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}
	t.Logf("Artist Name %s.", result.Tracks[0].ArtistName)
	t.Logf("Track  Name %s.", result.Tracks[0].Track)
	t.Logf("Track  Size %f.", result.Tracks[0].Size)
	t.Logf("Track  Bitrate %d.", result.Tracks[0].Bitrate)
	t.Logf("Track  Duration %s.", result.Tracks[0].Duration)
}

func TestZClient_Artist(t *testing.T) {
	result, err := client.Artist(70167)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}
	t.Logf("Artist Name %s.", result.Artist.Name)
	t.Logf("Artist ID %d.", result.Artist.ID)
}

func TestZClient_AutoComplete(t *testing.T) {
	result, err := client.AutoComplete("test")
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}
	t.Logf("Source Value %s.", "\"test\"")
	for i := 0; i < len(result.Terms); i++ {
		t.Logf("Founded: %s.", result.Terms[i])
	}
}

func TestZClient_Genre(t *testing.T) {
	result, err := client.Genre("pop", 1)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	t.Logf("Pages Count: %d.", result.PagesCount)
	t.Logf("Pages [1] Artist Name: %s.", result.Tracks[1].ArtistName)
	t.Logf("Pages [1] Track Name: %s.", result.Tracks[1].Track)
	t.Logf("Pages [1] Track ID: %d.", result.Tracks[1].ID)
}

func TestZClient_Track(t *testing.T) {
	result, err := client.Track(1950989)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	t.Logf("Rating : %f.", result.Rating)
	t.Logf("Artist Name: %s.", result.Track.ArtistName)
	t.Logf("Track Name: %s.", result.Track.Name)
	t.Logf("Music Author: %s.", result.Track.MusicAuthor)
	t.Logf("Right Possessors: %s.", result.Track.RightPossessors)
}

func TestZClient_Search(t *testing.T) {
	params := url.Values{}
	params.Add("query", "ZZ TOP")
	params.Add("page", strconv.Itoa(1))
	params.Add("type", "all")
	params.Add("sort", "")
	params.Add("style", "")

	result, err := client.Search(params)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	t.Logf("Pages Count: %d.", result.PagesCount)
	t.Logf("Pages [1] Artist Name: %s.", result.Tracks[1].ArtistName)
	t.Logf("Pages [1] Track Name: %s.", result.Tracks[1].Track)
	t.Logf("Pages [1] Track ID: %d.", result.Tracks[1].ID)
}

func TestZClient_Download(t *testing.T) {
	result, err := client.Download(1950989)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	t.Logf("URL : %s.", result.URL)
}

func TestZClient_Play(t *testing.T) {
	result, err := client.Play(1950989)
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	t.Logf("URL : %s.", result.URL)
}

func TestZClient_Options(t *testing.T) {
	result, err := client.Options()
	if err != nil {
		t.Errorf("Error %s.", err.Error())
	}

	options := strings.Split(result.Options, ";")
	for i := 0; i < len(options); i++ {
		t.Logf("Options : %s.", options[i])
	}
}

func TestMD5Hash(t *testing.T) {
	t.Logf("OK")
}
