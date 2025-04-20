package spoofer

import (
    "fmt"
    "strings"
    "syscall"
    "unsafe"

    "golang.org/x/sys/windows"
)

type StartupInfoEx struct {
    windows.StartupInfo
    AttributeList *PROC_THREAD_ATTRIBUTE_LIST
}

type PROC_THREAD_ATTRIBUTE_LIST struct {
    dwFlags  uint32
    size     uint64
    count    uint64
    reserved uint64
    unknown  *uint64
    entries  []*PROC_THREAD_ATTRIBUTE_ENTRY
}

type PROC_THREAD_ATTRIBUTE_ENTRY struct {
    attribute *uint32
    cbSize    uintptr
    lpValue   uintptr
}

type ProcessEntry32 struct {
    Size              uint32
    Usage             uint32
    ProcessID         uint32
    DefaultHeapID     uintptr
    ModuleID          uint32
    Threads           uint32
    ParentProcessID   uint32
    PriClassBase      int32
    Flags             uint32
    ExeFile          [windows.MAX_PATH]uint16
}

func FindProcessByName(name string) (uint32, error) {
    snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
    if err != nil {
        return 0, fmt.Errorf("impossible de créer un snapshot des processus: %v", err)
    }
    defer windows.CloseHandle(snapshot)

    var pe ProcessEntry32
    pe.Size = uint32(unsafe.Sizeof(pe))

    if err := windows.Process32First(snapshot, (*windows.ProcessEntry32)(&pe)); err != nil {
        return 0, fmt.Errorf("Process32First failed: %v", err)
    }

    name = strings.ToLower(name)
    for {
        exeName := windows.UTF16ToString(pe.ExeFile[:])
        if strings.ToLower(exeName) == name {
            return pe.ProcessID, nil
        }

        err = windows.Process32Next(snapshot, (*windows.ProcessEntry32)(&pe))
        if err != nil {
            if err == windows.ERROR_NO_MORE_FILES {
                break
            }
            return 0, fmt.Errorf("Process32Next failed: %v", err)
        }
    }

    return 0, fmt.Errorf("processus '%s' non trouvé", name)
}

func SpawnWithParentName(cmd string, parentName string) error {
    ppid, err := FindProcessByName(parentName)
    if err != nil {
        return fmt.Errorf("impossible de trouver le processus parent: %v", err)
    }

    kernel32 := windows.NewLazySystemDLL("kernel32.dll")
    initializeProcThreadAttributeList := kernel32.NewProc("InitializeProcThreadAttributeList")
    updateProcThreadAttribute := kernel32.NewProc("UpdateProcThreadAttribute")
    deleteProcThreadAttributeList := kernel32.NewProc("DeleteProcThreadAttributeList")


    var size uintptr
    ret, _, _ := initializeProcThreadAttributeList.Call(
        0,
        1,
        0,
        uintptr(unsafe.Pointer(&size)),
    )
    if ret != 0 {
        return fmt.Errorf("InitializeProcThreadAttributeList devrait échouer avec ERROR_INSUFFICIENT_BUFFER")
    }
    
    // allocate mem
    attributeList := make([]byte, size)
    startupInfoEx := &StartupInfoEx{}
    startupInfoEx.AttributeList = (*PROC_THREAD_ATTRIBUTE_LIST)(unsafe.Pointer(&attributeList[0]))
    
    // Initialyze attributes
    ret, _, err = initializeProcThreadAttributeList.Call(
        uintptr(unsafe.Pointer(startupInfoEx.AttributeList)),
        1,
        0,
        uintptr(unsafe.Pointer(&size)),
    )
    if ret == 0 {
        return fmt.Errorf("InitializeProcThreadAttributeList a échoué: %v", err)
    }

    parentHandle, err := windows.OpenProcess(windows.PROCESS_CREATE_PROCESS, false, ppid)
    if err != nil {
        return fmt.Errorf("OpenProcess a échoué: %v", err)
    }
    defer windows.CloseHandle(parentHandle)

    // update handle list
    ret, _, err = updateProcThreadAttribute.Call(
        uintptr(unsafe.Pointer(startupInfoEx.AttributeList)),
        0,
        0x00020000, // PROC_THREAD_ATTRIBUTE_PARENT_PROCESS
        uintptr(unsafe.Pointer(&parentHandle)),
        unsafe.Sizeof(parentHandle),
        0,
        0,
    )
    if ret == 0 {
        return fmt.Errorf("UpdateProcThreadAttribute a échoué: %v", err)
    }

    startupInfoEx.Cb = uint32(unsafe.Sizeof(*startupInfoEx))
    
    var procInfo windows.ProcessInformation
    creationFlags := windows.CREATE_NEW_CONSOLE | 0x00080000 // EXTENDED_STARTUPINFO_PRESENT

    err = windows.CreateProcess(
        nil,
        syscall.StringToUTF16Ptr(cmd),
        nil,
        nil,
        false,
        uint32(creationFlags),
        nil,
        nil,
        &startupInfoEx.StartupInfo,
        &procInfo,
    )
    
    // cleanup
    deleteProcThreadAttributeList.Call(uintptr(unsafe.Pointer(startupInfoEx.AttributeList)))
    
    if err != nil {
        return fmt.Errorf("CreateProcess a échoué: %v", err)
    }

    windows.CloseHandle(procInfo.Thread)
    windows.CloseHandle(procInfo.Process)
    
    return nil
}
