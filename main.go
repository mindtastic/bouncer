package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	uuid "github.com/hashicorp/go-uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var downstream = flag.String("downstream", "", "downstream Ory Kratos cluster")
var address = flag.String("addr", ":7666", "address to listen on")
var logging = flag.Bool("log", false, "enable logging of requests")

func main() {
	flag.Parse()

	if *downstream == "" {
		log.Println("downstream must not be empty")
		flag.PrintDefaults()
		return
	}

	downstreamURL, err := url.Parse(*downstream)
	if err != nil {
		log.Printf("could not parse downstream URL: %v", err)
		flag.PrintDefaults()
		return
	}

	mux := http.ServeMux{}
	mux.Handle("/health", handleHealthCheck())
	mux.Handle("/", handleEverythingElse(*downstreamURL))
	mux.Handle("/self-service/registration", handleRegistration(*downstreamURL))
	mux.Handle("/self-service/registration/", handleRegistration(*downstreamURL))

	proxy := &http.Server{
		Addr:           *address,
		Handler:        &mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("listening on port", *address)
	log.Fatal(proxy.ListenAndServe())
}

func handleHealthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}

func handleRegistration(url url.URL) *httputil.ReverseProxy {
	type requestBody struct {
		Method   string `json:"method"`
		Password string `json:"password"`
		Traits   struct {
			AccountKey string `json:"accountKey"`
		} `json:"traits"`
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme
			stamp, err := uuid.GenerateUUID()
			if err != nil {
				log.Printf("error generating UUID for request: %v", err)
				return
			}
			d := json.NewDecoder(req.Body)
			body := new(requestBody)
			if err := d.Decode(body); err != nil {
				log.Printf("error decoding request body: %v", err)
				return
			}
			body.Traits.AccountKey = stamp
			body.Password = fmt.Sprintf("%x", md5.Sum([]byte(stamp)))
			b := new(bytes.Buffer)
			e := json.NewEncoder(b)
			if err := e.Encode(body); err != nil {
				log.Printf("error encoding request body: %v", err)
				return
			}
			req.Body = ioutil.NopCloser(b)
			req.ContentLength = int64(b.Len())
			if *logging {
				log.Printf("proxied registration call to %q", req.URL.Path)
			}
		},
	}
	return proxy
}

func handleEverythingElse(url url.URL) http.Handler {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme
			if *logging {
				log.Printf("proxied other call to %q", req.URL.Path)
			}
		},
	}
	return proxy
}
