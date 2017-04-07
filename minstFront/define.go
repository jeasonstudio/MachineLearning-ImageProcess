package minstFront

/*
MINST数据库是一个手写图像数据库，里面
*/

// MinstImg Minst 数据集图像信息
type MinstImg struct {
	ImgWidth int         // 图像宽
	ImgHight int         // 图像高
	ImgData  [][]float32 // 图像数据二维动态数组
}

// MinstImgArr 存储图像数据的object
type MinstImgArr struct {
	ImgNum int       // 存储图像的数目
	ImgPtr *MinstImg // 存储图像数组指针
}

// MinstLabel 输出标记
type MinstLabel struct {
	L         int       // 输出标记的长
	LabelData []float32 // 输出标记数据
}

// MinstLabelArr 存储图像标记的数组
type MinstLabelArr struct {
	LabelNum int
	LabelPtr *MinstLabel
}
