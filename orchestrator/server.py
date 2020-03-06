"""Trivial implementation of Kinney orchestrator."""

# TODO(cody): move this into a common library for Python modules with __main__.
import rootpath
rootpath.append()

import concurrent.futures as futures
import grpc
import logging
import math
import orchestrator.api_pb2 as pb
import orchestrator.api_pb2_grpc as svc


class Orchestrator(svc.OrchestratorServicer):
    """gRPC service that exchanges status for commands."""

    def Charger(self, requests, context):
        """Receive a stream of updates from chargers, send back commands.

        Args:
            requests: iterable of pb.ChargerSession objects
            context: RPC tracing context
        Returns
            iterable of pb.ChargerCommand objects
        """
        # Just grab the first point for now.
        for r in requests:
            p = r.point
            break
        cmd = pb.ChargerCommand()
        cmd.point = p
        cmd.limit = math.inf
        cmd.lifetime.seconds = 300
        yield cmd


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    o = Orchestrator()
    svc.add_OrchestratorServicer_to_server(o, server)
    server.add_insecure_port('[::]:8191')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
