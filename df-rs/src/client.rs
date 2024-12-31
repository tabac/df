use std::panic;

use tokio_stream::StreamExt;
use tonic::transport::Channel;

use crate::{
    df::{data_fusion_executor_client::DataFusionExecutorClient, ExecuteQueryRequest},
    server::DataFusionExecutorNetwork,
};

#[derive(Debug, Clone)]
pub struct DataFusionExecutorClientImpl {
    client: DataFusionExecutorClient<Channel>,
}

impl DataFusionExecutorClientImpl {
    pub async fn new(
        network: DataFusionExecutorNetwork,
    ) -> Result<Self, Box<dyn std::error::Error>> {
        let dst = match network {
            DataFusionExecutorNetwork::Tcp => String::from("http://[::1]:50051"),

            DataFusionExecutorNetwork::Unix => {
                // Probably not possible through Unix Domain Socket.
                // Look at the following:
                //  - https://github.com/softprops/hyperlocal/tree/main
                //  - https://github.com/hyperium/tonic/issues/742
                panic!("df-rs: unimplemented")
            }
        };

        let client = DataFusionExecutorClient::connect(dst).await?;

        Ok(DataFusionExecutorClientImpl { client })
    }

    pub async fn run(&mut self, id: u64) -> Result<(), Box<dyn std::error::Error>> {
        log::info!("df-rs: Running client.");

        let request = ExecuteQueryRequest { id };

        let mut stream = self.client.execute_query(request).await?.into_inner();

        while let Some(response) = stream.next().await {
            let response = response?;

            log::info!(
                "df-rs: Client got response: ({}, {}, {})",
                id,
                response.id,
                response.request_id
            )
        }

        Ok(())
    }
}
