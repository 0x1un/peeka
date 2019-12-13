package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/text/encoding/unicode"
	"gopkg.in/ldap.v3"
)

const (
	FILENAME = "users.csv"
)

type UserProfile struct {
	Username           string
	Cn                 string
	Org                string
	Description        []string
	UnicodePwd         []string
	ObjectClass        []string
	UserAccountControl []string // 514 activate
	DisplayName        []string
	SAMAccountName     []string
}

func main() {
	ldapinfo := LDAPInfo{}
	ldapinfo.ReadInfo()
	conn := NewClient(&ldapinfo)
	defer conn.Conn.Close()
	CmdExecute(conn)
}

func (c *Client) DeleteUserFromCli(cmd *Flag) error {
	if err := c.DeleteSingleUser(cmd.RealName, cmd.Org); err != nil {
		return err
	}
	return nil
}
func (c *Client) CreateUserFromCli(cmd *Flag) error {
	err := c.AddSingleUser(UserProfile{
		Username:       cmd.UserName,
		Cn:             cmd.RealName,
		Org:            cmd.Org,
		SAMAccountName: []string{cmd.UserName},
		UserAccountControl: func() []string {
			if cmd.Disable {
				return []string{"514"}
			}
			return []string{"512"}
		}(),
		ObjectClass: []string{"top", "person", "organizationalPerson", "user"},
		UnicodePwd:  []string{EncodePwd(cmd.Password)},
		Description: []string{cmd.Description},
		DisplayName: []string{cmd.RealName},
	})
	if err != nil {
		return err
	}
	return nil
}

func FileUserToProfile(filename string, uprofiles *[]UserProfile) error {
	users, err := ReadUserFromCsvToStruct(filename)
	if err != nil {
		return err
	}
	var up UserProfile
	for _, user := range users {
		up.Cn = user.Name
		up.Username = user.Username
		up.Org = user.Org
		up.UnicodePwd = []string{EncodePwd(user.Password)}
		up.ObjectClass = []string{"top", "person", "organizationalPerson", "user"}
		up.UserAccountControl = []string{"512"}
		up.SAMAccountName = []string{user.Username}
		up.DisplayName = []string{user.Name}
		up.Description = []string{"批次: " + user.BatchNum}
		*uprofiles = append(*uprofiles, up)
	}
	return nil
}

func (c *Client) AddUsersFromFile(filename string) error {
	var uprofiles []UserProfile
	err := FileUserToProfile(filename, &uprofiles)
	if err != nil {
		return err
	}

	for _, user := range uprofiles {
		if err := c.AddSingleUser(user); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (c *Client) AddSingleUser(uprofile UserProfile) error {
	if err := c.Conn.Add(AddAttribute(uprofile)); err != nil {
		if ldap.IsErrorWithCode(err, 68) {
			log.Printf("用户\033[31m\033[01m\033[05m %s \033[0m已经存在\n", uprofile.Cn)
		} else {
			return errors.New(fmt.Sprintf("添加用户失败: %s", err))
		}
	} else {
		log.Printf("%s 添加成功\n", uprofile.Cn)
	}
	return nil
}

func EncodePwd(pwd string) string {
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM) // 使用小端编码
	pwdEncoded, err := utf16.NewEncoder().String("\"" + pwd + "\"")
	if err != nil {
		return ""
	}
	return pwdEncoded
}

func AddAttribute(profile UserProfile) *ldap.AddRequest {
	newreq := fmt.Sprintf("cn=%s,ou=%s,dc=contoso,dc=com", profile.Cn, profile.Org)
	sqlInsert := ldap.NewAddRequest(newreq, nil)
	sqlInsert.Attribute("objectClass", profile.ObjectClass)
	sqlInsert.Attribute("cn", []string{profile.Cn})
	sqlInsert.Attribute("userAccountControl", profile.UserAccountControl)
	sqlInsert.Attribute("displayName", profile.DisplayName)
	sqlInsert.Attribute("unicodePwd", profile.UnicodePwd)
	sqlInsert.Attribute("sAMAccountName", profile.SAMAccountName)
	sqlInsert.Attribute("description", profile.Description)
	sqlInsert.Attribute("pwdLastSet", []string{"0"})
	return sqlInsert
}

func (c *Client) DeleteSingleUser(username, ou string) error {
	userdn := fmt.Sprintf("cn=%s,ou=%s,%s", username, ou, c.BaseDN)
	delReq := ldap.NewDelRequest(userdn, nil)
	err := c.Conn.Del(delReq)
	if err != nil {
		return err
	}
	log.Printf("所属 %s 组的用户\033[31m\033[01m\033[05m %s \033[0m已删除\n", ou, username)
	return nil
}

// 删除一个文件中的所有用户，在ad上
func (c *Client) DeleteUserFromFile(filename string) error {
	users, err := ReadUserFromCsvToStruct(filename)
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := c.DeleteSingleUser(user.Name, user.Org); err != nil {
			if ldap.IsErrorWithCode(err, 32) {
				continue
			} else {
				return err
			}
		}
	}
	return nil
}
