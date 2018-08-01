package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// // open a file for read only
	// f1, err := os.Open("test1.txt")
	// PrintFatalError(err)
	// defer f1.Close()

	// // create a new file
	// f2, err := os.Create("test2.txt")
	// PrintFatalError(err)
	// defer f2.Close()

	// // open file for read & write
	// f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
	// PrintFatalError(err)
	// defer f3.Close()

	// // rename a file
	// err = os.Rename("text1.txt", "text1New.txt")
	// PrintFatalError(err)

	// // move a file
	// err = os.Rename("./test1.txt", "./testfolder/test1.txt")
	// PrintFatalError(err)

	// copy a file
	CopyFile("testfile.txt", "./testfolder/copytestfile.txt")

	// // delete a file
	// err = os.Remove("test2.txt")
	// PrintFatalError(err)

	// bytes, err := ioutil.ReadFile("test3.txt")
	// fmt.Println(string(bytes))

	// scanner := bufio.NewScanner(f3)
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// 	fmt.Println("Found line:", count, scanner.Text())
	// }

	// // buffered write, efficient store in memory, saves disk I/O
	// writebuffer := bufio.NewWriter(f3)
	// for i := 1; i <= 5; i++ {
	// 	writebuffer.WriteString(fmt.Sprintln("Added line:", i))
	// }
	// writebuffer.Flush()

}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file -", err)
	}
}

func CopyFile(fname1, fname2 string) {

	fOld, err := os.Open(fname1)
	PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	PrintFatalError(err)
	defer fNew.Close()

	// cope bytes from source to destination
	_, err = io.Copy(fNew, fOld)
	PrintFatalError(err)

	// flush file contents to disk
	err = fNew.Sync()
	PrintFatalError(err)

}
