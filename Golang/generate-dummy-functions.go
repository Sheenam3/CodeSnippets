//This program generates some dummy functions given in C program file in this folder. It was required to test a bcc tool/ebpf tool and x-tracer development
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 1; i < 40; i++ {

		for ch := 'a'; ch <= 'z'; ch++ {
			alpha := string(rune(ch))
			inttostr := strconv.Itoa(i)
			data := "void" + " " + alpha + inttostr + "();" + "\n"
			fmt.Print(data)
			file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				fmt.Println("Could not open example.txt")
				return
			}

			defer file.Close()

			_, err2 := file.WriteString(data)

			if err2 != nil {
				fmt.Println("Could not write text to example.txt")

			} else {
				fmt.Println("Operation successful! Text has been appended to example.txt")
			}

			nextalpha := ch + 1
			nalpha := string(rune(nextalpha))
			callingfunc := nalpha + inttostr
			if alpha == "z" {
				newseries := i + 1
				callingfunc = "a" + strconv.Itoa(newseries)
			}
			funcdef := "void" + " " + alpha + inttostr + "()" + "{" + "\n" + callingfunc + "();" + "\n" + "}" + "\n"
			fmt.Print(funcdef)
			file1, err := os.OpenFile("funcdef.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if err != nil {
				fmt.Println("Could not open example.txt")
				return
			}

			defer file1.Close()

			_, err2 = file1.WriteString(funcdef)

			if err2 != nil {
				fmt.Println("Could not write text to funcdef")

			} else {
				fmt.Println("Operation successful! Text has been appended to funcdef.txt")
			}
		}

	}
}
