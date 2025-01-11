package qdapi

import (
	"fmt"
	"strings"
)

const (
	Cookie_ywkey  = "ywkey"
	Cookie_ywguid = "ywguid"
	Cookie_QDInfo = "QDInfo"
)

type Cookie struct {
	Name  string
	Value string
}
type Cookies []*Cookie

func NewCookies(ywKey, ywGuid, qDInfo string) Cookies {
	return Cookies{
		&Cookie{Cookie_ywkey, ywKey},
		&Cookie{Cookie_ywguid, ywGuid},
		&Cookie{Cookie_QDInfo, qDInfo},
	}
}
func (c Cookies) String() string {
	sb := strings.Builder{}
	for _, cookie := range c {
		sb.WriteString(fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value))
	}
	return sb.String()
}
