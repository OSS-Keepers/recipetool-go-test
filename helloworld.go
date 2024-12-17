package main

import (
    "fmt"
    "os"
    "testing"

    "github.com/matryer/is"
    "github.com/godbus/dbus/v5"
)

func main() {
    // dbus and os dependency usage
    conn, err := dbus.ConnectSessionBus()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // print hello world
    fmt.Println("Hello World!")
}

// "is" dependency usage
func TestExample(t *testing.T) {
    is := is.New(t)

    is.True(true) // This test passes
    is.Equal(1, 1) // This test passes
    is.NoErr(nil) // No error
}
