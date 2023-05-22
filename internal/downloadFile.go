package internal

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func getFileName(URL string) string {
	split := strings.Split(URL, "/")

	return split[len(split)-1] + ".png"
}

func DownloadFile(URL string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := getFileName(URL)
	file, err := os.Create(workingDir + "/images/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
