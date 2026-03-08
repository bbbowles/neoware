package main

import (
    "fmt"
    "os"
    "os/exec"
    "os/signal"
	"syscall"
)

func main() {
    // volta o terminal no modo canonico normal
    defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

    sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    // disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
    // Now there are two things happening simultaneously.
    // Goroutine 2 (signal handler)
    go func() {
		<-sig
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		fmt.Println("\nRestored terminal. Exiting.")
		os.Exit(0)
	}()

    //Goroutine 1 (main loop)
    for {
        var b []byte = make([]byte, 3)
        os.Stdin.Read(b)

        switch{
        // caso seja a tecla enter
        case b[0] == 10 && b[1] == 0 && b[2] == 0:
            fmt.Print("selecionando")
        // caso nao seja escape character
        case b[2] == 0 && b[1] == 0:
            fmt.Print("nao escape character")
        // caso seja setas
        case b[0] == 27 && b[1] == 91:
            switch{
            case b[2] == 68:
                fmt.Print("esquerda")
            case b[2] == 66:
                fmt.Print("baixo")
            case b[2] == 67:
                fmt.Print("direita")
            case b[2] == 65:
                fmt.Print("cima")    
            }
        }
        fmt.Println("I got the byte", b, "("+string(b)+")")
    }
    
}
