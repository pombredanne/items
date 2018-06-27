package validator

import (
	"regexp"
)

var regexps = map[string]string{
	//"email":        `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,
	"email":       `^([0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*@([0-9a-zA-Z][-\w]*[0-9a-zA-Z]\.)+[a-zA-Z]{2,9})$`,
	"phone":        `^1[34578]\d{9}$`,
	"qq":           `[1-9]\d{4,12}$`,
	"link":         `(http|https)://\w+(-\w+)*(\.\w+(-\w+)*)*`,
	"telephone":    `^(1[3|4|5|8][0-9]\d{4,8})|0\d{2,3}-\d{7,8}$`,
	"company_name": `^[a-zA-Z\(\（\)\）\u4e00-\u9fa5]{1,100}$`,
	"real_name":    `^[a-zA-Z\u4e00-\u9fa5\-\.\·]{1,20}$`,
}
// 验证是否符合正则
func ValidRegexp(regexp_key string, values ...string) bool {
	for _, value := range values {
		if m, _ := regexp.MatchString(regexps[regexp_key], value); !m {
			return m
		}
	}
	return true
}
