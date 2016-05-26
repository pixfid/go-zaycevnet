/*
	Copyright (C) 2016  <Semchenko Aleksandr>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.
*/

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	apiURL string = "https://api.zaycev.net/external"
	helloURL string = apiURL + "/hello"
	authURL string = apiURL + "/auth?"
	topURL string = apiURL + "/top?"
	artistURL string = apiURL + "/artist/%d?"
	musicSetListURL string = apiURL + "/musicset/list?"
	musicSetDetileURL string = apiURL + "/musicset/detail?"
	genreURL string = apiURL + "/genre?"
	trackURL string = apiURL + "/track/%d?"
	autoCompleteURL string = apiURL + "/autocomplete?"
	searchURL string = apiURL + "/search?"
	optionsURL string = apiURL + "/options?"
	playURL string = apiURL + "/track/%d/play?"
	downloadURL string = apiURL + "/track/%d/download/?"
)

// Provides the client and associated elements for interacting with the
// Zaycev API
type ZClient struct {
	client      *http.Client //default http.Client
	helloToken  string       //set if stored before (optional)
	accessToken string       //set if stored before (optional)
	staticKey   string       //set required
}

// Generates a new client for the Zaycev API
func NewZClient(httpClient *http.Client, token, sKey string) *ZClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &ZClient{client: httpClient, accessToken: token, staticKey: sKey}
}

// ClientError is a generic error specific to the `api` package.
type ClientError struct {
	msg string
}

// Error returns a string representation of the error condition.
func (self ClientError) Error() string {
	return self.msg
}

// checkStaticKey ensures that the user configured her API key,
//   or returns an error.
func (self *ZClient) checkStaticKey() (err error) {
	if self.staticKey == "" {
		return ClientError{msg: "Empty Static Key is invalid"}
	} else {
		return nil
	}
}

// checkHelloToken ensures that the user configured her API key,
//   or returns an error.
func (self *ZClient) checkHelloToken() (err error) {
	if self.helloToken == "" {
		return ClientError{msg: "Empty Hello Token is invalid"}
	} else {
		return nil
	}
}

// checkAccessToken ensures that the user configured her API key,
//   or returns an error.
func (self *ZClient) checkAccessToken() (err error) {
	if self.accessToken == "" {
		return ClientError{msg: "Empty Access Token is invalid"}
	} else {
		return nil
	}
}

// fetchApiJson makes a request to the API and decodes the response.
// `actionUrl` is the final path component that specifies the API call
// `parameters` include the API key
// `result` is modified as an output parameter. It must be a pointer to a ZC JSON structure.
func (zc *ZClient) fetchApiJson(actionUrl string, values url.Values, result interface{}) (err error) {
	var resp *http.Response

	resp, err = zc.makeApiGetRequest(actionUrl, values)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err = dec.Decode(result); err != nil {
		return err
	}

	//TODO checkServiceError(body)

	return err
}

// makeApiGetRequest fetches a URL with querystring via HTTP GET and
//  returns the response if the status code is HTTP 200
// `parameters` should not include the apikey.
// The caller must call `resp.Body.Close()`.
func (zc *ZClient) makeApiGetRequest(fullUrl string, values url.Values) (resp *http.Response, err error) {

	req, err := http.NewRequest("GET", fullUrl + values.Encode(), nil)

	if err != nil {
		return resp, err
	}

	resp, err = zc.client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 200 {
		var msg string = fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)
		resp.Write(os.Stdout)
		return resp, ClientError{msg: msg}
	}

	return resp, nil
}

func (zc *ZClient) Auth() (err error) {
	if err = zc.checkStaticKey(); err != nil {
		return err
	}

	return zc.hello()
}

func (zc *ZClient) hello() (err error) {

	if err = zc.checkStaticKey(); err != nil {
		return err
	}

	t := &ZToken{}
	if err := zc.fetchApiJson(helloURL, url.Values{}, t); err != nil {
		return err
	}

	zc.helloToken = t.Token

	return zc.auth()
}

func (zc *ZClient) auth() (err error) {

	if err = zc.checkHelloToken(); err != nil {
		return err
	}
	r := &ZToken{}

	hash := MD5Hash(zc.helloToken + zc.staticKey)

	values := url.Values{}
	values.Add("code", zc.helloToken)
	values.Add("hash", hash)

	if err := zc.fetchApiJson(authURL, values, r); err != nil {
		return err
	}
	zc.accessToken = r.Token
	return err
}

func (zc *ZClient) Search(values url.Values) (r *ZSearch, err error) {
	r = &ZSearch{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values.Add("access_token", zc.accessToken)

	if err := zc.fetchApiJson(searchURL, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) AutoComplete(query string) (r *ZTerms, err error) {
	r = &ZTerms{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("query", query)

	if err := zc.fetchApiJson(autoCompleteURL, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Top(page int) (r *ZTop, err error) {
	r = &ZTop{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("access_token", zc.accessToken)

	if err := zc.fetchApiJson(topURL, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) MusicSetList(page int) (r *ZMusicSetList, err error) {
	r = &ZMusicSetList{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("page", strconv.Itoa(page))

	if err := zc.fetchApiJson(musicSetListURL, values, r); err != nil {
		return r, err
	}
	return r, err
}

func (zc *ZClient) MusicSetDetile(id int) (r *ZMusicSetDetile, err error) {
	r = &ZMusicSetDetile{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}
	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("id", strconv.Itoa(id))

	if err := zc.fetchApiJson(musicSetDetileURL, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Genre(genreName string, page int) (r *ZGenre, err error) {
	r = &ZGenre{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}
	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("page", strconv.Itoa(page))
	values.Add("genre", genreName)

	if err := zc.fetchApiJson(genreURL, values, r); err != nil {
		return r, err
	}

	return r, err

}

func (zc *ZClient) Artist(id int) (r *ZArtist, err error) {
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}
	r = &ZArtist{}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)

	url := fmt.Sprintf(artistURL, id)

	if err := zc.fetchApiJson(url, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Track(id int) (r *ZTrack, err error) {
	r = &ZTrack{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)

	url := fmt.Sprintf(trackURL, id)

	if err := zc.fetchApiJson(url, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Options() (r *ZOptions, err error) {
	r = &ZOptions{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)

	if err := zc.fetchApiJson(optionsURL, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Download(id int) (r *ZDownload, err error) {
	r = &ZDownload{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("encoded_identifier", "")

	url := fmt.Sprintf(downloadURL, id)
	if err := zc.fetchApiJson(url, values, r); err != nil {
		return r, err
	}

	return r, err
}

func (zc *ZClient) Play(id int) (r *ZPlay, err error) {
	r = &ZPlay{}
	if err = zc.checkAccessToken(); err != nil {
		return r, err
	}

	values := url.Values{}
	values.Add("access_token", zc.accessToken)
	values.Add("encoded_identifier", "")

	url := fmt.Sprintf(playURL, id)
	if err := zc.fetchApiJson(url, values, r); err != nil {
		return r, err
	}

	return r, err

}
