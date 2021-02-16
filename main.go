package main

import (
	"fmt"
	"strconv"
	"os"
	"log"
	"net/http"
	"io/ioutil"
)

func fileCreator(n int) {
	os.Chdir("./storage/posts")
	newFile, err := os.Create(strconv.Itoa(n) + ".txt")

	if err !=nil {
		log.Fatal(err)
	}

	log.Println(newFile)
	newFile.Close()

} 

func reqConn(n int) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(n))
	if err !=nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	msg := []byte(body)

	file, err := os.OpenFile(strconv.Itoa(n) + ".txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission danied")
		}
	}

	file.Write(msg)

	fmt.Printf("%v", string(body))
	defer res.Body.Close()
	defer file.Close()
}

func main() {
	for i:= 1; i <= 5; i++ {
		go fileCreator(i)
		go reqConn(i)
	}

	var input string
    fmt.Scanln(&input)
}