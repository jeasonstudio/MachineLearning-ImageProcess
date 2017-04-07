package minstFront

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ReadLable 读入图像标记
func ReadLable(fileName string) MinstLabelArr {
	a := MinstLabelArr{}
	return a
}

// ReadImg 读入图像
func ReadImg(filename string) MinstImgArr {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		defer file.Close()
		os.Exit(0)
	}
	file.Seek(8, 0)
	fmt.Println("Success Open File")
	var buffer bytes.Buffer
	io.CopyN(&buffer, file, 8)
	_bytes := buffer.Bytes()
	var magic []byte = []byte{0x12, 0x31, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22}
	for _, byte := range _bytes {
		fmt.Printf("%02X ", byte)
	}
	fmt.Println()
	if bytes.Compare(magic, _bytes) == 0 {
		fmt.Println("Equal")
	}

	a := MinstImgArr{}
	return a
}

// SaveImg 将图像数据保存成文件
func SaveImg(imgArrs MinstImgArr, filedir string) {

}

// ImgArr read_Img(const char* filename) // 读入图像
// {
// 	FILE  *fp=NULL;
// 	fp=fopen(filename,"rb");
// 	if(fp==NULL)
// 		printf("open file failed\n");
// 	assert(fp);

// 	int magic_number = 0;
// 	int number_of_images = 0;
// 	int n_rows = 0;
// 	int n_cols = 0;
// 	//从文件中读取sizeof(magic_number) 个字符到 &magic_number
// 	fread((char*)&magic_number,sizeof(magic_number),1,fp);
// 	magic_number = ReverseInt(magic_number);
// 	//获取训练或测试image的个数number_of_images
// 	fread((char*)&number_of_images,sizeof(number_of_images),1,fp);
// 	number_of_images = ReverseInt(number_of_images);
// 	//获取训练或测试图像的高度Heigh
// 	fread((char*)&n_rows,sizeof(n_rows),1,fp);
// 	n_rows = ReverseInt(n_rows);
// 	//获取训练或测试图像的宽度Width
// 	fread((char*)&n_cols,sizeof(n_cols),1,fp);
// 	n_cols = ReverseInt(n_cols);
// 	//获取第i幅图像，保存到vec中
// 	int i,r,c;

// 	// 图像数组的初始化
// 	ImgArr imgarr=(ImgArr)malloc(sizeof(MinstImgArr));
// 	imgarr->ImgNum=number_of_images;
// 	imgarr->ImgPtr=(MinstImg*)malloc(number_of_images*sizeof(MinstImg));

// 	for(i = 0; i < number_of_images; ++i)
// 	{
// 		imgarr->ImgPtr[i].r=n_rows;
// 		imgarr->ImgPtr[i].c=n_cols;
// 		imgarr->ImgPtr[i].ImgData=(float**)malloc(n_rows*sizeof(float*));
// 		for(r = 0; r < n_rows; ++r)
// 		{
// 			imgarr->ImgPtr[i].ImgData[r]=(float*)malloc(n_cols*sizeof(float));
// 			for(c = 0; c < n_cols; ++c)
// 			{
// 				unsigned char temp = 0;
// 				fread((char*) &temp, sizeof(temp),1,fp);
// 				imgarr->ImgPtr[i].ImgData[r][c]=(float)temp/255.0;
// 			}
// 		}
// 	}

// 	fclose(fp);
// 	return imgarr;
// }
