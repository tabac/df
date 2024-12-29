package server

import "github.com/tabac/df"

var _ df.DataFusionExecutor = DataFusionExecutorServer{}

type DataFusionExecutorServer struct {
}
