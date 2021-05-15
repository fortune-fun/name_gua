package name_gua

import (
	"github.com/godcong/yi"
)

type GuaType = int

const (
	DefaultBaGua GuaType = iota
	Ss1BaGua
	Ss2BaGua
	M1BaGua
	M2BaGua
)

// 周易八卦
// 姓总数取上卦，名总数取下卦，姓名总数取动爻
func (guaSanCai *GuaSanCai) BaGuaS1() *yi.Yi {
	return yi.NumberQiGua(guaSanCai.SanCaiNum.Tian, guaSanCai.SanCaiNum.Di, guaSanCai.SanCaiNum.Ren)
}

// 周神松卦象
// 姓名自下而上得三爻取下卦，自上而下得天人地三格，天格取上卦，天格取动爻
// 三四字名自下而上取卦，较为特殊
func (guaSanCai *GuaSanCai) BaGuaS2(given_strok2 int, given_strok1 int) *yi.Yi {
	return yi.NumberQiGua(yi.GetGua3Num(given_strok2, given_strok1, guaSanCai.SanCaiNum.Di), guaSanCai.SanCaiNum.Tian, guaSanCai.SanCaiNum.Ren)
}

// 周神松卦象2型
// 单名单姓，单名复姓
// 名总数取上卦，姓名总数取下卦，名总数取动爻
func (guaSanCai *GuaSanCai) BaGuaS3() *yi.Yi {
	return yi.NumberQiGua(guaSanCai.SanCaiNum.Ren, guaSanCai.SanCaiNum.Tian, guaSanCai.SanCaiNum.Tian)
}

// 民国卦象1
// 名总数取上卦，姓名总数取下卦，姓名总数取动爻
func (guaSanCai *GuaSanCai) BaGuaM1() *yi.Yi {
	return yi.NumberQiGua(guaSanCai.SanCaiNum.Ren, guaSanCai.SanCaiNum.Tian, guaSanCai.SanCaiNum.Ren)
}

// 民国卦象2
// 姓总数取上卦，姓名总数取下卦，姓名总数取动爻
func (guaSanCai *GuaSanCai) BaGuaM2() *yi.Yi {
	return yi.NumberQiGua(guaSanCai.SanCaiNum.Ren, guaSanCai.SanCaiNum.Di, guaSanCai.SanCaiNum.Ren)
}
