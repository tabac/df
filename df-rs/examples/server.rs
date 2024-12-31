use simple_logger::SimpleLogger;

use df_rs::server::DataFusionExecutorServerImpl;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    SimpleLogger::new().init()?;

    let server = DataFusionExecutorServerImpl::new();

    server.run().await?;

    Ok(())
}
