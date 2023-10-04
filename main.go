package main

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)


//FFMPEG must be installed on system and must be on PATH
//this checks ffmpeg
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

//getting ffmpeg via winget
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

/*

	getting video id from youtube link
for example

https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley

this video's id is dQw4w9WgXcQ

all youtube ids are has 11 characters in it and they coming right after 

				watchv?= 
			
			@Parameter

*/
func getVideoId(videoLink string) string {

	videoID := strings.Split(videoLink, "=")[1]

	if len(videoID) > 11{
		newID := videoID[:11]
		return newID
	}
	return videoID
}


//for clearing unused m4a and m4v files
//they contains video and sound files seperately 
func clearTemps(){

	deleteNec := exec.Command("cmd","/C","del *.m4a *.m4v")
	deleteNec.Run()
}

//downloading 1080p
func downloadHQ(videoID string){

	ctx := context.Background();
	hqDownload := downloader.Downloader{};
	client := youtube.Client{};

	con, err := client.GetVideoContext(context.Background(), videoID)
	if err != nil {
		panic(err)
	}
	hqDownload.DownloadComposite(ctx,"",con,"hd1080","mp4");
}

func main() {
	
	checkFFMPEG()
	clearTemps()

	myApp := app.New();
	myWindow := myApp.NewWindow("go-ytdl");


	//input
	input := widget.NewEntry();
	input.SetPlaceHolder("Paste Youtube Link Here");
	input.Resize(fyne.NewSize(250,40))
	input.Move(fyne.NewPos(75,60))


	//download button
	downButton := widget.NewButton("Download",func() {
		downloadHQ(getVideoId(input.Text))
	})
	downButton.Resize(fyne.NewSize(150,40));
	downButton.Move(fyne.NewPos(125,120))
	

	//clear temp files
	clearButton := widget.NewButton("Clear Temp Files",func() {
		clearTemps();
	})
	clearButton.Resize(fyne.NewSize(150,40));
	clearButton.Move(fyne.NewPos(125,170))


	//content
	content := container.NewWithoutLayout(input,downButton,clearButton)
	
	//run app
	myWindow.CenterOnScreen();
	myWindow.Resize(fyne.NewSize(400, 250));
	myWindow.SetContent(content);
	myWindow.ShowAndRun();
}
