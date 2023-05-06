package digitalclock

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is a text widget struct to hold info about the current widget
type Widget struct {
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new widget using settings
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
	}
	if settings.centerAlign {
		widget.View.SetTextAlign(tview.AlignCenter)
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.display()
}
