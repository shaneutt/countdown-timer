package main

import (
	"html/template"
	"net/http"
	"os"
	"time"
)

type CountdownInfo struct {
	DestinationTime  time.Time
	DestinationEvent string
}

func main() {
	// get env vars
	destinationEvent := os.Getenv(`COUNTDOWN_DESTINATION_EVENT`)
	destinationTime, err := time.Parse(`Jan 2, 2006 at 3:04pm (MST)`, os.Getenv(`COUNTDOWN_DESTINATION_TIME`))
	if err != nil {
		panic(err)
	}

	// compile the template
	template := template.Must(template.ParseFiles(`./countdown-timer.html`))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, CountdownInfo{destinationTime, destinationEvent})
	})

	http.ListenAndServe(":8080", nil)
}
