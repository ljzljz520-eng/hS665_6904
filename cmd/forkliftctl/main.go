package main

import (
	"flag"
	"fmt"
	"forkliftarchive/internal/api"
	"forkliftarchive/internal/service"
	"forkliftarchive/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	path := flag.String("db", "forklift.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	s := service.New(st, nil)
	fmt.Println("forklift archive listening", *addr)
	if e = http.ListenAndServe(*addr, api.New(s).Handler()); e != nil && !os.IsNotExist(e) {
		log.Fatal(e)
	}
}
