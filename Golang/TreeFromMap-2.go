package main

import (
	"encoding/json"
	"fmt"
)

var objdump [][]string
var funcs []string

type Nodes struct {
	Id       int
	ParentId int
	Fn       string
	Children []*Nodes
}

var Data []*Nodes
var storage = make(map[string][]string)

func main() {
	Data = []*Nodes{}
	objdump = [][]string{{"main", "fn1"}, {"main", "fn2"}, {"main", "fn3"}, {"main", "fn4"}, {"fn1", "fn5"}, {"fn2", "fn7"}, {"fn3", "fn9"}, {"fn3", "fn8"}, {"fn1", "fn6"}, {"fn5", "fn10"}, {"fn7", "fn11"}, {"fn8", "fn12"}, {"fn13", "fn14"}, {"fn8", "fn13"}, {"fn9", "fn23"}}

	funcs = []string{"fn5", "fn1", "fn2", "fn3", "fn4", "fn10", "fn11", "fn8", "fn9", "fn12", "fn13", "fn6", "fn7", "fn14"}
	//levl := getchildren("main")
	var root *Nodes = &Nodes{123, 0, "main", nil}
	iterateLevel([]string{"main"})
	for k, node := range Data {
		node.Id = k + 1

	}

	mp := make(map[string]int)
	for k, _ := range storage {

		for _, value := range Data {
			if value.Fn == k {
				mp[k] = value.Id
			}

		}

	}

	for _, v := range Data {
		for key, value := range mp {
			for k, vl := range storage {
				if k == "main" {
					for _, val := range vl {
						if v.Fn == val {
							v.ParentId = 123
						}
					}
				} else {
					if key == k {

						for _, val := range vl {
							if v.Fn == val {

								v.ParentId = value
							}
						}
					}
				}
			}
		}

	}
	/*for _, node := range Data{
		fmt.Println("l",node)
	}*/
	fmt.Println(root.Add(Data...), root.Size())
	bytes, _ := json.MarshalIndent(root, "", "\t") //formated output
	fmt.Println(string(bytes))
}

func (thi *Nodes) Size() int {

	var size int = len(thi.Children)
	for _, c := range thi.Children {
		size += c.Size()
	}
	return size
}

func (thi *Nodes) Add(nodes ...*Nodes) bool {
	var size = thi.Size()

	for _, n := range nodes {

		if n.ParentId == thi.Id {
			thi.Children = append(thi.Children, n)
		} else {
			for _, c := range thi.Children {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return thi.Size() == size+len(nodes)
}

func childexistsorno(fn []string, v string) bool {
	for _, val := range fn {
		if v == val {
			return true
		}
	}

	return false

}

func getchildren(fn string) []string {
	var children []string

	for _, v := range objdump {
		if v[0] == fn {
			yes := childexistsorno(funcs, v[1])
			if yes {
				storage[fn] = append(storage[fn], v[1])
				children = append(children, v[1])
				node := &Nodes{Fn: v[1], Children: nil}
				Data = append(Data, node)
			}

		}

	}

	return children
}

var child_at_levels [][]string

func iterateLevel(child []string) {
	var child_level []string
	for _, val := range child {

		children := getchildren(val)
		if len(children) > 0 {
			for _, chill := range children {
				//this creates all nodes in one level to iterate and search them for their child nodes
				child_level = append(child_level, chill)

			}
		}
	}

	if len(child_level) > 0 {
		child_at_levels = append(child_at_levels, [][]string{child_level}...)

		iterateLevel(child_level)
	}
}



/*Output:
true 14
{
	"Id": 123,
	"ParentId": 0,
	"Fn": "main",
	"Children": [
		{
			"Id": 1,
			"ParentId": 123,
			"Fn": "fn1",
			"Children": [
				{
					"Id": 5,
					"ParentId": 1,
					"Fn": "fn5",
					"Children": [
						{
							"Id": 10,
							"ParentId": 5,
							"Fn": "fn10",
							"Children": null
						}
					]
				},
				{
					"Id": 6,
					"ParentId": 1,
					"Fn": "fn6",
					"Children": null
				}
			]
		},
		{
			"Id": 2,
			"ParentId": 123,
			"Fn": "fn2",
			"Children": [
				{
					"Id": 7,
					"ParentId": 2,
					"Fn": "fn7",
					"Children": [
						{
							"Id": 11,
							"ParentId": 7,
							"Fn": "fn11",
							"Children": null
						}
					]
				}
			]
		},
		{
			"Id": 3,
			"ParentId": 123,
			"Fn": "fn3",
			"Children": [
				{
					"Id": 8,
					"ParentId": 3,
					"Fn": "fn9",
					"Children": null
				},
				{
					"Id": 9,
					"ParentId": 3,
					"Fn": "fn8",
					"Children": [
						{
							"Id": 12,
							"ParentId": 9,
							"Fn": "fn12",
							"Children": null
						},
						{
							"Id": 13,
							"ParentId": 9,
							"Fn": "fn13",
							"Children": [
								{
									"Id": 14,
									"ParentId": 13,
									"Fn": "fn14",
									"Children": null
								}
							]
						}
					]
				}
			]
		},
		{
			"Id": 4,
			"ParentId": 123,
			"Fn": "fn4",
			"Children": null
		}
	]
}


*/
