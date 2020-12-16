package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func youtubeDownloadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	youtubeURL := r.FormValue("youtubeURL")

	//Parsing URL
	_, err := url.ParseRequestURI(youtubeURL)

	if err != nil {
		fmt.Fprintf(w, "Invalid URL %s\n", youtubeURL)
	} else {
		fmt.Fprintf(w, "Downloading %s\n", youtubeURL)

		goExecutable, _ := exec.LookPath("youtube-dl")

		cmdGoVer := &exec.Cmd{
			Path:   goExecutable,
			Args:   []string{goExecutable, youtubeURL},
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}

		if err := cmdGoVer.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	}
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
