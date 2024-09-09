package diag

import (
	utility "go-http-server/utility/handlers"
	"log"
	"net/http"
)

type Diag struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func LogHandler(log *log.Logger) http.Handler {
	return utility.Pipe(func(d LogHandlerRequest) {
		log.Printf("%+v\n", d)
	})
}

type LogHandlerRequest struct {
	D Diag
}
