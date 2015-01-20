/*Copyright Joseph Sturtevant 1/11/15
Joseph Sturtevant
CSS 490 Tactical Programming
Assignment 1

This is a simple time server

Code modified from the wiki:
https://golang.org/doc/articles/wiki/

Advice on custom 404 from StackOverFlow user Mostafa
http://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
*/
package main

import (
	"fmt"
    "net/http"
    "time"
    "flag"
)

//Flag variables for port and version, as well as the current version
var (
	portFlag = flag.Int("port", 8080, "Defines the port number to listen on")
	versionFlag = flag.Bool("V", false, "Returns the version")
	version = "1.1"
)

//Handles calls to /time/
//Formats the time to HH:MM:SS AM/PM
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t := time.Now()
	curTime := t.Format("3:04:05 PM")
	UTCTime := t.Format("15:04:05 UTC")
    fmt.Fprintf(w, "<html><head><title>The Time</title></head>")
    fmt.Fprintf(w, "<body><p>The time is : <span style='color:red;font-size:2em'>%s</span>%s</p></body></html>", curTime, UTCTime)
}

//Handles calls to pretty much everywhere other than /time
func homeHandler(w http.ResponseWriter, r *http.Request){
	errorHandler(w,r, http.StatusNotFound)
}

//Error handler
//Prints a custom page on StatusNotFound error (404)
func errorHandler(w http.ResponseWriter, r *http.Request, status int){
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintf(w, "<html><head><title>You Dun Goofed</title></head>")
		fmt.Fprintf(w, "<body><p>These are not the URLs you're looking for.</p></body></html>")
	}
}

//Starts the server. 
//Doesn't run if the -V flag is set
func runServer(){
	http.HandleFunc("/time/", timeHandler)
	http.HandleFunc("/", homeHandler)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil); err != nil{
    	fmt.Printf("Port %v already in use", *portFlag)
    }
}

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Version: %v\n", version)
	} else {
		runServer()
	}
}