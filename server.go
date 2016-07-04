package fizzbuzz_server

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	number := r.FormValue("number")
	i, err := strconv.Atoi(number)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"error":"number parameter is not numerical"}`)
		return
	}
	fizzbuzz := NewFizzBuzz()
	result := fizzbuzz.Do(i)
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"message\":\"%s\"}", result)
}

func Run() {
	http.HandleFunc("/", handleFizzBuzz)
	http.ListenAndServe(":5000", nil)
}
