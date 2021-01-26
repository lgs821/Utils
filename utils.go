package utils

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"time"
)

const (
	regular = `^13[\d]{9}$|^14[\d]{9}$|^15[^4]{1}\d{8}$|^17[\d]{9}$|^18[\d]{9}$|^16[\d]{9}|^19[\d]{9}`
)

//IsMobile 判断是否为手机号
func IsMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	// fmt.Println(reg)
	return reg.MatchString(mobileNum)
}

//RandInt64 生成两个数字之间的随机数
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

//IsEmpty 判断是否为空值
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

// GetSettings 获取json设置
func GetSettings(path string, str string) interface{} {
	settingfile := path
	setinfo, _ := ReadFile(settingfile)
	setmap, err := JsonDecode(setinfo)
	if err != nil {
		fmt.Println(err)
	}
	setmaps := setmap.(map[string]interface{})
	return setmaps[str]
}

//FormatMobile 格式化手机号码
func FormatMobile(str string) string {
	if len(str) <= 10 {
		return str
	}
	return str[:3] + "****" + str[len(str)-4:]
}

//Decimal 保留两位小数点
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

//InArray 判断是否在数组中
func InArray(aaa []string, val string) bool {
	for _, v := range aaa {
		if v == val {
			return true
		}
	}
	return false
}
