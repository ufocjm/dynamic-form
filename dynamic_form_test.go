package dynamic_form

import (
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	//assert.NoError(t, err)

	dynamicForm := NewDynamicForm()
	dynamicForm.ReadJsonFile("../test.json")

}

func TestReadYamlFile(t *testing.T) {
	//assert.NoError(t, err)

	dynamicForm := NewDynamicForm()
	dynamicForm.ReadYamlFile("../test.yaml")
}
