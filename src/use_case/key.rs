use serde::Deserialize;
use std::io::{Error};

#[path = "../domain/key.rs"] pub mod key;

#[derive(Deserialize)]
pub struct CreateKey {
    pub user_id: String,
    pub key: String,
}

pub fn create_key(key: CreateKey) -> Result<i32, Error> {
    // let key = Key::new(key.user_id, key.key);
    Ok(2)
}


