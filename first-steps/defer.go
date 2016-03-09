package main

import "fmt"
import "os"

func main() {
	info := "Something to write to the file"
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f, info)

}

func createFile(p string) *os.File {
	fmt.Println("creating file", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, info string) {
	fmt.Println("writing to file" + f.Name())
	fmt.Fprintln(f, info)
}

func closeFile(f *os.File) {
	fmt.Println("closing file", f.Name())
	f.Close()
}
