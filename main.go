package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	"Tue": time.Thursday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
	"Sun": time.Sunday,
}

/*
func load_setting() {
	RunPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	RunDirPath := filepath.Dir(RunPath)

	f, err := os.Open(filepath.Join(RunDirPath, "setting.lua"))
	if err != nil {
		log.Fatalln("setting.lua does not exist.")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	setting_str = scanner.Text()
}
*/

func load_setting() {
	RunPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	RunDirPath := filepath.Dir(RunPath)

	setting_str, err := ioutil.ReadFile(filepath.Join(RunDirPath, "setting.yaml"))
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(setting_str)

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
								//fmt.Println("shutdown!!!")
								exec.Command("shutdown", "/s", "/f").Run()
								shutdown_running = true
								return
							}
						}
					}
				}
			})()

		}
		/*
			if now.Weekday() == time.Saturday {
				if start_time2 > end_time2 {
					if start_time2 <= now_temp || now_temp <= end_time2 {
						exec.Command("shutdown", "/s", "/f").Run()
						shutdown_running = true
					}
				} else {
					if start_time2 <= now_temp && now_temp <= end_time2 {
						exec.Command("shutdown", "/s", "/f").Run()
						shutdown_running = true
					}
				}
			} else {
				if start_time1 > end_time1 {
					if start_time1 <= now_temp || now_temp <= end_time1 {
						exec.Command("shutdown", "/s", "/f").Run()
						shutdown_running = true
					}
				} else {
					if start_time1 <= now_temp && now_temp <= end_time1 {
						exec.Command("shutdown", "/s", "/f").Run()
						shutdown_running = true
					}
				}
			}*/
		time.Sleep(time.Millisecond * 500)
	}
}
