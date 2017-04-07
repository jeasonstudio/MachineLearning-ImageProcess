package CNN

const (
	AvePool = 0
	MaxPool = 1
	MinPool = 2
)

// ConvolutionalLayer 卷积层
type ConvolutionalLayer struct {
	InputWidth  int //输入图像的宽
	InputHeight int //输入图像的长
	MapSize     int //特征模板的大小，模板一般都是正方形

	InChannels  int //输入图像的数目
	OutChannels int //输出图像的数目

	/* 关于特征模板的权重分布，这里是一个四维数组
	其大小为inChannels*outChannels*mapSize*mapSize大小
	这里用四维数组，主要是为了表现全连接的形式，实际上卷积层并没有用到全连接的形式
	这里的例子是DeapLearningToolboox里的CNN例子，其用到就是全连接 */
	MapData  [][][][]float32 //存放特征模块的数据
	DmapData [][][][]float32 //存放特征模块的数据的局部梯度

	BasicData     []float32 //偏置，偏置的大小，为outChannels
	IsFullConnect bool      //是否为全连接
	ConnectModel  []bool    //连接模式（默认为全连接）

	// 下面三者的大小同输出的维度相同
	V [][][]float32 // 进入激活函数的输入值
	Y [][][]float32 // 激活函数后神经元的输出

	// 输出像素的局部梯度
	D [][][]float32 // 网络的局部梯度,δ值
}

// PoolingLayer 采样层 pooling
type PoolingLayer struct {
	InputWidth  int //输入图像的宽
	InputHeight int //输入图像的长
	MapSize     int //特征模板的大小

	InChannels  int //输入图像的数目
	OutChannels int //输出图像的数目

	PoolType  int       //Pooling的方法
	BasicData []float32 //偏置

	Y [][][]float32 // 采样函数后神经元的输出,无激活函数
	D [][][]float32 // 网络的局部梯度,δ值
}

// OutLayer 输出层 全连接的神经网络
type OutLayer struct {
	InputNum  int //输入数据的数目
	OutputNum int //输出数据的数目

	WData     [][]float32 // 权重数据，为一个inputNum*outputNum大小
	BasicData []float32   //偏置，大小为outputNum大小

	// 下面三者的大小同输出的维度相同
	V []float32 // 进入激活函数的输入值
	Y []float32 // 激活函数后神经元的输出
	D []float32 // 网络的局部梯度,δ值

	IsFullConnect bool //是否为全连接
}

// CNN .
type CNN struct {
	LayerNum int
	C1       *ConvolutionalLayer
	S2       *PoolingLayer
	C3       *ConvolutionalLayer
	S4       *PoolingLayer
	O5       *OutLayer

	E []float32 // 训练误差
	L []float32 // 瞬时误差能量
}

// CnnOpts 训练相关
type CnnOpts struct {
	numepochs int     // 训练的迭代次数
	alpha     float32 // 学习速率
}

// void cnnsetup(CNN* cnn,nSize inputSize,int outputSize);
// /*
// 	CNN网络的训练函数
// 	inputData，outputData分别存入训练数据
// 	dataNum表明数据数目
// */
// void cnntrain(CNN* cnn,	ImgArr inputData,LabelArr outputData,CNNOpts opts,int trainNum);
// // 测试cnn函数
// float cnntest(CNN* cnn, ImgArr inputData,LabelArr outputData,int testNum);
// // 保存cnn
// void savecnn(CNN* cnn, const char* filename);
// // 导入cnn的数据
// void importcnn(CNN* cnn, const char* filename);

// // 初始化卷积层
// CovLayer* initCovLayer(int inputWidth,int inputHeight,int mapSize,int inChannels,int outChannels);
// void CovLayerConnect(CovLayer* covL,bool* connectModel);
// // 初始化采样层
// PoolLayer* initPoolLayer(int inputWidth,int inputHeigh,int mapSize,int inChannels,int outChannels,int poolType);
// void PoolLayerConnect(PoolLayer* poolL,bool* connectModel);
// // 初始化输出层
// OutLayer* initOutLayer(int inputNum,int outputNum);

// // 激活函数 input是数据，inputNum说明数据数目，bas表明偏置
// float activation_Sigma(float input,float bas); // sigma激活函数

// void cnnff(CNN* cnn,float** inputData); // 网络的前向传播
// void cnnbp(CNN* cnn,float* outputData); // 网络的后向传播
// void cnnapplygrads(CNN* cnn,CNNOpts opts,float** inputData);
// void cnnclear(CNN* cnn); // 将数据vyd清零

// /*
// 	Pooling Function
// 	input 输入数据
// 	inputNum 输入数据数目
// 	mapSize 求平均的模块区域
// */
// void avgPooling(float** output,nSize outputSize,float** input,nSize inputSize,int mapSize); // 求平均值

// /*
// 	单层全连接神经网络的处理
// 	nnSize是网络的大小
// */
// void nnff(float* output,float* input,float** wdata,float* bas,nSize nnSize); // 单层全连接神经网络的前向传播

// void savecnndata(CNN* cnn,const char* filename,float** inputdata); // 保存CNN网络中的相关数据

// #endif
