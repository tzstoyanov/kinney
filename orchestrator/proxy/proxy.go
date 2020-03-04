// proxy converts incoming HTTP REST calls into Orchestrator gRPCs to a backend.
package main

import (
  "context"
  "flag"
  "net/http"

  "github.com/golang/glog"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

  gw "github.com/CamusEnergy/kinney/orchestrator"
)

var backendFlag = flag.String("backend",  "localhost:8191", "Backend server network endpoint")

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterOrchestratorHandlerFromEndpoint(ctx, mux,  *backendFlag, opts)
  if err != nil {
    return err
  }

  return http.ListenAndServe(":8190", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}
