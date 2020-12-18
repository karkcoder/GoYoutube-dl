package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	downloadURL := r.FormValue("downloadUrl")

	//Parsing URL
	_, err := url.ParseRequestURI(downloadURL)

	if err != nil {
		fmt.Fprintf(w, "Invalid URL %s\n", downloadURL)
	} else {
		fmt.Fprintf(w, "Downloading %s\n", downloadURL)

		goExecutable, _ := exec.LookPath("youtube-dl")

		cmdGoVer := &exec.Cmd{
			Path:   goExecutable,
			Args:   []string{goExecutable, downloadURL},
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
	http.HandleFunc("/download", downloadHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
