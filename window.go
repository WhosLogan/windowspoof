package main

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

type Window struct {
	PID        uint   `json:"pid"`
	Name       string `json:"name"`
	WindowName string `json:"windowName"`
}

var (
	user32                       = syscall.MustLoadDLL("user32.dll")
	kernel32                     = syscall.MustLoadDLL("Kernel32.dll")
	psapi                        = syscall.MustLoadDLL("Psapi.dll")
	procEnumWindows              = user32.MustFindProc("EnumWindows")
	procGetWindowTextW           = user32.MustFindProc("GetWindowTextW")
	procIsWindowVisible          = user32.MustFindProc("IsWindowVisible")
	procGetWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	procOpenProcess              = kernel32.MustFindProc("OpenProcess")
	procGetModuleFileNameEx      = psapi.MustFindProc("GetModuleFileNameExW")
	procCloseHandle              = kernel32.MustFindProc("CloseHandle")
	windowCallback               = syscall.NewCallback(enumWindowCallback)
)

func GetWindows() []Window {
	var windows []Window
	_, _, _ = procEnumWindows.Call(windowCallback, uintptr(unsafe.Pointer(&windows)))
	return windows
}

func enumWindowCallback(hWindow uintptr, windows *[]Window) uintptr {
	buffer := make([]uint16, 256)
	_, _, _ = procGetWindowTextW.Call(hWindow, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	visible, _, _ := procIsWindowVisible.Call(hWindow)
	if isZero(buffer) || visible == 0 {
		return 1
	}
	title := syscall.UTF16ToString(buffer)
	pid := uint32(0)
	threadId, _, _ := procGetWindowThreadProcessId.Call(hWindow, uintptr(unsafe.Pointer(&pid)))
	if threadId == 0 {
		return 1
	}
	hProc, _, _ := procOpenProcess.Call(0x0400|0x0010, 0, uintptr(pid))
	if hProc == 0 {
		return 1
	}
	defer closeHandle(hProc)
	imageNameSize, _, _ := procGetModuleFileNameEx.Call(hProc, 0, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	if imageNameSize == 0 {
		return 1
	}
	imageName := syscall.UTF16ToString(buffer)
	*windows = append(*windows, Window{
		PID:        uint(pid),
		Name:       filepath.Base(imageName),
		WindowName: title,
	})
	return 1
}

func closeHandle(handle uintptr) {
	_, _, _ = procCloseHandle.Call(handle)
}

func isZero(buffer []uint16) bool {
	for _, b := range buffer {
		if b != 0 {
			return false
		}
	}
	return true
}
