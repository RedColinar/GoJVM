package references

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

type INVOKE_SPECIAL struct{ base.Index16Instruction }
//调用初始化方法，特殊处理父类的，私有的，和实例初始方法调用
func  (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame){
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	//方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	//拿到 解析后的 类 和方法
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	//如果解析的方法是构造函数，则声明的类必须是构造的类
	if resolvedMethod.Name() ==  "<init>" && resolvedMethod.Class() !=  resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil{
		panic("java.lang.NullPointerException")
	}
	//确保protected方法只能被声明的该方法的类或子类调用。
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass){
			panic("java.lang.IllegalAccessError")
		}
	//如果调用的是超类中的函数，但不是 构造函数，且当前类的ACC_SUPER标志被设置，
	//需要一个额外的 过程查找最终要调用的方法，否则前面从方法符号 引用解析出来的方法，就是要调用对的 方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>"{

			methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
				methodRef.Name(), methodRef.Descriptor())
		}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(farme, methodToBeInvoked)
}

