package main

import (
	"fmt"
	"time"

	"github.com/Lunarisnia/image-fetcher/internal"
)

func main() {
	start := time.Now()
	workReport := make(chan string)
	for workerCount := 0; workerCount < 4; workerCount++ {
		go func() {
			for i := 0; i < 25; i++ {
				generatedUrl := internal.GenerateImageURL()
				internal.DownloadFile(generatedUrl.URL)
			}
			workReport <- "finished"
		}()
	}
	w1, w2, w3, w4 := <-workReport, <-workReport, <-workReport, <-workReport
	if w1 != "" && w2 != "" && w3 != "" && w4 != "" {
		elapsed := time.Since(start)
		fmt.Println("Concurrent: ", elapsed)
	}
}
