package main

import (
	"log"
	"os"
	"unsafe"

	"github.com/angerige/go-dmsoft"
)

const (
	// 填写自己的注册码
	DmRegCode   = "" // 示例注册码
	DmExtraCode = "" // 示例附加码
)

type _bytes struct {
	Data int
	Len  int
}

func main() {
	// 方式1：只提供注册信息，使用默认dll路径
	dm, err := dmsoft.CreateDm(&dmsoft.DmConfig{
		RegCode:   DmRegCode,   // 必须提供
		ExtraCode: DmExtraCode, // 必须提供
	})
	if err != nil {
		log.Fatal(err)
	}
	defer dm.Release()

	// 方式2：提供完整配置
	dm, err = dmsoft.CreateDm(&dmsoft.DmConfig{
		DllPath:   "E:\\yx_go\\go-dmsoft\\example\\basic\\assets", // 可选，指定路径
		RegCode:   DmRegCode,                                      // 必须提供
		ExtraCode: DmExtraCode,                                    // 必须提供
	})
	if err != nil {
		log.Fatal(err)
	}
	defer dm.Release()

	// 开始使用...
	log.Printf("插件版本:%s", dm.Ver())

	var x, y int
	ret := dm.GetCursorPos(&x, &y)
	log.Printf("鼠标位置: %d, %d, %d", x, y, ret)

	var data, size int
	dm.GetScreenDataBmp(0, 0, 800, 800, &data, &size)
	bs := *(*[]byte)(unsafe.Pointer(&_bytes{
		data,
		size,
	}))
	os.WriteFile("test.bmp", bs, os.ModePerm)
}

// func CreateDmObj() *dmsoft.Dmsoft {
// 	// 获取当前工作目录
// 	dir, _ := os.Getwd()
// 	dllDir := dir + "\\dll"

// 	// 初始化 DmReg.dll
// 	if err := dmsoft.InitDll(dllDir + "\\DmReg.dll"); err != nil {
// 		log.Fatal("初始化 DmReg.dll 失败:", err)
// 	}

// 	// 设置 dm.dll 路径
// 	ret := dmsoft.SetDllPathW(dllDir+"\\dm.dll", 0)
// 	if ret {
// 		log.Println("插件注册成功！")
// 	} else {
// 		log.Println("插件注册失败！")
// 	}
// 	return dmsoft.NewDmsoft()
// }
