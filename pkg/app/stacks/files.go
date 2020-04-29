package stacks

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"
)

func CreateFiles(application *gtk.Application) (*gtk.Box, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/stacks/files.glade")
	if err != nil {
		return nil, err
	}

	// Get the object with the id of "account_settings_dialog".
	obj, err := builder.GetObject("stack_files")
	if err != nil {
		return nil, err
	}
	box, err := utils.IsBox(obj)
	if err != nil {
		return nil, err
	}

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{}
	builder.ConnectSignals(signals)

	return box, nil
}
