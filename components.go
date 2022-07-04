package main

import (
	"github.com/lxn/walk"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
	Edit   *walk.TextEdit
	OutPut *walk.TextEdit
}

// SelectFile 文件上传按纽
func (mw *MyMainWindow) SelectFile(attr *Attributes) {

	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "文本文件 (*.tar.gz)|*.gz|所有文件 (*.*)|*.*"

	mw.OutPut.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.OutPut.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		mw.OutPut.AppendText("Cancel\r\n")
		return
	}
	attr.PackagePath = dlg.FilePath
	mw.textEditAppend(CHOOSE_PACKAGE + dlg.FilePath + ENTER)
	mw.MouseDown()
}

// CheckParams 参数校验
func CheckParams(attr *Attributes) error {
	if strings.EqualFold("", attr.Host) || strings.EqualFold("", attr.Port) || strings.EqualFold("", attr.User) || strings.EqualFold("", attr.Password) {
		return PARAMS_ERROR
	}
	if strings.EqualFold("", attr.PackagePath) {
		return UPGRADE_FILE_ERROR
	}
	return nil
}

// Do 升级按纽
func (mw *MyMainWindow) Do(attr *Attributes) {
	// 参数检查
	mw.textEditAppend(CHECKPARAMS)
	err := CheckParams(attr)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	mw.textEditAppend(COMPLETE + ENTER + CONNECTED)
	// 获取SSH连接
	ssh, err := ConnectSsh(attr)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	// 读取升级文件
	mw.textEditAppend(COMPLETE + ENTER + READFILE)
	gzFile, err := Readfile(attr)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	// 获取FTP连接
	mw.textEditAppend(COMPLETE + ENTER + CONNECTFTP)
	ftp, err := ConnectFtp(ssh)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	// 上传升级文件
	mw.textEditAppend(COMPLETE + ENTER + UPLOAD)
	currentTime, err := Upload(ftp, gzFile)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	mw.textEditAppend(COMPLETE + ENTER + UNZIP)
	// 解压升级包
	err = Unzip(ssh, currentTime)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	// 执行备份脚本
	mw.textEditAppend(COMPLETE + ENTER + BACKUP)
	err = Backup(ssh, currentTime)
	if err != nil {
		mw.textEditAppend(err.Error())
		return
	}
	// 执行升级脚本
	mw.textEditAppend(COMPLETE + ENTER + UPGRADE)
	err = Upgrade(ssh, currentTime)
	if err != nil {
		mw.textEditAppend(err.Error())
		// 执行回滚脚本
		mw.textEditAppend(ENTER + ROLLBACK)
		err := Rollback(ssh, currentTime)
		if err != nil {
			mw.textEditAppend(err.Error())
			return
		}
		mw.textEditAppend(COMPLETE)
		return
	}
	mw.textEditAppend(COMPLETE)
}

// TextEditAppend 输出
func (mw *MyMainWindow) textEditAppend(outPut string) {
	mw.OutPut.AppendText(outPut)
	mw.MouseDown()
}
