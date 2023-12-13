use serde::Deserialize;
use std::io::{Error};
use mongodb::bson::{doc, Document};

#[path = "../domain/key.rs"] pub mod key;
#[path = "../infra/db.rs"] mod db;

#[derive(Deserialize)]
pub struct CreateKey {
    pub user_id: String,
    pub key: String,
}

pub async fn create_key(key: CreateKey) -> Result<i32, Error> {
    // let key = Key::new(key.user_id, key.key);
    println!("starting db");
    let client_db = db::new_db().await.unwrap();
    
    let key_collection = client_db.collection::<Document>("keys");

    let key = doc!{
        "user_id": key.user_id,
        "key": key.key
    };

    key_collection.insert_one(key, None).await.unwrap();

    Ok(1)
}


