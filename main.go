package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Grid struct {
	xsize uint32
	ysize uint32
}

type MarsRover struct {
	grid   Grid
	robots []Robot
}

type Orientation int

const (
	N Orientation = 1
	E Orientation = 2
	S Orientation = 3
	W Orientation = 4
)

func OrientationStr(s byte) Orientation {
	switch s {
	case 'N':
		return Orientation(1)
	case 'E':
		return Orientation(2)
	case 'S':
		return Orientation(3)
	case 'W':
		return Orientation(4)
	default:
		return Orientation(0)
	}
}

func getOrientation(o Orientation) byte {
	switch o {
	case Orientation(1):
		return 'N'
	case Orientation(2):
		return 'E'
	case Orientation(3):
		return 'S'
	case Orientation(4):
		return 'W'
	default:
		return '0'
	}
}

type Position struct {
	x           int
	y           int
	orientation Orientation
}

type Command int

const (
	F Command = 1
	L Command = 2
	R Command = 3
)

func CommandStr(s string) Command {
	switch s {
	case "F":
		return Command(1)
	case "L":
		return Command(2)
	case "R":
		return Command(3)
	default:
		return Command(0)
	}
}

type Robot struct {
	position       Position
	last_valid_pos Position
	lost           bool
	commands       []Command
}

func move_command(mr *MarsRover, robot *Robot) {
	robot.last_valid_pos = robot.position

	for _, command := range robot.commands {
		switch command {
		case F: // F
			switch robot.position.orientation {
			case N:
				robot.position.y += 1
			case E:
				robot.position.x += 1
			case S:
				robot.position.y -= 1
			case W:
				robot.position.x -= 1
			}
		case L: // L
			if robot.position.orientation == N {
				robot.position.orientation = W
			} else {
				robot.position.orientation -= 1
			}
		case R: // R
			if robot.position.orientation == W {
				robot.position.orientation = N
			} else {
				robot.position.orientation += 1
			}
		}

		if robot.position.x > int(mr.grid.xsize) || robot.position.x < 0 {
			robot.lost = true
		} else if robot.position.y > int(mr.grid.ysize) || robot.position.y < 0 {
			robot.lost = true
		} else {
			robot.last_valid_pos = robot.position
			robot.lost = false
		}
	}
}

func add_robot(mr *MarsRover, robot Robot) {
	mr.robots = append(mr.robots, robot)
}

func execute_move_commands(mr *MarsRover) {
	for index := range mr.robots {
		move_command(mr, &mr.robots[index])
	}
}

func print_final_state(mr *MarsRover) {
	for _, robot := range mr.robots {
		if robot.lost {
			fmt.Printf("(%d, %d, %c) LOST\n", robot.last_valid_pos.x, robot.last_valid_pos.y, getOrientation(robot.last_valid_pos.orientation))
		} else {
			fmt.Printf("(%d, %d, %c)\n", robot.last_valid_pos.x, robot.last_valid_pos.y, getOrientation(robot.last_valid_pos.orientation))
		}
	}
}

func main() {
	// var gridx_str string
	// var gridy_str string
	var gridx uint32
	var gridy uint32
	fmt.Print("Enter Grid Size: ")
	_, err := fmt.Scanln(&gridx, &gridy)
	if err != nil {
		panic(err)
	}

	// fmt.Println(gridx, gridy)

	grid := Grid{
		xsize: gridx,
		ysize: gridy,
	}

	var robots []Robot

	mr := MarsRover{
		grid:   grid,
		robots: robots,
	}

	var maxrobots int
	fmt.Print("Enter Number of Robots: ")
	fmt.Scanln(&maxrobots)

	r_cmd, _ := regexp.Compile("([A-Z])")

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < maxrobots; i++ {
		var robot_string string
		var robot_pos Position
		var robot_orientation byte
		var robot_cmd_str string
		var robot_cmd_arr []Command

		fmt.Print("Enter Robot String: ")
		scanner.Scan()
		robot_string = scanner.Text()

		_, err := fmt.Sscanf(robot_string, "(%d, %d, %c) %s", &robot_pos.x, &robot_pos.y, &robot_orientation, &robot_cmd_str)
		if err != nil {
			panic(err)
		}

		robot_pos.orientation = OrientationStr(robot_orientation)

		// fmt.Println(robot_string)
		matches := r_cmd.FindAllString(robot_cmd_str, -1)

		for _, match := range matches {
			// fmt.Println(match)
			robot_cmd_arr = append(robot_cmd_arr, CommandStr(match))
		}

		// fmt.Println(robot_cmd_arr)

		robot := Robot{
			position:       robot_pos,
			last_valid_pos: robot_pos,
			lost:           false,
			commands:       robot_cmd_arr,
		}
		// fmt.Println(robot)
		add_robot(&mr, robot)
	}

	execute_move_commands(&mr)
	print_final_state(&mr)
}
