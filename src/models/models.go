package models

type User struct {
	Name     string `bson:"name"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}

// User representa a estrutura de dados para um usuário
type AuthUser struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type UserData struct {
	UserID string `json:"user_id"`
}
