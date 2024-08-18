package dynamic_form

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"regexp"
)

type DynamicForm struct {
	FormMap map[string]Form
}

func NewDynamicForm() *DynamicForm {
	return &DynamicForm{
		FormMap: make(map[string]Form),
	}
}

func (df *DynamicForm) PutForm(form Form) {
	_, exist := df.FormMap[form.FormCode]
	if exist {
		panic(errors.New("form code already exists"))
	}
	df.FormMap[form.FormCode] = form
}

// 读取json文件
func (df *DynamicForm) ReadJsonFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	var forms []Form
	err = json.Unmarshal(data, &forms)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	for _, form := range forms {
		df.PutForm(form)
	}
}

// 读取yaml文件
func (df *DynamicForm) ReadYamlFile(path string) {
	// 1. 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 2. 解析 YAML 数据
	var forms []Form
	err = yaml.Unmarshal(data, &forms)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
	for _, form := range forms {
		df.PutForm(form)
	}
}

// 校验
func (df *DynamicForm) Validated(inputForm InputForm) error {
	form, exist := df.FormMap[inputForm.FormCode]
	if !exist {
		return errors.New(fmt.Sprintf("%s表单不存在", inputForm.FormCode))
	}

	inputFieldMap := make(map[string]InputField)
	for _, inputField := range inputForm.Fields {
		inputFieldMap[inputField.Id] = inputField
	}

	for _, field := range form.Fields {
		inputField, exist := inputFieldMap[field.Id]

		// 不必填 并且每天 直接进行下一个
		if !field.Required && !exist && len(inputField.Value) == 0 {
			continue
		}
		if field.Required && (!exist || len(inputField.Value) == 0) {
			return errors.New(fmt.Sprintf("%s不能为空", field.Text))
		}

		validations := field.Validations
		// 没有验证规则 就不必要校验了
		if len(validations) == 0 {
			continue
		}

		// 到这里肯定填写过或者必填的
		for _, validation := range validations {
			matched, err := regexp.MatchString(validation.Pattern, inputField.Value)
			if err != nil {
				return err
			}
			// 不匹配
			if !matched {
				return errors.New(validation.Message)
			}
		}
	}
	return nil
}
