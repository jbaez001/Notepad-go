/*
Copyright 2021 Juan Baez

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/jbaez001/Notepad-go/pkg/version"
)

// message boxes
const (
	MsgBoxAboutNotepad uint8 = 1 << iota
)

// ui strings
const (
	uiMasterWindowTitle string = "Notepad"

	// file
	uiMenuFile          string = "File"
	uiMenuFileNew       string = "New"
	uiMenuFileOpen      string = "Open..."
	uiMenuFileSave      string = "Save"
	uiMenuFileSaveAs    string = "Save As..."
	uiMenuFilePageSetup string = "Page Setup"
	uiMenuFilePrint     string = "Print..."
	uiMenuFileExit      string = "Exit"

	// edit
	uiMenuEdit          string = "Edit"
	uiMenuEditUndo      string = "Undo"
	uiMenuEditCut       string = "Cut"
	uiMenuEditCopy      string = "Copy"
	uiMenuEditPaste     string = "Paste"
	uiMenuEditDelete    string = "Delete"
	uiMenuEditFind      string = "Find"
	uiMenuEditFindNext  string = "Find Next"
	uiMenuEditReplace   string = "Replace"
	uiMenuEditGoTo      string = "Go To..."
	uiMenuEditSelectAll string = "Select All"
	uiMenuEditTimeDate  string = "Time/Date"

	// format
	uiMenuFormat               string = "Format"
	uiMenuFormatWordWrap       string = "Word Wrap"
	uiMenuFormatFont           string = "Font"
	uiMenuFormatShowWhitespace string = "View Whitespace"
	uiMenuFormatShowBorder     string = "View Border"

	// view
	uiMenuView                   string = "View"
	uiMenuViewZoom               string = "Zoom"
	uiMenuViewZoomIn             string = "Zoom In"
	uiMenuViewZoomOut            string = "Zoom Out"
	uiMenuViewZoomRestoreDefault string = "Restore Default Zoom"
	uiMenuViewStatusBar          string = "Status Bar"

	// help
	uiMenuHelp      string = "Help"
	uiMenuHelpView  string = "View Help"
	uiMenuHelpAbout string = "About Notepad"
)

// Application is our app type
type Application struct {
	masterWindow  *giu.MasterWindow
	currentMsgBox uint8

	// text editor
	textEditor        imgui.TextEditor
	currentFilename   string
	currentFileBuffer []byte
	hasUnsavedChanges bool

	// options
	optionWordWrap       bool //wordwrap
	optionShowWhiteSpace bool // white space
	optionShowBorder     bool // text editor border
	optionShowStatusBar  bool // show status bar
}

func New() *Application {
	// init the app
	_app := &Application{
		masterWindow: giu.NewMasterWindow(
			uiMasterWindowTitle,
			800,
			600,
			0,
		),
		textEditor: imgui.NewTextEditor(),
	}

	// initial textEditor settings
	_app.textEditor.SetTabSize(2)
	_app.textEditor.SetShowWhitespaces(false)

	return _app
}

// msgBoxResultAboutNotepad gets called whenever the user interacts
// with onClickMenuAboutNotepad
func (ptr *Application) msgBoxResultAboutNotepad(
	result giu.DialogResult,
) {
	if result == giu.DialogResultYes {
		os.Exit(0)
	} else {
		ptr.setCurrentMsgBox(0)
	}
}

// fileMenu will return the file menu
func (ptr *Application) fileMenu() *giu.MenuWidget {
	menuLayout := giu.Menu(uiMenuFile).Layout(
		giu.MenuItem(uiMenuFileNew).OnClick(ptr.onClickMenuFileNew),
		giu.MenuItem(uiMenuFileOpen).OnClick(ptr.onClickMenuFileOpen),
		giu.MenuItem(uiMenuFileSave).OnClick(ptr.onClickMenuFileSave),
		giu.MenuItem(uiMenuFileSaveAs).OnClick(ptr.onClickMenuFileSaveAs),
		giu.Separator(),
		giu.MenuItem(uiMenuFilePageSetup),
		giu.MenuItem(uiMenuFilePrint),
		giu.Separator(),
		giu.MenuItem(uiMenuFileExit).OnClick(ptr.onClickMenuFileExit),
	)

	return menuLayout
}

// editMenu will return the edit menu
func (ptr *Application) editMenu() *giu.MenuWidget {
	menuLayout := giu.Menu(uiMenuEdit).Layout(
		giu.MenuItem(uiMenuEditUndo),
		giu.Separator(),
		giu.MenuItem(uiMenuEditCut).OnClick(ptr.onClickMenuEditCut),
		giu.MenuItem(uiMenuEditCopy).OnClick(ptr.onClickMenuEditCopy),
		giu.MenuItem(uiMenuEditPaste).OnClick(ptr.onClickMenuEditPaste),
		giu.MenuItem(uiMenuEditDelete).OnClick(ptr.onClickMenuEditDelete),
		giu.Separator(),
		giu.MenuItem(uiMenuEditFind),
		giu.MenuItem(uiMenuEditFindNext),
		giu.MenuItem(uiMenuEditReplace),
		giu.MenuItem(uiMenuEditGoTo),
		giu.Separator(),
		giu.MenuItem(uiMenuEditSelectAll).OnClick(ptr.onClickSelectAll),
		giu.MenuItem(uiMenuEditTimeDate),
	)

	return menuLayout
}

// formatMenu will return the format menu
func (ptr *Application) formatMenu() *giu.MenuWidget {
	menuLayout := giu.Menu(uiMenuFormat).Layout(
		giu.Checkbox(uiMenuFormatWordWrap, &ptr.optionWordWrap),
		giu.MenuItem(uiMenuFormatFont),
		giu.Separator(),
		giu.Checkbox(uiMenuFormatShowWhitespace,
			&ptr.optionShowWhiteSpace).
			OnChange(ptr.onClickMenuFormatShowWhitespace),
		giu.Checkbox(uiMenuFormatShowBorder,
			&ptr.optionShowBorder).
			OnChange(ptr.onClickMenuFormatShowBorder),
	)

	return menuLayout
}

// viewMenu will return the view menu
func (ptr *Application) viewMenu() *giu.MenuWidget {

	zoomMenu := giu.Menu(uiMenuViewZoom).Layout(
		giu.MenuItem(uiMenuViewZoomIn),
		giu.MenuItem(uiMenuViewZoomOut),
		giu.MenuItem(uiMenuViewZoomRestoreDefault),
	)

	menuLayout := giu.Menu(uiMenuView).Layout(
		zoomMenu,
		giu.Checkbox(uiMenuViewStatusBar, &ptr.optionShowStatusBar),
	)

	return menuLayout
}

// helpMenu will return the help menu
func (ptr *Application) helpMenu() *giu.MenuWidget {
	menuLayout := giu.Menu(uiMenuHelp).Layout(
		giu.MenuItem(uiMenuHelpView),
		giu.MenuItem(uiMenuHelpAbout).OnClick(ptr.onClickMenuAboutNotepad),
	)

	return menuLayout
}

// handle application updates
func (ptr *Application) update() {
	// check to see if we have any unsaved changes
	if ptr.textEditor.IsTextChanged() && !ptr.hasUnsavedChanges {
		ptr.hasUnsavedChanges = true
	}

	giu.Update()
}

// renderTopMenu will return the top menu
func (ptr *Application) renderTopMenu() *giu.MenuBarWidget {
	// create the layout
	menuLayout := giu.Layout{
		// prepare message boxes
		giu.PrepareMsgbox(),
		ptr.fileMenu(),
		ptr.editMenu(),
		ptr.formatMenu(),
		ptr.viewMenu(),
		ptr.helpMenu(),
	}

	// create the menubar layout
	menuBarLayout := giu.MenuBar().Layout(menuLayout)
	menuBarLayout.Build()

	return menuBarLayout
}

// renderTextEditor will return the actual text editor
func (ptr *Application) renderTextEditor() giu.Widget {
	return giu.Custom(func() {
		// render text textEditor
		ptr.textEditor.Render(
			"",
			imgui.Vec2{
				X: 0,
				Y: 0,
			},
			ptr.optionShowBorder,
		)

		// show message boxes; if any
		ptr.showMessageBoxes()
	})
}

// render will render the Application
func (ptr *Application) render() {
	// render the main window
	giu.SingleWindowWithMenuBar().Layout(
		ptr.renderTopMenu(),
		ptr.renderTextEditor(),
	)

	// handle updates
	ptr.update()
}

// setCurrentMsgBox sets the current message box
func (ptr *Application) setCurrentMsgBox(msgboxId uint8) {
	ptr.currentMsgBox = msgboxId
}

func (ptr *Application) saveFile(filename string) {
	// update the file buffer
	ptr.currentFileBuffer = []byte(ptr.textEditor.GetText())

	if err := ioutil.WriteFile(filename,
		ptr.currentFileBuffer, 066); err != nil {
		log.Println("error saving file", err)
		return
	}

	ptr.currentFilename = filename
	log.Println(ptr.currentFilename, "saved...")

	if ptr.hasUnsavedChanges {
		ptr.hasUnsavedChanges = false
	}
}

// showMessageBoxes
func (ptr *Application) showMessageBoxes() {
	// render the current dialog only
	switch ptr.currentMsgBox {

	// about Notepad
	case MsgBoxAboutNotepad:
		giu.MsgboxV(
			uiMenuHelpAbout,
			fmt.Sprintf(
				"%s\n\n%s\n\n%s\n\n\n",
				uiMasterWindowTitle,
				version.Version,
				version.DateCompiled,
			),
			giu.MsgboxButtonsOk,
			ptr.msgBoxResultAboutNotepad,
		)

	// nothing to do
	default:
		break
	}
}

// Start will run the app
func (ptr *Application) Start() {
	ptr.masterWindow.Run(ptr.render)
}

// Notepad launches our app
func Notepad() {
	notitas := New()
	notitas.Start()
}
