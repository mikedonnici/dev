package string

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

type Service struct {
}

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
