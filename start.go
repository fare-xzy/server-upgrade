package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"os"
)

type MyMainWindow struct {
	*walk.MainWindow
	Edit *walk.TextEdit
}

func main() {
	mw := &MyMainWindow{}
	err := MainWindow{
		AssignTo: &mw.MainWindow,   //窗口重定向至mw，重定向后可由重定向变量控制控件
		Title:    "server-upgrade", //标题
		MinSize:  Size{Width: 400, Height: 300},
		Size:     Size{Width: 800, Height: 300},
		Layout:   HBox{}, //样式，纵向
		Children: []Widget{ //控件组
			// 组容器
			GroupBox{
				Layout:  VBox{},
				MaxSize: Size{Width: 400, Height: 400},
				Children: []Widget{
					GroupBox{
						Title:  "升级服务器配置",
						Layout: VBox{},
						Children: []Widget{
							Label{
								Text: "服务器IP:",
							},
							LineEdit{
								CueBanner: "127.0.0.1",
							},
							Label{
								Text: "SSH端口:",
							},
							LineEdit{
								CueBanner: "22",
							},
							Label{
								Text: "SSH用户名:",
							},
							LineEdit{
								CueBanner: "root",
							},
							Label{
								Text: "SSH密码:",
							},
							LineEdit{
								PasswordMode: true,
							},
							GroupBox{
								Title:  "升级内容",
								Layout: HBox{},
								Children: []Widget{
									CheckBox{
										Text:    "管理平台根链升级",
										Checked: Bind("Domesticated"),
									},
									CheckBox{
										Text:    "三级根证书升级",
										Checked: Bind("Domesticated"),
									},
									CheckBox{
										Text:    "SHA1算法升级到SHA256",
										Checked: Bind("Domesticated"),
									},
								},
							},
						},
					},
				},
			},
			GroupBox{
				Layout: VBox{},
				Children: []Widget{
					GroupBox{
						Title:  "工具说明",
						Layout: VBox{},
						Children: []Widget{
							TextEdit{
								MaxLength: int(^uint(0) >> 1),
								AssignTo:  &mw.Edit,
								ReadOnly:  true,
							},
						},
					},
					GroupBox{
						Title:  "操作",
						Layout: VBox{},
						Children: []Widget{
							PushButton{
								Text:      "打开文件",
								OnClicked: mw.selectFile, //点击事件响应函数
							},
							PushButton{
								Text: "一键升级",
								//OnClicked: mw.saveFile,
							},
						},
					},
				},
			},
		},
	}.Create() //创建

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	mw.Run() //运行
}

func (mw *MyMainWindow) selectFile() {

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
