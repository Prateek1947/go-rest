package main

import "encoding/json"

//MovieAlias alias of movie model
type MovieAlias Movie

//MovieJSON model for json
type MovieJSON struct {
	MovieAlias
	Actors    []string `json:"actors,omitempty"`
	PosterURI string   `json:"poster,omitempty"`
}

//MarshalJSON implementing marshaler interface
func (m Movie) MarshalJSON() ([]byte, error) {
	var actName []string
	var posterURI string
	for _, actor := range m.Actors {
		actName = append(actName, actor.Name)
	}
	if m.PosterURI != "" {
		posterURI = "media/posters/" + m.PosterURI
	}
	mj := MovieJSON{
		MovieAlias: MovieAlias(m),
		Actors:     actName,
		PosterURI:  posterURI,
	}
	return json.Marshal(mj)
}
