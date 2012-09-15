package ui

import "github.com/Nightgunner5/glut"

var window glut.Window

func Drive() {
	window = glut.CreateWindow("mmGo")
	glut.DisplayFunc(DisplayMenu)
	glut.ReshapeFunc(ReshapeMenu)
	glut.MainLoop()
}
