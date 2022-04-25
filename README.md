# server-upgrade

### 快速开始

- 开发环境
  * go 1.18

- 配置go代理
  
  ```
  控制台输入  go env 查看当前GOPROXY内容若为 GOPROXY=https://goproxy.cn,direct 可以不用进行修改
  使用go env -w GOPROXY=https://goproxy.cn,direct 进行代理变更
  也可以配置成其他的代理地址
  ```

- 包下载
  
  ```
  执行
  go get -u github.com/lxn/walk
  go get github.com/akavel/rsrc
  go get -u golang.org/x/crypto/ssh
  go get -u github.com/pkg/sftp
  ```
* 准备工作
  
  * 安装    tdm64-gcc-10.3.0-2.exe 并设置环境变量
  
  * 找到GOPATH中github.com/akavel/rsrc包 进入目录下执行go build 
  
  * 将执行结果.exe文件放置于GOROOT bin目录下（如果输出结果不是exe请按照下一步进行变更）

### 功能说明


### 打包

* 设置打包环境
  
  * go env -w GOOS=linux
  * go env -w GOOS=windows

* 打包
  
  * rsrc -manifest main.manifest -o main.syso 
  
  * go build -ldflags="-H windowsgui -w -s"

### 使用



### 异常
- 如遇被360杀毒软件拦截，请忽略拦截信任该软件