package database

import (
	"app-gerenciamento-financeiro/src/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createConnection() (*mongo.Collection, error) {
	// Configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongo1:27017/")

	// Conectando ao MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao MongoDB!")

	// Retornando a coleção
	collection := client.Database("gestao-financeira").Collection("users")
	return collection, nil
}

func RegisterUser(name string, username string, password string, email string) (primitive.ObjectID, error) {
	collection, err := createConnection()
	if err != nil {
		// log.Output(err)
		return primitive.NilObjectID, err
	}

	// Criando um novo documento de usuário
	newUser := models.User{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	resp, err_usr := FindUser(username)
	if resp == false {
		// Inserindo o novo usuário no banco de dados
		result, err := collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			return primitive.NilObjectID, err
		}

		// Obtendo o _id do usuário cadastrado
		objectID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			return primitive.NilObjectID, fmt.Errorf("Falha ao obter o _id do usuário cadastrado")
		}

		return objectID, nil
	} else {
		fmt.Println("Usuário já cadastrado, tente recuperar sua senha: ", err_usr)
		return primitive.NilObjectID, err_usr
	}

}

func FindUser(username string) (bool, error) {
	collection, err := createConnection()
	if err != nil {
		// log.Output(err)
		return false, err
	}

	filter := bson.M{"username": username}

	// Resultado da busca
	var result map[string]interface{}

	// Executando a busca
	err_search := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err_search != nil {
		// log.Fatal(err_search)
		return false, err_search
	} else {
		return true, nil
	}
}

func AuthenticateUser(username string, password string) (string, error) {

	collection, err := createConnection()
	if err != nil {
		// log.Output(err)
		return "None", err
	}

	filter := bson.M{"username": username, "password": password}

	// Resultado da busca
	var result map[string]interface{}

	// Executando a busca
	err_search := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err_search != nil {
		// log.Fatal(err_search)
		return "None", err_search
	}

	name, found := result["name"].(string)
	if !found {
		fmt.Println("Campo 'name' não encontrado no documento")
		// log.Fatal("Campo 'name' não encontrado no documento")
	}
	// Exibindo o documento encontrado
	fmt.Println("Documento encontrado: ", name)

	// Obtendo o valor do campo _id como um ObjectID
	objectID, found := result["_id"].(primitive.ObjectID)
	resp := objectID.Hex()
	if !found {
		fmt.Println("Campo '_id' não encontrado no documento ou não é um ObjectID")
		// log.Fatal("Campo '_id' não encontrado no documento ou não é um ObjectID")
		resp = "None"
	}
	return resp, nil
}
