package Interfaces

import scala.concurrent.ExecutionContextExecutor
import akka.http.scaladsl.server.Route
import akka.actor.ActorSystem
import akka.stream.Materializer

trait UserInterface {
    def forward_user(implicit
        system: ActorSystem,
        materializer: Materializer,
        executionContext: ExecutionContextExecutor
    ): Route
}
