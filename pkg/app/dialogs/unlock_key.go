package dialogs

import (
	"errors"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func CreateUnlockKey(application *gtk.Application) (*gtk.Dialog, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/dialogs/unlock_private_key.glade")
	if err != nil {
		return nil, err
	}

	// Get the object with the id of "unlock_private_key_dialog".
	obj, err := builder.GetObject("unlock_private_key_dialog")
	if err != nil {
		return nil, err
	}

	// Verify that the object is a pointer to a gtk.Dialog.
	dialog, ok := obj.(*gtk.Dialog);
	if !ok {
		return nil, errors.New("not a *gtk.Dialog")
	}
		
	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"unlock_private_key_apply_clicked_cb": func() { unlockKeyApplyCB(builder) },
		"unlock_private_key_cancel_clicked_cb": dialog.Close,
	}
	builder.ConnectSignals(signals)

	return dialog, nil
}

// ClickedApply is triggered as callback when
// apply button is clicked
func unlockKeyApplyCB(builder *gtk.Builder) {
	log.Println("apply!")
}