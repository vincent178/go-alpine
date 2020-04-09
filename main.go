package main

import (
	"fmt"
	"os/exec"
)

func main() {
	args := []string{"filter", "add", "dev", "eth0", "parent", "1:0", "basic", "match", "ipset(netwo_1029cfdb8807de284_tgt dst)", "classid", "1:1"}
	cmd := exec.Command("tc", args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(out))
}
