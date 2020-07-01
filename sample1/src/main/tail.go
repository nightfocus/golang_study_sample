package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	//"time"

	"github.com/fsnotify/fsnotify"
)

/*
func main() {

	fmt.Println("Usage: ./monitor logfile [key]")

	if len(os.Args) == 3 {
		TailFile(os.Args[1], os.Args[2])
	} else {
		TailFile(os.Args[1], "")
	}
}

*/

// 类似tail -f 的流式文件输出, 只匹配key关键字的行
func TailFile(filename string, key string) error {
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		return err
	}
	// fmt.Print(file.Seek(0, os.SEEK_END))

	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	if err = watcher.Add(filename); err != nil {
		return err
	}

	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}

		sby := string(by)
		if strings.Contains(sby, key) == true {
			fmt.Print(sby)
		}

		if err != io.EOF {
			continue
		}
		if err = waitForChange(watcher); err != nil {
			return err
		}
	}
}

func waitForChange(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}
