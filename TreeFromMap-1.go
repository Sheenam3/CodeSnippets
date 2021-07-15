
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var objdump [][]string
var funcs []string

type Node struct {
	Fn       string  `json:"Funcname"`
	Children []*Node `json:"Nodes"`
}

var storage = make(map[string][]string)

func main() {
	
	objdump = [][]string{{"main", "fn1"}, {"main", "fn2"}, {"main", "fn3"}, {"main", "fn4"}, {"fn1", "fn5"}, {"fn2", "fn7"}, {"fn3", "fn9"}, {"fn3", "fn8"}, {"fn1", "fn6"}, {"fn5", "fn10"}, {"fn7", "fn11"}, {"fn8", "fn12"}, {"fn8", "fn22"}, {"fn8", "fn13"}, {"fn9", "fn23"}}

	funcs = []string{"fn5", "fn1", "fn2", "fn3", "fn4", "fn10", "fn11", "fn8", "fn9", "fn12", "fn13", "fn6", "fn7"}
	//levl := getchildren("main")
	iterateLevel([]string{"main"})
	root_node := MakeTreeFromMap(storage, "main")
	
	bytes, err := json.Marshal(root_node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
	
}

func (node *Node) AddChild(child *Node) {
	node.Children = append(node.Children, child)
}

func CreateNewNode(fn string) *Node {
	newNode := new(Node)
	newNode.Fn = fn
	return newNode
}

func MakeTreeFromMap(treeMap map[string][]string, rootNodeFn string) *Node {
	cache := make(map[string]*Node)
	for fn, children := range treeMap {
		if _, nodeExists := cache[fn]; !nodeExists {
			node := CreateNewNode(fn)
			cache[fn] = node
		}
		for _, childFn := range children {
			if _, childExists := cache[childFn]; !childExists {
				child := CreateNewNode(childFn)
				cache[childFn] = child
			}
			cache[fn].AddChild(cache[childFn])
		}
	}
	return cache[rootNodeFn]
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





/* Output:
{
   "Funcname":"main",
   "Nodes":[
      {
         "Funcname":"fn1",
         "Nodes":[
            {
               "Funcname":"fn5",
               "Nodes":[
                  {
                     "Funcname":"fn10",
                     "Nodes":null
                  }
               ]
            },
            {
               "Funcname":"fn6",
               "Nodes":null
            }
         ]
      },
      {
         "Funcname":"fn2",
         "Nodes":[
            {
               "Funcname":"fn7",
               "Nodes":[
                  {
                     "Funcname":"fn11",
                     "Nodes":null
                  }
               ]
            }
         ]
      },
      {
         "Funcname":"fn3",
         "Nodes":[
            {
               "Funcname":"fn9",
               "Nodes":null
            },
            {
               "Funcname":"fn8",
               "Nodes":[
                  {
                     "Funcname":"fn12",
                     "Nodes":null
                  },
                  {
                     "Funcname":"fn13",
                     "Nodes":null
                  }
               ]
            }
         ]
      },
      {
         "Funcname":"fn4",
         "Nodes":null
      }
   ]
}

*/
