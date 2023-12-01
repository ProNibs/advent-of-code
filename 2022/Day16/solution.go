package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type Valves struct {
	name              string
	flow_rate         int
	connected_tunnels []string
}

func pressureRelease(current_time int, current_valve string, valve_list map[string]Valves) int {
	/* Given the current time, current valve, and the list
	figure out how much pressure could be released for the valve
	*/
	return current_time * valve_list[current_valve].flow_rate
}

// func determineBestRoute(current_time int, current_valve string, valve_list map[string]Valves, already_opened []string) []string {
// 	// Given the current valve and all current valves already opened
// 	// What's the next best valve to open?
// 	if len(already_opened) == 0 {
// 		// First time through, just ignore already opened conditions
// 		var current_path_attempt []string

// 	}

// }

func createGraph(valve_list map[string]Valves) {
	g := graph.New(graph.StringHash, graph.Acyclic())

	for _, v := range valve_list {
		_ = g.AddVertex(v.name, graph.VertexWeight(v.flow_rate))
		for _, item := range v.connected_tunnels {
			_ = g.AddEdge(v.name, item)
		}
	}
	file, _ := os.Create("my-graph.gv")
	_ = draw.DOT(g, file)
	_ = graph.DFS(g, "AA", func(value string) bool {
		fmt.Println(value)
		return false
	})
}

func main() {

	lines, _ := readLines("data.txt")
	valve_list := make(map[string]Valves)
	for _, line := range lines {
		line_split := strings.Split(line, " ")
		// Current Valve name
		current_valve_struct := new(Valves)
		current_valve_struct.name = line_split[1]
		// Get flow rate number
		flow_value := strings.Split(line_split[4], "=")[1]
		flow_value = strings.Trim(flow_value, ";")
		current_valve_struct.flow_rate, _ = strconv.Atoi(flow_value)
		// Generate tunnel list
		for _, valve_name := range line_split[9:] {
			current_valve_struct.connected_tunnels = append(
				current_valve_struct.connected_tunnels, strings.Trim(valve_name, ","))
		}
		valve_list[current_valve_struct.name] = *current_valve_struct
	}
	fmt.Println("Valve list", valve_list)
	createGraph(valve_list)

}
