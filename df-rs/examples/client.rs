use simple_logger::SimpleLogger;
use tokio::task::JoinHandle;

use df_rs::client::DataFusionExecutorClientImpl;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    SimpleLogger::new().init()?;

    let client = DataFusionExecutorClientImpl::new().await?;

    let handles: Vec<JoinHandle<()>> = (0..5)
        .map(|i| {
            let mut client = client.clone();

            tokio::spawn(async move {
                client.run(i).await.unwrap();
            })
        })
        .collect();

    for handle in handles {
        handle.await?;
    }

    Ok(())
}
