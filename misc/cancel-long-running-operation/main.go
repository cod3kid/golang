package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func cancellableHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	dsn := "host=localhost user=postgres password=postgres dbname=gorm-test port=5432"
	dbMain, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		http.Error(w, "Failed to open main connection to DB", http.StatusInternalServerError)
		return
	}
	dbCancel, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		http.Error(w, "Failed to open cancel connection to DB", http.StatusInternalServerError)
		return
	}

	// Get backend PID for main connection
	var pid int
	if err := dbMain.Raw("SELECT pg_backend_pid()").Scan(&pid).Error; err != nil {
		http.Error(w, "Failed to get backend PID", http.StatusInternalServerError)
		return
	}
	fmt.Println("Started processing request on PID:", pid)

	done := make(chan string, 1)

	go func() {
		var result int
		err := dbMain.Raw("SELECT 42 FROM (SELECT pg_sleep(10)) AS t").Scan(&result).Error

		if err != nil {
			done <- fmt.Sprintf("Query canceled or failed: %v", err)
		} else {
			done <- fmt.Sprintf("Query result: %d", result)
		}
	}()

	select {
	case msg := <-done:
		fmt.Fprintln(w, msg)
		fmt.Println("Finished processing request")
	case <-ctx.Done():
		dbCancel.Exec(fmt.Sprintf("SELECT pg_cancel_backend(%d)", pid))
		fmt.Println("Request canceled by the client:", ctx.Err())
	}
}

func main() {
	http.HandleFunc("/cancellable", cancellableHandler)
	fmt.Println("Server listening to port 3000")
	http.ListenAndServe(":3000", nil)
}
