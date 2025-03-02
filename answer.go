package dmsoft

// 答题 相关函数

func (com *Dmsoft) FaqCancel() int {
	ret, _ := com.dm.CallMethod("FaqCancel")
	return int(ret.Val)
}

func (com *Dmsoft) FaqCapture(x1, y1, x2, y2, quality, delay, time int) int {
	ret, _ := com.dm.CallMethod("FaqCapture", x1, y1, x2, y2, quality, delay, time)
	return int(ret.Val)
}

func (com *Dmsoft) FaqCaptureFromFile(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("FaqCaptureFromFile", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *Dmsoft) FaqCaptureString(str string) int {
	ret, _ := com.dm.CallMethod("FaqCaptureString", str)
	return int(ret.Val)
}

func (com *Dmsoft) FaqFetch() string {
	ret, _ := com.dm.CallMethod("FaqFetch")
	return ret.ToString()
}

func (com *Dmsoft) FaqGetSize(handle int) int {
	ret, _ := com.dm.CallMethod("FaqGetSize", handle)
	return int(ret.Val)
}

func (com *Dmsoft) FaqIsPosted() int {
	ret, _ := com.dm.CallMethod("FaqIsPosted")
	return int(ret.Val)
}

func (com *Dmsoft) FaqPost(server string, handle, requestType, timeOut int) int {
	ret, _ := com.dm.CallMethod("FaqPost", server, handle, requestType, timeOut)
	return int(ret.Val)
}

func (com *Dmsoft) FaqSend(server string, handle, requestType, timeOut int) string {
	ret, _ := com.dm.CallMethod("FaqSend", server, handle, requestType, timeOut)
	return ret.ToString()
}
