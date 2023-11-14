package main

import (
	"docker/warehouse"
	"docker/warehouse/db"
	"docker/warehouse/http/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// state := db.NewInMemoryState()
	state := db.NewRedisState()
	// state := db.NewGormState(db.NewGormConnection())

	// greg := warehouse.Customer{1, "Greg"}
	// bob := warehouse.Customer{2, "Bob"}

	// gregInterests := warehouse.CustomerSubscription{greg.ID, "blue jeans"}
	// bobInterests := warehouse.CustomerSubscription{bob.ID, "gray t-shirt"}

	// state.AddCustomers([]warehouse.Customer{greg, bob})
	// state.AddSubscriptions([]warehouse.CustomerSubscription{gregInterests, bobInterests})

	wh := warehouse.NewWarehouse(state)

	// customerNotifier := warehouse.NewCustomerNotifier(state, state)

	// wh.Register(customerNotifier)

	warehouseController := controllers.NewWarehouseController(wh)

	r := mux.NewRouter()
	r.HandleFunc("/warehouse/products/list", warehouseController.ListProducts).Methods("GET")
	r.HandleFunc("/warehouse/products/new", warehouseController.NewProduct).Methods("GET")
	r.HandleFunc("/warehouse/products", warehouseController.SubmitProduct).Methods("POST")

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}

}
