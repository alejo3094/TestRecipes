package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

// Application define Router for handler and database
type Application struct {
	Router *mux.Router
	BD     *sql.DB
}

const queryCreateTable = `CREATE TABLE IF NOT EXISTS recipe (id integer NOT NULL DEFAULT nextval('recipe_id_seq') PRIMARY KEY ,
name TEXT NOT NULL,
ingredients TEXT NOT NULL,
preparation TEXT NOT NULL,
time TEXT NOT NULL
)
`
const queryAlterTableSeq = `ALTER SEQUENCE recipe_id_seq OWNED BY recipe.id;`

// init data bd and Router with handler
func (app *Application) initialize() {

	// Connect to database.
	var err error
	app.BD, err = sql.Open("postgres", "postgresql://root@localhost:26257?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	} else {

		// create sequencues
		if _, err := app.BD.Exec("CREATE SEQUENCE IF NOT EXISTS recipe_id_seq"); err != nil {
			log.Fatal(err)
		}
		/*

			if _, err := app.BD.Exec(queryAlterTableSeq); err != nil {
				log.Fatal(err)
			} */
		// create table Receipe
		if _, err := app.BD.Exec(queryCreateTable); err != nil {
			log.Fatal(err)
		}
	}
	//app.Router = mux.NewRouter().StrictSlash(false)
	app.Router = mux.NewRouter()

	app.assignUrls()

}

// handlerFuncs for api

func (app *Application) obtainRecipes(w http.ResponseWriter, r *http.Request) {

	log.Println("--- all recipes ----")

	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	log.Println(count)
	log.Println(start)

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	recipes, err := obtainRecipes(app.BD, start, count)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, recipes)
}

func (app *Application) obtainRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("--- recipe by id ----")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, " Invalid ID for Recipe")
		return
	}

	re := []Recipe{Recipe{ID: id}}
	if err := re[0].obtainRecipe(app.BD); err != nil {
		switch err {
		case sql.ErrNoRows:
			responseWithError(w, http.StatusNotFound, "Recipe not found")
		default:
			responseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responseWithJSON(w, http.StatusOK, re)
}

func (app *Application) countRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("--- recipe count")
	count, err := countRecips(app.BD)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, count)
}

func (app *Application) obtainRecipeName(w http.ResponseWriter, r *http.Request) {

	log.Println("--- recipe by name ----")
	vars := mux.Vars(r)
	name := strings.ToLower(vars["name"])

	recipes, err := obtainRecipesByName(app.BD, name)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, recipes)

}

func (app *Application) createRecipe(w http.ResponseWriter, r *http.Request) {

	log.Println("--- create ----")
	var re Recipe
	//log.Println(r)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&re); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	if err := re.createRecipe(app.BD); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusCreated, re)
}

func (app *Application) updateRecipe(w http.ResponseWriter, r *http.Request) {

	log.Println("--- update recipe ----")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Recipe ID")
		return
	}

	var re Recipe
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&re); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	defer r.Body.Close()
	re.ID = id

	if err := re.updateRecipe(app.BD); err != "ok" {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}

	responseWithJSON(w, http.StatusOK, re)
}

func (app *Application) deleteRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("-------------delete recipe--------------")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Recipe Id")
		return
	}

	re := Recipe{ID: id}
	if err := re.deleteRecipe(app.BD); err != "ok" {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	setupResponse(&w)
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

func (app *Application) assignUrls() {
	app.Router.HandleFunc("/api/recipe/", app.obtainRecipes).Methods("GET")
	app.Router.HandleFunc("/api/recipe/{id:[0-9]+}", app.obtainRecipe).Methods("GET")       // by id
	app.Router.HandleFunc("/api/recipe/{name:[a-z]+}", app.obtainRecipeName).Methods("GET") // by name
	app.Router.HandleFunc("/api/recipe/", app.getOptions).Methods("OPTIONS")
	app.Router.HandleFunc("/api/recipe/{id:[0-9]+}", app.getOptions).Methods("OPTIONS")
	app.Router.HandleFunc("/api/recipe/", app.createRecipe).Methods("POST")
	app.Router.HandleFunc("/api/recipe/count/", app.countRecipe).Methods("GET")
	app.Router.HandleFunc("/api/recipe/{id:[0-9]+}", app.updateRecipe).Methods("PUT")
	app.Router.HandleFunc("/api/recipe/{id:[0-9]+}", app.deleteRecipe).Methods("DELETE")
}

func (app *Application) runServer() {
	server := &http.Server{
		Addr:           ":8090",
		Handler:        app.Router,
		ReadTimeout:    10 * time.Second, // 10 segundos
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //en megas operador binario para multiplicar 1 * 2 20 veces
	}

	log.Println("Listening....")
	log.Fatal(server.ListenAndServe())

}

func (app *Application) getOptions(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Methods, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Connection", "Keep-Alive")
}

func setupResponse(w *http.ResponseWriter) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	(*w).Header().Set("Content-Type", "application/json")
}
