use actix_web::{web, HttpResponse};
use bson::{Document, doc, Bson, oid::ObjectId};
use mongodb::{Client, options::ClientOptions, Collection};
use mongodb::error::Error as MongoError;
use log::{info, error};
use crate::routes::auth::UserData;
use crate::routes::register::RegisterData;
use dotenv::dotenv;

pub struct GetUser {
    pub success: bool,
    pub user: Document,
    pub message: String,
}

pub struct AuthResult {
    pub success: bool,
    pub user_id: Option<ObjectId>,
    pub name: String,
    pub message: String,
}

#[derive(Debug, serde::Serialize, serde::Deserialize)]
pub struct RegisterResult {
    pub success: bool,
    pub message: String,
}

async fn create_connection() -> Result<Collection<Document>, mongodb::error::Error> {
    dotenv().ok();

    let mongodb_url = std::env::var("MONGODB_URL").expect("MONGODB_URL não definida");
    let database_name = std::env::var("DATABASE_NAME").expect("DATABASE_NAME não definida");
    let users_collection = std::env::var("USERS_COLLECTION").expect("USERS_COLLECTION não definida");
    
    let client_options = ClientOptions::parse(&mongodb_url).await?;
    let client = Client::with_options(client_options)?;
    let db = client.database(&database_name);

    Ok(db.collection::<bson::Document>(&users_collection))
}

pub async fn validate_user(filter: Document) -> Result<(AuthResult), MongoError> {
    let collection = create_connection().await?;

    match collection.find_one(filter.clone(), None).await {
        Ok(Some(doc)) => {
            info!("Usuário validado com sucesso!");
            let user_id = doc.get_object_id("_id").ok();
            let name = doc.get_str("name").unwrap_or_default();
            Ok(AuthResult {
                success: true,
                user_id,
                name: name.to_string(),
                message: "Usuário validado com sucesso!".to_string(),
            })
        }
        Ok(None) => {
            info!("Usuário não encontrado ou senha incorreta");
            Ok(AuthResult {
                success: false,
                user_id: None,
                name: String::new(),
                message: "Usuário ou senha inválidos".to_string(),
            })
        }
        Err(e) => {
            error!("Erro durante a validação do usuário: {:?}", e);
            Err(e)
        }
    }
}


pub async fn create_user(user: RegisterData) -> Result<RegisterResult, MongoError> {
    let collection = create_connection().await?;

    // Verifica se o username já existe na base de dados
    if collection.find_one(doc! {"username": &user.username}, None).await?.is_some() {
        info!("Username já existente");
        return Ok(RegisterResult {
            success: false,
            message: "Username já existente".to_string(),
        });
    }

    // Cria um novo documento BSON para o usuário
    let user_doc = doc! {
        "name": Bson::String(user.name),
        "username": Bson::String(user.username),
        "password": Bson::String(user.password),
    };

    // Insere o novo usuário na coleção 'usuarios'
    collection.insert_one(user_doc, None).await?;

    info!("Usuário criado com sucesso!");
    Ok(RegisterResult {
        success: true,
        message: "Usuário criado com sucesso".to_string(),
    })
}

pub async fn get_user(filter: Document) -> Result<(GetUser), MongoError> {
    let collection = create_connection().await?;
    match collection.find_one(filter.clone(), None).await {
        Ok(Some(doc)) => {
            info!("Usuário encontrado com sucesso!");
            Ok(GetUser {
                success: true,
                user: doc,
                message: "Usuário validado com sucesso!".to_string(),
            })
        }
        Ok(None) => {
            info!("Usuário não encontrado ou senha incorreta");
            Ok(GetUser {
                success: false,
                user: Document::new(),
                message: "Usuário ou senha inválidos".to_string(),
            })
        }
        Err(e) => {
            error!("Erro durante a validação do usuário: {:?}", e);
            Err(e)
        }
    }
}