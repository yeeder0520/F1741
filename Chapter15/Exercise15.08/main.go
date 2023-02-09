package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type WorldTime struct {
	UTC      string `"json:utc"`
	Local    string `"json:local"`
	Timezone string `"json:timezone"`
}

func RestfulService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	now := time.Now()

	vl := r.URL.Query()
	tz, ok := vl["tz"]
	if ok {
		location, err := time.LoadLocation(strings.TrimSpace(tz[0]))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonBytes, _ := json.Marshal(WorldTime{Timezone: "Invalid timezone name"})
			w.Write(jsonBytes)
			return
		}
		now = now.In(location)
	}

	worldTime := WorldTime{}
	worldTime.UTC = now.UTC().Format(time.RFC3339)
	worldTime.Local = now.Format(time.RFC3339)
	worldTime.Timezone = now.Location().String()

	jsonBytes, _ := json.Marshal(worldTime)
	w.Write(jsonBytes)
	// json.NewEncoder(w).Encode(worldTime)
}

func main() {
	http.HandleFunc("/", RestfulService)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
