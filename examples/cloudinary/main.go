package main

import (
	"context"
	"fmt"
	goservice "github.com/tronglv92/ecommerce_go_common"
	"log"

	"github.com/tronglv92/ecommerce_go_common/plugin/cloudinary"
	"github.com/tronglv92/ecommerce_go_common/sdkcm"
)

func main() {
	service := goservice.New(
		goservice.WithName("cloudinary"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(cloudinary.New("cloudinary")),
	)

	_ = service.Init()

	videoFile := "videotest.mov" // put this file on project root to test

	cloudinary := service.MustGet("cloudinary").(cloudinary.Cloudinary)

	result, err := cloudinary.VideoUpload(context.Background(), videoFile, "video_preset", "test", "mp4")

	if err != nil {
		log.Fatalf("err: %+v", err.(sdkcm.AppError).Log)
	}

	fmt.Printf("%+v", result)
}
