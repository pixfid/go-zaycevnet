package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	apiURL            string = "https://api.zaycev.net/external"
	helloURL          string = apiURL + "/hello"
	authURL           string = apiURL + "/auth?"
	topURL            string = apiURL + "/top?"
	artistURL         string = apiURL + "/artist/%d?"
	musicSetListURL   string = apiURL + "/musicset/list?"
	musicSetDetileURL string = apiURL + "/musicset/detail?"
	genreURL          string = apiURL + "/genre?"
	trackURL          string = apiURL + "/track/%d?"
	feedbackURL       string = apiURL + "/feedback?"
	bugsURL           string = apiURL + "/bugs?"
	autoCompleteURL   string = apiURL + "/autocomplete?"
	searchURL         string = apiURL + "/search?"
	optionsURL        string = apiURL + "/options?"
	playURL           string = apiURL + "/track/%d/play?"
	downloadURL       string = apiURL + "/track/%d/download/?"
)

var (
	httpClient = http.DefaultClient
	errData    = []byte{'{', '}'}
)

type ZClient struct {
	accessToken string
	staticKey   string
}

func NewZClient(token, sKey string) *ZClient {
	return &ZClient{accessToken: token, staticKey: sKey}
}

func (zc *ZClient) Auth() {
	if zc.accessToken != "" {
		return
	}
	zc.hello()
}

//https://api.zaycev.net/external/hello
func (zc *ZClient) hello() {
	if zc.accessToken != "" {
		return
	}

	var t ZToken
	data, err := get(helloURL)
	if err != nil {

	}

	if err := t.parse(data); err != nil {

	}
	zc.auth(t.Token)
}

//https://api.zaycev.net/external/auth?code=%s&hash=%s
func (zc *ZClient) auth(token string) {

	var t ZToken
	hash := MD5Hash(token + zc.staticKey)

	params := url.Values{}
	params.Add("code", token)
	params.Add("hash", hash)

	uri := authURL + params.Encode()

	data, err := get(uri)

	if err != nil {

	}

	if err := t.parse(data); err != nil {

	}

	zc.accessToken = t.Token
}

//https://api.zaycev.net/external/search?query=%s&page=%s&type=%s&sort=%s&style=%s&access_token=%s
func (zc *ZClient) Search(params url.Values) (ZSearch, error) {
	return search(zc.accessToken, params)
}

func search(token string, params url.Values) (ZSearch, error) {
	var zSearch ZSearch
	params.Add("access_token", token)
	uri := searchURL + params.Encode()
	data, err := get(uri)
	if err != nil {
		return ZSearch{}, err
	}
	if err := zSearch.parse(data); err != nil {
		return ZSearch{}, err
	}
	return zSearch, nil
}

//https://api.zaycev.net/external/autocomplete?access_token=%s&code%s
func (zc *ZClient) AutoComplete(query string) (ZTerms, error) {
	return autoComplete(zc.accessToken, query)
}

func autoComplete(token, query string) (ZTerms, error) {
	var zTerms ZTerms

	params := url.Values{}
	params.Add("access_token", token)
	params.Add("query", query)

	uri := autoCompleteURL + params.Encode()
	data, err := get(uri)
	if err != nil {
		return ZTerms{}, err
	}
	if err := zTerms.parse(data); err != nil {
		return ZTerms{}, err
	}
	return zTerms, nil
}

//https://api.zaycev.net/external/top?page=%s&access_token=%s
func (zc *ZClient) TOP(page int) (ZTop, error) {
	return top(zc.accessToken, page)
}

func top(token string, page int) (ZTop, error) {
	var zTop ZTop
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("access_token", token)

	uri := topURL + params.Encode()
	data, err := get(uri)
	if err != nil {
		return ZTop{}, err
	}
	if err := zTop.parse(data); err != nil {
		return ZTop{}, err
	}
	return zTop, nil
}

//https://api.zaycev.net/external/musicset/list?page=%s&access_token=%s
func (zc *ZClient) MusicSetList(page int) (ZMusicSetList, error) {
	return musicSetList(zc.accessToken, page)
}

func musicSetList(token string, page int) (ZMusicSetList, error) {
	var zMusicSetList ZMusicSetList
	params := url.Values{}
	params.Add("access_token", token)
	params.Add("page", strconv.Itoa(page))

	uri := musicSetListURL + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZMusicSetList{}, err
	}
	if err := zMusicSetList.parse(data); err != nil {
		return ZMusicSetList{}, err
	}
	return zMusicSetList, nil
}

//https://api.zaycev.net/external/musicset/detail?id=%s&access_token=%s
func (zc *ZClient) MusicSetDetile(id int) (ZMusicSetDetile, error) {
	return musicSetDetile(zc.accessToken, id)
}

