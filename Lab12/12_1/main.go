package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func pageFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["mode"]
	id := vars["id"]
	fmt.Fprintf(w, "<html><body>PG: %v<br>ID: %v</body></html>", page, id)
}

func pageRand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, _ := strconv.Atoi(vars["a"])
	b, _ := strconv.Atoi(vars["b"])
	var liczbyLosowe []int
	for i := 0; i < a; i++ {
		liczbaLosowa := rand.Intn(b + 1) //b=20 to wtedy od 0 do 19 wiec jak dodamy 1 to bedzie od 0 do 20 tak jak w poleceniu
		liczbyLosowe = append(liczbyLosowe, liczbaLosowa)
	}
	fmt.Fprintf(w, "<html><body>%v liczby losowe < 0,%v >: %v", a, b, liczbyLosowe)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/page/{mode}/{id:[0-9]+}", pageFunc)
	r.HandleFunc("/losuj/{a:[0-9]+}/{b:[0-9]+}", pageRand)
	http.ListenAndServe(":8080", r)
}
