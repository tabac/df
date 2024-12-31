use std::pin::Pin;

use tokio_stream::Stream;
use tonic::transport::Server;
use tonic::{Request, Response, Status};

use crate::df::data_fusion_executor_server::{DataFusionExecutor, DataFusionExecutorServer};
use crate::df::{
    CreateSessionRequest, CreateSessionResponse, ExecuteQueryRequest, ExecuteQueryResponse,
};

#[derive(Debug, Default)]
pub struct DataFusionExecutorServerImpl {}

impl DataFusionExecutorServerImpl {
    pub fn new() -> Self {
        DataFusionExecutorServerImpl::default()
    }

    pub async fn run(self) -> Result<(), Box<dyn std::error::Error>> {
        log::info!("df-rs: Running server.");

        let addr = "[::1]:50051".parse()?;

        Server::builder()
            .add_service(DataFusionExecutorServer::new(self))
            .serve(addr)
            .await?;

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
