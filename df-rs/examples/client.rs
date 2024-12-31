use std::env;

use simple_logger::SimpleLogger;
use tokio::task::JoinHandle;

use df_rs::{client::DataFusionExecutorClientImpl, server::DataFusionExecutorNetwork};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    SimpleLogger::new().init()?;

    let args: Vec<String> = env::args().collect();

    let dst = if args.len() == 2 && args[1] == "unix" {
        DataFusionExecutorNetwork::Unix
    } else {
        DataFusionExecutorNetwork::Tcp
    };

    let client = DataFusionExecutorClientImpl::new(dst).await?;

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
