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
	pos        int
	intVals    [10]int
	structVals [10]VarGenStruct
	sliceVals  [10][]int
}

func (g *VarGenImpl) GenIntPtr() *int {
	g.pos = (g.pos + 1) % 10
	return &g.intVals[g.pos]
}

func (g *VarGenImpl) GenStructPtr() *VarGenStruct {
	g.pos = (g.pos + 1) % 10
	return &g.structVals[g.pos]
}

func (g *VarGenImpl) GenSlice() []int {
	return g.intVals[1:5]
}

func (g *VarGenImpl) GenSlicePtr() *[]int {
	g.pos = (g.pos + 1) % 10
	g.sliceVals[g.pos] = g.intVals[:g.pos]
	return &g.sliceVals[g.pos]
}
