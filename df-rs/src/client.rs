use tokio_stream::StreamExt;
use tonic::transport::Channel;

use crate::df::{data_fusion_executor_client::DataFusionExecutorClient, ExecuteQueryRequest};

#[derive(Debug, Clone)]
pub struct DataFusionExecutorClientImpl {
    client: DataFusionExecutorClient<Channel>,
}

impl DataFusionExecutorClientImpl {
    pub async fn new() -> Result<Self, Box<dyn std::error::Error>> {
        let client = DataFusionExecutorClient::connect("http://[::1]:50051").await?;

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
