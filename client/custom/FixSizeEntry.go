package custom

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// FixSizeEntry comment here
type FixSizeEntry struct {
	widget.Entry
	fixedSize fyne.Size
}

// NewFixSizeEntry create new television
func NewFixSizeEntry(s fyne.Size) *FixSizeEntry {
	e := &FixSizeEntry{}
	e.Wrapping = fyne.TextTruncate
	e.ExtendBaseWidget(e)
	e.SetFixSize(s)
	return e
}

// SetFixSize set fix size for television
func (e *FixSizeEntry) SetFixSize(s fyne.Size) {
	e.fixedSize = s
}

// MinSize set min size for Television
func (e *FixSizeEntry) MinSize() fyne.Size {
	return e.fixedSize
}
