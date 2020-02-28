package ads

import (
	"crypto/tls"
	"log"

	"github.com/0x1un/env"
	"github.com/joho/godotenv"
	"gopkg.in/ldap.v3"
)

type Client struct {
	Conn   *ldap.Conn
	BaseDN string
	Err    error
}

type LDAPInfo struct {
	BindUser   string `env:"BINDUSER"`
	BindPwd    string `env:"BINDPWD"`
	BaseDn     string `env:"BASEDN"`
	ServerHost string `env:"SERVER_HOST"`
}

func (l *LDAPInfo) ReadInfo() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(l); err != nil {
		log.Fatal(err)
	}
}

// func NewClient(binduser, bindpasswd, server, basedn string) *Client {
func NewClient(ldapinfo *LDAPInfo) *Client {
	log.Println("connect to ldap server...")
	conn, err := ldap.DialTLS("tcp", ldapinfo.ServerHost, &tls.Config{
		InsecureSkipVerify: true,
	})
	cli := &Client{}
	if err != nil {
		cli.Err = err
		return cli
	}
	err = conn.Bind(ldapinfo.BindUser, ldapinfo.BindPwd)
	if err != nil {
		cli.Err = err
		return cli
	}
	cli.Conn = conn
	cli.BaseDN = ldapinfo.BaseDn
	return cli
}
