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
		Size:     Size{Width: 800, Height: 600},
		Layout:   VBox{}, //样式，纵向
		Children: []Widget{ //控件组
			TextEdit{
				MaxLength: int(^uint(0) >> 1),
				AssignTo:  &mw.Edit,
				ReadOnly:  true,
				VScroll:   true,
			},
			//PushButton{
			//	Text:      "打开",
			//	OnClicked: mw.selectFile, //点击事件响应函数
			//},
			//PushButton{
			//	Text:      "另存为",
			//	OnClicked: mw.saveFile,
			//},
		},
	}.Create() //创建

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	mw.Run() //运行
}
