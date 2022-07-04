package main

import (
	"errors"
)

const (
	COMPLETE       = "成功"
	ENTER          = "\r\n"
	CHOOSE_PACKAGE = "选择升级包："
	CHECKPARAMS    = "检查输入参数："
	CONNECTED      = "获取SSH连接："
	READFILE       = "读取升级文件："
	CONNECTFTP     = "获取FTP连接："
	UPLOAD         = "升级包上传："
	BACKUP         = "执行备份："
	UPGRADE        = "执行升级："
	ROLLBACK       = "执行回滚："
)

var (
	PARAMS_ERROR       = errors.New("参数异常，请重新输入！")
	UPGRADE_FILE_ERROR = errors.New("请选择升级文件！")
	ERR_EOF            = errors.New("EOF")
	ERR_CLOSED_PIPE    = errors.New("io: read/write on closed pipe")
	ERR_NO_PROGRESS    = errors.New("multiple Read calls return no data or error")
	ERR_SHORT_BUFFER   = errors.New("short buffer")
	ERR_SHORT_WRITE    = errors.New("short write")
	ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")
)
