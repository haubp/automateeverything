package custom

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// FixSizeLabel comment here
type FixSizeLabel struct {
	widget.Label
	fixedSize fyne.Size
	rows int
}

// NewFixSizeLabel create new television
func NewFixSizeLabel() *FixSizeLabel {
	t := &FixSizeLabel{}
	t.Text = ""
	t.Alignment = fyne.TextAlignLeading
	t.TextStyle = fyne.TextStyle{}
	t.rows = 2

	t.ExtendBaseWidget(t)
	return t
}

// WriteAndExpand write to output, and expand height if needed
func (t *FixSizeLabel) WriteAndExpand(s string) {
	s = t.Text + "\n" + s
	if len(s) > 1000 {
		s = s[50:]
	}
	t.SetFixSize(fyne.NewSize(t.Size().Width, t.Size().Height + 20))
	t.SetText(s)
}

// SetFixSize set fix size for television
func (t *FixSizeLabel) SetFixSize(s fyne.Size) {
	t.fixedSize = s
}

// MinSize set min size for Television
func (t *FixSizeLabel) MinSize() fyne.Size {
	return t.fixedSize
}
