prepare:
	sudo apt install gcc libgl1-mesa-dev xorg-dev

win:
	fyne-cross windows -arch=*
