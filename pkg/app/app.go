package app

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"bytes"

	gopenpgp "github.com/ProtonMail/gopenpgp/v2/crypto"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Create a new ipfs-cloud application
func Create(version string, appID string, args []string) {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNewFromFile("ui/example.glade")
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"encrypt_clicked_cb": func() { encrypt_clicked_cb(builder) },
		}
		builder.ConnectSignals(signals)

		// Get the object with the id of "main_window".
		obj, err := builder.GetObject("main_window")
		errorCheck(err)

		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(obj)
		errorCheck(err)

		// Show the Window and all of its components.
		win.Show()
		application.AddWindow(win)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(args))
}

func isWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}

func isFileChooserButton(obj glib.IObject) (*gtk.FileChooserButton, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.FileChooserButton); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.FileChooserButton")
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
	pgpButton, err := isFileChooserButton(pgpObj)
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
	fileButton, err := isFileChooserButton(fileObj)
	errorCheck(err)

	filename := fileButton.GetFilename()
	b, err := ioutil.ReadFile(filename)
	errorCheck(err)

	return b
}

func getEncryptedFolderPath(builder *gtk.Builder) string {
	folderObj, err := builder.GetObject("chose-encrypted")
	errorCheck(err)
	folderButton, err := isFileChooserButton(folderObj)
	errorCheck(err)

	return folderButton.GetFilename()
}

func encrypt_clicked_cb(builder *gtk.Builder) {

	log.Println("reading private key...")
	key := getPgpKey(builder)

	log.Println("decrypting private key...")
	key, err := key.Unlock([]byte("banana"))
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
