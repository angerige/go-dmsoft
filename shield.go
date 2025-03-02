package dmsoft

// 防护盾 相关函数

func (com *Dmsoft) DmGuard(value int, type_ string) int {
	ret, _ := com.dm.CallMethod("DmGuard", value, type_)
	return int(ret.Val)
}

func (com *Dmsoft) DmGuardExtract(type_, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardExtract", type_, path)
	return int(ret.Val)
}

func (com *Dmsoft) DmGuardLoadCustom(type_, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardLoadCustom", type_, path)
	return int(ret.Val)
}

func (com *Dmsoft) DmGuardParams(cmd, subcmd, params string) string {
	ret, _ := com.dm.CallMethod("DmGuardParams", cmd, subcmd, params)
	return ret.ToString()
}

func (com *Dmsoft) UnLoadDriver() int {
	ret, _ := com.dm.CallMethod("UnLoadDriver")
	return int(ret.Val)
}
