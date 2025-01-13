package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/layout"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"time"
)

var paused bool
var remainingTime time.Duration

func startPomo() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Pomodoro Timer")
	myWindow.Resize(fyne.NewSize(300, 300))

	//app icon
	iconFile, err := os.Open("images/clock-icon.png")
	if err == nil {
		defer iconFile.Close()
		iconBytes, err := os.ReadFile("images/clock-icon.png")
		if err == nil {
			iconResource := fyne.NewStaticResource("clock-icon.png", iconBytes)
			myWindow.SetIcon(iconResource)
		} else {
			fmt.Println("Error reading icon file:", err)
		}
	} else {
		fmt.Println("Error opening icon file:", err)
	}

	//background image
	backgroundFile, err := os.Open("images/background.png")
	if err != nil {
		return
	}
	defer backgroundFile.Close()

	backgroundImg, _, err := image.Decode(backgroundFile)
	if err != nil {
		return
	}

	bgImage := canvas.NewImageFromImage(backgroundImg)
	bgImage.FillMode = canvas.ImageFillStretch

	pomodoroDuration := 25 * time.Minute
	breakDuration := 5 * time.Minute
	longBreakDuration := 15 * time.Minute
	intervals := 0

	timerText := canvas.NewText("25:00", color.White)
	timerText.TextSize = 72
	timerText.Alignment = fyne.TextAlignCenter
	timerText.TextStyle.Bold = true

	statusLabel := widget.NewLabel("Ready to be productive??")

	startButton := widget.NewButton("Start", nil)
	pauseButton := widget.NewButton("Pause", nil)
	pauseButton.Disable()

	startButton.OnTapped = func() {
		paused = false
		startButton.Disable()
		pauseButton.Enable()
		statusLabel.SetText("Pomodoro in progress...")
		remainingTime = pomodoroDuration
		runTimer(&remainingTime, timerText, func() {
			statusLabel.SetText("Pomodoro Complete! Go outside.")
			intervals++
			pauseButton.Disable()
			if intervals%4 == 0 {
				runTimer(&longBreakDuration, timerText, func() {
					statusLabel.SetText("Long break complete! Back to work.")
					startButton.Enable()
				})
			} else {
				runTimer(&breakDuration, timerText, func() {
					statusLabel.SetText("Break over! Start another Pomodoro.")
					startButton.Enable()
				})
			}
		})
	}

	pauseButton.OnTapped = func() {
		if paused {
			pauseButton.SetText("Pause")
			statusLabel.SetText("Pomodoro in progress...")
			paused = false
		} else {
			pauseButton.SetText("Resume")
			statusLabel.SetText("Paused")
			paused = true
		}
	}

	// myWindow.SetContent(container.NewVBox( //contents in order
	// 	statusLabel,
	// 	container.NewCenter(timerText), // Center the timer text
	// 	widget.NewSeparator(),
	// 	container.NewVBox(
	// 		startButton,
	// 		pauseButton,
	// 	),
	// ))

	//layout
	content := container.NewBorder(
		statusLabel, // Top
		container.NewVBox(
			container.NewMax( // Empty space as a gap
				canvas.NewRectangle(color.Transparent),
			),
			startButton,
			pauseButton,
		), // Bottom
		nil, nil, // Left and right (not used)
		timerText, // Centered timer with a gap
	)
	layeredContent := container.NewMax(
		bgImage,
		content,
	)
	myWindow.SetContent(layeredContent)

	myWindow.ShowAndRun()
}

func runTimer(remaining *time.Duration, timerText *canvas.Text, onComplete func()) {
	go func() {
		for *remaining > 0 {
			if paused {
				time.Sleep(time.Second)
				continue
			}
			*remaining -= time.Second
			minutes := int(remaining.Minutes())
			seconds := int(remaining.Seconds()) % 60
			timerText.Text = fmt.Sprintf("%02d:%02d", minutes, seconds)
			timerText.Refresh()
			time.Sleep(time.Second)
		}
		timerText.Text = "00:00"
		timerText.Refresh()
		onComplete()
	}()
}
