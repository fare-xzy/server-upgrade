package util

import (
	"errors"
)

const (
	CONNECTED = "获取SSH连接："
	READFILE  = "读取升级文件："
	UPLOAD    = "升级包上传："
	BACKUP    = "执行备份："
	UPGRADE   = "执行升级："
	ROLLBACK  = "执行回滚："
)

const (
	ALL_MODE   = 0777
	RWX_U_MODE = 0755
)

var (
	ERR_EOF            = errors.New("EOF")
	ERR_CLOSED_PIPE    = errors.New("io: read/write on closed pipe")
	ERR_NO_PROGRESS    = errors.New("multiple Read calls return no data or error")
	ERR_SHORT_BUFFER   = errors.New("short buffer")
	ERR_SHORT_WRITE    = errors.New("short write")
	ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")
)
