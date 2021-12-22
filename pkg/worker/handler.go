package worker

import (
	"fmt"
	"net/http"
	"strconv"
)

var WorkQueue = make(chan WorkRequest, 100)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Not Allowed")
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rTypeString := r.FormValue("worktype")
	if rTypeString == "" {
		http.Error(w, "Work Type Undefined", http.StatusBadRequest)
		return
	}
	rType, err := strconv.Atoi(rTypeString)
	if err != nil {
		http.Error(w, "Please Input Number in Work Type", http.StatusBadRequest)
		return
	}
	work := WorkRequest{rType, nil, nil, nil}

	WorkQueue <- work

}
