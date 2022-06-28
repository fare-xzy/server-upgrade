package main

import (
	"fmt"
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
	Edit   *walk.TextEdit
	OutPut *walk.TextEdit
}

// SelectFile 文件上传按纽
func (mw *MyMainWindow) SelectFile() {

	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "文本文件 (*.tar.gz)|所有文件 (*.*)|*.*"
	mw.Edit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.Edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		mw.Edit.AppendText("Cancel\r\n")
		return
	}

	mw.MouseDown()
}

// Do 升级按纽
func Do(attr *Attributes) {
	fmt.Print(attr.Port)
	fmt.Print(attr.Host)
	fmt.Print(attr.User)
	fmt.Print(attr.Password)
	//NetworkTest()
	//ConnectSsh()
	//Readfile()
	//ConnectFtp()
	//Upload()
	//Backup()
	//Upgrade()
	//Rollback()
}
