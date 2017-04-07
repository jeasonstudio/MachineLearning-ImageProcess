package mat

// 二维矩阵数组的操作

const (
	full  = 0
	same  = 1
	valid = 2
)

type nSize struct {
	C int // 列（宽）
	R int // 行（高）
}

// float** rotate180(float** mat, nSize matSize);// 矩阵翻转180度

// void addmat(float** res, float** mat1, nSize matSize1, float** mat2, nSize matSize2);// 矩阵相加

// float** correlation(float** map,nSize mapSize,float** inputData,nSize inSize,int type);// 互相关

// float** cov(float** map,nSize mapSize,float** inputData,nSize inSize,int type); // 卷积操作

// // 这个是矩阵的上采样（等值内插），upc及upr是内插倍数
// float** UpSample(float** mat,nSize matSize,int upc,int upr);

// // 给二维矩阵边缘扩大，增加addw大小的0值边
// float** matEdgeExpand(float** mat,nSize matSize,int addc,int addr);

// // 给二维矩阵边缘缩小，擦除shrinkc大小的边
// float** matEdgeShrink(float** mat,nSize matSize,int shrinkc,int shrinkr);

// void savemat(float** mat,nSize matSize,const char* filename);// 保存矩阵数据

// void multifactor(float** res, float** mat, nSize matSize, float factor);// 矩阵乘以系数

// float summat(float** mat,nSize matSize);// 矩阵各元素的和

// char * combine_strings(char *a, char *b);

// char* intTochar(int i);
