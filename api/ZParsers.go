package api

import "encoding/json"

type parser interface {
	parse(data []byte) error
}

func (t *ZToken) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZTop) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZSearch) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZTerms) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZArtist) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZDownload) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZPlay) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZError) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZGenre) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZTrack) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZMusicSetDetile) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZMusicSetList) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZOptions) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}

func (t *ZSettings) parse(data []byte) error {
	return json.Unmarshal(data, &t)
}
