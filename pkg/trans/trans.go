package trans

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT)
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("GetTranslator(%s) failed", locale)
		}

		// 注册国际化翻译器
		switch locale {
		case "en":
			err = enTrans.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTrans.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTrans.RegisterDefaultTranslations(v, Trans)
		}
		if err != nil {
			return
		}

		// 注册自定义国际化翻译
		// if err := v.RegisterTranslation(
		// 	"checkDate",
		// 	Trans,
		// 	registerTranslator("checkDate", "{0}必须要晚于当前日期"),
		// 	translate,
		// ); err != nil {
		// 	return err
		// }
		return
	}
	return
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
