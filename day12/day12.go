package day12

import (
	"math"
	"strconv"
)

// Waypoint creates a waypoint
type Waypoint struct {
	Movements map[Direction]int
}

// NewWaypoint creates a new way point
func NewWaypoint() *Waypoint {
	return &Waypoint{
		Movements: map[Direction]int{
			North: 1,
			East:  10,
		},
	}
}

// Move moves the waypoint a distance given a direction
func (w *Waypoint) Move(direction Direction, distance int) {
	w.Movements[direction] += distance
}

// RotateRight rotates the waypoint right respective to the boat
func (w *Waypoint) RotateRight(degrees int) {
	times := degrees / 90
	var north, south, east, west int
	for i := 0; i < times; i++ {
		north, south, east, west = w.Movements[North], w.Movements[South], w.Movements[East], w.Movements[West]
		w.Movements[North] = west
		w.Movements[East] = north
		w.Movements[South] = east
		w.Movements[West] = south
	}
}

// RotateLeft rotates the waypoint left respective to the boat
func (w *Waypoint) RotateLeft(degrees int) {
	w.RotateRight(360 - degrees)
}

// Boat represents the boat we want to move
type Boat struct {
	Direction Direction
	Movements map[Direction]int
	Waypoint  *Waypoint
}

// NewBoat initializes a new boat
func NewBoat() *Boat {
	return &Boat{
		Direction: East,
		Movements: make(map[Direction]int),
		Waypoint:  NewWaypoint(),
	}
}

// MoveBoatForward moves a specific distance in
// the boat's direction
func (b *Boat) MoveBoatForward(distance int) {
	b.Movements[b.Direction] += distance
}

// MoveForwardRelativeToWaypoint moves the boat a number of times
// the waypoint's direction
func (b *Boat) MoveForwardRelativeToWaypoint(times int) {
	for direction, distance := range b.Waypoint.Movements {
		b.Movements[direction] += (distance * times)
	}
}

// RotateRight rotates the boat in the right direction
func (b *Boat) RotateRight(degrees int) {
	times := degrees / 90
	for i := 0; i < times; i++ {
		switch b.Direction {
		case North:
			b.Direction = East
		case East:
			b.Direction = South
		case South:
			b.Direction = West
		case West:
			b.Direction = North
		}
	}
}

// RotateLeft rotates the boat in left direction
func (b *Boat) RotateLeft(degrees int) {
	b.RotateRight(360 - degrees)
}

// RotateWaypointRight rotates the waypoint right
func (b *Boat) RotateWaypointRight(degrees int) {
	b.Waypoint.RotateRight(degrees)
}

// RotateWaypointLeft rotates the waypoint left
func (b *Boat) RotateWaypointLeft(degrees int) {
	b.Waypoint.RotateLeft(degrees)
}

// MoveBoat moves the boat a specific distance in the
// given direction
func (b *Boat) MoveBoat(direction Direction, distance int) {
	b.Movements[direction] += distance
}

// MoveWaypoint moves the waypoint at a given direction
func (b *Boat) MoveWaypoint(direction Direction, distance int) {
	b.Waypoint.Move(direction, distance)
}

// ManhattanDistance calculates the manhattan distance
// between (0,0) and the point where the boat is
func (b *Boat) ManhattanDistance() int {
	north, south := b.Movements[North], b.Movements[South]
	west, east := b.Movements[West], b.Movements[East]

	y := math.Abs(float64(north - south))
	x := math.Abs(float64(west - east))

	return int(x + y)
}

// RunActionsWithoutWaypoint runs the actions without taking into account
// the waypoint
func (b *Boat) RunActionsWithoutWaypoint(actions []Action) {
	for _, a := range actions {
		switch a.Action {
		case "N":
			b.MoveBoat(North, a.Value)
		case "E":
			b.MoveBoat(East, a.Value)
		case "S":
			b.MoveBoat(South, a.Value)
		case "W":
			b.MoveBoat(West, a.Value)
		case "F":
			b.MoveBoatForward(a.Value)
		case "R":
			b.RotateRight(a.Value)
		case "L":
			b.RotateLeft(a.Value)
		}
	}
}

// RunActionsWithWaypoint runs the actions taking into account the waypoint
func (b *Boat) RunActionsWithWaypoint(actions []Action) {
	for _, a := range actions {
		switch a.Action {
		case "N":
			b.MoveWaypoint(North, a.Value)
		case "E":
			b.MoveWaypoint(East, a.Value)
		case "S":
			b.MoveWaypoint(South, a.Value)
		case "W":
			b.MoveWaypoint(West, a.Value)
		case "F":
			b.MoveForwardRelativeToWaypoint(a.Value)
		case "R":
			b.RotateWaypointRight(a.Value)
		case "L":
			b.RotateWaypointLeft(a.Value)
		}
	}
}

// FirstPart runs the instructions and calcs the manhattan distance
// between the boat's initial position and the final one
func FirstPart(lines []string) (int, error) {
	actions, err := parseLines(lines)
	if err != nil {
		return 0, err
	}
	b := NewBoat()
	b.RunActionsWithoutWaypoint(actions)
	return b.ManhattanDistance(), nil
}

// SecondPart runs the instructions and calcs the manhattan distance
// between the boat's initial position and the final one
// taking into account the waypoint's relative position
func SecondPart(lines []string) (int, error) {
	actions, err := parseLines(lines)
	if err != nil {
		return 0, err
	}
	b := NewBoat()
	b.RunActionsWithWaypoint(actions)
	return b.ManhattanDistance(), nil
}

// Action represents an action
// e.g. R90, L180, F10
type Action struct {
	Action string
	Value  int
}

func parseLines(lines []string) ([]Action, error) {
	actions := make([]Action, len(lines))
	for i, l := range lines {
		action, strValue := l[:1], l[1:]
		v, err := strconv.Atoi(strValue)
		if err != nil {
			return nil, err
		}
		actions[i] = Action{
			Action: action,
			Value:  v,
		}
	}
	return actions, nil
}
