package application

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var builder gtk.Builder

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
			"chose-pgp-key_file_set_cb": chose_pgp_key_file_set_cb,
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

func chose_pgp_key_file_set_cb() {
	_, err := builder.GetObject("chose-pgp-key")
	errorCheck(err)
	//button, err := isFileChooserButton(obj)
	//errorCheck(err)

	//filename := button.GetFilename()

	//log.Printf("Filename changed: %s", filename)
}
