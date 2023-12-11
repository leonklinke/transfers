use axum::{
    routing::{get, post},
    Router,
};
pub mod key;

pub async fn serve() {
    let app: Router = Router::new()
        .route("/keys", post(key::create_key));
    
    println!("Listening on http://localhost:8080");
    
    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

