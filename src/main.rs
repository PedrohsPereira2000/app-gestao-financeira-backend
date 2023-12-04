use actix_web::{web, App, HttpServer, HttpResponse, Responder};
mod routes;
mod database;
use routes::auth::auth_user;
use routes::register::register_user;
use log::info;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init();

    HttpServer::new(|| {
        App::new()
            .route("/register", web::post().to(register_user))
            .route("/user", web::post().to(auth_user))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}