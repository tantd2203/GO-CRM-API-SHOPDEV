package main

import "GO-CRM-API-SHOPDEV/internal/routers"

func main() {

	r := routers.NewServer()
	r.Run()

}
