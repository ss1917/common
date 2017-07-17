package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"time"
)

//var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
var emailPattern = regexp.MustCompile(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`)

//md5加密
func GenMD5(pwd string) string {
	md5Password := md5.New()
	io.WriteString(md5Password, pwd)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	return fmt.Sprintf("%s", buffer.String())
}

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

func Isphone(p []byte) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, p); !m {
		return false
	} else {
		return true
	}
}
