package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/text/encoding/unicode"
	"gopkg.in/ldap.v3"
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
	conn := NewClient("administrator@contoso.com", "goodluck@123", "172.20.6.10:636", "dc=contoso,dc=com")
	fmt.Println(conn)
	defer conn.Conn.Close()

	// err := AddUsersFromFile("./users.csv", conn.Conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// DeleteUserFromFile("./users.csv", conn.Conn)
	conn.DeleteUserFromFile("./users.csv")
}

func AddUsersFromFile(filename string, conn *ldap.Conn) error {
	users, err := ReadUserFromCsvToStruct(filename)
	if err != nil {
		return err
	}

	var profile UserProfile
	for _, user := range users {
		profile = UserProfile{
			Cn:                 user.Name,
			Username:           user.Username,
			UnicodePwd:         []string{EncodePwd(user.Password)},
			ObjectClass:        []string{"top", "person", "organizationalPerson", "user"},
			UserAccountControl: []string{"512"}, // 激活状态，514为禁用状态
			DisplayName:        []string{user.Name},
			SAMAccountName:     []string{user.Username}, // 登录账户名
			Description:        []string{"批次: " + user.BatchNum},
			Org:                user.Org,
		}
		if err := conn.Add(AddAttribute(profile)); err != nil {
			if ldap.IsErrorWithCode(err, 68) {
				return errors.New("User is already exist")
			} else {
				return errors.New(fmt.Sprintf("User insert err: %s", err))
			}
		}
		log.Printf("Added user: %s\t初始密码: %s\n", profile.Cn, profile.UnicodePwd[0])
	}

	return nil
}

func (c *Client) AddSingleUser(uprofile UserProfile) error {
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
		log.Println(err)
		return err
	}
	log.Printf("Deleted user: %s of %s\n", username, ou)
	return nil
}

// 删除一个文件中的所有用户，在ad上
func (c *Client) DeleteUserFromFile(filename string) {
	users, err := ReadUserFromCsvToStruct(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		if err := c.DeleteSingleUser(user.Name, user.Org); err != nil {
			log.Fatal(err)
		}
		log.Printf("Deleted user: %s of %s\n", user.Name, user.Org)
	}
}
