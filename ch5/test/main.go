package main

import (
	"fmt"
	"os"
)

func main() {

	var rmdirs []func()

	// TODO: Need to be understanded!
	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	fmt.Println("do some work...")
	for _, rmdir := range rmdirs {
		rmdir()
	}

}

func tempDirs() []string {
	return []string{
		"./test1",
		"./test2",
		"./test3",
		"./test4",
	}
}

// func tempDirs() []string {
//         content := []byte("temporay file's content")
//         dirs := []string{}
//         for i := 0; i < 4; i++ {
//                 dir, err := ioutil.TempDir("", "example_*")
//                 if err != nil {
//                         log.Fatal(err)
//                 }
//                 dirs = append(dirs, dir)
//                 // defer os.RemoveAll(dir)
//                 tmpfn := filepath.Join(dir, "tmpfile")
//                 if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
//                         log.Fatal(err)
//                 }
//         }
//         return dirs
// }
