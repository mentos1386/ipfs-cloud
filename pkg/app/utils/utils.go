package utils

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

func IsApplicationWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.ApplicationWindow); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}

func IsDialog(obj glib.IObject) (*gtk.Dialog, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.Dialog); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.Dialog")
}

func IsFileChooserButton(obj glib.IObject) (*gtk.FileChooserButton, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.FileChooserButton); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.FileChooserButton")
}
