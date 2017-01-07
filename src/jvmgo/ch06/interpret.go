package main

import "fmt"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
//解释器
func interpret(methodInfo *classfile.MemberInfo){
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread,bytecode)
}

func catchErr(frame *rtda.Frame){
	if r := recover(); r != nil{
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}
//loop()函数循环执行“计算pc，解码指令  ，执行指令”
func  loop(thread *rtda.Thread, bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc :=  frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(bytecode,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		fmt.Printf("pc:%2d inst:%T %v\n",pc,inst,inst)
		inst.Execute(frame)
	}
}