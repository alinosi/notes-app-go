package main

import "fmt"

type dummy interface {
	read()
}

type subUser struct {
	name string
}

type user struct {
	another subUser
	age     int
}

func (u subUser) read() {
	u.name = "ganti nama"
}

func changename(a dummy) {
	a.read()
}

func main() {
	b := subUser{"yanti"}
	a := &user{b, 12}

	changename(a.another)
	fmt.Println(b.name)
	fmt.Println(a.another.name)
}

// err := json.NewDecoder(r.Body).Decode(&newNote)

// melakukan simulasi terhadap field suatu struct, apakah mengubah value asli atau tidak
