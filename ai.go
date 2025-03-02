package dmsoft

import (
	"github.com/go-ole/go-ole"
)

// AI 相关函数

func (com *Dmsoft) AiYoloDetectObjects(x1, y1, x2, y2 int, prob, iou float32) string {
	ret, _ := com.dm.CallMethod("AiYoloDetectObjects", x1, y1, x2, y2, prob, iou)
	return ret.ToString()
}

func (com *Dmsoft) AiYoloDetectObjectsToDataBmp(x1, y1, x2, y2 int, prob, iou float32, data, size *int, mode int) int {
	// 创建 COM 变量
	d := ole.NewVariant(ole.VT_I4, int64(*data))
	s := ole.NewVariant(ole.VT_I4, int64(*size))

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("AiYoloDetectObjectsToDataBmp", x1, y1, x2, y2, prob, iou, &d, &s, mode)

	// 更新输出参数
	*data = int(d.Val)
	*size = int(s.Val)

	// 清理 COM 变量
	d.Clear()
	s.Clear()

	return int(ret.Val)
}

func (com *Dmsoft) AiYoloDetectObjectsToFile(x1, y1, x2, y2 int, prob, iou float32, file string, mode int) int {
	ret, _ := com.dm.CallMethod("AiYoloDetectObjectsToFile", x1, y1, x2, y2, prob, iou, file, mode)
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloFreeModel(index int) int {
	ret, _ := com.dm.CallMethod("AiYoloFreeModel", index)
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloObjectsToString(objects string) string {
	ret, _ := com.dm.CallMethod("AiYoloObjectsToString", objects)
	return ret.ToString()
}

func (com *Dmsoft) AiYoloSetModel(index int, file, pwd string) int {
	ret, _ := com.dm.CallMethod("AiYoloSetModel", index, file, pwd)
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSetModelMemory(index, data, size int, pwd string) int {
	ret, _ := com.dm.CallMethod("AiYoloSetModelMemory", index, data, size, pwd)
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSetVersion(ver string) int {
	ret, _ := com.dm.CallMethod("AiYoloSetVersion", ver)
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSortsObjects(objects string, height int) string {
	ret, _ := com.dm.CallMethod("AiYoloSortsObjects", objects, height)
	return ret.ToString()
}

func (com *Dmsoft) AiYoloUseModel(index int) int {
	ret, _ := com.dm.CallMethod("AiYoloUseModel", index)
	return int(ret.Val)
}

func (com *Dmsoft) LoadAi(file string) int {
	ret, _ := com.dm.CallMethod("LoadAi", file)
	return int(ret.Val)
}

func (com *Dmsoft) LoadAiMemory(data, size int) int {
	ret, _ := com.dm.CallMethod("LoadAiMemory", data, size)
	return int(ret.Val)
}

func (com *Dmsoft) AiEnableFindPicWindow(enable int) int {
	ret, _ := com.dm.CallMethod("AiEnableFindPicWindow", enable)
	return int(ret.Val)
}

func (com *Dmsoft) AiFindPic(x1, y1, x2, y2 int, picName string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("AiFindPic", x1, y1, x2, y2, picName, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiFindPicEx(x1, y1, x2, y2 int, picName string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("AiFindPicEx", x1, y1, x2, y2, picName, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) AiFindPicMem(x1, y1, x2, y2 int, picInfo string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("AiFindPicMem", x1, y1, x2, y2, picInfo, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiFindPicMemEx(x1, y1, x2, y2 int, picInfo string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("AiFindPicMemEx", x1, y1, x2, y2, picInfo, sim, dir)
	return ret.ToString()
}
