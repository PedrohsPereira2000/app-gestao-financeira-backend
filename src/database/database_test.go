package database

import (
    "testing"
	"fmt"
    "github.com/stretchr/testify/assert"
)

// func TestRegisterUser(t *testing.T) {
//     // Chame a função RegisterUser
//     name := "John Doe"
//     username := "johndoe"
//     password := "password123"
//     email := "johndoe@example.com"
//     objectID, err := RegisterUser(name, username, password, email)
//     assert.NoError(t, err)

    
// }

func TestCreateConnection(t *testing.T) {
	// Chame a função createConnection
	collection, err := createConnection()
	fmt.Println("Falha ao conectar: ", err)
	fmt.Println("Falha no Collection: ", collection)
	assert.NoError(t, err)
	assert.NotNil(t, collection)
}