package main

import (
    "fmt"
)

type User struct {
    Name       string
    Occupation string
}

func (u User) String() string {

    return fmt.Sprintf("%s is a(n) %s", u.Name, u.Occupation)
}

func main() {

    u1 := User{"Kamlesh Kumar", "engineer"}
    u2 := User{"Durgesh Nand", "doctor"}

    fmt.Println(u1)
    fmt.Println(u2)
}
