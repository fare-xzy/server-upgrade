package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os"
)

func main() {
	var db *walk.DataBinder
	attr := new(Attributes)
	mw := &MyMainWindow{}
	err := MainWindow{
		AssignTo: &mw.MainWindow,   //窗口重定向至mw，重定向后可由重定向变量控制控件
		Title:    "Server-Upgrade", //标题
		MinSize:  Size{Width: 400, Height: 300},
		Size:     Size{Width: 800, Height: 300},
		Layout:   HBox{}, //样式，纵向
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "attr",
			DataSource:     attr,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{ //控件组
			// 组容器
			GroupBox{
				Layout:  VBox{},
				MaxSize: Size{Width: 410, Height: 400},
				Children: []Widget{
					GroupBox{
						Title:  "产品名称",
						Layout: HBox{},
						Children: []Widget{
							ComboBox{
								//Value:         Bind("SpeciesId", SelRequired{}),
								BindingMember: "Id",
								DisplayMember: "Name",
								Model:         knownSpecies(),
							},
						},
					},
					//GroupBox{
					//	Title:  "产品版本",
					//	Layout: VBox{},
					//	Children: []Widget{
					//		ComboBox{
					//			Value:         Bind("SpeciesId", SelRequired{}),
					//			BindingMember: "Id",
					//			DisplayMember: "Name",
					//			Model:         version(),
					//		},
					//	},
					//},
					GroupBox{
						Title:  "升级内容",
						Layout: HBox{},
						Children: []Widget{
							CheckBox{
								Text: "管理平台根链升级",
								//Checked: Bind("Domesticated"),
							},
							CheckBox{
								Text: "三级根证书升级",
								//Checked: Bind("Domesticated"),
							},
							CheckBox{
								Text: "SHA1算法升级到SHA256",
								//Checked: Bind("Domesticated"),
							},
						},
					},
					GroupBox{
						Title:  "升级服务器配置",
						Layout: VBox{},
						Children: []Widget{
							Label{
								Text: "服务器IP:",
							},
							LineEdit{
								Text: Bind("Host"),
							},
							Label{
								Text: "SSH端口:",
							},
							LineEdit{
								Text: Bind("Port"),
							},
							Label{
								Text: "SSH用户名:",
							},
							LineEdit{
								Text: Bind("User"),
							},
							Label{
								Text: "SSH密码:",
							},
							LineEdit{
								PasswordMode: true,
								Text:         Bind("Password"),
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
						Title:  "执行结果",
						Layout: VBox{},
						Children: []Widget{
							TextEdit{
								MaxLength: int(^uint(0) >> 1),
								AssignTo:  &mw.OutPut,
								ReadOnly:  true,
								VScroll:   true,
							},
						},
					},
					GroupBox{
						Title:  "操作",
						Layout: VBox{},
						Children: []Widget{
							PushButton{
								Text:      "选择对应版本升级包",
								OnClicked: mw.SelectFile, //点击事件响应函数
							},
							PushButton{
								Text: "一键升级",
								OnClicked: func() {
									if err := db.Submit(); err != nil {
										log.Print(err)
										return
									}
									Do(attr)
								},
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

func knownSpecies() []*Species {
	return []*Species{
		{1, "手写信息数字签名系统"},
		{2, "PDF签章服务器"},
	}
}

func version() []*Species {
	return []*Species{
		{1, "1.3.6"},
		{2, "1.3.7"},
	}
}

type Species struct {
	Id   int
	Name string
}
