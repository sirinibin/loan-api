package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirinibin/loan-api/controller"
	"github.com/sirinibin/loan-api/env"
)

func main() {
	fmt.Println("Loan API")

	httpPort := env.Getenv("API_PORT", "4000")
	httpsPort, err := strconv.Atoi(httpPort)
	if err != nil {
		log.Print(err)
		return
	}
	httpsPort = httpsPort + 1

	router := mux.NewRouter()

	// Get balance sheet
	router.HandleFunc("/v1/balance-sheet", controller.GetBalanceSheet).Methods("POST")
	// Get outcome
	router.HandleFunc("/v1/outcome", controller.GetOutcome).Methods("POST")

	go func() {
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(httpsPort), "localhost.cert.pem", "localhost.key.pem", router))

	}()

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			log.Printf("Serving @ https://" + ip.String() + ":" + strconv.Itoa(httpsPort) + " /\n")
			log.Printf("Serving @ http://" + ip.String() + ":" + httpPort + " /\n")
		}
	}
	log.Fatal(http.ListenAndServe(":"+httpPort, router))

}
