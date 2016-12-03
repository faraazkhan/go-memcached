package client

import (
	"github.com/olekukonko/tablewriter"
	stats "github.com/youtube/vitess/go/memcache"
	"log"
	"os"
	"strings"
	"time"
)

func cleanNameForStat(s string) (str string) {
	return strings.Title(strings.Replace(s, "_", " ", -1))
}

func Stats(hosts []string) {
	var data [][]string
	connectionString := strings.Join(hosts, ",")
	timeOut := time.Duration(5000000)
	interestingStats := [...]string{"pid", "uptime", "curr_connections", "total_connections", "get_hits", "get_misses", "curr_items"}
	connection, err := stats.Connect(connectionString, timeOut)
	if err != nil {
		log.Printf("There was a problem establishing connection with memcached at %v\n %s", connectionString, err)
	}
	result, err := connection.Stats("")
	if err != nil {
		log.Printf("There was a problem getting stats from memcached\n %v", err)
	}
	statistics := strings.Split(string(result), "\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Stat", "Value"})

	for _, statistic := range statistics {
		for _, interestingStat := range interestingStats {
			if strings.Contains(statistic, interestingStat) {
				tableRow := strings.Split(statistic, " ")[1:]
				tableRow[0] = cleanNameForStat(tableRow[0])
				data = append(data, tableRow)
			}
		}
	}

	for _, v := range data {
		table.Append(v)
	}

	table.Render()

	connection.Close()
}
