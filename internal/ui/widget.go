package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)


type hoverIcon struct {
	widget.Icon
	tooltip *fyne.Container 
	canvas  fyne.Canvas
}

func newHoverIcon(res fyne.Resource, msg string, c fyne.Canvas) *hoverIcon {
	h := &hoverIcon{canvas: c}
	h.SetResource(res)
	h.ExtendBaseWidget(h)

	bg := canvas.NewRectangle(color.RGBA{R: 30, G: 30, B: 30, A: 255})
	text := widget.NewLabel(msg)
	text.Alignment = fyne.TextAlignCenter
	
	content := container.NewStack(bg, container.NewPadded(text))
	
	h.tooltip = container.NewWithoutLayout(content)

	return h
}

func (h *hoverIcon) MouseIn(e *desktop.MouseEvent) {
	content := h.tooltip.Objects[0]
	minSize := content.MinSize()
	
	content.Resize(minSize)

	newX := e.AbsolutePosition.X - minSize.Width - 10
	newY := e.AbsolutePosition.Y - minSize.Height - 10
	content.Move(fyne.NewPos(newX, newY))

	h.canvas.Overlays().Add(h.tooltip)
}

func (h *hoverIcon) MouseMoved(e *desktop.MouseEvent) {
	content := h.tooltip.Objects[0]
	minSize := content.MinSize()
	
	newX := e.AbsolutePosition.X - minSize.Width - 10
	newY := e.AbsolutePosition.Y - minSize.Height - 10
	content.Move(fyne.NewPos(newX, newY))
}

func (h *hoverIcon) MouseOut() {
	h.canvas.Overlays().Remove(h.tooltip)
}