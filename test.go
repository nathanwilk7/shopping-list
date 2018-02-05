package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Name string `json:"name"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Quantity int `json:"quantity"`
}

func itemsHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	f, err := os.Open("items.txt")
	defer f.Close()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		fmt.Printf("%v", err)
	}
	reader := bufio.NewReader(f)
	items := []Item{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Split(line, ",")
		name, latStr, lonStr := fields[0], fields[1], strings.TrimSpace(fields[2])
		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			break
		}
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			break
		}
		items = append(items, Item{
			Name: name,
			Lat: lat,
			Lon: lon,
		})
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			continue
		}
	}
	j, err := json.Marshal(items)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		fmt.Printf("%v", err)
	}
	w.Write(j)
}

func listHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	f, err := os.Open("list.txt")
	defer f.Close()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		fmt.Printf("%v", err)
	}
	reader := bufio.NewReader(f)
	items := []Item{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Split(line, ",")
		name, latStr, lonStr, quantityStr := fields[0], fields[1], fields[2], strings.TrimSpace(fields[3])
		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			break
		}
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			break
		}
		quantityInt64, err := strconv.ParseInt(quantityStr, 10, 32)
                if err != nil {
                       	fmt.Fprintf(w, "%v", err)
                        fmt.Printf("%v", err)
                        break
               	}
		quantity := int(quantityInt64)
		items = append(items, Item{
			Name: name,
			Lat: lat,
			Lon: lon,
			Quantity: quantity,
		})
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			fmt.Printf("%v", err)
			continue
		}
	}
	j, err := json.Marshal(items)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		fmt.Printf("%v", err)
	}
	w.Write(j)
}

func saveHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
    items := []Item{}
    err := decoder.Decode(&items)
    if err != nil {
        fmt.Println(err)
		fmt.Fprintf(w, "%v", err)
    }
	f, err := os.Create("list.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%v", err)
	}
	for _, item := range items {
		f.Write([]byte(fmt.Sprintf("%s,%f,%f,%d\n", item.Name, item.Lat, item.Lon, item.Quantity)))
	}
}

func addHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
    item := Item{}
    err := decoder.Decode(&item)
    if err != nil {
        fmt.Println(err)
		fmt.Fprintf(w, "%v", err)
    }
	f, err := os.OpenFile("items.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%v", err)
	}
	f.Write([]byte(fmt.Sprintf("%s,%f,%f\n", item.Name, item.Lat, item.Lon)))
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Fprintf(w, "%s", f)
}

func main () {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    	if err != nil {
        	fmt.Println("ListenAndServe: ", err)
    	}
	//http.ListenAndServe(":443", nil)
}
