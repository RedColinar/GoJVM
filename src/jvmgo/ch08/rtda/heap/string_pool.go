package heap

import "unicode/utf16"
	//这种被共享的字符串被称为“限定字符串”（interned string）。
	//当你觉得某个特定的文本会被创建并且反复在程序中使用多次
	//如果字符串池中有对应的字符串实例，返回，否则，创建一个
var internedStrings = map[string]*Object{}
//把go字符串转化成java字符串对象
func JString(loader *ClassLoader, goStr string) *Object{

	if internerStr,ok := internedStrings[goStr];ok{
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func GoString() string{
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
