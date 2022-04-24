package connect

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"server-upgrade/util"
	"strings"
	"time"
)

type Attributes struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Specify  bool   `yaml:"specify"`
}

//var conn *ssh.Client

// 测试SSH连接
func Connect_test(attr Attributes) error {
	joinHostPort := net.JoinHostPort(attr.Host, attr.Port)
	_, err := net.DialTimeout("tcp", joinHostPort, 3*time.Second)
	return err
}

// SSH连接
func Ssh_connect(attr Attributes) (*ssh.Client, error) {
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(attr.Password))

	clientConfig := &ssh.ClientConfig{
		User:    attr.User,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr := fmt.Sprintf("%s:%s", attr.Host, attr.Port)
	sshClient, err := ssh.Dial("tcp", addr, clientConfig)
	return sshClient, err
}

// FTP连接
func Sftp_connect(sshClient *ssh.Client) (*sftp.Client, error) {

	client, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

// FTP上传文件
func Sftp_upload(client *sftp.Client, gz []byte) (string, error) {
	w := client.Walk("~/")
	//w := client.Walk("~/upgrade/")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
	}
	currentTime := time.Now().Format("20060102150405")
	client.Mkdir(currentTime)
	fileName := fmt.Sprintf("%s/%s", currentTime, currentTime+".tar.gz")
	f, err := client.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(gz); err != nil {
		log.Fatal(err)
	}
	f.Close()
	// check it's there
	fi, err := client.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)
	return currentTime, err
}

// 执行升级操作
func doExecution(sshClient *ssh.Client, currentTime string, attr Attributes) error {
	// 解压文件
	combo, err := execute(sshClient, fmt.Sprintf("tar -zxPvf ~/%s/%s -C ./%s/", currentTime, currentTime+".tar.gz", currentTime))
	if err != nil {
		return err
	}
	log.Println("解压命令输出:", string(combo))
	// 执行备份脚本
	combo, err = execute(sshClient, fmt.Sprintf("find ~/%s/upgrade/backup.sh -type f | wc -l", currentTime))
	if strings.EqualFold(strings.TrimSpace(string(combo)), "1") {
		combo, err = execute(sshClient, fmt.Sprintf("sh ~/%s/upgrade/backup.sh", currentTime))
		log.Println("备份命令输出:", string(combo))
		if err != nil {
			return err
		}
	} else {
		log.Println("命令输出:", "程序无备份脚本，跳过备份步骤")
	}
	// 执行升级脚本
	if strings.EqualFold(attr.User, util.ADMIN) && attr.Specify {
		combo, err = execute(sshClient, fmt.Sprintf("echo %s | sudo -S sh ~/%s/upgrade/upgrade.sh %s", attr.Password, currentTime, "~/"+currentTime+"/upgrade"))
	} else {
		combo, err = execute(sshClient, fmt.Sprintf("sh ~/%s/upgrade/upgrade.sh %s", currentTime, "~/"+currentTime+"/upgrade"))
	}
	if err != nil {
		return err
	}
	log.Println("升级命令输出:", string(combo))
	return err
}

// 远程执行并返回执行结果
func execute(sshClient *ssh.Client, cmd string) ([]byte, error) {
	session, _ := sshClient.NewSession()
	defer session.Close()
	combo, err := session.CombinedOutput(cmd)
	return combo, err
}

func Execution(attr Attributes, gz []byte) error {
	// 获取ssh连接
	sshConnect, err := Ssh_connect(attr)
	if err != nil {
		log.Println("服务器SSH连接失败：", err)
		panic(err)
	}
	defer sshConnect.Close()
	// 获取sftp连接
	sftpConnect, err := Sftp_connect(sshConnect)
	if err != nil {
		log.Println("服务器SFTP连接失败：", err)
		panic(err)
	}
	defer sftpConnect.Close()

	// sftp上传文件
	currentTime, err := Sftp_upload(sftpConnect, gz)
	if err != nil {
		log.Println("sftp上传文件失败：", err)
		panic(err)
	}
	// 执行
	err = doExecution(sshConnect, currentTime, attr)
	if err != nil {
		log.Println("执行升级脚本失败：", err)
		panic(err)
	}
	return err
}
