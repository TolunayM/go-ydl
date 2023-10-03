package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/kkdai/youtube/v2"
	//  "github.com/u2takey/ffmpeg-go"
	"github.com/kkdai/youtube/v2/downloader"
)

/*
func ffmpeg(videoid string,soundId string){
	video := ffmpeg_go.Input(videoid);
	sound := ffmpeg_go.Input(soundId);
	ffmpeg_go.
	//KwArgs format = ffmpeg_go.KwArgs{"f":"mp4"}


	file,err := os.Create("videoWithSound.mp4");
	if err != nil {
		panic(err);
	}

	defer file.Close();

	_,err = io.Copy(file,output);
	if err != nil{
		panic(err)
	}
}*/

func checkFFMPEG() {
	os := runtime.GOOS
	if os == "windows" {
		cmd := exec.Command("ffmpeg", "-version")
		_, err := cmd.Output()
		if err != nil {
			fmt.Print("FFMPEG must be installed first!")
			getFFMPEG()
		}
		fmt.Println("FFMPEG is good to go")
	}
}
func getFFMPEG() {
	os := runtime.GOOS
	if os == "windows" {
		cmd := exec.Command("cmd", "/C", "winget install ffmpeg")
		output, _ := cmd.Output()
		cmd.Stdin = strings.NewReader("Y")
		cmd.Run()
		fmt.Print(string(output))
	}
	if os == "linux" {
		fmt.Print("FFMPEG MUST BE INSTALLED")
	}

	if os == "darwin" {
		fmt.Print("FFMPEG MUST BE INSTALLED")
	}
}

func getVideoId(videoLink string) string {
	videoID := strings.Split(videoLink, "=")[1]
	return videoID
}

func main() {

	checkFFMPEG()
	fmt.Println("Video linki")
	var videoLink string
	fmt.Scanln(&videoLink)

	videoID := getVideoId(videoLink)
	client := youtube.Client{}
	hqDownload := downloader.Downloader{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	//HQ DOWNLOAD
	con, err := client.GetVideoContext(context.Background(), videoID)
	if err != nil {
		panic(err)
	}

	hqDownload.DownloadComposite(context.Background(), "", con, "hd1080", "mp4")

	//TODO
	//ffmpeg -i video.mp4 -i audio.mp3 -c:v copy -c:a aac -strict experimental output.mp4

	//add ffmpeg installer for 1 time use first get os then use command script for installing like winget install ffmpeg
	// or use ffmpeg wrapper for who don't want to install ffmpeg

	//videoInfo,err := client.GetVideoContext(context.Background(),videoID);
	if err != nil {
		panic(err)
	}

	description, err := os.Create(videoID + ".txt")
	if err != nil {
		panic(err)
	}
	defer description.Close()
	fmt.Fprintf(description, video.Description)
	//fmt.Println(videoInfo);

	//0 720p 128kb/s
	//1 360p
	//2
	//3
	//4 144p 24kb/s
	//5

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])

	if err != nil {
		panic(err)
	}

	defer stream.Close()

	file2, err := os.Create("video.mp4")

	if err != nil {
		panic(err)
	}

	defer file2.Close()

	_, err = io.Copy(file2, stream)
	if err != nil {
		panic(err)
	}

}
