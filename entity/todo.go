package entity

type Todo struct {
	ID   uint32
	Text string
	Done bool
}

type Todos []Todo
