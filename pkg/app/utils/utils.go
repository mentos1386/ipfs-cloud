package utils

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func IsWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if window, ok := obj.(*gtk.Window); ok {
		return window, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func IsApplicationWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if window, ok := obj.(*gtk.ApplicationWindow); ok {
		return window, nil
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

func IsButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.Button); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.IsButton")
}

func IsFileChooserButton(obj glib.IObject) (*gtk.FileChooserButton, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.FileChooserButton); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.FileChooserButton")
}

func IsEntry(obj glib.IObject) (*gtk.Entry, error) {
	// Make type assertion (as per gtk.go).
	if entry, ok := obj.(*gtk.Entry); ok {
		return entry, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func IsStack(obj glib.IObject) (*gtk.Stack, error) {
	// Make type assertion (as per gtk.go).
	if stack, ok := obj.(*gtk.Stack); ok {
		return stack, nil
	}
	return nil, errors.New("not a *gtk.Stack")
}

func IsStackSwitcher(obj glib.IObject) (*gtk.StackSwitcher, error) {
	// Make type assertion (as per gtk.go).
	if stackSwitcher, ok := obj.(*gtk.StackSwitcher); ok {
		return stackSwitcher, nil
	}
	return nil, errors.New("not a *gtk.StackSwitcher")
}

func IsFixed(obj glib.IObject) (*gtk.Fixed, error) {
	// Make type assertion (as per gtk.go).
	if fixed, ok := obj.(*gtk.Fixed); ok {
		return fixed, nil
	}
	return nil, errors.New("not a *gtk.Fixed")
}

func IsBox(obj glib.IObject) (*gtk.Box, error) {
	// Make type assertion (as per gtk.go).
	if box, ok := obj.(*gtk.Box); ok {
		return box, nil
	}
	return nil, errors.New("not a *gtk.Box")
}
