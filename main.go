package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mgw2007/golang-metric/inmemory"
	"github.com/mgw2007/golang-metric/metric"
)

var inMemory metric.HandleMetrics

func init() {
	inMemory = inmemory.NewMetric()
}
func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Usage")
	flag.Parse()
	fmt.Println("App run on url http://localhost:" + port)

	router := httprouter.New()
	router.POST("/metric/:key", HandlePostMetric)
	router.GET("/metric/:key/sum", HandleGetMetric)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

// HandlePostMetric for handle add metric
func HandlePostMetric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inMemory.AddMetric(ps.ByName("key"), time.Now())
	js, _ := json.Marshal(metric.PostMetricResponse{})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// HandleGetMetric for hadle get metric count
func HandleGetMetric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	count, err := inMemory.GetMetricCount(ps.ByName("key"), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	js, _ := json.Marshal(metric.GetMetricCountResponse{Value: count})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
