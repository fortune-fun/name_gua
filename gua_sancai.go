package name_gua

import (
	"github.com/godcong/yi"
)

type GuaSanCai struct {
	SanCaiNum *yi.SanCaiNum
}

//SanCaiWuGe 三才五格
func GetGuaSanCai(tian, ren, di int) *GuaSanCai {
	sanCaiNum := yi.SanCaiNum{
		Tian: tian,
		Ren:  ren,
		Di:   di,
	}
	guaSanCai := GuaSanCai{
		SanCaiNum: &sanCaiNum,
	}

	return &guaSanCai
}
