package api

import (
	"encoding/json"
	"fmt"
	"go-fib/fib"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	server   *http.Server
	httpPort int
}

type healthApiImpl struct {
	resource string
}

// fibApiImpl The implementation for the operations and resources of the Fib REST API
// especified at ../fib_openapi3.yaml
type fibApiImpl struct{}

type Response struct {
	FibonacciNthPosition uint16 `json:"fibonacciNthPosition"`
	Value                string `json:"value"`
}

// fibResourcesHandler is a handler to address 'the challenge' requirements
func (api fibApiImpl) fibResourcesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Printf("\nEndpoint Hit: GET /fib; Params (%v)\n", r.URL.Query())

		parameters := r.URL.Query()
		if (parameters["n"] == nil) || (len(parameters["n"]) < 1) {
			msg := "ERROR: required query parameted 'N' was not provided"
			log.Println(msg)

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(msg))
			return
		}

		parameterN := parameters["n"][0]

		if positionN, err := strconv.Atoi(parameterN); err == nil {
			position := uint16(positionN)

			if (position < uint16(0)) || (position > 1474) {
				msg := fmt.Sprintf("ERROR: invalid parameter range. Should be between 0 - 1474, but was: %v", positionN)
				log.Println(msg)

				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(msg))
				return
			}

			value, _ := fib.RetrieveNthFibonacci(position)

			w.Write([]byte(value))
			log.Println()

			return
		}

		msg := fmt.Sprintf("ERROR: invalid input value %v", parameterN)
		log.Println(msg)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))

	case http.MethodPost:
		methodNotAllowed(w)
	case http.MethodPut:
		methodNotAllowed(w)
	case http.MethodDelete:
		methodNotAllowed(w)
	default:
		methodNotAllowed(w)
	}
}

// fibonacciResourcesHandler is a handle to implements my custom format
func (api fibApiImpl) fibonacciResourcesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Printf("\nEndpoint Hit: GET /fibonacci; Params (%v)\n", r.URL.Query())

		parameters := r.URL.Query()
		if (parameters["n"] == nil) || (len(parameters["n"]) < 1) {
			msg := "ERROR: required query parameted 'N' was not provided"
			log.Println(msg)

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(msg))
			return
		}

		parameterN := parameters["n"][0]

		if positionN, err := strconv.Atoi(parameterN); err == nil {
			position := uint16(positionN)

			value, err := fib.RetrieveNthFibonacci(position)
			if err != nil {
				msg := fmt.Sprintf("ERROR: %v", err)
				log.Println(msg)

				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(msg))
				return
			}

			jData, err := json.Marshal(Response{FibonacciNthPosition: position, Value: value})
			if err != nil {
				log.Printf("ERROR: %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)

			log.Println()

			return
		}

		msg := fmt.Sprintf("ERROR: invalid input value %v", parameterN)
		log.Println(msg)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))

	case http.MethodPost:
		methodNotAllowed(w)
	case http.MethodPut:
		methodNotAllowed(w)
	case http.MethodDelete:
		methodNotAllowed(w)
	default:
		methodNotAllowed(w)
	}
}

func NewFibApiImpl() *fibApiImpl {
	return &fibApiImpl{}
}

// healthResourcesHandler is a handler created to address health check ends
func (api healthApiImpl) healthResourcesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nEndpoint Hit: GET /health_check; Params (N): %v\n", r.URL.Query())

	switch r.Method {
	case http.MethodGet:

		fmt.Fprintf(w, "Health check")

		log.Println()
	default:
		methodNotAllowed(w)
	}
}

func NewHealthApiImpl() *healthApiImpl {
	return &healthApiImpl{resource: "/health_check"}
}

func NewApiServer(httpPort int) *ApiServer {

	mux := http.NewServeMux()

	healthApiImpl := NewHealthApiImpl()
	mux.HandleFunc(healthApiImpl.resource, healthApiImpl.healthResourcesHandler)

	fibApiImpl := NewFibApiImpl()
	mux.HandleFunc("/fib", fibApiImpl.fibResourcesHandler)
	mux.HandleFunc("/fibonacci", fibApiImpl.fibonacciResourcesHandler)

	s := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", httpPort),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &ApiServer{server: s, httpPort: httpPort}
}

func (server *ApiServer) StartApiServer() error {

	log.Printf("\nApiServer started at http port %d\n", server.httpPort)

	err := server.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("unable to start the server.\ncause: %v", err)
	}

	return nil
}

func methodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)

	w.Write([]byte("Method not allowed"))
}
