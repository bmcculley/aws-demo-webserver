package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	hostname string
	html     bool
)

func setHostName() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func hostNameServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	if html {
		fmt.Fprintf(w, "<h1>Hello from %s!</h1>", hostname)
	} else {
		w.Header().Set("Content-Type", "application/json")

		r.Header.Set("hostname", hostname)

		jsonHeader, err := json.Marshal(r.Header)
		if err != nil {
			log.Println(err)
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, jsonHeader, "", "\t")
		if err != nil {
			log.Println("JSON parse error: ", err)
		}

		fmt.Fprintf(w, string(prettyJSON.Bytes()))
	}
}

func parseFlag() {
	htmlArgPtr := flag.Bool("html", false, "Serve hello from <hostname>")
	flag.Parse()
	html = *htmlArgPtr
}

func main() {
	parseFlag()
	setHostName()
	http.HandleFunc("/", hostNameServer)
	http.ListenAndServe(":80", nil)
}
