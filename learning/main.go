package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)

func main(){
	videoID := "t6J-m7I056E";
	client := youtube.Client{};
	hqDownload := downloader.Downloader{}


	video,err := client.GetVideo(videoID);
	if err != nil{
		panic(err);
	}



	//HQ DOWNLOAD
	con,err := client.GetVideoContext(context.Background(),videoID);
	if err != nil{
		panic(err)
	}

	hqDownload.DownloadComposite(context.Background(),"" , con,"hd1080","mp4");

	//TODO
	//ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -strict experimental output.mp4

	


	//videoInfo,err := client.GetVideoContext(context.Background(),videoID);
	if err != nil{
		panic(err);
	}

	description,err := os.Create(videoID + ".txt");
	if err != nil{
		panic(err)
	}
	defer description.Close();
	fmt.Fprintf(description,video.Description);
	//fmt.Println(videoInfo);

	//0 720p 128kb/s
	//1 360p
	//2
	//3
	//4 144p 24kb/s
	//5

	formats := video.Formats.WithAudioChannels();
	stream,_,err := client.GetStream(video,&formats[0]);

	if err != nil{
		panic(err);
	}

	defer stream.Close();

	file2,err := os.Create("video.mp4");

	if err != nil{
		panic(err);
	}

	defer file2.Close();

	_,err = io.Copy(file2,stream);
	if err != nil{
		panic(err);
	}

}