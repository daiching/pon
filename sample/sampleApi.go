package main

import (
	"bytes"
	"encoding/json"
	"github.com/daiching/pon"
	"net/htp"
)

type Sample struct {
	Id   int
	Name string
}

func init() {
	api := pon.NewApi("/sample/")
	setGetSampleAPI(api)
	setPostSampleAPI(api)
	api.Map()
}

func setGetSampleAPI(api *pon.Api) error {
	api.SetGet(func(w http.ResponseWriter, r *http.Request) (interface{}, int) {
		s := Sample{1, "GET"}
		return s, pon.StatusOK
	})
	return nil
}

func setPostSampleAPI(api *pon.Api) error {
	api.SetPost(func(w http.ResponseWriter, r *http.Request) (interface{}, int) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		s := Sample{}
		json.Unmarshal(buf.Bytes(), &s)
		return s, pon.StatusOK
	})
	return nil
}
