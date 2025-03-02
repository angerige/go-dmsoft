package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/angerige/go-dmsoft"
)

const (
	// 大漠插件注册码
	DmRegCode   = "XXXXX" // 示例注册码
	DmExtraCode = "XXXXX" // 示例附加码
)

// 全局变量定义
var (
	globalDm   *dmsoft.Dmsoft   // 全局大漠对象
	dmColl     []*dmsoft.Dmsoft // 多个大漠对象集合
	handleColl []int            // 窗口句柄集合
)

func main() {
	var err error
	// 创建全局大漠对象
	globalDm, err = dmsoft.CreateDm(&dmsoft.DmConfig{
		DllPath:   "E:\\yx_go\\go-dmsoft\\example\\multi_window\\dll", // 使用绝对路径
		RegCode:   DmRegCode,
		ExtraCode: DmExtraCode,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer globalDm.Release()

	log.Printf("插件版本:%s", globalDm.Ver())

	EnumWindowHwnd() // 枚举所有窗口句柄
	BindMultipleWindowsAsync()
}

// BindMultipleWindowsAsync 异步绑定多个窗口
func BindMultipleWindowsAsync() {
	res := make(chan int) // 创建结果通道
	defer close(res)

	// 为每个窗口创建一个大漠对象和工作线程
	for i := 0; i < len(handleColl); i++ {
		// 使用相同的配置创建新对象
		dm, err := dmsoft.CreateDm(&dmsoft.DmConfig{
			DllPath:   "E:\\yx_go\\go-dmsoft\\example\\multi_window\\dll",
			RegCode:   DmRegCode,   // 添加注册码
			ExtraCode: DmExtraCode, // 添加附加码
		})
		if err != nil {
			log.Printf("创建大漠对象失败: %v", err)
			continue
		}
		dmColl = append(dmColl, dm)
		go BindMultipleWindows(dmColl[len(dmColl)-1], handleColl[i], res)
	}

	// 等待所有线程完成并输出结果
	for i := 0; i < len(dmColl); i++ {
		v := <-res
		log.Printf("窗口句柄:%d,截图结果:%v", handleColl[i], v != 0)
	}
}

// BindMultipleWindows 绑定多个窗口
func BindMultipleWindows(dm *dmsoft.Dmsoft, handle int, result chan<- int) {
	// 绑定窗口
	ret := dm.BindWindowEx(handle, "dx2", "windows3", "windows", "", 0)
	if ret != 0 {
		log.Printf("绑定窗口成功,句柄:%d", handle)
	}
	// 截图
	r1 := dm.Capture(0, 0, 2000, 2000, "test_"+strconv.Itoa(handle)+".bmp")
	result <- r1 // 将结果发送到 channel

	log.Printf("释放大漠对象:%v,窗口句柄:%d", dm, handle)
	dm.UnBindWindow() // 解除绑定
	dm.Release()
}

// EnumWindowHwnd 枚举出逍遥模拟器的窗口句柄
func EnumWindowHwnd() {
	handleStr := globalDm.EnumWindow(0, "*无标题 - 记事本", "Notepad", 1+2)
	handleStrSlice := strings.Split(handleStr, ",")
	// 将string类型转换为int类型
	for _, str := range handleStrSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}
		handleColl = append(handleColl, num)
	}
	log.Printf("获取到 %d 个窗口句柄:%v", len(handleColl), handleColl)
}
