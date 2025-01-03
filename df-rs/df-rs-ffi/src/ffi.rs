use tokio::runtime;

use simple_logger::SimpleLogger;

use df_rs::server::{DataFusionExecutorNetwork, DataFusionExecutorServerImpl};

#[no_mangle]
pub extern "C" fn start_tcp_ffi() {
    start_tcp();
}

#[no_mangle]
pub extern "C" fn start_unix_ffi() {
    start_unix();
}

pub fn start_tcp() {
    start_server(DataFusionExecutorNetwork::Tcp);
}

pub fn start_unix() {
    start_server(DataFusionExecutorNetwork::Unix);
}

fn start_server(network: DataFusionExecutorNetwork) {
    SimpleLogger::new().init().unwrap();

    let server = DataFusionExecutorServerImpl::new(network);

    let rt = runtime::Runtime::new().unwrap();

    rt.block_on(async {
        server.run().await.unwrap();
    });
}
