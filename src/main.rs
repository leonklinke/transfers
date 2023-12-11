#[path = "handler/router.rs"] mod router;

#[tokio::main]
async fn main() {
    router::serve().await;
}