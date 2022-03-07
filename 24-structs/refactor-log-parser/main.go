package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func main() {
	var ()

	p := parser{
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++
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

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		p.total += visits

		r := result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
		p.sum[domain] = r
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println((strings.Repeat("-", 45)))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Println(p.total)
}
