package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/charlottepincher/cpincher-task/calculate_pack"
	"github.com/gorilla/mux"
)

// Set the pack sizes from the example
// TODO: Could make this so it's read in from a config json instead?
var pack_sizes = []int{250, 500, 1000, 2000, 5000}

type OrderAmount struct {
	Items int `json:"order"`
}

func IncomingOrder(w http.ResponseWriter, r *http.Request) {
	// Get order amount from json input to API
	var order_amount OrderAmount
	_ = json.NewDecoder(r.Body).Decode(&order_amount)
	ordered := order_amount.Items

	// Ensure pack sizes are correctly ordered before passing to function
	sort.Ints(pack_sizes)
	packs := calculate_pack.CalculatePacks(ordered, pack_sizes)
	fmt.Fprint(w, "\n", packs)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", IncomingOrder)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, router)
}
