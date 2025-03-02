package dmsoft

// 内存 相关函数

func (com *Dmsoft) ReadString(hwnd int, addr string, type_, len int) string {
	ret, _ := com.dm.CallMethod("ReadString", hwnd, addr, type_, len)
	return ret.ToString()
}

func (com *Dmsoft) ReadInt(hwnd int, addr string, type_ int) int64 {
	ret, _ := com.dm.CallMethod("ReadInt", hwnd, addr, type_)
	return ret.Val
}

func (com *Dmsoft) ReadIntAddr(hwnd int, addr int64, type_ int) int64 {
	ret, _ := com.dm.CallMethod("ReadIntAddr", hwnd, addr, type_)
	return ret.Val
}

func (com *Dmsoft) VirtualAllocEx(hwnd int, addr int64, size, type_ int) int64 {
	ret, _ := com.dm.CallMethod("VirtualAllocEx", hwnd, addr, size, type_)
	return ret.Val
}

func (com *Dmsoft) WriteData(hwnd int, addr string, data string) int {
	ret, _ := com.dm.CallMethod("WriteData", hwnd, addr, data)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDataAddr(hwnd int, addr int64, data string) int {
	ret, _ := com.dm.CallMethod("WriteDataAddr", hwnd, addr, data)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualFreeEx(hwnd int, addr int64) int {
	ret, _ := com.dm.CallMethod("VirtualFreeEx", hwnd, addr)
	return int(ret.Val)
}

func (com *Dmsoft) DoubleToData(value float64) string {
	ret, _ := com.dm.CallMethod("DoubleToData", value)
	return ret.ToString()
}

func (com *Dmsoft) FindData(hwnd int, addrRange, data string) string {
	ret, _ := com.dm.CallMethod("FindData", hwnd, addrRange, data)
	return ret.ToString()
}

func (com *Dmsoft) FindDataEx(hwnd int, addrRange, data string, step, multiThread, mode int) string {
	ret, _ := com.dm.CallMethod("FindDataEx", hwnd, addrRange, data, step, multiThread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindDouble(hwnd int, addrRange string, doubleValueMin, doubleValueMax float64) string {
	ret, _ := com.dm.CallMethod("FindDouble", hwnd, addrRange, doubleValueMin, doubleValueMax)
	return ret.ToString()
}

func (com *Dmsoft) FindDoubleEx(hwnd int, addrRange string, doubleValueMin, doubleValueMax float64, step, multiThread, mode int) string {
	ret, _ := com.dm.CallMethod("FindDoubleEx", hwnd, addrRange, doubleValueMin, doubleValueMax, step, multiThread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindFloat(hwnd int, addrRange string, floatValueMin, floatValueMax float32) string {
	ret, _ := com.dm.CallMethod("FindFloat", hwnd, addrRange, floatValueMin, floatValueMax)
	return ret.ToString()
}

func (com *Dmsoft) FindFloatEx(hwnd int, addrRange string, floatValueMin, floatValueMax float32, step, multiThread, mode int) string {
	ret, _ := com.dm.CallMethod("FindFloatEx", hwnd, addrRange, floatValueMin, floatValueMax, step, multiThread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindInt(hwnd int, addrRange string, intValueMin, intValueMax int64, type_ int) string {
	ret, _ := com.dm.CallMethod("FindInt", hwnd, addrRange, intValueMin, intValueMax, type_)
	return ret.ToString()
}

func (com *Dmsoft) FindIntEx(hwnd int, addrRange string, intValueMin, intValueMax int64, type_, step, multiThread, mode int) string {
	ret, _ := com.dm.CallMethod("FindIntEx", hwnd, addrRange, intValueMin, intValueMax, type_, step, multiThread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FindString(hwnd int, addrRange, stringValue string, type_ int) string {
	ret, _ := com.dm.CallMethod("FindString", hwnd, addrRange, stringValue, type_)
	return ret.ToString()
}

func (com *Dmsoft) FindStringEx(hwnd int, addrRange, stringValue string, type_, step, multiThread, mode int) string {
	ret, _ := com.dm.CallMethod("FindStringEx", hwnd, addrRange, stringValue, type_, step, multiThread, mode)
	return ret.ToString()
}

func (com *Dmsoft) FloatToData(value float32) string {
	ret, _ := com.dm.CallMethod("FloatToData", value)
	return ret.ToString()
}

func (com *Dmsoft) FreeProcessMemory(hwnd int) int {
	ret, _ := com.dm.CallMethod("FreeProcessMemory", hwnd)
	return int(ret.Val)
}

func (com *Dmsoft) GetCommandLine(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetCommandLine", hwnd)
	return ret.ToString()
}

func (com *Dmsoft) GetModuleBaseAddr(hwnd int, module string) int64 {
	ret, _ := com.dm.CallMethod("GetModuleBaseAddr", hwnd, module)
	return ret.Val
}

func (com *Dmsoft) GetModuleSize(hwnd int, module string) int {
	ret, _ := com.dm.CallMethod("GetModuleSize", hwnd, module)
	return int(ret.Val)
}

func (com *Dmsoft) GetRemoteApiAddress(hwnd int, baseAddr int64, funName string) int64 {
	ret, _ := com.dm.CallMethod("GetRemoteApiAddress", hwnd, baseAddr, funName)
	return ret.Val
}

func (com *Dmsoft) Int64ToInt32(value int64) int {
	ret, _ := com.dm.CallMethod("Int64ToInt32", value)
	return int(ret.Val)
}

func (com *Dmsoft) IntToData(value int64, type_ int) string {
	ret, _ := com.dm.CallMethod("IntToData", value, type_)
	return ret.ToString()
}

func (com *Dmsoft) OpenProcess(pid int) int {
	ret, _ := com.dm.CallMethod("OpenProcess", pid)
	return int(ret.Val)
}

func (com *Dmsoft) ReadData(hwnd int, addr string, length int) string {
	ret, _ := com.dm.CallMethod("ReadData", hwnd, addr, length)
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddr(hwnd int, addr int64, length int) string {
	ret, _ := com.dm.CallMethod("ReadDataAddr", hwnd, addr, length)
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddrToBin(hwnd int, addr int64, length int) int {
	ret, _ := com.dm.CallMethod("ReadDataAddrToBin", hwnd, addr, length)
	return int(ret.Val)
}

func (com *Dmsoft) ReadDataToBin(hwnd int, addr string, length int) int {
	ret, _ := com.dm.CallMethod("ReadDataToBin", hwnd, addr, length)
	return int(ret.Val)
}

func (com *Dmsoft) ReadDouble(hwnd int, addr string) float64 {
	ret, _ := com.dm.CallMethod("ReadDouble", hwnd, addr)
	return float64(ret.Val)
}

func (com *Dmsoft) ReadDoubleAddr(hwnd int, addr int64) float64 {
	ret, _ := com.dm.CallMethod("ReadDoubleAddr", hwnd, addr)
	return float64(ret.Val)
}

func (com *Dmsoft) ReadFloat(hwnd int, addr string) float32 {
	ret, _ := com.dm.CallMethod("ReadFloat", hwnd, addr)
	return float32(ret.Val)
}

func (com *Dmsoft) ReadFloatAddr(hwnd int, addr int64) float32 {
	ret, _ := com.dm.CallMethod("ReadFloatAddr", hwnd, addr)
	return float32(ret.Val)
}

func (com *Dmsoft) ReadStringAddr(hwnd int, addr int64, type_, len int) string {
	ret, _ := com.dm.CallMethod("ReadStringAddr", hwnd, addr, type_, len)
	return ret.ToString()
}

func (com *Dmsoft) SetMemoryFindResultToFile(file string) int {
	ret, _ := com.dm.CallMethod("SetMemoryFindResultToFile", file)
	return int(ret.Val)
}

func (com *Dmsoft) SetMemoryHwndAsProcessId(en int) int {
	ret, _ := com.dm.CallMethod("SetMemoryHwndAsProcessId", en)
	return int(ret.Val)
}

func (com *Dmsoft) SetParam64ToPointer() int {
	ret, _ := com.dm.CallMethod("SetParam64ToPointer")
	return int(ret.Val)
}

func (com *Dmsoft) StringToData(value string, type_ int) string {
	ret, _ := com.dm.CallMethod("StringToData", value, type_)
	return ret.ToString()
}

func (com *Dmsoft) TerminateProcess(pid int) int {
	ret, _ := com.dm.CallMethod("TerminateProcess", pid)
	return int(ret.Val)
}

func (com *Dmsoft) TerminateProcessTree(pid int) int {
	ret, _ := com.dm.CallMethod("TerminateProcessTree", pid)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualProtectEx(hwnd int, addr int64, size, type_, oldProtect int) int {
	ret, _ := com.dm.CallMethod("VirtualProtectEx", hwnd, addr, size, type_, oldProtect)
	return int(ret.Val)
}

func (com *Dmsoft) VirtualQueryEx(hwnd int, addr int64, pmbi int) string {
	ret, _ := com.dm.CallMethod("VirtualQueryEx", hwnd, addr, pmbi)
	return ret.ToString()
}

func (com *Dmsoft) WriteDataAddrFromBin(hwnd int, addr int64, data int, length int) int {
	ret, _ := com.dm.CallMethod("WriteDataAddrFromBin", hwnd, addr, data, length)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDataFromBin(hwnd int, addr string, data int, length int) int {
	ret, _ := com.dm.CallMethod("WriteDataFromBin", hwnd, addr, data, length)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDouble(hwnd int, addr string, value float64) int {
	ret, _ := com.dm.CallMethod("WriteDouble", hwnd, addr, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteDoubleAddr(hwnd int, addr int64, value float64) int {
	ret, _ := com.dm.CallMethod("WriteDoubleAddr", hwnd, addr, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteFloat(hwnd int, addr string, value float32) int {
	ret, _ := com.dm.CallMethod("WriteFloat", hwnd, addr, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteFloatAddr(hwnd int, addr int64, value float32) int {
	ret, _ := com.dm.CallMethod("WriteFloatAddr", hwnd, addr, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteInt(hwnd int, addr string, type_ int, value int64) int {
	ret, _ := com.dm.CallMethod("WriteInt", hwnd, addr, type_, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteIntAddr(hwnd int, addr int64, type_ int, value int64) int {
	ret, _ := com.dm.CallMethod("WriteIntAddr", hwnd, addr, type_, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteString(hwnd int, addr string, type_ int, value string) int {
	ret, _ := com.dm.CallMethod("WriteString", hwnd, addr, type_, value)
	return int(ret.Val)
}

func (com *Dmsoft) WriteStringAddr(hwnd int, addr int64, type_ int, value string) int {
	ret, _ := com.dm.CallMethod("WriteStringAddr", hwnd, addr, type_, value)
	return int(ret.Val)
}
