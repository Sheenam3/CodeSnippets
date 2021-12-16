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
	TableData []SingleTable
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

	entry := SingleTable{Name: name, Pid: pid, HostPid: "hostpid", Status: "status", MarkStatus: "NM"}
	data.TableData = append(data.TableData, entry)
	//module.Gtable.TableData = data.TableData
	m.Unlock()
	fmt.Println("Table-------", data.TableData)

}

/*Output- NM didn't change to M
Table------- [{sheenam 1 hostpid status NM}]
Table------- [{sheenam 1 hostpid status NM} {Krishna 2 hostpid status NM}]
Table------- [{sheenam 1 hostpid status NM} {Krishna 2 hostpid status NM} {red 3 hostpid status NM}]
Table------- [{sheenam 1 hostpid status NM} {Krishna 2 hostpid status NM} {red 3 hostpid status NM} {yellow 5 hostpid status NM}]
Table------- [{sheenam 1 hostpid status NM} {Krishna 2 hostpid status NM} {red 3 hostpid status NM} {yellow 5 hostpid status NM} {hello 4 hostpid status NM}]
0 {sheenam 1 hostpid status NM}
1 {Krishna 2 hostpid status NM}
2 {red 3 hostpid status NM}
3 {yellow 5 hostpid status NM}
4 {hello 4 hostpid status NM}*/
