package dmsoft

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// 全局变量，用于加载 DLL 和函数
var (
	dmReg32      *syscall.LazyDLL  // 延迟加载 DmReg.dll
	_SetDllPathA *syscall.LazyProc // ASCII版函数指针
	_SetDllPathW *syscall.LazyProc // Unicode版函数指针
)

// InitDll 初始化 DLL 路径
func InitDll(dmRegPath string) error {
	// 加载 DmReg.dll
	dmReg32 = syscall.NewLazyDLL(dmRegPath)

	// 获取函数指针
	_SetDllPathA = dmReg32.NewProc("SetDllPathA")
	_SetDllPathW = dmReg32.NewProc("SetDllPathW")

	return nil
}

// Dmsoft 大漠插件的主结构体
type Dmsoft struct {
	dm       *ole.IDispatch // COM 调度接口，用于调用方法
	IUnknown *ole.IUnknown  // COM 基础接口
}

// NewDmsoft 创建一个新的大漠对象实例
func NewDmsoft() *Dmsoft {
	var com Dmsoft
	var err error
	// 初始化 COM 环境
	ole.CoInitialize(0)
	// 创建 COM 对象
	com.IUnknown, err = oleutil.CreateObject("dm.dmsoft")
	if err != nil {
		panic(err)
	}
	// 获取调度接口
	com.dm, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	return &com
}

// Release 释放 COM 对象和资源
func (com *Dmsoft) Release() {
	com.IUnknown.Release()
	com.dm.Release()
	ole.CoUninitialize()
}

// SetDllPathA 设置 DLL 路径（ASCII版本）
func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathA.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

// SetDllPathW 设置 DLL 路径（Unicode版本）
func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathW.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

// DmConfig 大漠对象配置
type DmConfig struct {
	DllPath   string // DLL 目录路径
	RegCode   string // 注册码
	ExtraCode string // 附加码
}

// CreateDm 创建大漠对象的工厂方法
func CreateDm(config *DmConfig) (*Dmsoft, error) {
	// 如果没有提供DLL路径，使用默认路径（当前目录下的dll文件夹）
	if config == nil || config.DllPath == "" {
		// 获取当前工作目录
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("获取当前目录失败: %v", err)
		}

		if config == nil {
			config = &DmConfig{
				DllPath: dir + "\\dll", // 使用默认路径
			}
		} else {
			config.DllPath = dir + "\\dll" // 只更新路径
		}
	}
	// 验证必要的注册信息
	if config.RegCode == "" || config.ExtraCode == "" {
		return nil, fmt.Errorf("注册码和附加码不能为空")
	}

	// 初始化 DmReg.dll
	if err := InitDll(config.DllPath + "\\DmReg.dll"); err != nil {
		return nil, fmt.Errorf("初始化 DmReg.dll 失败: %v", err)
	}

	// 设置 dm.dll 路径
	if !SetDllPathW(config.DllPath+"\\dm.dll", 0) {
		return nil, fmt.Errorf("设置 dm.dll 路径失败")
	}

	// 创建对象
	dm := NewDmsoft()

	// 如果提供了注册码，进行注册
	if config.RegCode != "" {
		ret := dm.Reg(config.RegCode, config.ExtraCode)
		if ret == 1 {
			// 注册成功
			log.Println("VIP注册成功")
		} else {
			// 注册失败，直接返回错误码
			return nil, fmt.Errorf("注册失败: 错误码 %d", ret)
		}
	}

	return dm, nil
}
