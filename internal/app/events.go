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

// onClickMenuFileNew is called when user clicks uiMenuFileNew
func (ptr *Application) onClickMenuFileNew() {
}

// onClickMenuFileOpen is called when user clicks uiMenuFileOpen
func (ptr *Application) onClickMenuFileOpen() {
}

// onClickMenuFileExit is called when user clicks uiMenuFileExit
func (ptr *Application) onClickMenuFileExit() {
}

// onClickMenuFileSave is called when user clicks uiMenuFileSave
func (ptr *Application) onClickMenuFileSave() {
}

// onClickMenuFileSaveAs is called when user clicks uiMenuFileSaveAs
func (ptr *Application) onClickMenuFileSaveAs() {
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
