package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var usageDoc = `
识别各种解压缩工具将目标文件解药到目录.
前提是你要已经安装各种解压缩工具.
可识别zip rar 7z gz,目前不支持密码解压
`

func usage() {
	fmt.Println(usageDoc)
}

func cmdHelp(args []string) {
	if len(args) == 0 {
		usage()
		return
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || len(args) > 1 {
		usage()
		os.Exit(2)
		return
	}
	if args[0] == "help" {
		cmdHelp(args[1:])
		return
	}

	toUN(args[0])

}

//toUN 解压缩
func toUN(unName string) {
	unExt := path.Ext(unName)
	switch unExt {
	case ".zip":
		unZIP(unName)
	case ".rar":
		unRAR(unName)
	case ".gz":
		unGZ(unName)
	case ".7z":
		un7Z(unName)

	}

}

//unZIP zip 解压缩
func unZIP(unName string) {
	cmd := exec.Command("unzip", unName, "-d", strings.TrimSuffix(filepath.Base(unName), path.Ext(unName)))
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

//unRAR rar 解压
func unRAR(unName string) {
	//创建解压文件夹
	cmd := exec.Command("mkdir", strings.TrimSuffix(filepath.Base(unName), path.Ext(unName)))
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	//移动要解压的文件
	cmd = exec.Command("mv", unName, strings.TrimSuffix(filepath.Base(unName), path.Ext(unName)))
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	//解压
	cmd = exec.Command("unrar", "x", strings.TrimSuffix(filepath.Base(unName), path.Ext(unName))+"/"+unName, strings.TrimSuffix(filepath.Base(unName), path.Ext(unName)))
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	//移动回来
	cmd = exec.Command("mv", strings.TrimSuffix(filepath.Base(unName), path.Ext(unName))+"/"+unName, "./")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

//unGZ 解压GZ
//tar zxvf xxx.tar.gz
func unGZ(unName string) {
	cmd := exec.Command("tar", "zxvf", unName)
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

//un7Z 解压un7z
func un7Z(unName string) {
	cmd := exec.Command("7z", "x", unName)
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
