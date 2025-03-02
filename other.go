package dmsoft

// 其他杂项函数

func (com *Dmsoft) ActiveInputMethod(hwnd int, inputMethod string) int {
	ret, _ := com.dm.CallMethod("ActiveInputMethod", hwnd, inputMethod)
	return int(ret.Val)
}

func (com *Dmsoft) CheckInputMethod(hwnd int, inputMethod string) int {
	ret, _ := com.dm.CallMethod("CheckInputMethod", hwnd, inputMethod)
	return int(ret.Val)
}

func (com *Dmsoft) EnterCri() int {
	ret, _ := com.dm.CallMethod("EnterCri")
	return int(ret.Val)
}

func (com *Dmsoft) ExecuteCmd(cmd, currentDir string, timeOut int) string {
	ret, _ := com.dm.CallMethod("ExecuteCmd", cmd, currentDir, timeOut)
	return ret.ToString()
}

func (com *Dmsoft) FindInputMethod(inputMethod string) int {
	ret, _ := com.dm.CallMethod("FindInputMethod", inputMethod)
	return int(ret.Val)
}

func (com *Dmsoft) InitCri() int {
	ret, _ := com.dm.CallMethod("InitCri")
	return int(ret.Val)
}

func (com *Dmsoft) LeaveCri() int {
	ret, _ := com.dm.CallMethod("LeaveCri")
	return int(ret.Val)
}

func (com *Dmsoft) ReleaseRef() int {
	ret, _ := com.dm.CallMethod("ReleaseRef")
	return int(ret.Val)
}

func (com *Dmsoft) SetExitThread(mode int) int {
	ret, _ := com.dm.CallMethod("SetExitThread", mode)
	return int(ret.Val)
}
