package diff

import (
	"github.com/go-git/go-git/v6/plumbing/color"
)

type ColorKey string

const (
	Context                   ColorKey = "context"
	Meta                      ColorKey = "meta"
	Frag                      ColorKey = "frag"
	Old                       ColorKey = "old"
	New                       ColorKey = "new"
	Commit                    ColorKey = "commit"
	Whitespace                ColorKey = "whitespace"
	Func                      ColorKey = "func"
	OldMoved                  ColorKey = "oldMoved"
	OldMovedAlternative       ColorKey = "oldMovedAlternative"
	OldMovedDimmed            ColorKey = "oldMovedDimmed"
	OldMovedAlternativeDimmed ColorKey = "oldMovedAlternativeDimmed"
	NewMoved                  ColorKey = "newMoved"
	NewMovedAlternative       ColorKey = "newMovedAlternative"
	NewMovedDimmed            ColorKey = "newMovedDimmed"
	NewMovedAlternativeDimmed ColorKey = "newMovedAlternativeDimmed"
	ContextDimmed             ColorKey = "contextDimmed"
	OldDimmed                 ColorKey = "oldDimmed"
	NewDimmed                 ColorKey = "newDimmed"
	ContextBold               ColorKey = "contextBold"
	OldBold                   ColorKey = "oldBold"
	NewBold                   ColorKey = "newBold"
)

type ColorConfig map[ColorKey]string

type ColorConfigOption func(ColorConfig)

func WithColor(key ColorKey, color string) ColorConfigOption {
	_ = "STUB: not implemented"
	return *new(ColorConfigOption)
}

var defaultColorConfig = ColorConfig{
	Context:                   color.Normal,
	Meta:                      color.Bold,
	Frag:                      color.Cyan,
	Old:                       color.Red,
	New:                       color.Green,
	Commit:                    color.Yellow,
	Whitespace:                color.BgRed,
	Func:                      color.Normal,
	OldMoved:                  color.BoldMagenta,
	OldMovedAlternative:       color.BoldBlue,
	OldMovedDimmed:            color.Faint,
	OldMovedAlternativeDimmed: color.FaintItalic,
	NewMoved:                  color.BoldCyan,
	NewMovedAlternative:       color.BoldYellow,
	NewMovedDimmed:            color.Faint,
	NewMovedAlternativeDimmed: color.FaintItalic,
	ContextDimmed:             color.Faint,
	OldDimmed:                 color.FaintRed,
	NewDimmed:                 color.FaintGreen,
	ContextBold:               color.Bold,
	OldBold:                   color.BoldRed,
	NewBold:                   color.BoldGreen,
}

func NewColorConfig(options ...ColorConfigOption) ColorConfig {
	_ = "STUB: not implemented"
	return *new(ColorConfig)
}

func (cc ColorConfig) Reset(key ColorKey) string { _ = "STUB: not implemented"; return "" }
