package native

import "jvmgo/ch09/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod){
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod{
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok{
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNative"{
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame){
	//啥也不做
}