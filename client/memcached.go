package client

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strings"
)

func conn(connectionString []string) *memcache.Client {
	connection := memcache.New(strings.Join(connectionString, ","))
	return connection
}

func Get(keys []string, connectionString []string) {
	connection := conn(connectionString)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value", "Expiration"})

	var data [][]string
	results, err := connection.GetMulti(keys)
	if err != nil {
		log.Printf("Error retrieving key", err)
	}
	for _, result := range results {
		value := string(result.Value)
		key := string(result.Key)
		expiration := string(result.Expiration)
		tableRow := []string{key, value, expiration}
		data = append(data, tableRow)
	}
	for _, v := range data {
		table.Append(v)

	}
	table.Render()

}

func Set(key, value string, connectionString []string) {
	connection := conn(connectionString)
	err := connection.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Println("There was an error setting key: ", err)
	}
	log.Println("Successfully set value")
	fmt.Printf("%s was set to %s", key, value)
}

func Flush(keys []string, connectionString []string) {
	connection := conn(connectionString)
	for _, key := range keys {
		if key == "all" {
			err := connection.DeleteAll()
			if err != nil {
				fmt.Println("There was an error deleting keys:", err)
				return
			}
			fmt.Println("Successfully deleted all keys!")
			return
		}
		err := connection.Delete(key)
		if err != nil {
			fmt.Println("There was an error deleting keys: ", err)
			return
		}
		log.Printf("Deleted %s", key)
	}
}
