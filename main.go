package main

import (
	"log"
	"media/video"
)

func main() {
	vInfo := video.Stat("video.mp4")
	log.Printf("vInfo=%+v\n", vInfo)
}
