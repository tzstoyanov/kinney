"""Trivial implementation of Kinney orchestrator."""

# TODO(cody): move this into a common library for Python modules with __main__.
import rootpath
rootpath.append()

from concurrent import futures
import grpc
from grpc_reflection.v1alpha import reflection
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
        # Pro tempore: just grab point from first request.
        for r in requests:
            p = r.point
            break
        cmd = pb.ChargerCommand()
        cmd.point = p
        cmd.limit = math.inf
        cmd.lifetime.seconds = 300
        yield cmd


class _LoggingInterceptor(grpc.ServerInterceptor):
    """gRPC interceptor that logs request contents for debugging."""

    def intercept_service(self, continuation, req):
        """Log one incoming gRPC request.

        Args:
            continuation: callable receiving request info
            req: request info (not payload)
        Returns:
            value returned from continuation
        """
        logging.info("Received call: %r", req)
        return continuation(req)


def main():
    logging.basicConfig(level=logging.INFO)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10),
                         interceptors=[_LoggingInterceptor()])
    o = Orchestrator()
    svc.add_OrchestratorServicer_to_server(o, server)

    reflectable = [
        pb.DESCRIPTOR.services_by_name['Orchestrator'].full_name,
        reflection.SERVICE_NAME,
    ]
    reflection.enable_server_reflection(reflectable, server)

    server.add_insecure_port('[::]:8191')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    main()
