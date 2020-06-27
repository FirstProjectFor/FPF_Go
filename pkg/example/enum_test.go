package demo

import "testing"
import "github.com/stretchr/testify/assert"

var (
	ACTIVE  = NewStatusEnum("10", "启用")
	DISABLE = NewStatusEnum("20", "禁用")
	Values  = []StatueEnum{ACTIVE, DISABLE}
)

type StatueEnum struct {
	Code string
	Name string
}

func NewStatusEnum(code string, name string) StatueEnum {
	return StatueEnum{Name: name, Code: code}
}

func StatusCodeToName(code string) string {
	if "" == code {
		return ""
	}
	for _, v := range Values {
		if v.Code == code {
			return v.Name
		}
	}
	return ""
}

func TestEnum(t *testing.T) {
	assert.Equal(t, "", StatusCodeToName(""))
	assert.Equal(t, "启用", StatusCodeToName("10"))
	assert.Equal(t, "禁用", StatusCodeToName("20"))
	assert.Equal(t, "", StatusCodeToName("00"))
}
