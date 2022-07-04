package main

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
	Host        string
	Port        string
	User        string
	Password    string
	PackagePath string
}

func NetworkTest(attr Attributes) error {
	joinHostPort := net.JoinHostPort(attr.Host, attr.Port)
	_, err := net.DialTimeout("tcp", joinHostPort, 3*time.Second)
	return err
}

func ConnectSsh(attr *Attributes) (*ssh.Client, error) {
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

func Readfile(attr *Attributes) ([]byte, error) {
	return util.ReadFileOnce(attr.PackagePath)
}
func ConnectFtp(sshClient *ssh.Client) (*sftp.Client, error) {
	client, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}
func Upload(client *sftp.Client, gz []byte) (string, error) {
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

func Unzip(sshClient *ssh.Client, currentTime string) error {
	// 解压文件
	combo, err := execute(sshClient, fmt.Sprintf("tar -zxPvf ~/%s/%s -C ./%s/", currentTime, currentTime+".tar.gz", currentTime))
	if err != nil {
		return err
	}
	log.Println("解压命令输出:", string(combo))
	return nil
}

func Backup(sshClient *ssh.Client, currentTime string) error {
	// 执行备份脚本
	combo, err := execute(sshClient, fmt.Sprintf("find ~/%s/upgrade/backup.sh -type f | wc -l", currentTime))
	if strings.EqualFold(strings.TrimSpace(string(combo)), "1") {
		combo, err = execute(sshClient, fmt.Sprintf("sh ~/%s/upgrade/backup.sh", currentTime))
		log.Println("备份命令输出:", string(combo))
		if err != nil {
			return err
		}
	} else {
		log.Println("命令输出:", "程序无备份脚本，跳过备份步骤")
	}
	return nil
}

func Upgrade(sshClient *ssh.Client, currentTime string) error {
	// 执行升级脚本
	combo, err := execute(sshClient, fmt.Sprintf("sh ~/%s/upgrade/upgrade.sh %s", currentTime, "~/"+currentTime+"/upgrade"))
	if err != nil {
		return err
	}
	log.Println("升级命令输出:", string(combo))
	return err
}

func Rollback() {

}

// 远程执行并返回执行结果
func execute(sshClient *ssh.Client, cmd string) ([]byte, error) {
	session, _ := sshClient.NewSession()
	defer session.Close()
	combo, err := session.CombinedOutput(cmd)
	return combo, err
}
