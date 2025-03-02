package dmsoft

import "github.com/go-ole/go-ole"

// 窗口 相关函数
func (com *Dmsoft) FindWindow(class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowByProcessId", processId, class, title)
	return int(ret.Val)
}

func (com *Dmsoft) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.dm.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *Dmsoft) EnumProcess(name string) string {
	ret, _ := com.dm.CallMethod("EnumProcess", name)
	return ret.ToString()
}

func (com *Dmsoft) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowByProcess(processName, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindowByProcess", processName, title, className, filter)
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.dm.CallMethod("EnumWindowByProcessId", pid, title, className, filter)
	return ret.ToString()
}

func (com *Dmsoft) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.dm.CallMethod("EnumWindowSuper", spec1, flag1, type1, spec2, flag2, type2, sort)
	return ret.ToString()
}

func (com *Dmsoft) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.dm.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

func (com *Dmsoft) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.dm.CallMethod("FindWindowSuper", spec1, flag1, type1, spec2, flag2, type2)
	return int(ret.Val)
}

// GetClientRect 获取窗口客户区域
// 参数:
//
//	hwnd: 窗口句柄
//	x1, y1, x2, y2: 返回的坐标
//
// 返回值:
//
//	1: 成功
//	0: 失败
func (com *Dmsoft) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	// 创建 COM 变量
	x1Var := ole.NewVariant(ole.VT_I4, 0)
	y1Var := ole.NewVariant(ole.VT_I4, 0)
	x2Var := ole.NewVariant(ole.VT_I4, 0)
	y2Var := ole.NewVariant(ole.VT_I4, 0)

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("GetClientRect", hwnd, &x1Var, &y1Var, &x2Var, &y2Var)

	// 获取返回值
	*x1 = int(x1Var.Val)
	*y1 = int(y1Var.Val)
	*x2 = int(x2Var.Val)
	*y2 = int(y2Var.Val)

	// 清理 COM 变量
	x1Var.Clear()
	y1Var.Clear()
	x2Var.Clear()
	y2Var.Clear()

	return int(ret.Val)
}

// GetClientSize 获取窗口客户区域大小
// 参数:
//
//	hwnd: 窗口句柄
//	width, height: 返回的宽高
//
// 返回值:
//
//	1: 成功
//	0: 失败
func (com *Dmsoft) GetClientSize(hwnd int, width, height *int) int {
	// 创建 COM 变量
	widthVar := ole.NewVariant(ole.VT_I4, 0)
	heightVar := ole.NewVariant(ole.VT_I4, 0)

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("GetClientSize", hwnd, &widthVar, &heightVar)

	// 获取返回值
	*width = int(widthVar.Val)
	*height = int(heightVar.Val)

	// 清理 COM 变量
	widthVar.Clear()
	heightVar.Clear()

	return int(ret.Val)
}

func (com *Dmsoft) GetForegroundFocus() int {
	ret, _ := com.dm.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

func (com *Dmsoft) GetForegroundWindow() int {
	ret, _ := com.dm.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

func (com *Dmsoft) GetMousePointWindow() int {
	ret, _ := com.dm.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *Dmsoft) GetPointWindow(x, y int) int {
	ret, _ := com.dm.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *Dmsoft) GetProcessInfo(pid int) string {
	ret, _ := com.dm.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

func (com *Dmsoft) GetSpecialWindow(flag int) int {
	ret, _ := com.dm.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

func (com *Dmsoft) GetWindow(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowClass(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *Dmsoft) GetWindowProcessId(hwnd int) int {
	ret, _ := com.dm.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *Dmsoft) GetWindowState(hwnd, flag int) int {
	ret, _ := com.dm.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowThreadId(hwnd int) int {
	ret, _ := com.dm.CallMethod("GetWindowThreadId", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) GetWindowTitle(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

// ScreenToClient 屏幕坐标转窗口坐标
func (com *Dmsoft) ScreenToClient(hwnd int, x, y *int) int {
	// 创建 COM 变量
	xVar := ole.NewVariant(ole.VT_I4, 0)
	yVar := ole.NewVariant(ole.VT_I4, 0)

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("ScreenToClient", hwnd, &xVar, &yVar)

	// 获取返回值
	*x = int(xVar.Val)
	*y = int(yVar.Val)

	// 清理 COM 变量
	xVar.Clear()
	yVar.Clear()

	return int(ret.Val)
}

func (com *Dmsoft) SendPaste(hwnd int) int {
	ret, _ := com.dm.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) SendString(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod("SendString", hwnd, str)
	return int(ret.Val)
}

func (com *Dmsoft) SendString2(hwnd int, str string) int {
	ret, _ := com.dm.CallMethod("SendString2", hwnd, str)
	return int(ret.Val)
}

func (com *Dmsoft) SendStringIme(str string) int {
	ret, _ := com.dm.CallMethod("SendStringIme", str)
	return int(ret.Val)
}

func (com *Dmsoft) SendStringIme2(hwnd int, str string, mode int) int {
	ret, _ := com.dm.CallMethod("SendStringIme2", hwnd, str, mode)
	return int(ret.Val)
}

func (com *Dmsoft) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.dm.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowText(hwnd int, title string) int {
	ret, _ := com.dm.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *Dmsoft) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.dm.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}

func (com *Dmsoft) SetSendStringDelay(delay int) int {
	ret, _ := com.dm.CallMethod("SetSendStringDelay", delay)
	return int(ret.Val)
}

// ClientToScreen 窗口坐标转屏幕坐标
func (com *Dmsoft) ClientToScreen(hwnd int, x, y *int) int {
	// 创建 COM 变量
	xVar := ole.NewVariant(ole.VT_I4, 0)
	yVar := ole.NewVariant(ole.VT_I4, 0)

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("ClientToScreen", hwnd, &xVar, &yVar)

	// 获取返回值
	*x = int(xVar.Val)
	*y = int(yVar.Val)

	// 清理 COM 变量
	xVar.Clear()
	yVar.Clear()

	return int(ret.Val)
}
