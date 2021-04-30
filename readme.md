### KCD With the power of mux


Install `github.com/alexisvisco/kcd-mux`
```shell
go get github.com/alexisvisco/kcd-mux
```

Use `kcdmux.Setup()` to register path extractor for mux.

### Example


```go
package main

import (
	"fmt"
	"github.com/alexisvisco/kcd"
	"github.com/alexisvisco/kcd-mux/pkg/kcdmux"
	"github.com/alexisvisco/kcd/pkg/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	kcdmux.Setup() // Do not forget this part otherwise you will not be able to recover the path parameters

	r.HandleFunc("/{name}", kcd.Handler(YourHttpHandler, http.StatusOK))
	//                          ^ Here the magic happen this is the only thing you need
	//                            to do. Adding kcd.Handler(your handler)

	_ = http.ListenAndServe(":3000", r)
}

// CreateCustomerInput is an example of input for an http request.
type CreateCustomerInput struct {
	Name   string   `path:"name"`
	Emails []string `query:"emails" exploder:","`
}

// CustomerOutput is the output type of your handler it contain the input for simplicity.
type CustomerOutput struct {
	Name string `json:"name"`
}

// YourHttpHandler is your http handler but in a shiny version.
// You can add *http.ResponseWriter or http.Request in params if you want.
func YourHttpHandler(in *CreateCustomerInput) (CustomerOutput, error) {
	// do some stuff here

	fmt.Printf("%+v", in)

	return CustomerOutput{}, errors.NewWithKind(errors.KindInternal, "c'est fini !")
}

```
