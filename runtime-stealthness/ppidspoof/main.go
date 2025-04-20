package main

import (
    "fmt"
    "os"

    "ppidspoof/spoofer"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: ppidspoof.exe <command> <parent_process_name>")
        fmt.Println("Example: example.exe \"cmd.exe\" explorer.exe")
        os.Exit(1)
    }

    cmd := os.Args[1]
    parentName := os.Args[2]

    err := spoofer.SpawnWithParentName(cmd, parentName)
    if err != nil {
        fmt.Printf("error:", err)
        os.Exit(1)
    }

}
