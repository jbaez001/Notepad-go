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
	"io/ioutil"
	"os"

	"github.com/sqweek/dialog"
)

// onClickMenuFileNew is called when user clicks uiMenuFileNew
func (ptr *Application) onClickMenuFileNew() {

	// reset the text editor
	ptr.textEditor.SetText("")
}

// onClickMenuFileOpen is called when user clicks uiMenuFileOpen
func (ptr *Application) onClickMenuFileOpen() {
	fileName, err := dialog.File().Filter(
		"Open file",
	).Load()
	if err != nil || fileName == "" {
		return
	}

	// open the file
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	// update current file information
	ptr.currentFilename = fileName
	ptr.currentFileBuffer = fileContents

	// display the contents
	ptr.textEditor.SetText(string(ptr.currentFileBuffer))
}

// onClickMenuFileExit is called when user clicks uiMenuFileExit
func (ptr *Application) onClickMenuFileExit() {

	// confirm whether the user intended to exit Notepad or not
	confirmationDialog := dialog.MsgBuilder{
		Dlg: dialog.Dlg{Title: "Exit Notepad"},
		Msg: "Exit Notepad?",
	}

	// TODO: should really only do this if we have unsaved changes
	if confirmationDialog.YesNo() {
		os.Exit(0)
	}
}

// onClickMenuFileSave is called when user clicks uiMenuFileSave
func (ptr *Application) onClickMenuFileSave() {
	// sanity check
	if ptr.currentFilename == "" || ptr.currentFileBuffer == nil {
		return
	}

	// save the file
	ptr.saveFile(ptr.currentFilename)
}

// onClickMenuFileSaveAs is called when user clicks uiMenuFileSaveAs
func (ptr *Application) onClickMenuFileSaveAs() {
	// if file exists, the below method will also prompt the
	// user for confirmation whether or not the intent is to
	// overwrite the intended file
	filename, err := dialog.File().Filter(
		"Select",
	).Save()
	if err != nil || filename == "" {
		return
	}

	// save the file
	ptr.saveFile(filename)
}

// onClickMenuEditCut is called when user clicks uiMenuEditCut
func (ptr *Application) onClickMenuEditCut() {
	ptr.textEditor.Cut()
}

// onClickMenuEditCopy is called when user clicks uiMenuEditCopy
func (ptr *Application) onClickMenuEditCopy() {
	ptr.textEditor.Copy()
}

// onClickMenuEditPaste is called when user clicks uiMenuEditPaste
func (ptr *Application) onClickMenuEditPaste() {
	ptr.textEditor.Paste()
}

// onClickMenuEditDelete is called when user clicks uiMenuEditDelete
func (ptr *Application) onClickMenuEditDelete() {
	ptr.textEditor.Delete()
}

// onClickSelectAll is called when user clicks uiMenuEditSelectAll
func (ptr *Application) onClickSelectAll() {
	// https://github.com/AllenDang/imgui-go/pull/7
	ptr.textEditor.SelectAll()
}

// onClickMenuFormatShowWhitespace is called when user clicks
// uiMenuFormatShowWhitespace
func (ptr *Application) onClickMenuFormatShowWhitespace() {
}

// onClickMenuFormatShowBorder is called when user clicks
// uiMenuFormatShowBorder
func (ptr *Application) onClickMenuFormatShowBorder() {
	ptr.textEditor.SetShowWhitespaces(ptr.optionShowWhiteSpace)
}

// onClickMenuAboutNotepad is called when user clicks uiMenuHelpAbout
func (ptr *Application) onClickMenuAboutNotepad() {
	ptr.setCurrentMsgBox(MsgBoxAboutNotepad)
}
