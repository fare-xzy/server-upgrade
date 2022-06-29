package main

import (
	"github.com/lxn/walk"
	"server-upgrade/util"
	"strings"
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

// CheckParams 参数校验
func CheckParams(attr *Attributes) error {
	if strings.EqualFold("", attr.Host) || strings.EqualFold("", attr.Port) || strings.EqualFold("", attr.User) || strings.EqualFold("", attr.Password) {
		return util.PARAMS_ERROR
	}
	return nil
}

// Do 升级按纽
func Do(attr *Attributes) {
	CheckParams(attr)
	//NetworkTest()
	//ConnectSsh()
	//Readfile()
	//ConnectFtp()
	//Upload()
	//Backup()
	//Upgrade()
	//Rollback()
}
