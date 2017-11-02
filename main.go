package main

import (
	. "backend/commons"
)

func main() {
	r := initRouter()
	r.Run(HOST + PORT)
}
