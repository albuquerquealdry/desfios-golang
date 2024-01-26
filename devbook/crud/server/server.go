package server

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("[ERROR-INTERFACE][Erro body json]"))
		return
	}

	var user user

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro struct Unmarshal]"))
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro connect database]"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome,email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro build statement]"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro insert data]"))
		return

	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro get id the insert]"))
		return

	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("[INFO][Success insert user %d]", idInsert)))
	fmt.Println(user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Connect()
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Error Close Connect Database]"))
		return
	}
	defer db.Close()

	line, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Error Exec Query into Database]"))
		return
	}

	defer line.Close()

	var users []user

	for line.Next() {
		var user user
		if erro := line.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("[ERROR-SYSTEM][Error Scan the user]"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Erro json encode]"))
		return
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, erro := strconv.ParseUint(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Error Parameter covert to int]"))
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Error Close Connection Database]"))
		return

	}
	defer db.Close()

	line, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("[ERROR-SYSTEM][Error search user in database]"))
		return

	}

	var user user
	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("[ERROR-SYSTEM][Erro Scan user]"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		if user.Name == "" {
			w.Write([]byte("[ERROR-SYSTEM][User not found]"))
			return
		}
		w.Write([]byte("[ERROR-SYSTEM][Erro Convert user in json]"))
		return
	}

}
