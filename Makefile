
prepare:
	sudo apt install gcc libgl1-mesa-dev xorg-dev

win:
	go build -ldflags -H=windowsgui yogocut.go