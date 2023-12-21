package router

import (
	"app-gerenciamento-financeiro/src/database"
	"app-gerenciamento-financeiro/src/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	// Configurar rotas aqui
	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/authenticate", AuthenticateHandler).Methods("POST")

	return router
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n<--------------------------------------------------------------->\n")
	fmt.Println("👋 Uma conexão foi efetuada -> POST /register.")
	// Verificando se a requisição é do tipo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificando o corpo JSON da requisição
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	resp, err := database.RegisterUser(newUser.Name, newUser.Username, newUser.Password, newUser.Email)
	if resp == primitive.NilObjectID {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Falha no registro do Usuário: ", err)
		fmt.Fprintln(w, "Falha no registro do Usuário: ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Println("👋 Usuário cadastrado com sucesso!", resp)
		fmt.Fprintln(w, "Usuário cadastrado com sucesso!")
	}
}

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	userData := models.UserData{
		UserID: "123",
	}

	fmt.Println("\n<--------------------------------------------------------------->\n")
	fmt.Println("👋 Uma conexão foi efetuana -> POST /authenticate.")
	// Verificando se a requisição é do tipo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w)
	}
	var newAuthUser models.AuthUser
	err := json.NewDecoder(r.Body).Decode(&newAuthUser)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "None")
	}

	userID, err := database.AuthenticateUser(newAuthUser.Username, newAuthUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("👋 Falha no registro do Usuário.")
		// fmt.Fprintln(w, "Falha na registro do Usuário")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println("👋 Autenticação bem-sucedida!")
		userData.UserID = userID
		// fmt.Fprintln(w, "Usuário autenticado com sucesso")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
	resp, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, "Erro ao converter para JSON", http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
