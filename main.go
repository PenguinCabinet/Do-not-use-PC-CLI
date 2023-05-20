package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var start_time int
var end_time int

func load_setting() {
	RunPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	RunDirPath := filepath.Dir(RunPath)

	f, err := os.Open(filepath.Join(RunDirPath, "setting.txt"))
	if err != nil {
		log.Fatalln("setting.txt does not exist.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	temp := ""
	temp1 := 0
	temp2 := 0

	scanner.Scan()
	temp = scanner.Text()
	fmt.Sscanf(temp, "%d:%d", &temp1, &temp2)
	start_time = 60*temp1 + temp2

	scanner.Scan()
	temp = scanner.Text()
	fmt.Sscanf(temp, "%d:%d", &temp1, &temp2)
	end_time = 60*temp1 + temp2

	if err = scanner.Err(); err != nil {
		log.Fatalln("Error reading setting.txt.")
	}
}

func main() {
	shutdown_running := false
	load_setting()
	for {
		now := time.Now()
		now_temp := 60*now.Hour() + now.Minute()
		if !shutdown_running {
			if start_time > end_time {
				if start_time <= now_temp || now_temp <= end_time {
					exec.Command("shutdown", "/s", "/f").Run()
					shutdown_running = true
				}
			} else {
				if start_time <= now_temp && now_temp <= end_time {
					exec.Command("shutdown", "/s", "/f").Run()
					shutdown_running = true
				}
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}
