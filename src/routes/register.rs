use actix_web::{web, App, HttpServer, HttpResponse, Responder};
use crate::database::user::create_user;

#[derive(Debug, serde::Serialize, serde::Deserialize)]
pub struct RegisterData {
    pub name: String,
    pub username: String,
    pub password: String,
}

pub async fn register_user(data: web::Json<RegisterData>) -> HttpResponse {
    let user = RegisterData {
        name: data.name.clone(),
        username: data.username.clone(),
        password: data.password.clone(),
    };
    // Chama a função create_user do arquivo connection.rs
    match create_user(user).await {
        Ok(result) => {
            if result.success {
                HttpResponse::Ok().json(result)
            } else {
                HttpResponse::InternalServerError().json(result)
            }
        }
        Err(_) => HttpResponse::InternalServerError().body("Erro ao registrar usuário"),
    }
}