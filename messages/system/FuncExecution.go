// automatically generated by the FlatBuffers compiler, do not modify

package system

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FuncExecution struct {
	_tab flatbuffers.Table
}

func GetRootAsFuncExecution(buf []byte, offset flatbuffers.UOffsetT) *FuncExecution {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FuncExecution{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FuncExecution) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FuncExecution) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FuncExecution) Id(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FuncExecution) IdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FuncExecution) FuncId(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FuncExecution) FuncIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FuncExecution) UserId(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FuncExecution) UserIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FuncExecution) Context() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FuncExecution) Stdout() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FuncExecution) Stderr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FuncExecution) Start() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FuncExecution) MutateStart(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *FuncExecution) End() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FuncExecution) MutateEnd(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func (rcv *FuncExecution) Duration() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FuncExecution) MutateDuration(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func FuncExecutionStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func FuncExecutionAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func FuncExecutionStartIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FuncExecutionAddFuncId(builder *flatbuffers.Builder, funcId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(funcId), 0)
}
func FuncExecutionStartFuncIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FuncExecutionAddUserId(builder *flatbuffers.Builder, userId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(userId), 0)
}
func FuncExecutionStartUserIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FuncExecutionAddContext(builder *flatbuffers.Builder, context flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(context), 0)
}
func FuncExecutionAddStdout(builder *flatbuffers.Builder, stdout flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(stdout), 0)
}
func FuncExecutionAddStderr(builder *flatbuffers.Builder, stderr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(stderr), 0)
}
func FuncExecutionAddStart(builder *flatbuffers.Builder, start uint32) {
	builder.PrependUint32Slot(6, start, 0)
}
func FuncExecutionAddEnd(builder *flatbuffers.Builder, end uint32) {
	builder.PrependUint32Slot(7, end, 0)
}
func FuncExecutionAddDuration(builder *flatbuffers.Builder, duration uint32) {
	builder.PrependUint32Slot(8, duration, 0)
}
func FuncExecutionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}