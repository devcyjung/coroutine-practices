package main

import (
	"examples/ch3/svg"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	bp1 := svg.BaseBlueprint()
	bp2 := svg.BaseBlueprint()
	bp1.ProjectionAngle = math.Pi

	bpArray := []svg.Blueprint{bp1, bp2}

	logfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	http.HandleFunc("/plot", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		idStr := params.Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || !(0 <= id && id <= len(bpArray)) {
			http.Error(w, "plot not found", http.StatusBadRequest)
			logger.Printf("Bad request /plot?id=%v", idStr)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		err = bpArray[id].PlotSVG(w)
		if err != nil {
			logger.Printf("%+v", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
