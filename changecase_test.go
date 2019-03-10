package changecase_test

import (
	"testing"

	"github.com/josa42/go-changecase"
	"github.com/stretchr/testify/assert"
)

// should convert to Lower case
func Test_ToLower(t *testing.T) {
	assert.Equal(t, "test", changecase.ToLower("TEST"))
	assert.Equal(t, "test", changecase.ToLower("test"))
}

// should convert to upper case
func Test_ToUpper(t *testing.T) {
	assert.Equal(t, "TEST", changecase.ToUpper("TEST"))
	assert.Equal(t, "TEST", changecase.ToUpper("test"))
}

// should upper case the first character
func Test_ToUpperFirst(t *testing.T) {
	assert.Equal(t, "TEST", changecase.ToUpperFirst("TEST"))
	assert.Equal(t, "Test", changecase.ToUpperFirst("test"))
}

// should Lower case the first character
func Test_ToLowerFirst(t *testing.T) {
	assert.Equal(t, "tEST", changecase.ToLowerFirst("TEST"))
	assert.Equal(t, "test", changecase.ToLowerFirst("test"))
}

// should determine whether a string is all upper case or not
func Test_IsUpper(t *testing.T) {
	assert.Equal(t, true, changecase.IsUpper("ALLUPPERCASE"))
	assert.Equal(t, false, changecase.IsUpper("NotAllUpperCase"))
	assert.Equal(t, false, changecase.IsUpper("alllowercase"))
}

// should determine whether a string is all lower case or not
func Test_IsLower(t *testing.T) {
	assert.Equal(t, false, changecase.IsLower("ALLLowerCASE"))
	assert.Equal(t, false, changecase.IsLower("NotAllLowerCase"))
	assert.Equal(t, true, changecase.IsLower("alllowercase"))
}

// should convert to title case
func Test_ToTitle(t *testing.T) {
	assert.Equal(t, "A Test Sentence", changecase.ToTitle("a test sentence"))
	assert.Equal(t, "I Found A Bug", changecase.ToTitle("i found a bug"))
	assert.Equal(t, "Some Things", changecase.ToTitle("some things"))
	assert.Equal(t, "Quotes", changecase.ToTitle(`"quotes"`))
	assert.Equal(t, "Hyphen Ness", changecase.ToTitle("hyphen-ness"))
}

// should convert to no case
func Test_ToNo(t *testing.T) {
	assert.Equal(t, "a simple test", changecase.ToNo("a-simple-test"))
	assert.Equal(t, "this is a test", changecase.ToNo("this is a test"))
	assert.Equal(t, "this is a test", changecase.ToNo("this_is_a_test"))
	assert.Equal(t, "this is a test", changecase.ToNo("this-is-a-test"))
	assert.Equal(t, "test api", changecase.ToNo("testAPI"))
	assert.Equal(t, "test api", changecase.ToNo("Test API"))
	assert.Equal(t, "some things", changecase.ToNo("some (things)"))

}

// should convert to header case
func Test_ToHeader(t *testing.T) {
	assert.Equal(t, "A-Simple-Test", changecase.ToHeader("a-simple-test"))
	assert.Equal(t, "This-Is-A-Test", changecase.ToHeader("this is a test"))
	assert.Equal(t, "This-Is-A-Test", changecase.ToHeader("this_is_a_test"))
	assert.Equal(t, "This-Is-A-Test", changecase.ToHeader("this-is-a-test"))
}

// should convert to sentence case
func Test_ToSentence(t *testing.T) {
	assert.Equal(t, "A simple test", changecase.ToSentence("a-simple-test"))
	assert.Equal(t, "This is a test", changecase.ToSentence("this is a test"))
	assert.Equal(t, "This is a test", changecase.ToSentence("this_is_a_test"))
	assert.Equal(t, "This is a test", changecase.ToSentence("this-is-a-test"))
}

