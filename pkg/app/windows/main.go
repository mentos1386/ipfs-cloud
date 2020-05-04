package windows

import (
	"context"
	"log"

	"github.com/mentos1386/ipfs-cloud/pkg/app/dialogs"
	"github.com/mentos1386/ipfs-cloud/pkg/app/stacks"
	"github.com/mentos1386/ipfs-cloud/pkg/app/utils"

	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/gotk3/gotk3/gtk"
)

func CreateMain(ctx context.Context, application *gtk.Application, ipfs icore.CoreAPI) (*gtk.ApplicationWindow, error) {
	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	if err != nil {
		return nil, err
	}

	stackUpload, err := stacks.CreateUpload(ctx, application, ipfs)
	if err != nil {
		return nil, err
	}
	stackFiles, err := stacks.CreateFiles(application, ipfs)
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
	stackStack.AddTitled(stackUpload, "upload", "Upload")
	stackStack.AddTitled(stackFiles, "files", "Files")

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"account_settings_clicked_cb": func() { accountSettingsClicked(application) },
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

func accountSettingsClicked(application *gtk.Application) {
	accountSettingsDialog, err := dialogs.CreateAccountSettings(application)
	errorCheck(err)

	accountSettingsDialog.Show()
}
