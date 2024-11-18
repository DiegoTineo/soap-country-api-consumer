package main

import (
	"fmt"
	"sync"

	"github.com/DiegoTineo/soap-country-api-consumer/cmd"
)

func execServer(wg *sync.WaitGroup){
	defer wg.Done()
	defer fmt.Println("server: end")
	cmd.ListAndServe()
}

func execCli(wg *sync.WaitGroup){
	defer wg.Done()
	defer fmt.Println("cli: end")
	cmd.RunCli()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go execServer(&wg)
	go execCli(&wg)
	wg.Wait()

	fmt.Println("end")
}