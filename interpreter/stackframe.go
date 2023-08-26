package interpreter

import (
	"slices"
	"sync/atomic"

	"github.com/charmbracelet/log"

	"arc/ast"
	"arc/utilities"
)

var frameId = atomic.Int64{}

type StackFrame struct {
	id     int64
	parent *StackFrame

	vars map[string]*ast.RuntimeValue

	deferStack *utilities.Stack[*ast.DeferStatement]

	// The function that this stack frame belongs to
	function *ast.FunctionDeclaration

	// The environment that this stack frame belongs to
	env *Environment

	markers []ast.Node

	callSites []*ast.CallExpression
	errs      []FrameErrorInfo
}

type FrameErrorInfo struct {
	Error any
	Node  ast.Node
}

var Frames = make([]*StackFrame, 0)

func NewStackFrame(parent ...*StackFrame) *StackFrame {
	inst := &StackFrame{
		id:         frameId.Add(1),
		vars:       make(map[string]*ast.RuntimeValue),
		deferStack: utilities.NewStack[*ast.DeferStatement](),
	}

	if len(parent) > 0 {
		inst.parent = parent[0]
	}

	Frames = append(Frames, inst)

	return inst
}

func getFrame() *StackFrame {
	if len(Frames) == 0 {
		log.Warnf("No stack frames found")
		return nil
	}
	return Frames[len(Frames)-1]
}

func (self *StackFrame) addExecutionMarker(node any) {
	if node == nil {
		return
	}
	if n, ok := node.(ast.Node); ok {
		self.markers = append(self.markers, n)
		// Cap the markers at 30
		if len(self.markers) > 30 {
			self.markers = self.markers[len(self.markers)-30:]
		}
	} else {
		log.Debugf("Cannot add execution marker for node, it doesn't implement ast.Node: %#v", node)
	}
}
func (self *StackFrame) addCallSite(node *ast.CallExpression) {
	if node == nil {
		return
	}
	self.callSites = append(self.callSites, node)
}

func (self *StackFrame) errored(err any) {
	e := FrameErrorInfo{
		Error: err,
	}

	if len(self.markers) > 0 {
		e.Node = self.markers[len(self.markers)-1]
	} else {
		e.Node = self.function
	}
	self.errs = append(self.errs, e)
}

func (self *StackFrame) didError() bool {
	return len(self.errs) > 0 // || self.parent != nil && self.parent.didError()
}

func (self *StackFrame) unrollCallSites() []*StackFrame {
	frames := make([]*StackFrame, 0)

	frame := self
	for frame != nil {
		frames = append(frames, frame)
		frame = frame.parent
	}

	slices.Reverse(frames)

	return frames
}

func (self *StackFrame) pop() {
	if len(Frames) == 0 {
		return
	}
	Frames = Frames[:len(Frames)-1]
}
