package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type Message struct {
// 	Text string `json:"textx"`
// }

// func main() {
// 	fmt.Println("Yooh")

// 	http.HandleFunc("/", test)
// 	http.HandleFunc("/greet", greet)
// 	http.HandleFunc("/json", send)

// 	//log.Fatal(http.ListenAndServe(":3000", nil))
// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil) //starts server at port 8080

// }

// // func helloUser(writer http.ResponseWriter, request *http.Request) {
// // 	var greeting = "Yooh, sup bro"
// // 	fmt.Fprintln(writer, greeting)

// // }

// func test(w http.ResponseWriter, r *http.Request) {
// 	//w http.ResponseWriter-sends responses to the client
// 	//r *http.Request-contains all incoming request data ie body,headers etc
// 	fmt.Println("Sup bra", r.URL.Path)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("naem")

// 	if name == "" {
// 		name = "Guest"
// 	}

// 	fmt.Fprintf(w, "sup %v", name)
// }

// func send(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	msg := Message{Text: "Yooh!"}
// 	json.NewEncoder(w).Encode(msg)
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Before modification:", r.URL.Path)//r.Url.Path-read the request path

// 	// modify the request directly (since r is a pointer)
// 	r.URL.Path = "/changed-path"

// 	fmt.Println("After modification:", r.URL.Path)

// 	// read the request body (works because r is a pointer — body is a stream)
// 	body, _ := io.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	fmt.Fprintf(w, "Modified Path: %s\n", r.URL.Path)
// 	fmt.Fprintf(w, "Request Body: %s\n", string(body))
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Server running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type List struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Episodes int    `json:"episodes"`
}

var lists []List

func main() {
	// http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
	// 	// if r.Method == http.MethodPost{
	// 	// 	addTo(w,r)
	// 	// } else if r.Method == http.MethodGet{
	// 	// 	getList(w,r)
	// 	// } else {
	// 	// 	http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	// 	// }

	// 	switch r.Method {
	// 	case http.MethodPost:
	// 		addTo(w, r)
	// 	case http.MethodGet:
	// 		getList(w, r)

	// 	default:
	// 		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	http.HandleFunc("/add", addTo)
	http.HandleFunc("/get", getList)
	http.HandleFunc("/delete", deleteItem)
	http.HandleFunc("/update", update)

	fmt.Println("Server Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func addTo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post request expected", http.StatusMethodNotAllowed)
		return
	}
	var list List
	err := json.NewDecoder(r.Body).Decode(&list)
	//r.Body-incoming http request body,
	//NewDecoder-creates a decoder that can read from the request stream
	//Decode-take json an fills in var
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	list.ID = len(lists)
	//fmt.Println(list.ID)
	lists = append(lists, list)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("Something added to list: %v", list.Title)})
	//w-http response writer-sends data back to client
	//NewEncoder-creates encoder that can write json in res
	//Encode()-takes go data and converts to json, then sends ti response body

}

func getList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only get is allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only delete Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i, t := range lists {
		if t.ID == id {
			lists = append(lists[:i], lists[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("%v deleted from list", id)})
			return
		}
	}

	http.Error(w, "Error while deleting itme", http.StatusBadRequest)

}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only put Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id param", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var updated List
	if err = json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
	}

	for i, l := range lists {
		if l.ID == id {
			if updated.Title != "" {
				lists[i].Title = updated.Title
			}
			if updated.Episodes != 0 {
				lists[i].Episodes = updated.Episodes
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"Message": fmt.Sprintf("%v updated ", updated.Title)})

			return

		}
	}
}
