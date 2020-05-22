package agent

const (
	topSide  = "top"
	downSide = "down"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Position) Equal(pos Position) bool {
	if p.X == pos.X && p.Y == pos.Y {
		return true
	}
	return false
}

type Move struct {
	Player Player
	From   Position `json:"from"`
	To     Position `json:"to"`
}

func (m *Move) Equal(move Move) bool {
	if m.From.Equal(move.From) && m.To.Equal(move.To) {
		return true
	}
	if m.From.Equal(move.To) && m.To.Equal(move.From) {
		return true
	}
	return false
}

type Player struct {
	Name string `json:"name"`
	Side string `json:"side"`
}

type World struct {
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Moves       []Move   `json:"moves"`
	FilledMoves []Move   `json:"filled-moves"`
	Turn        string   `json:"turn"`
	BallPos     Position `json:"ball-pos"`
	Player1     Player   `json:"player1"`
	Player2     Player   `json:"player2"`
}