package server

import (
	"PaymentGatewayDemo/models"
	"encoding/json"
	"html/template"
	"net/http"
)

type PageVars struct {
	OrderId string
}

func (srv *Server) Payment(w http.ResponseWriter, r *http.Request) {
	var paymentInfo models.PaymentInfo
	err := json.NewDecoder(r.Body).Decode(&paymentInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := map[string]interface{}{
		"amount":   paymentInfo.Amount,
		"currency": paymentInfo.Currency,
		"receipt":  paymentInfo.Receipt,
	}
	body, err := srv.Client.Order.Create(data, nil)

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

func (srv *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerInfo models.CustomerInfo
	err := json.NewDecoder(r.Body).Decode(&customerInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var notes map[string]interface{}
	for i, ele := range customerInfo.Notes {
		notes["notes_key_"+string(rune(i+1))] = ele
	}
	data := map[string]interface{}{
		"name":          customerInfo.Name,
		"contact":       customerInfo.Contact,
		"email":         customerInfo.Email,
		"fail_existing": customerInfo.FailExisting,
		"gstin":         customerInfo.Gstin,
		"notes":         notes,
	}
	body, err := srv.Client.Customer.Create(data, nil)
	jsondata, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsondata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (srv *Server) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerInfo models.UpdateCustomerInfo
	err := json.NewDecoder(r.Body).Decode(&customerInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data map[string]interface{}
	if customerInfo.Name != "" {
		data["name"] = customerInfo.Name
	}
	if customerInfo.Email != "" {
		data["email"] = customerInfo.Email
	}
	if customerInfo.Contact != 0 {
		data["contact"] = customerInfo.Contact
	}
	body, err := srv.Client.Customer.Edit(customerInfo.ID, data, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsondata, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsondata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
