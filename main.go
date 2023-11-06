package main

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler("https://freshman.tech", http.StatusSeeOther))

	http.ListenAndServe(":6000", mux)
}
