package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var categories = []Categories{
	{ID: 1, Name: "Pakaian Wanita"},
	{ID: 2, Name: "Pakaian Pria"},
	{ID: 3, Name: "Pakaian Anak & Bayi"},
}

func main() {
	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// METHOD GET
	http.HandleFunc("GET /categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	})

	// METHOD CREATE
	http.HandleFunc("POST /categories", func(w http.ResponseWriter, r *http.Request) {
		var newCategory Categories
		if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
			http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
			return
		}

		for _, category := range categories {
			if category.Name == newCategory.Name {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				response := map[string]string{
					"message": "Kategori sudah ada",
					"status":  "fail",
				}
				json.NewEncoder(w).Encode(response)
				return
			}
		}

		if newCategory.Name == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			response := map[string]string{
				"message": "Name tidak boleh kosong",
				"status":  "fail",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		newID := 1
		if len(categories) > 0 {
			newID = categories[len(categories)-1].ID + 1
		}

		newCategory.ID = newID
		categories = append(categories, newCategory)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCategory)
	})

	// METHOD UPDATE
	http.HandleFunc("PUT /categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID harus angka", http.StatusBadRequest)
			return
		}

		var updateData Categories
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			http.Error(w, "Data JSON tidak valid", http.StatusBadRequest)
			return
		}

		if updateData.Name == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			response := map[string]string{
				"message": "Name tidak boleh kosong",
				"status":  "fail",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		for i, cat := range categories {
			if cat.ID == id {
				categories[i].Name = updateData.Name
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(categories[i])
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Kategori tidak ditemukan",
		})
	})

	// METHOD GET ONE
	http.HandleFunc("GET /categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID harus angka", http.StatusBadRequest)
			return
		}

		for _, cat := range categories {
			if cat.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(cat)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Kategori tidak ditemukan",
		})
	})

	// METHOD DELETE
	http.HandleFunc("DELETE /categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID harus angka", http.StatusBadRequest)
			return
		}

		for i, cat := range categories {
			if cat.ID == id {
				categories = append(categories[:i], categories[i+1:]...)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Kategori berhasil dihapus",
				})
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Kategori tidak ditemukan",
		})
	})

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
