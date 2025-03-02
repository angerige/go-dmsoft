package dmsoft

// Foobar 相关函数

func (com *Dmsoft) CreateFoobarCustom(hwnd, x, y int, picName, transColor string, sim float32) int {
	ret, _ := com.dm.CallMethod("CreateFoobarCustom", hwnd, x, y, picName, transColor, sim)
	return int(ret.Val)
}

func (com *Dmsoft) CreateFoobarEllipse(hwnd, x, y, w, h int) int {
	ret, _ := com.dm.CallMethod("CreateFoobarEllipse", hwnd, x, y, w, h)
	return int(ret.Val)
}

func (com *Dmsoft) CreateFoobarRect(hwnd, x, y, w, h int) int {
	ret, _ := com.dm.CallMethod("CreateFoobarRect", hwnd, x, y, w, h)
	return int(ret.Val)
}

func (com *Dmsoft) CreateFoobarRoundRect(hwnd, x, y, w, h, rw, rh int) int {
	ret, _ := com.dm.CallMethod("CreateFoobarRoundRect", hwnd, x, y, w, h, rw, rh)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarClearText(hwnd int) int {
	ret, _ := com.dm.CallMethod("FoobarClearText", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarClose(hwnd int) int {
	ret, _ := com.dm.CallMethod("FoobarClose", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarDrawLine(hwnd, x1, y1, x2, y2 int, color string, style, width int) int {
	ret, _ := com.dm.CallMethod("FoobarDrawLine", hwnd, x1, y1, x2, y2, color, style, width)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarDrawPic(hwnd, x, y int, picName, transColor string) int {
	ret, _ := com.dm.CallMethod("FoobarDrawPic", hwnd, x, y, picName, transColor)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarDrawText(hwnd, x, y, w, h int, text, color string, align int) int {
	ret, _ := com.dm.CallMethod("FoobarDrawText", hwnd, x, y, w, h, text, color, align)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarFillRect(hwnd, x1, y1, x2, y2 int, color string) int {
	ret, _ := com.dm.CallMethod("FoobarFillRect", hwnd, x1, y1, x2, y2, color)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarLock(hwnd int) int {
	ret, _ := com.dm.CallMethod("FoobarLock", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarPrintText(hwnd int, text, color string) int {
	ret, _ := com.dm.CallMethod("FoobarPrintText", hwnd, text, color)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarSetFont(hwnd int, fontName string, size, flag int) int {
	ret, _ := com.dm.CallMethod("FoobarSetFont", hwnd, fontName, size, flag)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarSetSave(hwnd int, file string, enable int, header string) int {
	ret, _ := com.dm.CallMethod("FoobarSetSave", hwnd, file, enable, header)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarSetTrans(hwnd, isTrans int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("FoobarSetTrans", hwnd, isTrans, color, sim)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarStartGif(hwnd, x, y int, picName string, repeatLimit, delay int) int {
	ret, _ := com.dm.CallMethod("FoobarStartGif", hwnd, x, y, picName, repeatLimit, delay)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarStopGif(hwnd, x, y int, picName string) int {
	ret, _ := com.dm.CallMethod("FoobarStopGif", hwnd, x, y, picName)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarTextLineGap(hwnd, lineGap int) int {
	ret, _ := com.dm.CallMethod("FoobarTextLineGap", hwnd, lineGap)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarTextPrintDir(hwnd, dir int) int {
	ret, _ := com.dm.CallMethod("FoobarTextPrintDir", hwnd, dir)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarTextRect(hwnd, x, y, w, h int) int {
	ret, _ := com.dm.CallMethod("FoobarTextRect", hwnd, x, y, w, h)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarUnlock(hwnd int) int {
	ret, _ := com.dm.CallMethod("FoobarUnlock", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) FoobarUpdate(hwnd int) int {
	ret, _ := com.dm.CallMethod("FoobarUpdate", hwnd)
	return int(ret.Val)
}
