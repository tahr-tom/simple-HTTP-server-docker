package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"sort"
	"time"
)

const ISO8601 = "2006-01-02T15:04:05Z"

// global map storage
var messagesMap = make(map[string]*Value)

// for slice that will be converted to json response
type jsonRes struct {
	Timestamp string `json:"timestamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type jsonReq struct {
	Key   string
	Value string
}

// for map storage
type Entry struct {
	key   string
	value *Value
}

type Value struct {
	Timestamp time.Time
	Value     string
}

// for sorting by timestamp
type byTimestamp []Entry

func (d byTimestamp) Len() int {
	return len(d)
}

func (d byTimestamp) Less(i, j int) bool {
	return d[i].value.Timestamp.After(d[j].value.Timestamp)
}

func (d byTimestamp) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func mapToJSON() []byte {
	// convert the map to slice then sort by timestamp in descending order
	slice := make(byTimestamp, 0, len(messagesMap))
	for key, value := range messagesMap {
		slice = append(slice, Entry{key, value})
	}
	sort.Sort(slice)

	// flatten the Entry struct
	messageSlice := make([]jsonRes, 0)
	for _, entry := range slice {
		messageSlice = append(messageSlice, jsonRes{
			Timestamp: entry.value.Timestamp.Format(ISO8601),
			Key:       entry.key,
			Value:     entry.value.Value,
		})
	}

	response, err := json.Marshal(messageSlice)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func addMessage(req jsonReq) {
	messagesMap[req.Key] = &Value{
		Timestamp: time.Now(),
		Value:     req.Value,
	}
}

// using hardcoded timestamp for debugging (adding init entries)
func addMessageDebug(key string, value string, timestamp time.Time) {
	messagesMap[key] = &Value{
		Timestamp: timestamp,
		Value:     value,
	}
}

// helper function for addMessageDebug
func parseTime(timeString string) time.Time {
	t, err := time.Parse(ISO8601, timeString)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func main() {
	addMessageDebug("a", "same Value", parseTime("2019-12-02T06:53:32Z"))
	addMessageDebug("asdf", "some other Value", parseTime("2019-12-02T06:53:35Z"))

	r := chi.NewRouter()
	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mapToJSON())
	})

	r.Post("/add", func(writer http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		var req jsonReq
		err := decoder.Decode(&req)
		if err != nil {
			log.Fatal(err)
		}
		addMessage(req)
	})

	http.ListenAndServe(":80", r)
}
