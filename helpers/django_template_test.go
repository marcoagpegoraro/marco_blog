package helpers_test

import (
	"testing"

	"github.com/marcoagpegoraro/marco_blog/helpers"
)

func TestCamelCaseToCapitalizeFirstWorldAndAddSpaces(t *testing.T) {

	str := "thisIsACamelCaseExample"

	strConverted := helpers.CamelCaseToCapitalizeFirstWorldAndAddSpaces(str)

	strExpected := "This is a camel case example"

	if strConverted != strExpected {
		t.Errorf("Error converting string camel case to normal text")
	}
}
