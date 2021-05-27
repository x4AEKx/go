package main

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"flag"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	stringsLibrary()
	inputOutput()
	filesAndDirectories()
	listGo()

	kids := []Person{
		{"Jill", 9},
		{"Jack", 12},
		{"Alex", 19},
	}

	sort.Sort(ByName(kids))

	fmt.Println(kids)

	hash()
	hashSHA1()

	// Определение флагов
	maxp := flag.Int("max", 6, "the max value")
	// Парсинг
	flag.Parse()
	// Генерация числа от 0 до max
	fmt.Println(rand.Intn(*maxp))

	// Mutex
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)

	// go server()
	// go client()

	// var input string
	// fmt.Scanln(&input)

	// http.HandleFunc("/hello", hello)
	// http.ListenAndServe(":9000", nil)
}

func stringsLibrary() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "test"
		strings.ToLower("TEST"),

		// "TEST"
		strings.ToUpper("test"),
	)
}

func inputOutput() {
	var buf bytes.Buffer
	buf.Write([]byte("test"))
}

func filesAndDirectories() {
	filePath := "test.txt"
	newFile := "test1.txt"

	readFile := func(path string) string {
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			return err.Error()
		}

		str := string(bs)
		return str
	}

	createFile := func(newFile string) {
		file, err := os.Create(newFile)

		if err != nil {
			// здесь перехватывается ошибка
			return
		}
		defer file.Close()

		file.WriteString("newFile: Hello Golang")
	}

	createFile(newFile)
	fmt.Println(readFile(newFile))

	fmt.Println(readFile(filePath))

	dir := "."

	ls := func(dirPath string) {
		dir, err := os.Open(dirPath)
		if err != nil {
			return
		}
		defer dir.Close()

		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}
		for _, fi := range fileInfos {
			fmt.Println(fi.Name())
		}
	}

	ls(dir)

	// allDirectories := func(dirPath string) {
	// 	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
	// 		fmt.Println(path)
	// 		return nil
	// 	})
	// }

	// allDirectories(dir)
}

func listGo() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func hash() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}

func hashSHA1() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

// func server() {
// 	// слушать порт
// 	ln, err := net.Listen("tcp", ":9999")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for {
// 		// принятие соединения
// 		c, err := ln.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		// обработка соединения
// 		go handleServerConnection(c)
// 	}
// }

// func handleServerConnection(c net.Conn) {
// 	// получение сообщения
// 	var msg string
// 	err := gob.NewDecoder(c).Decode(&msg)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Received", msg)
// 	}

// 	c.Close()
// }

// func client() {
// 	// соединиться с сервером
// 	c, err := net.Dial("tcp", "127.0.0.1:9999")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// послать сообщение
// 	msg := "Hello World"
// 	fmt.Println("Sending", msg)
// 	err = gob.NewEncoder(c).Encode(msg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	c.Close()
// }

// func hello(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set(
// 		"Content-Type",
// 		"text/html",
// 	)
// 	io.WriteString(
// 		res,
// 		`<doctype html>
// <html>
// 	<head>
// 			<title>Hello World</title>
// 	</head>
// 	<body>
// 			Hello World!
// 	</body>
// </html>`,
// 	)
// }

// type Server struct{}

// func (this *Server) Negate(i int64, reply *int64) error {
// 	*reply = -i
// 	return nil
// }

// func server() {
// 	rpc.Register(new(Server))
// 	ln, err := net.Listen("tcp", ":9999")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for {
// 		c, err := ln.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		go rpc.ServeConn(c)
// 	}
// }

// func client() {
// 	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var result int64
// 	err = c.Call("Server.Negate", int64(999), &result)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Server.Negate(999) =", result)
// 	}
// }
