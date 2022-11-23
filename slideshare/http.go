package slideshare

import (
	"io"
	"net/http"
)

func GetSlideShareLink(uri string) string {
	resp, err := http.Get(uri)
	if err != nil {
		panic("request error")
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("read error")
	}
	return string(bytes)
}
