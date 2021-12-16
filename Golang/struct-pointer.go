package main

import (
	"fmt"
	"sync"
)

type SingleTable struct {
	Name       string
	Pid        int
	HostPid    string
	Status     string
	MarkStatus string
}

var m sync.Mutex

type Data Table

var data Data

type Table struct {
	TableData []*SingleTable
}

func main() {

	addToTable("sheenam", 1, "kuch", &m)
	addToTable("Krishna", 2, "kuch", &m)
	addToTable("red", 3, "kuch", &m)
	addToTable("yellow", 5, "kuch", &m)
	addToTable("hello", 4, "kuch", &m)

	MarkTabledata("sheenam")
	for k, v := range data.TableData {
		fmt.Println(k, v)
	}
}

func MarkTabledata(funcname string) {

	for _, v := range data.TableData {
		if v.Name == funcname {
      //this status will be reflected in table data because its pointer(reference)
			v.MarkStatus = "M"

		}

	}

}

func deleteProcessEntry(pid int) {

	for k, v := range data.TableData {
		fmt.Println("v.Pid:", v.Pid)
		fmt.Println(pid)
		if v.Pid == pid {
			data.TableData = append(data.TableData[:k], data.TableData[k+1:]...)

		}
	}

}

func addToTable(name string, pid int, process string, m *sync.Mutex) {
	m.Lock()
	entry := &SingleTable{Name: name, Pid: pid, HostPid: "hostpid", Status: "status", MarkStatus: "NM"}
	data.TableData = append(data.TableData, entry)
	m.Unlock()
	fmt.Println("Table-------", data.TableData)

}
