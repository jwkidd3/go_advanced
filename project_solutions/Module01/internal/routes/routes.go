package routes

import "net/http"

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to go!!!"))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go!!!"))
}
func Handlers() {
	http.HandleFunc("/welcome", welcome)
	http.Handle("/greet", http.HandlerFunc(greet))
}
