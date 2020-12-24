package main

/*
#include<windows.h>
#include <conio.h>

int SetTitle()
{
	SetConsoleTitle("GYJ Remote Shutdown");
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
	fmt.Println("状态: 关机执行成功！")
	command := exec.Command("cmd", "/C", "shutdown -s -t 10")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func termination(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 终止执行成功！")
	command := exec.Command("cmd", "/C", "shutdown -a")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func restartcc(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("状态: 重启执行成功！")
	command := exec.Command("cmd", "/C", "shutdown -r -t 10")
	err := command.Run()
	if err != nil {
		//fmt.Println(err.Error())
	}

	_, _ = w.Write([]byte("success"))
}

func main() {
	C.SetTitle()
	fmt.Println("GYJ 远程关机")
	fmt.Println("\n日志:")
	fmt.Println("状态：启动成功")
	mux := http.NewServeMux()
	mux.HandleFunc("/shutdown-jcegbCFFGRWYwYq1Y4bH", shutdown)
	mux.HandleFunc("/restart-jcegbCFFGRWYwYq1Y4bH", restartcc)
	mux.HandleFunc("/stop-jcegbCFFGRWYwYq1Y4bH", termination)
	err := http.ListenAndServe(":25250", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}