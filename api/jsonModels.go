package main

import "encoding/json"

//MovieAlias alias of movie model
type MovieAlias Movie

//MovieJSON model for json
type MovieJSON struct {
	MovieAlias
	Actors []string `json:"actors"`
}

//MarshalJSON implementing marshaler interface
func (m Movie) MarshalJSON() ([]byte, error) {
	var actName []string
	for _, actor := range m.Actors {
		actName = append(actName, actor.Name)
	}
	mj := MovieJSON{
		MovieAlias: MovieAlias(m),
		Actors:     actName,
	}
	return json.Marshal(mj)
}
