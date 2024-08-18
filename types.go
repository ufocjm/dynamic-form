package dynamic_form

type (
	Form struct {
		FormCode  string  `json:"formCode" yaml:"formCode"`   // 表单编码
		FormTitle string  `json:"formTitle" yaml:"formTitle"` // 表单名称
		Fields    []Field `json:"fields" yaml:"fields"`       // 字段列表
	}

	Field struct {
		Id            string       `json:"id" yaml:"id"`                       // 唯一标识
		Type          string       `json:"type" yaml:"type"`                   // 类型
		Text          string       `json:"text" yaml:"text"`                   // 展示值
		ValueType     string       `json:"valueType" yaml:"valueType"`         // 值类型 string float int bool
		Placeholder   string       `json:"placeholder" yaml:"placeholder"`     // Placeholder
		DefaultValue  string       `json:"defaultValue" yaml:"defaultValue"`   // 默认值
		Required      bool         `json:"required" yaml:"required"`           // 是否必填
		Validations   []Validation `json:"validations" yaml:"validations"`     // 验证规则
		OptionsSource string       `json:"optionsSource" yaml:"optionsSource"` // 选项里的字典值
		Col           int          `json:"col" yaml:"col"`                     // 每行几列
		Disabled      bool         `json:"disabled" yaml:"disabled"`           // 是否不可编辑
	}

	Validation struct {
		Pattern string `json:"pattern" yaml:"pattern"` // 正则
		Message string `json:"message" yaml:"message"` // 不匹配正则的提示
	}

	// 用户输入的
	InputForm struct {
		FormCode string       `json:"formCode"` // 表单编码
		Fields   []InputField `json:"fields" `  // 字段列表
	}

	// 用户输入的字段
	InputField struct {
		Id    string `json:"id" `    // 唯一标识
		Value string `json:"value" ` // 值
	}
)
