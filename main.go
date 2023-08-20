package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

type user struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Tasks []task `json:"tasks"`
}

type task struct {
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
}

func main() {
	http.HandleFunc("/run", run)
	server := http.Server{
		Addr: ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func run(w http.ResponseWriter, r *http.Request) {
	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Fatalln(err)
		return
	}
	u.Tasks = sortExpire(u.Tasks)
	prettyJSON, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	encodeJSON := encodeBase64(prettyJSON)
	fmt.Println(encodeJSON)
	decodeJSON := decodeBase64(encodeJSON)
	fmt.Println(decodeJSON)
}

func encodeBase64(jsonData []byte) string {
	return base64.StdEncoding.EncodeToString(jsonData)
}

func decodeBase64(data string) string {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatalln(err)
	}
	return string(dec)
}

func sortExpire(t []task) []task {
	dateFormat := "2006/01/02"
	sort.Slice(t, func(i, j int) bool {
		dateI, err := time.Parse(dateFormat, t[i].Deadline)
		if err != nil {
			return false
		}
		dateJ, err := time.Parse(dateFormat, t[j].Deadline)
		if err != nil {
			return true
		}
		return dateI.Before(dateJ)
	})
	return t
}
