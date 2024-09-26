package api

import (
	"bytes"
	"encoding/json"
	"etp/pkg/db"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	r  *mux.Router // маршрутизатор запросов
	db *db.DB      // база данных
}

func New(db *db.DB) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.Use(api.headersMiddleware)
	api.r.HandleFunc("/pre_search/{article}", api.preSearchesHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/pre_search", api.newPreSearchHandler).Methods(http.MethodPost)
	//api.r.HandleFunc("/orders/{id}", api.updateOrderHandler).Methods(http.MethodPatch)
	//api.r.HandleFunc("/orders/{id}", api.deleteOrderHandler).Methods(http.MethodDelete)
}

func (api *API) preSearchesHandler(w http.ResponseWriter, r *http.Request) {
	//preSearches := api.db.PreSearches()
	article := mux.Vars(r)["article"]
	var jsonStr = []byte(fmt.Sprintf(`{"Request":{"RequestData":{"article":"%s"}}}`, article))
	body := bytes.NewBuffer(jsonStr)
	url := fmt.Sprintf("%s/%s", os.Getenv("QWEP_URL"), "preSearch")
	req, _ := http.NewRequest("POST", url, body)
	token := fmt.Sprintf("Bearer %s", os.Getenv("QWEP_TOKEN"))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var j map[string]map[string]map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&j)
	if err != nil {
		panic(err)
	}
	responseErrors := j["Response"]["errors"]
	if responseErrors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body error: %s", err.Error())
		//errorMessage := responseErrors[0]["message"]
	} else {
		json.NewEncoder(w).Encode(j["Response"]["entity"]["shortNumbers"])
	}
}

func (api *API) newPreSearchHandler(w http.ResponseWriter, r *http.Request) {
	var ps db.PreSearch
	err := json.NewDecoder(r.Body).Decode(&ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := api.db.NewPreSearch(ps)
	w.Write([]byte(strconv.Itoa(id)))
}
