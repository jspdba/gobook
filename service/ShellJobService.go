package service

import (
	"encoding/json"
	"os"
	"path"
	"gobook/models"
	"github.com/robfig/cron"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
	"strings"
)

var (
	config []*models.ShellJob
)

func init() {
	LoadConfig("")
	c := cron.NewWithLocation(time.FixedZone("Shanghai",800))

	for _,shell :=range config{
		c.AddFunc(shell.Cron, func() {
			if shell.Cmd==""{
				return
			}
			cmd := exec.Command("/bin/sh", "-c", strings.Replace(shell.Cmd,"{time}",time.Now().Format("2006-01-02_15-04-05"),-1))
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Println("StdoutPipe: " + err.Error())
				return
			}

			stderr, err := cmd.StderrPipe()
			if err != nil {
				fmt.Println("StderrPipe: ", err.Error())
				return
			}

			if err := cmd.Start(); err != nil {
				fmt.Println("Start: ", err.Error())
				return
			}

			bytesErr, err := ioutil.ReadAll(stderr)
			if err != nil {
				fmt.Println("ReadAll stderr: ", err.Error())
				return
			}

			if len(bytesErr) != 0 {
				fmt.Printf("stderr is not nil: %s", bytesErr)
				return
			}

			bytes, err := ioutil.ReadAll(stdout)
			if err != nil {
				fmt.Println("ReadAll stdout: ", err.Error())
				return
			}

			if err := cmd.Wait(); err != nil {
				fmt.Println("Wait: ", err.Error())
				return
			}

			fmt.Printf("stdout: %s", bytes)
		})
	}
	c.Start()
}

func ListShellJobs() []*models.ShellJob{
	return config
}

func LoadConfig(cfg string){
	if cfg==""{
		dir,e:=os.Getwd()
		if e!=nil{
			return
		}
		cfg=path.Join(dir,"conf/shell_jobs.json")
	}

	fd, err := os.Open(cfg)
	if err != nil {
		panic("无法打开配置文件 shell_jobs.json: " + err.Error())
	}
	defer fd.Close()

	err = json.NewDecoder(fd).Decode(&config)
	if err != nil {
		panic("解析配置文件出错: " + err.Error())
	}
}