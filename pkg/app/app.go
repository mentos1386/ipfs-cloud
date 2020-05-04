package app

import (
	"context"
	"log"

	"github.com/mentos1386/ipfs-cloud/pkg/app/windows"
	"github.com/mentos1386/ipfs-cloud/pkg/ipfs"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// Create a new ipfs-cloud application
func Create(version string, appID string) (*gtk.Application, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}

	// Connect function to application activate event
	_, err = application.Connect("activate", func() {
		log.Println("application activate")

		ipfs, err := ipfs.StartNode(ctx)
		if err != nil {
			log.Panicf("Failed starting ipfs node! %v", err)
		}

		mainWindow, err := windows.CreateMain(ctx, application, ipfs)
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
		cancel()
	})
	if err != nil {
		return nil, err
	}

	return application, nil
}
