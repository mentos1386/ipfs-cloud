package stacks

import (
	"io/ioutil"
	"log"

	gopenpgp "github.com/ProtonMail/gopenpgp/v2/crypto"
	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mentos1386/ipfs-cloud/pkg/app/state"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"
	"github.com/mentos1386/ipfs-cloud/pkg/ipfs"
)

func CreateUpload(application *gtk.Application, node icore.CoreAPI) (*gtk.Box, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/stacks/upload.glade")
	if err != nil {
		return nil, err
	}

	// Get the object with the id of "account_settings_dialog".
	obj, err := builder.GetObject("stack_upload")
	if err != nil {
		return nil, err
	}
	box, err := utils.IsBox(obj)
	if err != nil {
		return nil, err
	}

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"upload_file_file_set_cb": func() { uploadFile(builder, node) },
	}
	builder.ConnectSignals(signals)

	return box, nil
}

func getDecryptedFile(builder *gtk.Builder) (*[]byte, error) {
	fileObj, err := builder.GetObject("upload_file")
	if err != nil {
		return nil, err
	}

	fileButton, err := utils.IsFileChooserButton(fileObj)
	if err != nil {
		return nil, err
	}

	filename := fileButton.GetFilename()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func uploadFile(builder *gtk.Builder, node icore.CoreAPI) {
	state := state.GetState()

	log.Println("reading file to encrypt...")
	decryptedFile, err := getDecryptedFile(builder)
	if err != nil {
		log.Panic(err)
	}

	log.Println("geting folder to store encrypted file to...")
	encryptedFolderPath := "/home/tinej-personal/Downloads"
	log.Println(encryptedFolderPath)

	log.Println("encrypting...")
	keyRing, err := gopenpgp.NewKeyRing(state.OpenPGPDecryptedKey)
	if err != nil {
		log.Panic(err)
	}

	pgpMessage, err := keyRing.Encrypt(gopenpgp.NewPlainMessage(*decryptedFile), nil)
	if err != nil {
		log.Panic(err)
	}

	pgpMessageArmored, err := pgpMessage.GetArmored()
	if err != nil {
		log.Panic(err)
	}

	log.Println(pgpMessageArmored)

	path, err := ipfs.Store(pgpMessageArmored, node)
	if err != nil {
		log.Panic(err)
	}

	log.Println(path)
}