func musicSetDetile(token string, id int) (ZMusicSetDetile, error) {
	var zMusicSetDetile ZMusicSetDetile
	params := url.Values{}
	params.Add("access_token", token)
	params.Add("id", strconv.Itoa(id))

	uri := musicSetDetileURL + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZMusicSetDetile{}, err
	}
	if err := zMusicSetDetile.parse(data); err != nil {
		return ZMusicSetDetile{}, err
	}
	return zMusicSetDetile, nil
}

//https://api.zaycev.net/external/genre?genre=%s&page=%s&access_token=%s
func (zc *ZClient) Genre(genreName string, page int) (ZGenre, error) {
	return genre(zc.accessToken, genreName, page)
}

func genre(token, genre string, page int) (ZGenre, error) {
	var zGenre ZGenre
	params := url.Values{}
	params.Add("access_token", token)
	params.Add("page", strconv.Itoa(page))
	params.Add("genre", genre)

	uri := genreURL + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZGenre{}, err
	}
	if err := zGenre.parse(data); err != nil {
		return ZGenre{}, err
	}
	return zGenre, nil
}

//https://api.zaycev.net/external/artist/%d?access_token=%s
func (zc *ZClient) Artist(id int) (ZArtist, error) {
	return artist(zc.accessToken, id)
}

func artist(token string, id int) (ZArtist, error) {
	var zArtist ZArtist
	params := url.Values{}
	params.Add("access_token", token)

	u := fmt.Sprintf(artistURL, id)

	uri := u + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZArtist{}, err
	}
	if err := zArtist.parse(data); err != nil {
		return ZArtist{}, err
	}
	return zArtist, nil
}

//https://api.zaycev.net/external/track/%d?access_token=%s
func (zc *ZClient) Track(id int) (ZTrack, error) {
	return track(zc.accessToken, id)
}

func track(token string, id int) (ZTrack, error) {
	var zTrack ZTrack
	params := url.Values{}
	params.Add("access_token", token)

	u := fmt.Sprintf(trackURL, id)

	uri := u + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZTrack{}, err
	}
	if err := zTrack.parse(data); err != nil {
		return ZTrack{}, err
	}
	return zTrack, nil
}

//https://api.zaycev.net/external/options?access_token=%s
func (zc *ZClient) Options() (ZOptions, error) {
	return options(zc.accessToken)
}

func options(token string) (ZOptions, error) {
	var zOptions ZOptions
	params := url.Values{}
	params.Add("access_token", token)

	uri := optionsURL + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZOptions{}, err
	}
	if err := zOptions.parse(data); err != nil {
		return ZOptions{}, err
	}
	return zOptions, nil
}

//https://api.zaycev.net/external/track/%d/download/?access_token=%s&encoded_identifier=%s"
func (zc *ZClient) Download(id int) (ZDownload, error) {
	return download(zc.accessToken, id)
}

func download(token string, id int) (ZDownload, error) {
	var zDownload ZDownload
	params := url.Values{}
	params.Add("access_token", token)
	params.Add("encoded_identifier", "")

	u := fmt.Sprintf(downloadURL, id)

	uri := u + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZDownload{}, err
	}
	if err := zDownload.parse(data); err != nil {
		return ZDownload{}, err
	}
	return zDownload, nil
}

//https://api.zaycev.net/external/track/%s/play?access_token=%s&encoded_identifier=%s
func (zc *ZClient) Play(id int) (ZPlay, error) {
	return play(zc.accessToken, id)
}

func play(token string, id int) (ZPlay, error) {
	var zPlay ZPlay
	params := url.Values{}
	params.Add("access_token", token)
	params.Add("encoded_identifier", "")

	u := fmt.Sprintf(playURL, id)

	uri := u + params.Encode()

	data, err := get(uri)
	if err != nil {
		return ZPlay{}, err
	}
	if err := zPlay.parse(data); err != nil {
		return ZPlay{}, err
	}
	return zPlay, nil
}

//get data
func get(uri string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return errData, err
	}

	res, err := client.Do(req)
	if err != nil {
		return errData, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errData, fmt.Errorf("not found")
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errData, err
	}
	se := checkServiceError(bodyBytes)
	if se != nil {
		return errData, se
	}
	return bodyBytes, nil
}

func checkServiceError(data []byte) error {
	var zError ZError
	err := zError.parse(data)
	if err != nil || zError.Error.Text != "" {
		return fmt.Errorf("Message %s, code %d", zError.Error.Text, zError.Error.Code)
	}
	return nil
}
