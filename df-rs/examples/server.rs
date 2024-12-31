use std::env;

use simple_logger::SimpleLogger;

use df_rs::server::{DataFusionExecutorNetwork, DataFusionExecutorServerImpl};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    SimpleLogger::new().init()?;

    let args: Vec<String> = env::args().collect();

    let network = if args.len() == 2 && args[1] == "unix" {
        DataFusionExecutorNetwork::Unix
    } else {
        DataFusionExecutorNetwork::Tcp
    };

    let server = DataFusionExecutorServerImpl::new(network);

    server.run().await?;

    Ok(())
}
