package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	DFS = iota
	BFS
	GBFS
	ASTAR
	DIJKSTRA
)

type Point struct {
	Row int
	Col int
}

type Wall struct {
	State Point
	wall  bool
}

type Maze struct {
	Height int
	Width  int
	Start  Point
	Goal   Point
	Walls  [][]Wall
}

func main() {
	var m Maze
	var maze, searchType string

	flag.StringVar(&maze, "file", "maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("maze height/width:", m.Height, m.Width)
}

func (g *Maze) Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening %s: %s\n", fileName, err)
	}
	defer f.Close()

	var fileContents []string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("cannot open file %s: %s", fileName, err)
		}

		fileContents = append(fileContents, line)
	}

	foundStart, foundEnd := false, false
	for _, line := range fileContents {
		if strings.Contains(line, "A") {
			foundStart = true
		}

		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}

	if !foundStart {
		return fmt.Errorf("starting location not found")
	}

	if !foundEnd {
		return fmt.Errorf("ending location not found")
	}

	g.Height = len(fileContents)
	g.Width = len(fileContents[0])

	var rows [][]Wall

	for i, row := range fileContents {
		var cols []Wall

		for j, col := range row {
			curLetter := fmt.Sprintf("%c", col)
			var wall Wall

			switch curLetter {
			case "A":
				g.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "B":
				g.Goal = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = true
			default:
				continue
			}

			cols = append(cols, wall)
		}

		rows = append(rows, cols)
	}

	g.Walls = rows

	return nil
}
