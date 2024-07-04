package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "booking-website/database"
    "booking-website/handlers"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    database.Connect()

    r := mux.NewRouter()
    r.HandleFunc("/bookings", handlers.GetBookings).Methods("GET")
    r.HandleFunc("/booking", handlers.CreateBooking).Methods("POST")

    http.Handle("/", r)
    log.Println("Starting server on port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
