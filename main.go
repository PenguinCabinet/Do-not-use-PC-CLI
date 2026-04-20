package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"runtime"
	"gopkg.in/yaml.v3"
)

type HM_t struct {
	Hours   int `yaml:"hours"`
	Minutes int `yaml:"minutes"`
}

type Setting_t struct {
	Rules []struct {
		If struct {
			Weeks []string `yaml:"weeks"`
		} `yaml:"if"`
		Apply struct {
			Allow_times []struct {
				Start HM_t `yaml:"start"`
				End   HM_t `yaml:"end"`
			} `yaml:"allowtimes"`
		} `yaml:"apply"`
	} `yaml:"rules"`
}

var start_time1 int
var end_time1 int
var start_time2 int
var end_time2 int
var setting Setting_t

var weeks_map = map[string]time.Weekday{
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
	"Sun": time.Sunday,
}


func get_HOME_path()string{
	switch runtime.GOOS {
	case "windows":
		home := os.Getenv("USERPROFILE")
		if home != "" {
			return home
		}
		return os.Getenv("HOME")

	default:
		// Linux,macOS
		home := os.Getenv("HOME")
		if home != "" {
			return home
		}
		return os.Getenv("USERPROFILE")
	}
}

func generate_setting_path1()string{
	home_path:=get_HOME_path()
	return filepath.Join(home_path, ".Do-not-use-PC.yaml")
}

func generate_setting_path2()string{
	RunPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	RunDirPath := filepath.Dir(RunPath)

	return filepath.Join(RunDirPath, "setting.yaml")
}

func poweroff(){
	switch runtime.GOOS {
	case "windows":
		err:=exec.Command("shutdown", "/s", "/f").Run()
		if err!=nil{
			panic(err)
		}
	case "linux":
		err:=exec.Command("systemctl", "poweroff").Run()
		if err!=nil{
			panic(err)
		}
	default:
		panic("Not supporting OS")	
	}
}
func load_setting() {
	path1:=generate_setting_path1()
	path2:=generate_setting_path2()
	path:=""

	_, err := os.Stat(path1)
	if err == nil {
		path=path1	
	}else{
		_, err := os.Stat(path2)
		if err == nil {
			path=path2
		}else{
			panic("Noting setting file.")
		}
	}

	setting_str, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal([]byte(setting_str), &setting)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

func main() {
	shutdown_running := false

	start_time1 = 60*21 + 0
	end_time1 = 60*6 + 0

	start_time2 = 60*22 + 45
	end_time2 = 60*6 + 0

	load_setting()

	for {
		now := time.Now()
		now_temp := 60*now.Hour() + now.Minute()
		if !shutdown_running {
			(func() {
				for _, e := range setting.Rules {
					dofunc := func() bool {
						flag := true
						for _, e2 := range e.Apply.Allow_times {
							if 60*e2.Start.Hours+e2.Start.Minutes <= now_temp && now_temp <= 60*e2.End.Hours+e2.End.Minutes {
								flag = false
							}
						}
						return flag
					}

					for _, e2 := range e.If.Weeks {
						if e2 == "All" || now.Weekday() == weeks_map[e2] {
							if dofunc() {
								poweroff()

								shutdown_running = true
								return
							}
						}
					}
				}
			})()

		}
		time.Sleep(time.Millisecond * 500)
	}
}
