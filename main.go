package main

//rzp_test_XkDWTHT1Xa5vGz
//BR7TjcTcMH0rxdPNK19UCzAV
import (
	"github.com/go-chi/chi"
	"github.com/razorpay/razorpay-go"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	router := chi.NewRouter()
	router.Get("/pay", Payment)
	httpPort := "8082"
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}

type PageVars struct {
	OrderId string
}

func Payment(w http.ResponseWriter, r *http.Request) {
	client := razorpay.NewClient(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	data := map[string]interface{}{
		"amount":   50000,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := body["id"]
	str := value.(string)
	HomePageVars := PageVars{OrderId: str}
	t, err := template.ParseFiles("app.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
