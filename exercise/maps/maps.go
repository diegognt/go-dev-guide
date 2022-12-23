//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func setServerStatus(servers []string) map[string]int {
	serversMap := make(map[string]int)
	for _, server := range servers {
		serversMap[server] = Online
	}

	return serversMap
}

// * Create a function to print server status displaying:
func printServerStatus(servers map[string]int) {
	fmt.Println("\n--- Servers status ---")
	//  - number of servers
	fmt.Println("There are", len(servers), "servers")

	// - number of servers for each status (Online, Offline, Maintenance, Retired)
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Unhandled server status")
		}
	}

	fmt.Println(stats[Online], "servers are Online")
	fmt.Println(stats[Offline], "servers are Offline")
	fmt.Println(stats[Maintenance], "servers are Maintenance")
	fmt.Println(stats[Retired], "servers are Retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serversStatus := setServerStatus(servers)
	printServerStatus(serversStatus)

	// - change server status of `darkstar` to `Retired`
	serversStatus["darkstar"] = Retired
	// - change server status of `aiur` to `Offline`
	serversStatus["aiur"] = Offline
	//  - call display server info function
	printServerStatus(serversStatus)
	// - change server status of all servers to `Maintenance`
	for server := range serversStatus {
		serversStatus[server] = Maintenance
	}
	// - call display server info function
	printServerStatus(serversStatus)

}
