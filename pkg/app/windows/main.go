package windows

import (
	"io/ioutil"
	"log"
	"bytes"


	"github.com/mentos1386/ipfs-cloud/pkg/app/dialogs"
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

	chosePgpKeyObj, err := builder.GetObject("chose-pgp-key")
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
		"encrypt_clicked_cb": func() { encryptClickedCB(builder) },
		"chose-pgp-key_file_set_cb": func() { chosePgpKeyFileSetCB(application, chosePgpButton) },
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

	log.Println("reading private key...")
	key := getPgpKey(builder)

	log.Println("decrypting private key...")
	key, err := key.Unlock([]byte("bananana"))
	errorCheck(err)

	log.Println("reading file to encrypt...")
	decryptedFile := getDecryptedFile(builder)

	log.Println("geting folder to store encrypted file to...")
	encryptedFolderPath := getEncryptedFolderPath(builder)
	log.Println(encryptedFolderPath)

	log.Println("encrypting...")
	keyRing, err := gopenpgp.NewKeyRing(key)
	errorCheck(err)

	pgpMessage, err := keyRing.Encrypt(gopenpgp.NewPlainMessage(decryptedFile), nil)
	errorCheck(err)

	pgpMessageArmored, err := pgpMessage.GetArmored()
	errorCheck(err)

	log.Println(pgpMessageArmored)
}

func chosePgpKeyFileSetCB(application *gtk.Application, fileChooserbutton *gtk.FileChooserButton) {
	unlockKeyDialog, err := dialogs.CreateUnlockKey(application)
	errorCheck(err)

	unlockKeyDialog.Show()
}
