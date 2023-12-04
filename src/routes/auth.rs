use actix_web::{web, App, HttpServer, HttpResponse, Responder};
use crate::database::user::validate_user;
use bson::{doc, Bson};

#[derive(Debug, serde::Serialize, serde::Deserialize)]
pub struct UserData {
    username: String,
    password: String,
}

pub async fn auth_user(data: web::Json<UserData>) -> HttpResponse {
    let filter = doc! {
        "username": Bson::String(data.username.clone()),
        "password": Bson::String(data.password.clone()),
    };

    match validate_user(filter).await {
        Ok(result) => {
            if result.success {
                let user_id_message = result.user_id.map_or(String::new(), |id| format!(" User ID: {}", id));
                HttpResponse::Ok().body(format!("{}{}", result.message, user_id_message))
            } else {
                HttpResponse::Unauthorized().body(result.message)
            }
        }
        Err(_) => HttpResponse::InternalServerError().body("Erro ao validar usuário"),
    }
}
