package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mccutchen/go-httpbin/v2/httpbin"

	wasm "github.com/http-wasm/http-wasm-host-go/handler/nethttp"
)

var (
	port int
)

func getEnvInt(name string, defaultValue int) int {
	if val := os.Getenv(name); val != "" {
		intVal, _ := strconv.Atoi(val)
		return intVal
	}

	return defaultValue
}

var mws sliceFlags

func main() {
	flag.IntVar(&port, "port", getEnvInt("PORT", 8080), "Port to listen on")
	flag.Var(&mws, "middleware", "Middleware to use")

	// parse flags from command line
	flag.Parse()

	ctx := context.Background()
	var w http.Handler = httpbin.New().Handler()

	for i := range mws {
		guest, err := retrieveGuest(mws[len(mws)-i-1])
		if err != nil {
			log.Panicln(err)
		}

		h, err := wasm.NewMiddleware(ctx, guest)
		if err != nil {
			log.Panicln(err)
		}
		defer h.Close(ctx)

		w = h.NewHandler(ctx, w)
	}

	// handle route using handler function
	http.Handle("/", w)

	// listen to port
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func retrieveGuest(location string) ([]byte, error) {
	switch {
	case strings.HasPrefix(location, "http://"), strings.HasPrefix(location, "https://"):
		resp, err := http.Get(location)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return io.ReadAll(resp.Body)
	default:
		return os.ReadFile(location)
	}
}
