package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Define a struct to match your JSON data
type EventData struct {
	Title         string `json:"title"`
	StartDateTime string `json:"start_date_time"`
	Link          string `json:"link"`
}

func main() {
	http.HandleFunc("/create_ics", createICSHandler)
	log.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func createICSHandler(w http.ResponseWriter, r *http.Request) {

	// Add these headers to handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request for CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	//Curl post with different parameters:
	// title := r.FormValue("title")
	// startDateTime := r.FormValue("start_date_time")
	// link := r.FormValue("link")

	// Decode the JSON request body into the EventData struct
	var eventData EventData
	err := json.NewDecoder(r.Body).Decode(&eventData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(eventData)
	// Validating input
	if eventData.Title == "" || eventData.StartDateTime == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Parse startDateTime
	//startTime, err := time.Parse("2006-01-02T15:04:05", eventData.StartDateTime)
	// Attempt to parse the start date and time
	startTime, err := parseDateToISO(eventData.StartDateTime)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DDTHH:MM:SS or Jan 20, 15:04 pm", http.StatusBadRequest)
		return
	}

	// Create ICS content
	icsContent := buildICSContent(eventData.Title, startTime, eventData.Link)

	// Set headers for ICS file download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.ics", strings.ReplaceAll(eventData.Title, " ", "_")))
	w.Header().Set("Content-Type", "text/calendar")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(icsContent))
}

func buildICSContent(title string, startTime time.Time, link string) string {
	// // Now isoStartTime is a time.Time object, and you can format it to an ISO string
	// isoFormattedStartTime := startTime.Format(time.RFC3339)
	endTime := startTime.Add(time.Hour) // Assuming 1-hour event

	// ICS Date format: 20240117T121500Z
	const layout = "20060102T150405Z"
	icsFormat := "BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nBEGIN:VEVENT\nSUMMARY:%s\nDTSTART:%s\nDTEND:%s\nDESCRIPTION:%s\nEND:VEVENT\nEND:VCALENDAR\n"

	return fmt.Sprintf(icsFormat, title, startTime.UTC().Format(layout), endTime.UTC().Format(layout), link)
}
func parseDateToISO(dateStr string) (time.Time, error) {
	// Define the layout matching the ISO format
	const layout = "2006-01-02T15:04:05"

	// Parse the time in given layout assuming it is in PST
	loc, err := time.LoadLocation("America/Los_Angeles") // PST location
	if err != nil {
		return time.Time{}, fmt.Errorf("error loading timezone: %v", err)
	}

	parsedTime, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}

	// The parsedTime is now in PST
	return parsedTime, nil
}

// // parseDateToISO tries to parse a date string in various formats and returns a time.Time object.
// func parseDateToISO(dateStr string) (time.Time, error) {
//     // Define multiple formats to try
//     formats := []string{
//         "01/02/06 at 03:04pm",
//         "2006-01-02T15:04:05",
//         "Jan 2, 2006 at 15:04 pm",
//     }

//     for _, format := range formats {
//         parsedTime, err := time.Parse(format, dateStr)
//         if err == nil {
//             // Format successfully parsed, return time.Time object
//             return parsedTime, nil
//         }
//     }

//     // None of the formats matched, return zero time and an error
//     return time.Time{}, fmt.Errorf("invalid date format")
// }
