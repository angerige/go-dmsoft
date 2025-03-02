package dmsoft

// 文件 相关函数

func (com *Dmsoft) CopyFile(srcFile, dstFile string, over int) int {
	ret, _ := com.dm.CallMethod("CopyFile", srcFile, dstFile, over)
	return int(ret.Val)
}

func (com *Dmsoft) CreateFolder(folder string) int {
	ret, _ := com.dm.CallMethod("CreateFolder", folder)
	return int(ret.Val)
}

func (com *Dmsoft) DecodeFile(file, pwd string) int {
	ret, _ := com.dm.CallMethod("DecodeFile", file, pwd)
	return int(ret.Val)
}

func (com *Dmsoft) DeleteFile(file string) int {
	ret, _ := com.dm.CallMethod("DeleteFile", file)
	return int(ret.Val)
}

func (com *Dmsoft) DeleteFolder(folder string) int {
	ret, _ := com.dm.CallMethod("DeleteFolder", folder)
	return int(ret.Val)
}

func (com *Dmsoft) DeleteIni(section, key, file string) int {
	ret, _ := com.dm.CallMethod("DeleteIni", section, key, file)
	return int(ret.Val)
}

func (com *Dmsoft) DeleteIniPwd(section, key, file, pwd string) int {
	ret, _ := com.dm.CallMethod("DeleteIniPwd", section, key, file, pwd)
	return int(ret.Val)
}

func (com *Dmsoft) DownloadFile(url, saveFile string, timeout int) int {
	ret, _ := com.dm.CallMethod("DownloadFile", url, saveFile, timeout)
	return int(ret.Val)
}

func (com *Dmsoft) EncodeFile(file, pwd string) int {
	ret, _ := com.dm.CallMethod("EncodeFile", file, pwd)
	return int(ret.Val)
}

func (com *Dmsoft) EnumIniKey(section, file string) string {
	ret, _ := com.dm.CallMethod("EnumIniKey", section, file)
	return ret.ToString()
}

func (com *Dmsoft) EnumIniKeyPwd(section, file, pwd string) string {
	ret, _ := com.dm.CallMethod("EnumIniKeyPwd", section, file, pwd)
	return ret.ToString()
}

func (com *Dmsoft) EnumIniSection(file string) string {
	ret, _ := com.dm.CallMethod("EnumIniSection", file)
	return ret.ToString()
}

func (com *Dmsoft) EnumIniSectionPwd(file, pwd string) string {
	ret, _ := com.dm.CallMethod("EnumIniSectionPwd", file, pwd)
	return ret.ToString()
}

func (com *Dmsoft) GetFileLength(file string) int {
	ret, _ := com.dm.CallMethod("GetFileLength", file)
	return int(ret.Val)
}

func (com *Dmsoft) GetRealPath(path string) string {
	ret, _ := com.dm.CallMethod("GetRealPath", path)
	return ret.ToString()
}

func (com *Dmsoft) IsFileExist(file string) int {
	ret, _ := com.dm.CallMethod("IsFileExist", file)
	return int(ret.Val)
}

func (com *Dmsoft) IsFolderExist(folder string) int {
	ret, _ := com.dm.CallMethod("IsFolderExist", folder)
	return int(ret.Val)
}

func (com *Dmsoft) MoveFile(srcFile, dstFile string) int {
	ret, _ := com.dm.CallMethod("MoveFile", srcFile, dstFile)
	return int(ret.Val)
}

func (com *Dmsoft) ReadFile(file string) string {
	ret, _ := com.dm.CallMethod("ReadFile", file)
	return ret.ToString()
}

func (com *Dmsoft) ReadIni(section, key, file string) string {
	ret, _ := com.dm.CallMethod("ReadIni", section, key, file)
	return ret.ToString()
}

func (com *Dmsoft) ReadIniPwd(section, key, file, pwd string) string {
	ret, _ := com.dm.CallMethod("ReadIniPwd", section, key, file, pwd)
	return ret.ToString()
}

func (com *Dmsoft) SelectDirectory() string {
	ret, _ := com.dm.CallMethod("SelectDirectory")
	return ret.ToString()
}

func (com *Dmsoft) SelectFile() string {
	ret, _ := com.dm.CallMethod("SelectFile")
	return ret.ToString()
}

func (com *Dmsoft) WriteFile(file, content string) int {
	ret, _ := com.dm.CallMethod("WriteFile", file, content)
	return int(ret.Val)
}

func (com *Dmsoft) WriteIni(section, key, value, file string) int {
	ret, _ := com.dm.CallMethod("WriteIni", section, key, value, file)
	return int(ret.Val)
}

func (com *Dmsoft) WriteIniPwd(section, key, value, file, pwd string) int {
	ret, _ := com.dm.CallMethod("WriteIniPwd", section, key, value, file, pwd)
	return int(ret.Val)
}
