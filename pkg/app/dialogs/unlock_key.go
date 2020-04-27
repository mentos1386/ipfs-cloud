package dialogs

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mentos1386/ipfs-cloud/pkg/app/state"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"

	gopenpgp "github.com/ProtonMail/gopenpgp/v2/crypto"
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
	dialog, ok := obj.(*gtk.Dialog)
	if !ok {
		return nil, errors.New("not a *gtk.Dialog")
	}

	passwordObj, err := builder.GetObject("password")
	if err != nil {
		return nil, err
	}
	passwordEntry, err := utils.IsEntry(passwordObj)
	if err != nil {
		return nil, err
	}

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"unlock_private_key_apply_clicked_cb":  func() { unlockKeyApplyCB(builder, passwordEntry, dialog) },
		"unlock_private_key_cancel_clicked_cb": dialog.Close,
	}
	builder.ConnectSignals(signals)

	return dialog, nil
}

// ClickedApply is triggered as callback when
// apply button is clicked
func unlockKeyApplyCB(builder *gtk.Builder, entry *gtk.Entry, dialog *gtk.Dialog) {
	state := state.GetState()

	pgpFile, err := ioutil.ReadFile(state.OpenPGPPrivateKeyPath)
	if err != nil {
		log.Panic(err)
	}

	key, err := gopenpgp.NewKeyFromArmoredReader(bytes.NewReader(pgpFile))
	if err != nil {
		log.Panic(err)
	}

	password, err := entry.GetBuffer()
	if err != nil {
		log.Panic(err)
	}

	passwordText, err := password.GetText()
	if err != nil {
		log.Panic(err)
	}

	unlockedKey, err := key.Unlock([]byte(passwordText))
	if err != nil {
		log.Panic(err)
	}

	state.OpenPGPDecryptedKey = unlockedKey

	dialog.Close()
}