// should convert to camel case
func Test_ToCamel(t *testing.T) {
	assert.Equal(t, "testString", changecase.ToCamel("TestString"))
	assert.Equal(t, "testString", changecase.ToCamel("Test String"))
	assert.Equal(t, "testString", changecase.ToCamel("Test_String"))
	assert.Equal(t, "testString", changecase.ToCamel("Test-String"))
	assert.Equal(t, "testApi", changecase.ToCamel("Test API"))
	assert.Equal(t, "webkitTransform", changecase.ToCamel("-webkit-transform"))
	assert.Equal(t, "fooBarBaz", changecase.ToCamel("fooBarBaz"))
	assert.Equal(t, "someThings", changecase.ToCamel("some (things)"))
}

// should convert to pascal case
func Test_ToPascal(t *testing.T) {
	assert.Equal(t, "TestString", changecase.ToPascal("testString"))
	assert.Equal(t, "TestString", changecase.ToPascal("Test String"))
	assert.Equal(t, "TestString", changecase.ToPascal("Test_String"))
	assert.Equal(t, "TestString", changecase.ToPascal("Test-String"))
	assert.Equal(t, "FacebookApi", changecase.ToPascal("Facebook API"))
	assert.Equal(t, "ATestAgain", changecase.ToPascal("a-test-again"))
	assert.Equal(t, "ABetterTest", changecase.ToPascal("a---better__test"))
}

// should convert to snake case
func Test_ToSnake(t *testing.T) {
	assert.Equal(t, "test_string", changecase.ToSnake("testString"))
	assert.Equal(t, "test_string", changecase.ToSnake("Test String"))
	assert.Equal(t, "test_string", changecase.ToSnake("Test_String"))
	assert.Equal(t, "test_string", changecase.ToSnake("Test-String"))
	assert.Equal(t, "a_better_test", changecase.ToSnake("a---better__test"))
}

// should convert to param case
func Test_ToParam(t *testing.T) {
	assert.Equal(t, "test-string", changecase.ToParam("testString"))
	assert.Equal(t, "test-string", changecase.ToParam("Test String"))
	assert.Equal(t, "test-string", changecase.ToParam("Test_String"))
	assert.Equal(t, "test-string", changecase.ToParam("Test-String"))
	assert.Equal(t, "a-better-test", changecase.ToParam("a---better__test"))
}

// should convert to constant case
func Test_ToConstant(t *testing.T) {
	assert.Equal(t, "TEST_STRING", changecase.ToConstant("testString"))
	assert.Equal(t, "TEST_STRING", changecase.ToConstant("Test String"))
	assert.Equal(t, "TEST_STRING", changecase.ToConstant("Test_String"))
	assert.Equal(t, "TEST_STRING", changecase.ToConstant("Test-String"))
	assert.Equal(t, "A_BETTER_TEST", changecase.ToConstant("a---better__test"))
}

// should convert to dot case
func Test_ToDot(t *testing.T) {
	assert.Equal(t, "test.string", changecase.ToDot("testString"))
	assert.Equal(t, "test.string", changecase.ToDot("Test String"))
	assert.Equal(t, "test.string", changecase.ToDot("Test_String"))
	assert.Equal(t, "test.string", changecase.ToDot("Test-String"))
	assert.Equal(t, "a.better.test", changecase.ToDot("a---better__test"))
}

// should convert to path case
func Test_ToPath(t *testing.T) {
	assert.Equal(t, "test/string", changecase.ToPath("testString"))
	assert.Equal(t, "test/string", changecase.ToPath("Test String"))
	assert.Equal(t, "test/string", changecase.ToPath("Test_String"))
	assert.Equal(t, "test/string", changecase.ToPath("Test-String"))
	assert.Equal(t, "a/better/test", changecase.ToPath("a---better__test"))
}

// should swap the cases
func Test_Swap(t *testing.T) {
	assert.Equal(t, "rAnDomCASe", changecase.Swap("RaNdOMcasE"))
	assert.Equal(t, "Mix iT.dOWN?", changecase.Swap("mIX It.Down?"))
}
