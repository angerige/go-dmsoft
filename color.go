package dmsoft

import (
	"github.com/go-ole/go-ole"
)

// 图色 相关函数
func (com *Dmsoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

// FindColor 找色
func (com *Dmsoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

// FindMultiColor 多点找色
func (com *Dmsoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColorE(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorE", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) GetColor(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColor", x, y)
	return ret.ToString()
}

func (com *Dmsoft) GetColorBGR(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColorBGR", x, y)
	return ret.ToString()
}

func (com *Dmsoft) GetColorHSV(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColorHSV", x, y)
	return ret.ToString()
}

func (com *Dmsoft) GetPicSize(picName string) string {
	ret, _ := com.dm.CallMethod("GetPicSize", picName)
	return ret.ToString()
}

func (com *Dmsoft) GetScreenData(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod("GetScreenData", x1, y1, x2, y2)
	return int(ret.Val)
}

// GetScreenDataBmp 获取屏幕区域的图像数据，返回位图格式数据
func (com *Dmsoft) GetScreenDataBmp(x1, y1, x2, y2 int, data, size *int) int {
	// 创建 COM 变量
	d := ole.NewVariant(ole.VT_I4, int64(*data))
	s := ole.NewVariant(ole.VT_I4, int64(*size))

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("GetScreenDataBmp", x1, y1, x2, y2, &d, &s)

	// 更新输出参数
	*data = int(d.Val)
	*size = int(s.Val)

	// 清理 COM 变量
	d.Clear()
	s.Clear()

	return int(ret.Val)
}

func (com *Dmsoft) IsDisplayDead(x1, y1, x2, y2, t int) int {
	ret, _ := com.dm.CallMethod("IsDisplayDead", x1, y1, x2, y2, t)
	return int(ret.Val)
}

func (com *Dmsoft) LoadPic(picName string) int {
	ret, _ := com.dm.CallMethod("LoadPic", picName)
	return int(ret.Val)
}

func (com *Dmsoft) FreePic(picName string) int {
	ret, _ := com.dm.CallMethod("FreePic", picName)
	return int(ret.Val)
}

func (com *Dmsoft) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.dm.CallMethod("AppendPicAddr", picInfo, addr, size)
	return ret.ToString()
}

func (com *Dmsoft) BGR2RGB(bgrColor string) string {
	ret, _ := com.dm.CallMethod("BGR2RGB", bgrColor)
	return ret.ToString()
}

func (com *Dmsoft) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.dm.CallMethod("CaptureGif", x1, y1, x2, y2, file, delay, time)
	return int(ret.Val)
}

func (com *Dmsoft) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("CaptureJpg", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *Dmsoft) CapturePre(file string) int {
	ret, _ := com.dm.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *Dmsoft) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("CapturePng", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *Dmsoft) CmpColor(x, y int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.dm.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

func (com *Dmsoft) EnableFindPicMultithread(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFindPicMultithread", enable)
	return int(ret.Val)
}

func (com *Dmsoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.dm.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

// FindColorBlock 找色块
func (com *Dmsoft) FindColorBlock(x1, y1, x2, y2 int, color string, sim float32, count, width, height int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColorBlock", x1, y1, x2, y2, color, sim, count, width, height, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float32, count, width, height int) string {
	ret, _ := com.dm.CallMethod("FindColorBlockEx", x1, y1, x2, y2, color, sim, count, width, height)
	return ret.ToString()
}

func (com *Dmsoft) FindColorE(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorE", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindMulColor(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("FindMulColor", x1, y1, x2, y2, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

// FindPic 找图
func (com *Dmsoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicE(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicE", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicExS", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

// FindPicMem 从内存中找图
func (com *Dmsoft) FindPicMem(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicMem", x1, y1, x2, y2, picInfo, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicMemE(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicMemE", x1, y1, x2, y2, picInfo, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicMemEx(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicMemEx", x1, y1, x2, y2, picInfo, deltaColor, sim, dir)
	return ret.ToString()
}

// FindPicS 找图并返回坐标描述
func (com *Dmsoft) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicS", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

// FindPicSim 找相似图片
func (com *Dmsoft) FindPicSim(x1, y1, x2, y2 int, picName string, deltaColor string, sim, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicSim", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicSimE(x1, y1, x2, y2 int, picName string, deltaColor string, sim, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicSimE", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicSimEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicSimEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

// FindPicSimMem 从内存中找相似图片
func (com *Dmsoft) FindPicSimMem(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicSimMem", x1, y1, x2, y2, picInfo, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindPicSimMemE(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicSimMemE", x1, y1, x2, y2, picInfo, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindPicSimMemEx(x1, y1, x2, y2 int, picInfo string, deltaColor string, sim, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicSimMemEx", x1, y1, x2, y2, picInfo, deltaColor, sim, dir)
	return ret.ToString()
}

// FindShape 找形状
func (com *Dmsoft) FindShape(x1, y1, x2, y2 int, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindShape", x1, y1, x2, y2, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) FindShapeE(x1, y1, x2, y2 int, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindShapeE", x1, y1, x2, y2, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) FindShapeEx(x1, y1, x2, y2 int, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindShapeEx", x1, y1, x2, y2, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *Dmsoft) GetAveHSV(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveHSV", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *Dmsoft) GetAveRGB(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveRGB", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *Dmsoft) GetColorNum(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("GetColorNum", x1, y1, x2, y2, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) ImageToBmp(picName, bmpName string) int {
	ret, _ := com.dm.CallMethod("ImageToBmp", picName, bmpName)
	return int(ret.Val)
}

func (com *Dmsoft) LoadPicByte(addr, size int, picName string) int {
	ret, _ := com.dm.CallMethod("LoadPicByte", addr, size, picName)
	return int(ret.Val)
}

func (com *Dmsoft) MatchPicName(picName string) string {
	ret, _ := com.dm.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *Dmsoft) RGB2BGR(rgbColor string) string {
	ret, _ := com.dm.CallMethod("RGB2BGR", rgbColor)
	return ret.ToString()
}

func (com *Dmsoft) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", mode, info)
	return int(ret.Val)
}
