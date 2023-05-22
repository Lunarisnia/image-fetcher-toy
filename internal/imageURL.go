package internal

import (
	"fmt"
	"math/rand"
)

type ImageName struct {
	URL      string
	FileName string
}

func GenerateImageURL() ImageName {
	host := "http://fpoimg.com/"

	fileName := fmt.Sprintf("%vx%v", rand.Intn(1000), rand.Intn(1000))
	return ImageName{URL: host + fileName, FileName: fileName}
}
