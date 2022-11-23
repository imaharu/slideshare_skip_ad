package main

import (
	"fmt"
	"slideshare_skip_ad/slideshare"
)

func main() {
	resp := slideshare.GetSlideShareLink("https://www.slideshare.net/maruyama097/twitter-49709690")
	fmt.Println(resp)
}
