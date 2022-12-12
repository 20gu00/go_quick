package common

import (
	"fmt"
	//"go_forum/model/param"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 获取json tag,validator默认是返回对应的结构体字段
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 非必要 为RegisterParam注册自定义校验方法
		//v.RegisterStructValidation(RegisterInputStructLevelValidation, param.RegisterInput{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// 非必要

// validator默认的输出key会包含对应的结构体(前端输入的参数对应的结构体)名称,去除提示信息中的结构体名称
// 这里只是一层没有结构体嵌套
//func RemoveTopStruct(fields map[string]string) map[string]string {
//	res := map[string]string{}
//	for field, v := range fields {
//		res[field[strings.Index(field, ".")+1:]] = v
//	}
//	return res
//}

// 自定义RegisterInput结构体校验函数
//func RegisterInputStructLevelValidation(sl validator.StructLevel) {
//	su := sl.Current().Interface().(param.RegisterInput)
//	if su.Password != su.RePassword {
//		// 输出错误提示信息，最后一个参数就是传递的param
//		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
//	}
//}
