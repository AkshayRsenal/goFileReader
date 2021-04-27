package main

import (
    "bufio"
    "fmt"
    "log"
    "strconv"
    "os"
)

type Student struct {
    FirstName string
    LastName  string
}

func main() {
    fmt.Println("What is the name of your file?\n") 
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
		sum, err := strconv.Atoi(line)
		if err != nil {
					log.Fatal(err)
				}
		if len(line) == 0 {
			// skip blank lines
			continue
		}
		
		
		addition += sum;
		
		fmt.Println(addition )
	}
	
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    // for scanner.Scan() {
    //     line := scanner.Text()
    //     if len(line) == 0 {
    //         // skip blank lines
    //         continue
    //     }
    //     if '0' <= line[0] && line[0] <= '9' {
    //         sum := 0
    //         for _, field := range strings.Fields(line) {
    //             n, err := strconv.Atoi(field)
    //             if err != nil {
    //                 log.Fatal(err)
    //             }
    //             sum += n
    //         }
    //         fmt.Println(sum)
    //     } else {
    //         fields := strings.Fields(line)
    //         if len(fields) != 2 {
    //             log.Fatal("don't know how to get first name last name")
    //         }
    //         fmt.Println("First:", fields[0], "Last:", fields[1])
    //     }
    // }
    // if err := scanner.Err(); err != nil {
    //     log.Fatal(err)
    // }
}