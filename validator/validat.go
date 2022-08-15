package validator

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func RegMatcher(target string, pattern string) (bool, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return false, errors.New(fmt.Sprintf("invalid regex: %s", pattern))
	}
	return reg.MatchString(target), nil
}

func IsAllAlph(s string) bool {

	res, _ := RegMatcher(s, `^[a-zA-Z]+$`)
	return res
}

func IsAllUpper(s string) bool {
	res, _ := RegMatcher(s, `^[A-Z]+$`)
	return res
}

func IsAllLower(s string) bool {
	res, _ := RegMatcher(s, `^[a-z]+$`)
	return res
}

func IsAllNum(s string) bool {
	res, _ := RegMatcher(s, `^[0-9]+$`)
	return res
}

func ContainLetter(s string) bool {
	res, _ := RegMatcher(s, `[a-zA-Z]+`)
	return res
}

func ContainUpper(s string) bool {
	res, _ := RegMatcher(s, `([A-Z]+)`)
	return res
}

func ContainLower(s string) bool {
	res, _ := RegMatcher(s, `([a-z]+)`)
	return res
}

func ContainNumber(s string) bool {
	res, _ := RegMatcher(s, `([0-9]+)`)
	return res
}

func IsIntegerStr(s string) bool {
	res, _ := RegMatcher(s, `^([0-9]+)$`)
	return res
}

func IsFloatStr(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsIp(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

func IsIpV4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return strings.Contains(s, ".")
}

func IsIpV6(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	return strings.Contains(s, ":")
}

func IsPort(portStr string) bool {
	if port, err := strconv.Atoi(portStr); err == nil && port > 0 && port < 65536 {
		return true
	}
	return false
}

func IsAbousoluteUrl(urlStr string) bool {
	if _, err := url.Parse(urlStr); err == nil {
		return true
	}
	return false
}

func IsEmail(emailStr string) bool {
	emailReg := `^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`
	res, _ := RegMatcher(emailStr, emailReg)
	return res
}

func IsChinesePhone(phoneStr string) bool {
	phoneReg := `^1[3|4|5|7|8][0-9]{9}$`
	res, _ := RegMatcher(phoneStr, phoneReg)
	return res
}

func ContainChinese(s string) bool {
	reg := "[\u4e00-\u9fa5]+"
	res, _ := RegMatcher(s, reg)
	return res
}
