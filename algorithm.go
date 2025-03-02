package dmsoft

// 算法相关函数

func (com *Dmsoft) ExcludePos(allPos string, type_, x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("ExcludePos", allPos, type_, x1, y1, x2, y2)
	return ret.ToString()
}

func (com *Dmsoft) FindNearestPos(allPos string, type_, x, y int) string {
	ret, _ := com.dm.CallMethod("FindNearestPos", allPos, type_, x, y)
	return ret.ToString()
}

func (com *Dmsoft) SortPosDistance(allPos string, type_, x, y int) string {
	ret, _ := com.dm.CallMethod("SortPosDistance", allPos, type_, x, y)
	return ret.ToString()
}
