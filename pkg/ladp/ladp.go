package ladp

import (
	"errors"
	"fmt"
	"strings"
)

// ldap登录验证
func LdapLoginAuth(username string, passowrd string, cfg map[string]string) (bool, error) {
	server := cfg["server"]
	port := cfg["port"]
	prefix := cfg["prefix"]

	username = prefix + "\\" + username
	if server == "" || prefix == "" {
		return false, errors.New("ErrUnknownIssuer")
	}

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		logs.Error(err)
		return false, err
	}

	defer l.Close()
	controls := []ldap.Control{}
	controls = append(controls, ldap.NewControlBeheraPasswordPolicy())

	bindRequest := ldap.NewSimpleBindRequest(username, strings.TrimSpace(passowrd), controls)
	_, err = l.SimpleBind(bindRequest)
	if err != nil {
		return false, err
	}
	return true, err
}
