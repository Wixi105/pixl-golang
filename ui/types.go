package ui

import (
	"fyne.io/fyne/v2"
	"github.com/Wixi105/pixl-golang/apptype"
	"github.com/Wixi105/pixl-golang/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
