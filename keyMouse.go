package dmsoft

import (
	"fmt"

	"github.com/go-ole/go-ole"
)

// keyMouse 相关函数

func (com *Dmsoft) KeyPress(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) KeyDown(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) KeyUp(vkCode int) int {
	ret, _ := com.dm.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) LeftClick() int {
	ret, _ := com.dm.CallMethod("LeftClick")
	return int(ret.Val)
}

func (com *Dmsoft) LeftDown() int {
	ret, _ := com.dm.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *Dmsoft) LeftUp() int {
	ret, _ := com.dm.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *Dmsoft) RightClick() int {
	ret, _ := com.dm.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *Dmsoft) RightDown() int {
	ret, _ := com.dm.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *Dmsoft) RightUp() int {
	ret, _ := com.dm.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *Dmsoft) MoveTo(x, y int) int {
	ret, _ := com.dm.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseAccuracy(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseAccuracy", enable)
	return int(ret.Val)
}

// GetCursorPos 获取鼠标位置
// 参数:
//
//	x, y: 返回的鼠标坐标
//
// 返回值:
//
//	1: 成功
//	0: 失败
func (com *Dmsoft) GetCursorPos(x, y *int) int {
	// 创建 COM 变量
	x_ret := ole.NewVariant(ole.VT_I4, 0)
	y_ret := ole.NewVariant(ole.VT_I4, 0)

	// 调用 COM 方法
	ret, _ := com.dm.CallMethod("GetCursorPos", &x_ret, &y_ret)

	// 获取返回值
	*x = int(x_ret.Val)
	*y = int(y_ret.Val)

	// 清理 COM 变量
	x_ret.Clear()
	y_ret.Clear()

	return int(ret.Val)
}

func (com *Dmsoft) GetCursorShape() string {
	ret, _ := com.dm.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *Dmsoft) GetCursorShapeEx(type_ int) string {
	ret, _ := com.dm.CallMethod("GetCursorShapeEx", type_)
	return ret.ToString()
}

func (com *Dmsoft) GetCursorSpot() (x, y int, ok bool) {
	ret, _ := com.dm.CallMethod("GetCursorSpot")
	str := ret.ToString()
	if str == "" {
		return 0, 0, false
	}

	fmt.Sscanf(str, "%d,%d", &x, &y)
	return x, y, true
}

func (com *Dmsoft) GetKeyState(vkCode int) int {
	ret, _ := com.dm.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *Dmsoft) GetMouseSpeed() int {
	ret, _ := com.dm.CallMethod("GetMouseSpeed")
	return int(ret.Val)
}

func (com *Dmsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *Dmsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

func (com *Dmsoft) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.dm.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *Dmsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.dm.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

func (com *Dmsoft) LeftDoubleClick() int {
	ret, _ := com.dm.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleClick() int {
	ret, _ := com.dm.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleDown() int {
	ret, _ := com.dm.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *Dmsoft) MiddleUp() int {
	ret, _ := com.dm.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *Dmsoft) MoveR(rx, ry int) int {
	ret, _ := com.dm.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *Dmsoft) MoveToEx(x, y, w, h int) (targetX, targetY int, ok bool) {
	ret, _ := com.dm.CallMethod("MoveToEx", x, y, w, h)
	str := ret.ToString()
	if str == "" {
		return 0, 0, false
	}

	fmt.Sscanf(str, "%d,%d", &targetX, &targetY)
	return targetX, targetY, true
}

func (com *Dmsoft) SetKeypadDelay(type_ string, delay int) int {
	ret, _ := com.dm.CallMethod("SetKeypadDelay", type_, delay)
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseDelay(type_ string, delay int) int {
	ret, _ := com.dm.CallMethod("SetMouseDelay", type_, delay)
	return int(ret.Val)
}

func (com *Dmsoft) SetMouseSpeed(speed int) int {
	ret, _ := com.dm.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *Dmsoft) WaitKey(vkCode, timeout int) int {
	ret, _ := com.dm.CallMethod("WaitKey", vkCode, timeout)
	return int(ret.Val)
}

func (com *Dmsoft) WheelDown() int {
	ret, _ := com.dm.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *Dmsoft) WheelUp() int {
	ret, _ := com.dm.CallMethod("WheelUp")
	return int(ret.Val)
}
