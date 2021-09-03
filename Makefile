prepare:
	sudo apt install gcc libgl1-mesa-dev xorg-dev

lx:
	fyne-cross linux

win:
	fyne-cross windows -arch=*

icon:
	fyne bundle Icon.png > icon.go

lx-start:
	fyne-cross/bin/linux-amd64/yogocut