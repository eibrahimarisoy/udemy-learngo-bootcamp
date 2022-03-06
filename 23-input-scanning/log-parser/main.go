package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		if len(fields) != 2 {
			fmt.Printf("wrong input %v\n", fields)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("wrong input %v\n", fields[1])
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		sum[domain] += visits
		total += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println((strings.Repeat("-", 45)))

	sort.Strings(domains)
	for _, domain := range domains {
		fmt.Printf("%-30s %10d\n", domain, sum[domain])
	}
	fmt.Println(total)
}
