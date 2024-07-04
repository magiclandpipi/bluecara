package handlers

import (
    "encoding/json"
    "net/http"
    "booking-website/database"
    "booking-website/models"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
    var bookings []models.Booking
    database.DB.Find(&bookings)
    json.NewEncoder(w).Encode(&bookings)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
    var booking models.Booking
    json.NewDecoder(r.Body).Decode(&booking)
    database.DB.Create(&booking)
    json.NewEncoder(w).Encode(&booking)
}
