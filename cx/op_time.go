package base

import (
	"time"
)

func op_time_UnixMilli (expr *CXExpression, stack *CXStack, fp int) {
	out1 := expr.Outputs[0]
	outB1 := FromI64(time.Now().UnixNano() / int64(1000000))
	WriteMemory(stack, GetFinalOffset(stack, fp, out1, MEM_WRITE), out1, outB1)
}

func op_time_UnixNano (expr *CXExpression, stack *CXStack, fp int) {
	out1 := expr.Outputs[0]
	outB1 := FromI64(time.Now().UnixNano())
	WriteMemory(stack, GetFinalOffset(stack, fp, out1, MEM_WRITE), out1, outB1)
}

func op_time_Sleep (expr *CXExpression, stack *CXStack, fp int) {
	inp1 := expr.Inputs[0]
	time.Sleep(time.Duration(ReadI32(stack, fp, inp1)) * time.Millisecond)
}
