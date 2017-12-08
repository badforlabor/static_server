package main
import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"fmt"
)

func staticResource(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	path := r.URL.Path
	request_type := path[strings.LastIndex(path, "."):]
	switch request_type {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	case ".json":
		w.Header().Set("content-type", "application:json;charset=utf8")
	default:
	}

	/*
	header('content-type:application:json;charset=utf8');
header('Access-Control-Allow-Origin:*');
header('Access-Control-Allow-Methods:POST');
header('Access-Control-Allow-Headers:x-requested-with,content-type');
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")


	fullpath := "." + path
	fmt.Println("加载资源：", fullpath)

	fin, err := os.Open(fullpath)
	defer fin.Close()
	if err != nil {
		fmt.Println("static resource:", err)
		w.Write([]byte(""))
		return
	}
	fd, _ := ioutil.ReadAll(fin)
	w.Write(fd)
}
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("start.")

	// realPath = flag.String("path", "", "static resource path")
	//flag.Parse()
	http.HandleFunc("/", staticResource)
	err := http.ListenAndServe(":2888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end.")
}