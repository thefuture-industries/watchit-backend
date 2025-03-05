package Interfaces

import akka.actor.ActorSystem
import akka.stream.Materializer
import akka.http.scaladsl.model.{ HttpRequest, HttpResponse }
import scala.concurrent.{ ExecutionContext, Future }

trait ForwardInterface {
    def forward_to_backend(request: HttpRequest, targetURL: String)(implicit
        system: ActorSystem,
        mat: Materializer,
        ec: ExecutionContext
    ): Future[HttpResponse]
}
