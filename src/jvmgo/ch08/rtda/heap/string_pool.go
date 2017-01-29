package heap

import "unicode/utf16"
//字符串池
var internedStrings = map[string]*Object{}
//如果字符串池中有对应的字符串实例，返回，否则，创建一个
func JString(loader *ClassLoader, goStr string) *Object{
  if internedStr, ok := internedStrings[goStr]; ok{
    return internedStr
  }
  chars := stringToUtf16(gostr)
  jChars := &Object{loader.LoadClass("[C"), chars}
  
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16{
	runes := []rune(s) //utf32
	return utf16.Encode(runes)
}
