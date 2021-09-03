//go:generate fyne bundle -o data.go Icon.png

package main

import (
	"fmt"
	"log"
	"os/user"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type userInput struct {
	link 	string
	start 	string
	finish 	string
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Yogocut v.0.3")
	myWindow.SetIcon(resourceLogoPng)
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(600, 300))

	// Получаем данные текущего пользователя.
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Youtube Link input
	linkInput := widget.NewEntry()
	linkInput.SetPlaceHolder("Enter Youtube link here...")

	// Start time input
	startInput := widget.NewEntry()
	startInput.SetPlaceHolder("Enter start time of video...")

	// Finish time input
	finishInput := widget.NewEntry()
	finishInput.SetPlaceHolder("Enter finish time of video...")

	// Result output path
	resOutput := widget.NewEntry()
	resOutput.SetText("Donwload to: " + usr.HomeDir)

	downloadBtn := widget.NewButton(
		"Download", 
		func(){
			// Get work data
			data := userInput{
				link: linkInput.Text,
				start: startInput.Text,
				finish: finishInput.Text,
			}
			downloadVideo(resOutput, &data)
		},
	)

	content := container.NewVBox(
		linkInput,
		startInput,
		finishInput, 
		resOutput, 
		downloadBtn,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func downloadVideo(outObj *widget.Entry, data *userInput) {
	fmt.Println(data.link)
	fmt.Println(data.start, data.finish)
	outObj.SetText(data.link + " completed!")
}