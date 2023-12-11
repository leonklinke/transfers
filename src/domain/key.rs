use serde::Deserialize;

#[derive(Deserialize)]
pub struct Key {
    pub id: String,
    pub user_id: String,
    pub key: String,
}