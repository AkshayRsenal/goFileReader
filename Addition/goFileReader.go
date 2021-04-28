package main

import (
    "bufio"
    "fmt"
    "log"
    "strings"
    "strconv"
    "os"
)

type Student struct {
    FirstName string
    LastName  string
}

func main() {
    fmt.Println("What is the name of the file which contains the numbers?\n") 
    var filename string 
    fmt.Scan(&filename)

    file, err := os.Open(filename)
    if err != nil {
     log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
	addition :=0
	for scanner.Scan() {
		
		line := scanner.Text()
		
		if len(line) == 0 {
			// skip blank lines
			continue
		}
		if '0' <= line[0] && line[0] <= '9' {
		sum, err := strconv.Atoi(line)
		if err != nil {
					log.Fatal(err)
				}
		addition += sum;
		}else{
			fields := strings.Fields(line)
			fmt.Println("First Line:", fields[0])
		}
	}
	fmt.Println(addition)
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    
}