package validator_trans

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

func SetZh(vali *validator.Validate) ut.Translator{
	zh_ch := zh.New()
	uni := ut.New(zh_ch)

	trans, _ := uni.GetTranslator("zh")
	//　验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(vali,trans)
	return trans;
}