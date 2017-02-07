package  lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

const jlClass = "java/lang/Class"

func init(){
	native.Register(jlClass, "getPrimitiveClass","(Ljava/lang/String;)Ljava/lang/Class;",getPrimitiveClass)
	native.Register(jlClass, "getName0","()Ljava/lang/String")
}