gopl.io/chl/echo3
//Echo3 выводит аргументы командной строки
package main
import (
 "fmt"
 "os"
)
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
