package Interfaces

import scala.concurrent.ExecutionContextExecutor
import akka.http.scaladsl.server.Route
import akka.actor.ActorSystem
import akka.stream.Materializer

trait MovieInterface {
    def forward_movie(implicit
        system: ActorSystem,
        materializer: Materializer,
        executionContext: ExecutionContextExecutor
    ): Route
}
