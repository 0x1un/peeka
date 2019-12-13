package main

import (
	"errors"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

type Flag struct {
	Disable     bool
	Remove      bool
	Create      bool
	FileName    string
	RealName    string
	UserName    string
	Password    string
	Description string
	Org         string
}

func CmdExecute(client *Client) {
	f := Flag{}

	flag.ErrHelp = errors.New("")

	run := flag.NewFlagSet("run", flag.ExitOnError)
	run.StringVar(&f.FileName, "file", "example.csv", "指定文件")
	run.BoolVar(&f.Remove, "remove", false, "删除选项")
	run.BoolVar(&f.Create, "create", false, "创建选项")

	create := flag.NewFlagSet("create", flag.ExitOnError)
	create.StringVar(&f.RealName, "name", "张三", "指定创建的用户名")
	create.StringVar(&f.UserName, "uname", "wb-zs888888", "指定帐号")
	create.StringVar(&f.Password, "pwd", "abc@123", "指定密码")
	create.StringVar(&f.Org, "org", "dev", "指定该用户所在组织")
	create.StringVar(&f.Description, "bn", "0", "指定批次号")
	create.BoolVar(&f.Disable, "disable", false, "是否禁用该帐户")

	remove := flag.NewFlagSet("remove", flag.ExitOnError)
	remove.StringVar(&f.RealName, "name", "张三", "指定删除的用户名")
	remove.StringVar(&f.Org, "org", "dev", "指定删除用户的所属组织")

	if !(len(os.Args) == 1) {
		switch os.Args[1] {
		case "run":
			run.Parse(os.Args)
			if f.Create {
				if err := client.AddUsersFromFile(f.FileName); err != nil {
					log.Fatal(err)
				}
			}
			if f.Remove {
				if err := client.DeleteUserFromFile(f.FileName); err != nil {
					log.Fatal(err)
				}
			}
		case "create":
			create.Parse(os.Args)
			if err := client.CreateUserFromCli(&f); err != nil {
				log.Fatal(err)
			}
		case "remove":
			remove.Parse(os.Args)
			if err := client.DeleteUserFromCli(&f); err != nil {
				log.Fatal(err)
			}
		default:
			flag.Usage()
		}
	}
}
