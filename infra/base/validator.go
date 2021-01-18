package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"moyutec.top/resk/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Translator() ut.Translator {
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("Not found translator: zh")
	}
}

func ValidateStruct(s interface{}) (err error) {
	err = Validate().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(Translator()))
			}
		}
		return err
	}
	return nil
}
