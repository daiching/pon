# pon
pon is simple Go Rest API library

## Sample
```go
func main() {
	// make new API
	api := pon.NewApi("/sample")

	// set handler for GET
	api.SetGet(func(w http.ResponseWriter, r *http.Request) (interface{}, int) {
		var s struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
		s.Id = 1
		s.Name = "GET"

		// return status code and interface{}
		return s, pon.StatusOK
	})
	
	//mapping
	api.Map()
	// start api
	pon.Start("8080")
}
```

You can find practical usage of this library if read /sample.

## Basic usage 
You can start this library by next:
1. Declare pon.Api by pon.NewApi().
2. Set handler for each HTTP method by pon.Api.SetGet(), SetPost(), and so on.
3. Determine mapping of api handler by pon.Api.Map().
4. Start api by pon.Start().

## Very simple
As mentioned above "Basic usage", to finish, this library hope you only set http hadler for GET / POST / PUT / DELETE. So this library prepare "func SetMethod(func(){}) (interface{}, int)" and "SetJsonMethod(func(){}) (string, int)". 

Specifically:
- func SetMethod(func(){}) (interface{}, int) : SetGet / SetPost / SetPut / SetDelete 
- func SetJsonMethod(func(){}) (string, int) : SetJsonGet / SetJsonPost / SetJsonPut / SetJsonDelete 

(interface{}, int) is for example "any struct" and "status code". 
(string, int) is for example "json" and "status code".  
