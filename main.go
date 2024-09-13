package main

//rzp_test_XkDWTHT1Xa5vGz
//BR7TjcTcMH0rxdPNK19UCzAV
import (
	"PaymentGatewayDemo/server"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	srv := server.InitServices()
	router := chi.NewRouter()
	router.Get("/pay", srv.Payment)
	router.Post("/customer", srv.CreateCustomer)
	httpPort := "8082"
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
