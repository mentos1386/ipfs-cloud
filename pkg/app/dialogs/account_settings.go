package dialogs

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mentos1386/ipfs-cloud/pkg/app/state"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"
)

func CreateAccountSettings(application *gtk.Application) (*gtk.ApplicationWindow, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/dialogs/account_settings.glade")
	if err != nil {
		return nil, err
	}

	// Get the object with the id of "account_settings_dialog".
	obj, err := builder.GetObject("account_settings_dialog")
	if err != nil {
		return nil, err
	}
	window, err := utils.IsApplicationWindow(obj)
	if err != nil {
		return nil, err
	}

	chosePgpKeyObj, err := builder.GetObject("chose_pgp_key")
	if err != nil {
		return nil, err
	}
	chosePgpButton, err := utils.IsFileChooserButton(chosePgpKeyObj)
	if err != nil {
		return nil, err
	}

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"chose_pgp_key_file_set_cb":         func() { chosePgpKeyFileSetCB(application, chosePgpButton) },
		"account_settings_apply_clicked_cb": window.Close,
	}
	builder.ConnectSignals(signals)

	return window, nil
}

func chosePgpKeyFileSetCB(application *gtk.Application, button *gtk.FileChooserButton) {
	unlockKeyDialog, err := CreateUnlockKey(application)
	if err != nil {
		log.Panic(err)
	}

	state := state.GetState()
	state.OpenPGPPrivateKeyPath = button.GetFilename()

	unlockKeyDialog.Show()
}
