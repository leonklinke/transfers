use serde::Deserialize;
use std::io::{Error};

#[path = "../domain/key.rs"] pub mod key;
#[path = "../infra/db.rs"] mod db;

#[derive(Deserialize)]
pub struct CreateKey {
    pub user_id: String,
    pub key: String,
}

pub async fn create_key(key: CreateKey) -> Result<i32, Error> {
    // let key = Key::new(key.user_id, key.key);
    println!("starting db {}", key.user_id);
    let db = db::new_db().await.unwrap();
    println!("listing dbs");
    db::list_databases(db).await;
    println!("end");
    Ok(2)
}


