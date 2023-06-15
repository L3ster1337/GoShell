package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func executarTrojan() {
	fmt.Println("Executando Trojan...")

	cmd := exec.Command("/path/to/trojan")
	cmd.Run()
}

func ocultarTrojan() {
	fmt.Println("Ocultando Trojan...")

	cmd := exec.Command("bash", "-c", "echo 0 > /proc/self/status")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}

func main() {
	executarTrojan()
	ocultarTrojan()
}
