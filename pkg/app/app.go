package app

import (
	"log"

	"github.com/mentos1386/ipfs-cloud/pkg/app/windows"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Create a new ipfs-cloud application
func Create(version string, appID string) (*gtk.Application, error) {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}

	// Connect function to application startup event, this is not required.
	_, err = application.Connect("startup", func() {
		log.Println("application startup")
	})
	if err != nil {
		return nil, err
	}

	// Connect function to application activate event
	_, err = application.Connect("activate", func() {
		log.Println("application activate")

		mainWindow, err := windows.CreateMain(application)
		if err != nil {
			log.Panicf("Failed creating main window! %v", err)
		}

		mainWindow.Show()
		application.AddWindow(mainWindow)
	})
	if err != nil {
		return nil, err
	}

	// Connect function to application shutdown event, this is not required.
	_, err = application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})
	if err != nil {
		return nil, err
	}

	return application, nil
}


