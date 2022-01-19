package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Sign(i interface{}, secret string, ex ...string) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	var m = make(map[string]interface{})
	if err = json.Unmarshal(b, &m); err != nil {
		return "", err
	}
	return MapSign(m, secret, ex...), err
}

func MapSign(m map[string]interface{}, secret string, ex ...string) string {
	var index []string
	for k := range m {
		if contains(k, ex) {
			continue
		}
		index = append(index, k)
	}
	sort.Strings(index)
	var texts []string
	for _, k := range index {
		texts = append(texts, fmt.Sprintf("%s=%v", k, m[k]))
	}
	text := strings.Join(texts, "&")
	return CommonSign(text, secret)
}

func CommonSign(text string, secret string) string {
	b := md5.Sum([]byte(secret + text + secret))
	return hex.EncodeToString(b[:])
}

func contains(s string, array []string) bool {
	for _, v := range array {
		if v == s {
			return true
		}
	}
	return false
}
