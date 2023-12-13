use mongodb::{Client, Database, options::ClientOptions};
use std::env;

// List the names of the databases in that deployment.

pub async fn new_db() -> Result<Database, mongodb::error::Error> {
    // Parse a connection string into an options struct.
    let url = env::var("MONGODB_URI").expect("MONGODB_URL env not set");
    let mut client_options = ClientOptions::parse(url).await?;

    // Manually set an option.
    client_options.app_name = Some("transfers".to_string());

    // Get a handle to the deployment.
    let client = Client::with_options(client_options)?;

    let db = client.database("transfers");
    
    return Ok(db)
}
pub async fn list_databases(db: Client) -> Result<(), mongodb::error::Error> {
    let dbs = db.list_database_names(None, None).await?;
    for db_name in dbs {
        println!("{}", db_name);
    }
    Ok(())
}