package main

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"

	zbar "github.com/PeterCxy/gozbar"
)

func main() {
	scanner := zbar.NewScanner()
	scanner.SetConfig(0, zbar.CFG_ENABLE, 1)
	defer scanner.Destroy()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		src, _, err := image.Decode(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		img := zbar.FromImage(src)
		scanner.Scan(img)
		res := []string{}
		if img.First() != nil {
			img.First().Each(func(text string) {
				res = append(res, text)
			})
		}
		r.Header.Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
