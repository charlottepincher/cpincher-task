package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/charlottepincher/cpincher-task/calculate_pack"
	"github.com/charlottepincher/cpincher-task/html"
	"github.com/gorilla/mux"
)

type Config struct {
	PackSizes []int `json:"packs"`
}

// Used if the API is directly queried
type OrderAmount struct {
	Items int `json:"order"`
}

var pack_sizes []int

// gcd = Greatest Common Divisor
var gcd int

func GetConfig(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func IncomingOrder(w http.ResponseWriter, r *http.Request) {
	// Get order amount from json input to API
	var order_amount OrderAmount
	_ = json.NewDecoder(r.Body).Decode(&order_amount)
	ordered := order_amount.Items

	// Ensure pack sizes are correctly ordered before passing to function
	sort.Ints(pack_sizes)
	packs := calculate_pack.CalculatePacks(ordered, pack_sizes, gcd)
	for _, value := range pack_sizes {
		fmt.Fprintf(w, "%d: %d\n", value, packs[value])
	}
}

func init() {
	// Get pack sizes from the json config file
	pack_sizes = GetConfig("config.json").PackSizes
	gcd = calculate_pack.FindGCD(pack_sizes)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/apitest", IncomingOrder)
	router.HandleFunc("/", Webpage)
	http.Handle("/image/", http.StripPrefix("/image", http.FileServer(http.Dir("./image"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, router)
}

func Webpage(w http.ResponseWriter, r *http.Request) {
	p := html.DashboardParams{}
	if r.Method == http.MethodPost {
		orderAmount := r.FormValue("order")
		// Ensure pack sizes are correctly ordered before passing to function
		sort.Ints(pack_sizes)
		ordered, _ := strconv.Atoi(orderAmount)
		packs := calculate_pack.CalculatePacks(ordered, pack_sizes, gcd)
		for _, value := range pack_sizes {
			p.Packs = append(p.Packs, fmt.Sprintf("%d: %d\n", value, packs[value]))
		}
		p.Ordered = ordered
	}
	html.Dashboard(w, p)
}
