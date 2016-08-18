package main

type VarGenStruct struct {
	a, b, c int
}

type VarGen interface {
	GenIntPtr() *int
	GenStructPtr() *VarGenStruct
	GenSlice() []int
	GenSlicePtr() *[]int
}

// VarGen implementation
type VarGenImpl struct {
	intVals   [10]int
	structVal VarGenStruct
	sliceVal  []int
}

func (g *VarGenImpl) GenIntPtr() *int {
	return &g.intVals[0]
}

func (g *VarGenImpl) GenStructPtr() *VarGenStruct {
	return &g.structVal
}

func (g *VarGenImpl) GenSlice() []int {
	return g.intVals[1:5]
}

func (g *VarGenImpl) GenSlicePtr() *[]int {
	g.sliceVal = g.intVals[1:3]
	return &g.sliceVal
}
