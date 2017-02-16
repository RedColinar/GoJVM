package lang

import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

func inti(){
	native.Register("java/lang/Throwable","fillIntStackTrace",
		"(I)Ljava/lang/Throwable;",fillIntStackTrace)
}

func fillIntStackTrace(frame *rtda.Frame){
	
}