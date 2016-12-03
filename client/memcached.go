package client

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func Get(keys []string, connectionString string) {
	connection := memcache.New(connectionString)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key", "Value", "Expiration"})

	var data [][]string
	for _, key := range keys {
		log.Printf("Looking for %s", key)
		val, err := connection.Get(key)
		if err != nil {
			fmt.Printf("Error looking for %s\n %s\n", key, err)
			continue
		}
		value := string(val.Value)
		expiration := string(val.Expiration)
		tableRow := []string{key, value, expiration}
		data = append(data, tableRow)
	}

	for _, v := range data {
		table.Append(v)

	}
	table.Render()

}

func Set(key, value string, connectionString string) {
	connection := memcache.New(connectionString)
	err := connection.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Println("There was an error setting key: ", err)
	}
	log.Println("Successfully set value")
	fmt.Printf("%s was set to %s", key, value)
}

func Flush(keys []string, connectionString string) {
	connection := memcache.New(connectionString)
	for _, key := range keys {
		if key == "all" {
			err := connection.DeleteAll()
			if err != nil {
				fmt.Println("There was an error deleting keys:", err)
				return
			}
		}
		err := connection.Delete(key)
		if err != nil {
			fmt.Println("There was an error deleting keys: ", err)
			return
		}
		log.Printf("Deleted %s", key)
	}
}
