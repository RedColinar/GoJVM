package main

import "fmt"
import "jvmgo/ch07/instructions"
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"
//解释器
//logInst 控制是否把指令执行信息打印到控制台
func interpret(method *heap.Method, logInst bool){
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread){
	if r := recover(); r != nil{
		logFrames(thread)
		panic(r)
	}
}
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty(){
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(),className,method.Name(),method.Descriptor())
	}
}
//loop()函数循环执行“计算pc，解码指令,执行指令”
func  loop(thread *rtda.Thread, logInst bool){
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc :=  frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(frame.Method().Code(),pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if (logInst) {
			logInstruction(frame, inst)
		}
		//execute
		inst.Execute(frame)
		//判断java虚拟机栈中是否还有帧
		if thread.IsStackEmpty(){
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className,methodName,pc,inst,inst)
}