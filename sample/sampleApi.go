package main

import (
	"github.com/daiching/pon"

	"net/http"
)

type Sample struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	api := pon.NewApi("/sample")
	setGetSampleAPI(api)
	setJsonPostSampleAPI(api)
	api.Map()
}

func setGetSampleAPI(api *pon.Api) error {
	api.SetGet(func(w http.ResponseWriter, r *http.Request) (interface{}, int) {
		s := Sample{1, "GET"}
		return s, pon.StatusOK
	})
	return nil
}

func setJsonPostSampleAPI(api *pon.Api) error {
	api.SetJsonPost(func(w http.ResponseWriter, r *http.Request) (string, int) {
		return `{"id":2,"name":"POST"}`, pon.StatusOK
	})
	return nil
}
