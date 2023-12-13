use mongodb::{Client, options::ClientOptions};

// List the names of the databases in that deployment.

pub async fn new_db() -> Result<Client, mongodb::error::Error> {
    // Parse a connection string into an options struct.
    let mut client_options = ClientOptions::parse("mongodb://root:notasecret@localhost:27017").await?;

    // Manually set an option.
    client_options.app_name = Some("transfers".to_string());

    // Get a handle to the deployment.
    let client = Client::with_options(client_options)?;

    for db_name in client.list_database_names(None, None).await? {
        println!("{}", db_name);
    }
    
    return Ok(client)
}
pub async fn list_databases(db: Client) -> Result<(), mongodb::error::Error> {
    let dbs = db.list_database_names(None, None).await?;
    println!("dbs {}", dbs.len());
    for db_name in dbs {
        println!("{}", db_name);
    }
    Ok(())
}