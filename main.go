package main

/*
#include<windows.h>
#include <conio.h>

int SetTitle()
{
	SetConsoleTitle("GYJ Remote Control");
	return 0;
}
*/
import "C"
import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func shutdown(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 关机执行成功")
	command := exec.Command("cmd", "/C", "shutdown -s -t 10")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func termination(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 终止执行成功")
	command := exec.Command("cmd", "/C", "shutdown -a")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func restartcc(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 重启执行成功")
	command := exec.Command("cmd", "/C", "shutdown -r -t 10")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func lock(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 锁屏执行成功")
	command := exec.Command("cmd", "/C", "rundll32.exe user32.dll LockWorkStation")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><title>GYJ</title><h1><center><br><br><br><br><br><br>GYJ Remote Control Services<br><br><br><br><br><br></h1></center><hr><h3><center>&copy; 2021 GYJ Corporation</center></h3></html>\n")
}

func main() {
	C.SetTitle()
	fmt.Println("GYJ 远程控制服务 [版本 1.0.0.0]")
	fmt.Println("(C) 2021 GYJ Corporation. 保留所有权利。")
	fmt.Println("\n\n运行日志:")
	fmt.Println("———————————")
	fmt.Println("状态：服务启动成功")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/shutdown-jcegbCFFGRWYwYq1Y4bH", shutdown)
	mux.HandleFunc("/restart-jcegbCFFGRWYwYq1Y4bH", restartcc)
	mux.HandleFunc("/stop-jcegbCFFGRWYwYq1Y4bH", termination)
	mux.HandleFunc("/lock-jcegbCFFGRWYwYq1Y4bH", lock)
	err := http.ListenAndServe(":25250", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
