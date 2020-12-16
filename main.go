package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func youtubeDownloadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	youtubeURL := r.FormValue("youtubeUrl")
	fmt.Fprintf(w, "Downloading %s\n", youtubeURL)
	out, err := exec.Command("youtube-dl " + youtubeURL).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Log %s\n", out)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/downloadyoutube", youtubeDownloadHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
