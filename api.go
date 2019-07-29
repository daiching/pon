package pon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// This structure needs to be prepared according to API path.
type Api struct {
	path        string
	get         func(w http.ResponseWriter, r *http.Request) (string, int)
	post        func(w http.ResponseWriter, r *http.Request) (string, int)
	put         func(w http.ResponseWriter, r *http.Request) (string, int)
	del         func(w http.ResponseWriter, r *http.Request) (string, int)
	isSetGet    bool
	isSetPost   bool
	isSetPut    bool
	isSetDelete bool
}

func NewApi(path string) *Api {
	a := Api{}
	a.path = path
	return &a
}

// Path And process in API are tied.
func (a *Api) Map() {
	mapping(a.path, func(w http.ResponseWriter, r *http.Request) (string, int) {
		fmt.Println("** Mapped function is run")
		if r.Method == "GET" && a.isSetGet {
			fmt.Println("GET function is run")
			return a.get(w, r)
		}
		if r.Method == "POST" && a.isSetPost {
			fmt.Println("POST function is run")
			return a.post(w, r)
		}
		if r.Method == "PUT" && a.isSetPut {
			fmt.Println("PUT function is run")
			return a.put(w, r)
		}
		if r.Method == "DELETE" && a.isSetDelete {
			fmt.Println("DELETE function is run")
			return a.del(w, r)
		}
		setAllowHeader(a, w)

		fmt.Println("Default Error Response")
		return getDefaultErrorJson(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed
	})
}

// SetGet, SetPost.. register behavior corresponding to method to API.
// Also, they convert interface which is the return value from the argument's function to json
func (a *Api) SetGet(f func(w http.ResponseWriter, r *http.Request) (interface{}, int)) {
	a.get = setMethod(f, "GET")
	a.isSetGet = true
}

func (a *Api) SetPost(f func(w http.ResponseWriter, r *http.Request) (interface{}, int)) {
	a.post = setMethod(f, "POST")
	a.isSetPost = true
}

func (a *Api) SetPut(f func(w http.ResponseWriter, r *http.Request) (interface{}, int)) {
	a.put = setMethod(f, "PUT")
	a.isSetPut = true
}

func (a *Api) SetDelete(f func(w http.ResponseWriter, r *http.Request) (interface{}, int)) {
	a.del = setMethod(f, "DELETE")
	a.isSetDelete = true
}

// SetJsonGet, SetJsonPost.. register behavior corresponding to method to API.
// Because the return value from the argument's function is already json, convert is not required.
func (a *Api) SetJsonGet(f func(w http.ResponseWriter, r *http.Request) (string, int)) {
	a.get = setJsonMethod(f, "GET")
	a.isSetGet = true
}

func (a *Api) SetJsonPost(f func(w http.ResponseWriter, r *http.Request) (string, int)) {
	a.post = setJsonMethod(f, "POST")
	a.isSetPost = true
}

func (a *Api) SetJsonPut(f func(w http.ResponseWriter, r *http.Request) (string, int)) {
	a.put = setJsonMethod(f, "PUT")
	a.isSetPut = true
}

func (a *Api) SetJsonDelete(f func(w http.ResponseWriter, r *http.Request) (string, int)) {
	a.del = setJsonMethod(f, "DELETE")
	a.isSetDelete = true
}

// 1. Api struct methods "SetGet, SetPost, .." is filtered by this method.
// 2. interface is translated to json(string).
func setMethod(f func(w http.ResponseWriter, r *http.Request) (interface{}, int), m string) func(w http.ResponseWriter, r *http.Request) (string, int) {
	return func(w http.ResponseWriter, r *http.Request) (string, int) {
		inf, sts := f(w, r)
		fmt.Println("OriginData:", inf)
		fmt.Println("OriginStatus:", sts)
		if b, e := json.Marshal(inf); e == nil {
			sts = getDefaultStatusWithMethod(m, sts)
			json := string(b)
			fmt.Println("Status:", sts)
			fmt.Println("Json:", json)
			return permeatedDefaultFilter(json, sts)
		}
		fmt.Println("CantMarshal to json")
		return getDefaultErrorJson(http.StatusInternalServerError), http.StatusInternalServerError
	}
}

// Api struct methods "SetJsonGet, SetJsonPost, .." is filtered by this method.
func setJsonMethod(f func(w http.ResponseWriter, r *http.Request) (string, int), m string) func(w http.ResponseWriter, r *http.Request) (string, int) {
	return func(w http.ResponseWriter, r *http.Request) (string, int) {
		js, sts := f(w, r)
		sts = getDefaultStatusWithMethod(m, sts)
		return permeatedDefaultFilter(js, sts)
	}
}

// Http Response Header "Allow" is set according to method set on API struct
func setAllowHeader(a *Api, w http.ResponseWriter) {
	var allowMedthods []string
	if a.isSetGet {
		allowMedthods = append(allowMedthods, "GET")
	}
	if a.isSetPost {
		allowMedthods = append(allowMedthods, "POST")
	}
	if a.isSetPut {
		allowMedthods = append(allowMedthods, "PUT")
	}
	if a.isSetDelete {
		allowMedthods = append(allowMedthods, "DELETE")
	}
	var allowValue string
	for i, m := range allowMedthods {
		allowValue += m
		if (i + 1) < len(allowMedthods) {
			allowValue += ", "
		}
	}
	w.Header().Set("Allow", allowValue)
}

// make and return handler for http.HandleFunc.
func getHandler(f func(w http.ResponseWriter, r *http.Request) (string, int)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		msg, sts := f(w, r)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if sts == StatusConflict {
			w.Header().Set("Location", r.URL.String())
		}
		w.WriteHeader(sts)
		fmt.Fprintf(w, msg)
		return
	}
}

// mapping http handler with path of url.
func mapping(path string, f func(w http.ResponseWriter, r *http.Request) (string, int)) {
	handler := getHandler(f)
	http.HandleFunc(path, handler)
}

// start api server
func Start(portNum string) error {
	err := http.ListenAndServe(":"+portNum, nil)
	return err
}
