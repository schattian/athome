package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/athomecomar/athome/backend/users/userconf"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("action not given (up|down|drop|force|goto)")
	}
	action := os.Args[1]

	var opt string
	if len(os.Args) > 2 {
		opt = os.Args[2]
	}
	cmd := exec.Command("bin/migrate", userconf.GetDATABASE_SRC(), action, opt)
	stdout, err := cmd.Output()
	log.Println(action)
	if err != nil {
		log.Fatalf("cmd.Run: %v", err)
	}
	log.Println(stdout)
}
