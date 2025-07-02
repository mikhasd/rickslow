package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"io"
	"log/slog"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

//go:embed rickroll.txt.gz
var rickrollGz []byte

const frameLen = 2025

func main() {
	runtime.GOMAXPROCS(1)
	logger := slog.Default()
	logger.Info("starting at port :10888")
	rickroll := decompressResponse()

	server := http.Server{
		Addr: ":10888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("rickrolling", slog.String("path", r.URL.Path))
			flusher := w.(http.Flusher)
			start := time.Now()
			w.Header().Add("Content-Type", "text/html")
			w.Header().Add("Content-Length", strconv.Itoa(len(rickroll)))
			w.Header().Add("Server", "Apache")
			w.Header().Add("X-Powered-By", "PHP/8.3")
			w.WriteHeader(200)
			flusher.Flush()

			var err error = nil
			rick := rickroll
			for err == nil || len(rick) < frameLen {
				w.Write([]byte("\033[H\033[J"))
				_, err = w.Write(rick[0:frameLen])
				flusher.Flush()
				rick = rick[frameLen:]
				time.Sleep(18 * time.Millisecond)
			}

			dur := time.Since(start)
			logger.Info("request successfully rickrolled", slog.Duration("duration", dur))
		}),
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Error("error on http server", slog.String("error", err.Error()))
	}
	logger.Info("terminating")
}

func decompressResponse() []byte {
	logger := slog.Default()
	logger.Info("decompressing rickroll.txt.gz", slog.Int("length", len(rickrollGz)))
	reader := bytes.NewReader(rickrollGz)
	greader, err := gzip.NewReader(reader)
	if err != nil {
		panic(err)
	}

	defer greader.Close()
	rickroll, err := io.ReadAll(greader)
	if err != nil {
		panic(err)
	}
	logger.Info("rickroll.txt decompressed", slog.Int("length", len(rickroll)))
	return rickroll
}
