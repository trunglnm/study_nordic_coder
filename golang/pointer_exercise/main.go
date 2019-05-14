package main

type Ward struct {
	Id   uint8
	Name string
}

type District struct {
	Id    uint8
	Name  string
	Wards []Ward
}

type City struct {
	Id        uint8
	Name      string
	Districts []District
}

func main() {

}
