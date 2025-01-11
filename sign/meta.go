package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Meta struct {
	QDInfoRaw string
	qdInfos   []string
	InfosRW   MetaRW

	SDKSignRaw string
	sdkSign    []string
	SdkRW      MetaRW
}

func NewMeta(QDInfoRaw string, SDKSignRaw string) (*Meta, error) {
	m := &Meta{QDInfoRaw: QDInfoRaw, SDKSignRaw: SDKSignRaw}
	if err := m.parse(); err != nil {
		return nil, err
	}
	m.SdkRW = MetaRW{
		data:       m.sdkSign,
		dataStruct: SDKSignStruct,
	}
	m.InfosRW = MetaRW{
		data:       m.qdInfos,
		dataStruct: QDInfoStruct,
	}
	return m, nil
}
func (m *Meta) String() string {
	sb := strings.Builder{}
	sb.WriteString("qdInfos:\n")
	for i, info := range m.qdInfos {
		sb.WriteString(fmt.Sprintf("%s=%s\n", QDInfoStruct[i], info))
	}
	sb.WriteString("sdkSign:\n")
	for i, info := range m.sdkSign {
		sb.WriteString(fmt.Sprintf("%s=%s\n", SDKSignStruct[i], info))
	}
	return sb.String()
}
func (m *Meta) ModifyTimeStamp() error {
	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if err := m.InfosRW.Modify(FiledTimestamp, ts); err != nil {
		return err
	}
	return m.SdkRW.Modify(FiledTimestamp, ts)
}

func (m *Meta) QDInfo() (string, error) {
	info := strings.Join(m.qdInfos, "|")
	encrypt, err := EncryptQDInfo(info)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}
func (m *Meta) SDKSign(uri string) (string, error) {
	m.SdkRW.Modify(FiledHashUrl, hash(normParams(uri)))
	sdkSign := strings.Join(m.sdkSign, "|")
	encrypt, err := EncryptSDKSign(sdkSign)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}
func normParams(urlStr string) string {
	idx := strings.Index(urlStr, "?")
	if idx == -1 {
		return ""
	}

	queryParts := urlStr[idx+1:]
	if queryParts == "" {
		return ""
	}

	parts := strings.Split(queryParts, "&")
	queries := make([]string, 0, len(parts))

	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos == -1 {
			queries = append(queries, strings.ToLower(part)+"=")
		} else {
			key := strings.ToLower(part[:pos])
			value := url.QueryEscape(part[pos+1:])
			queries = append(queries, key+"="+value)
		}
	}

	sort.Strings(queries)
	return strings.Join(queries, "&")
}
func hash(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func (m *Meta) parse() error {
	qdInfo, err := DecryptQDInfo(m.QDInfoRaw)
	if err != nil {
		return err
	}
	m.qdInfos = strings.Split(qdInfo, "|")[:len(QDInfoStruct)-5]
	sign, err := DecryptSDKSign(m.SDKSignRaw)
	if err != nil {
		return err
	}
	m.sdkSign = strings.Split(sign, "|")
	if len(m.qdInfos) < 2 {
		return fmt.Errorf("输入的qdInfos有问题")
	}
	if len(m.sdkSign) < 2 {
		return fmt.Errorf("输入的sdkSign有问题")
	}
	return nil
}
