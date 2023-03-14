package custom

import (
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// FixSizeLabel comment here
type FixSizeLabel struct {
	widget.Label
	fixedSize fyne.Size
	rows      int
	linecount int
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

func deleteFirstLine(s string) string {
	newLineIndex := strings.Index(s, "\n")
	return s[newLineIndex+1:]
}

// WriteAndExpand write to output, and expand height if needed
func (t *FixSizeLabel) WriteAndExpand(s string) {
	t.linecount++
	if t.linecount > 10 {
		t.Text = deleteFirstLine(t.Text)
		t.Text = deleteFirstLine(t.Text)
		t.Text = deleteFirstLine(t.Text)
		t.Text = deleteFirstLine(t.Text)
		t.Text = deleteFirstLine(t.Text)
		t.linecount -= 5
	}
	currentText := t.Text + "\n" + s
	t.SetText(currentText)
}

// SetFixSize set fix size for television
func (t *FixSizeLabel) SetFixSize(s fyne.Size) {
	t.fixedSize = s
}

// MinSize set min size for Television
func (t *FixSizeLabel) MinSize() fyne.Size {
	return t.fixedSize
}
