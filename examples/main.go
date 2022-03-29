package main

import (
	"log"

	"github.com/sikang99/gstreamer-go"
)

func main() {

	// err := gstreamer.ListPlugins()
	// if err != nil {
	// 	log.Println("list plugin error", err)
	// 	return
	// }

	// plugins := []string{"videotestsrc", "autovideosink"}

	// err = gstreamer.CheckPlugins(plugins)
	// if err != nil {
	// 	log.Println("check plugin error", err)
	// 	return
	// }

	pipeline, err := gstreamer.New("videotestsrc  ! capsfilter name=filter ! autovideosink")
	if err != nil {
		log.Println("pipeline create error", err)
		return
	}

	filter := pipeline.FindElement("filter")
	if filter == nil {
		log.Println("pipeline find element error ")
		return
	}

	filter.SetCap("video/x-raw,width=1280,height=720")

	pipeline.Start()
}
