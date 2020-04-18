package app

import (
	"log"
	"os"

	"github.com/mentos1386/ipfs-cloud/pkg/app/windows"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Create a new ipfs-cloud application
func Create(version string, appID string, args []string) {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Panicf("Failed creating gtk application! %v", err)
	}

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		mainWindow, err := windows.CreateMain(application)
		if err != nil {
			log.Panicf("Failed creating main window! %v", err)
		}

		mainWindow.Show()
		application.AddWindow(mainWindow)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(args))
}


