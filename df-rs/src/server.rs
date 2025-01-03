use std::pin::Pin;

use tokio::net::UnixListener;
use tokio_stream::wrappers::UnixListenerStream;
use tokio_stream::Stream;
use tonic::transport::Server;
use tonic::{Request, Response, Status};

use crate::df::data_fusion_executor_server::{DataFusionExecutor, DataFusionExecutorServer};
use crate::df::{
    CreateSessionRequest, CreateSessionResponse, ExecuteQueryRequest, ExecuteQueryResponse,
};

#[derive(Debug)]
pub enum DataFusionExecutorNetwork {
    Tcp,
    Unix,
}

#[derive(Debug)]
pub struct DataFusionExecutorServerImpl {
    network: DataFusionExecutorNetwork,
    address: String,
}

impl DataFusionExecutorServerImpl {
    pub fn new(network: DataFusionExecutorNetwork) -> Self {
        let address = match network {
            DataFusionExecutorNetwork::Tcp => String::from("127.0.0.1:50051"),
            DataFusionExecutorNetwork::Unix => String::from("/tmp/df.sock"),
        };

        DataFusionExecutorServerImpl { network, address }
    }

    pub async fn run(self) -> Result<(), Box<dyn std::error::Error>> {
        log::info!("df-rs: Running server ({:?}).", self.network);

        match self.network {
            DataFusionExecutorNetwork::Tcp => {
                let addr = self.address.parse()?;

                Server::builder()
                    .add_service(DataFusionExecutorServer::new(self))
                    .serve(addr)
                    .await?;
            }
            DataFusionExecutorNetwork::Unix => {
                let listener = UnixListener::bind(&self.address)?;

                let stream = UnixListenerStream::new(listener);

                Server::builder()
                    .add_service(DataFusionExecutorServer::new(self))
                    .serve_with_incoming(stream)
                    .await?;
            }
        }

        Ok(())
    }
}

type ResponseStream = Pin<Box<dyn Stream<Item = Result<ExecuteQueryResponse, Status>> + Send>>;

#[tonic::async_trait]
impl DataFusionExecutor for DataFusionExecutorServerImpl {
    type ExecuteQueryStream = ResponseStream;

    async fn create_session(
        &self,
        _: Request<CreateSessionRequest>,
    ) -> Result<Response<CreateSessionResponse>, Status> {
        let response = CreateSessionResponse {};

        Ok(Response::new(response))
    }

    async fn execute_query(
        &self,
        request: Request<ExecuteQueryRequest>,
    ) -> Result<Response<Self::ExecuteQueryStream>, Status> {
        let request_id = request.get_ref().id;

        let stream =
            tokio_stream::iter((0..5).map(move |id| Ok(ExecuteQueryResponse { id, request_id })));

        Ok(Response::new(Box::pin(stream) as Self::ExecuteQueryStream))
    }
}
