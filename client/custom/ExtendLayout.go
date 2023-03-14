package custom

import (
	"fyne.io/fyne/v2"
)

// ExtendLayout ...
type ExtendLayout struct {
	isVertical bool
}

// CreateExtendLayout ...
func CreateExtendLayout(isVertical bool) *ExtendLayout {
	return &ExtendLayout{isVertical}
}

// Layout ...
func (g *ExtendLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	count := 0
	for _, child := range objects {
		if child.Visible() {
			count++
		}
	}

	var x float32
	var y float32

	i := 0
	for ; i < len(objects) - 1; i++ {
		objects[i].Move(fyne.NewPos(x, y))
		if g.isVertical {
			y +=  objects[i].MinSize().Height
		} else {
			x += objects[i].MinSize().Width
		}
	}

	if g.isVertical {
		objects[i].Resize(fyne.NewSize(objects[i].MinSize().Width, size.Height-y-5))
		objects[i].Move(fyne.NewPos(x, y))
	} else {
		objects[i].Resize(fyne.NewSize(size.Width-x-5, objects[i].MinSize().Height))
		objects[i].Move(fyne.NewPos(x, y))
	}
}

// MinSize ...
func (g *ExtendLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var width float32
	var height float32
	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		if g.isVertical {
			height += child.MinSize().Height
			if width < child.MinSize().Width {
				width = child.MinSize().Width
			} 
		} else {
			width += child.MinSize().Width
			if height < child.MinSize().Height {
				height = child.MinSize().Height
			}
		}
	}
	return fyne.NewSize(width, height)
}
