package windows

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/mentos1386/ipfs-cloud/pkg/app/dialogs"
	"github.com/mentos1386/ipfs-cloud/pkg/app/state"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"

	gopenpgp "github.com/ProtonMail/gopenpgp/v2/crypto"

	"github.com/gotk3/gotk3/gtk"
)

func CreateMain(application *gtk.Application) (*gtk.ApplicationWindow, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	if err != nil {
		return nil, err
	}

	stackUploadObj, err := builder.GetObject("stack_upload")
	if err != nil {
		return nil, err
	}
	stackUploadFixed, err := utils.IsFixed(stackUploadObj)
	if err != nil {
		return nil, err
	}
	stackSharedFilesObj, err := builder.GetObject("stack_shared_files")
	if err != nil {
		return nil, err
	}
	stackSharedFilesFixed, err := utils.IsFixed(stackSharedFilesObj)
	if err != nil {
		return nil, err
	}

	stackSwitcherObj, err := builder.GetObject("stack_switcher")
	if err != nil {
		return nil, err
	}
	stackSwitcherSwitcher, err := utils.IsStackSwitcher(stackSwitcherObj)
	if err != nil {
		return nil, err
	}

	stackObj, err := builder.GetObject("stack")
	if err != nil {
		return nil, err
	}
	stackStack, err := utils.IsStack(stackObj)
	if err != nil {
		return nil, err
	}
	stackStack.AddTitled(stackUploadFixed, "upload", "Upload")
	stackStack.AddTitled(stackSharedFilesFixed, "shared-files", "Shared Files")
	stackSwitcherSwitcher.SetStack(stackStack)

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"account_settings_clicked_cb": func() { accountSettingsClicked(application) },
		"encrypt_clicked_cb":          func() { encryptClickedCB(builder) },
	}
	builder.ConnectSignals(signals)

	// Get the object with the id of "main_window".
	obj, err := builder.GetObject("main_window")
	if err != nil {
		return nil, err
	}
	mainWindow, err := utils.IsApplicationWindow(obj)
	if err != nil {
		return nil, err
	}

	mainWindow.SetApplication(application)

	return mainWindow, nil
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

func getPgpKey(builder *gtk.Builder) *gopenpgp.Key {
	pgpObj, err := builder.GetObject("chose-pgp-key")
	errorCheck(err)
	pgpButton, err := utils.IsFileChooserButton(pgpObj)
	errorCheck(err)

	pgpFilename := pgpButton.GetFilename()
	pgpFile, err := ioutil.ReadFile(pgpFilename)
	errorCheck(err)

	key, err := gopenpgp.NewKeyFromArmoredReader(bytes.NewReader(pgpFile))
	errorCheck(err)

	return key
}

func getDecryptedFile(builder *gtk.Builder) []byte {
	fileObj, err := builder.GetObject("chose-decrypted")
	errorCheck(err)
	fileButton, err := utils.IsFileChooserButton(fileObj)
	errorCheck(err)

	filename := fileButton.GetFilename()
	b, err := ioutil.ReadFile(filename)
	errorCheck(err)

	return b
}

func getEncryptedFolderPath(builder *gtk.Builder) string {
	folderObj, err := builder.GetObject("chose-encrypted")
	errorCheck(err)
	folderButton, err := utils.IsFileChooserButton(folderObj)
	errorCheck(err)

	return folderButton.GetFilename()
}

func encryptClickedCB(builder *gtk.Builder) {
	state := state.GetState()

	log.Println("reading file to encrypt...")
	decryptedFile := getDecryptedFile(builder)

	log.Println("geting folder to store encrypted file to...")
	encryptedFolderPath := getEncryptedFolderPath(builder)
	log.Println(encryptedFolderPath)

	log.Println("encrypting...")
	keyRing, err := gopenpgp.NewKeyRing(state.OpenPGPDecryptedKey)
	errorCheck(err)

	pgpMessage, err := keyRing.Encrypt(gopenpgp.NewPlainMessage(decryptedFile), nil)
	errorCheck(err)

	pgpMessageArmored, err := pgpMessage.GetArmored()
	errorCheck(err)

	log.Println(pgpMessageArmored)
}

func accountSettingsClicked(application *gtk.Application) {
	accountSettingsDialog, err := dialogs.CreateAccountSettings(application)
	errorCheck(err)

	accountSettingsDialog.Show()
}
