package main
import (
	"db"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)
func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage")
}
func varticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, db.Displayarticles())
}
func varticlesid(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	fmt.Fprintln(w, db.Displayarticlebyid(id))
}
func carticles(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	title := mux.Vars(r)["Title"]
	description := mux.Vars(r)["Description"]
	price := mux.Vars(r)["Price"]
	key, _ := strconv.Atoi(id)
	key1, _ := strconv.Atoi(price)
	fmt.Fprintln(w, db.CreateArticles(key, title, description, key1))
}
func createarticles(w http.ResponseWriter, r *http.Request) {
	id, _ := r.URL.Query()["ID"]
	title, _ := r.URL.Query()["Title"]
	description, _ := r.URL.Query()["Description"]
	price, _ := r.URL.Query()["Price"]
	key := id[0]
	key1 := price[0]
	id1, _ := strconv.Atoi(key)
	id4, _ := strconv.Atoi(key1)
	id2 := strings.Join(title, " ")
	id3 := strings.Join(description, " ")
	fmt.Fprintln(w, db.CreateArticles(id1, id2, id3, id4))
}
func delarticles(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["Title"]
	fmt.Fprintln(w, db.Deletearticles(title))
}
func updarticles(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	description := mux.Vars(r)["Description"]
	key, _ := strconv.Atoi(id)
	fmt.Fprintln(w, db.Updatearticles(key, description))
}
func handleFunction() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeFunc)
	myRouter.HandleFunc("/varticles", varticles).Methods("GET")
	myRouter.HandleFunc("/varticles/{ID}/{Title}/{Description}/{Price}", carticles).Methods("POST")
	myRouter.HandleFunc("/varticles/{ID}", varticlesid).Methods("GET")
	myRouter.HandleFunc("/darticles/{Title}", delarticles).Methods("DELETE")
	myRouter.HandleFunc("/uarticles/{ID}/{Description}", updarticles).Methods("PUT")
	myRouter.HandleFunc("/carticles", createarticles).Methods("POST")
	myRouter.PathPrefix("/calculator").Handler(http.FileServer(rice.MustFindBox("../calculator").HTTPBox()))
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}
func main() {
	handleFunction()
}