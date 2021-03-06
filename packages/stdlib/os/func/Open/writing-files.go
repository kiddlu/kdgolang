package main
import (
    "bufio"
    "fmt"
    "runtime"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var path string;
    go_os := runtime.GOOS;

    //To start, here’s how to dump a string (or just bytes) into a file.
    d1 := []byte("hello\ngo\n")    
    if go_os == "windows" {
        path = os.Getenv("TMP") + "\\dat1"
    } else {
        path = "/tmp/dat1"
    }
    err := ioutil.WriteFile(path, d1, 0644)
    check(err)

    //For more granular writes, open a file for writing.
    if go_os == "windows" {
        path = os.Getenv("TMP") + "\\dat2"
    } else {
        path = "/tmp/dat2"
    }
    f, err := os.Create(path)
    check(err)

    //It’s idiomatic to defer a Close immediately after opening a file.
    defer f.Close()

    //You can Write byte slices as you’d expect.
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    //A WriteString is also available.
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    //Issue a Sync to flush writes to stable storage.
    f.Sync()

    //bufio provides buffered writers in addition to the buffered readers we saw earlier.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    //Use Flush to ensure all buffered operations have been applied to the underlying writer.
    w.Flush()
}