package agent

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/Gimulator/client-go"
)

type Agent struct {
	controller

	ch chan client.Object
}

func NewAgent() (Agent, error) {
	ch := make(chan client.Object, 16)
	c, err := newController(ch)
	if err != nil {
		return Agent{}, err
	}

	agent := Agent{
		ch:         ch,
		controller: c,
	}
	agent.load()

	return agent, nil
}

func (a *Agent) load() {
	a.setRegister()
}

func (a *Agent) Listen() {
	for {
		fmt.Println("Start listening")
		obj := <-a.ch
		if obj.Key.Type != typeVerdict {
			continue
		}

		world := &World{}
		err := json.Unmarshal([]byte(obj.Value.(string)), world)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if world.Turn != name {
			continue
		}

		a.Action(world)
	}
}
func (a *Agent) Action(world *World) {
	time.Sleep(time.Millisecond * 500)
	validMoves := a.validMoves(world)

	if len(validMoves) == 0 {
		panic("there is no move")
	}

	i := rand.Intn(len(validMoves))
	for {
		err := a.setAction(validMoves[i])
		if err == nil {
			break
		}
	}
}

var (
	dirX = []int{1, 1, 0, -1, -1, -1, 0, 1}
	dirY = []int{0, 1, 1, 1, 0, -1, -1, -1}
)

func (a *Agent) validMoves(w *World) []Move {
	var validMoves []Move

	for ind := 0; ind < 8; ind++ {
		x := w.BallPos.X + dirX[ind]
		y := w.BallPos.Y + dirY[ind]
		if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
			continue
		}

		validMove := Move{
			From: w.BallPos,
			To: Position{
				X: x,
				Y: y,
			},
		}

		isValid := true
		for _, m := range w.Moves {
			if validMove.Equal(m) {
				isValid = false
			}
		}

		for _, m := range w.FilledMoves {
			if validMove.Equal(m) {
				isValid = false
			}
		}

		if isValid {
			validMoves = append(validMoves, validMove)
		}
	}
	return validMoves
}
