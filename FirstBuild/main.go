package main

import (
	"net/http"
	"utkarshsandeep/firstgo/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
