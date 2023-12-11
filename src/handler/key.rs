use axum::{
    http::StatusCode,
    Json,
};

#[path = "../use_case/key.rs"] pub mod key;

pub async fn create_key(Json(payload): Json<key::CreateKey>,
) -> StatusCode {    
    let _ = key::create_key(payload);

    StatusCode::CREATED
}