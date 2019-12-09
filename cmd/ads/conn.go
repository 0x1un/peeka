package main

import (
	"crypto/tls"
	"fmt"
	"reflect"

	"gopkg.in/ldap.v3"
)

type Client struct {
	Conn   *ldap.Conn
	BaseDN string
	Err    error
}

func NewClient(binduser, bindpasswd, server, basedn string) *Client {
	conn, err := ldap.DialTLS("tcp", server, &tls.Config{
		InsecureSkipVerify: true,
	})
	cli := &Client{}
	if err != nil {
		cli.Err = err
		return cli
	}
	err = conn.Bind(binduser, bindpasswd)
	if err != nil {
		cli.Err = err
		return cli
	}
	fmt.Println(reflect.TypeOf(conn))
	cli.Conn = conn
	cli.BaseDN = basedn
	return cli
}
