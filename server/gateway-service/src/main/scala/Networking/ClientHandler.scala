package Networking

import akka.actor.ActorSystem
import akka.http.scaladsl.model._
import akka.http.scaladsl.Http
import akka.stream.Materializer
import akka.http.scaladsl.settings.ConnectionPoolSettings

import scala.concurrent.{ ExecutionContext, Future }
import scala.concurrent.duration._

import Packages.Logger

object ClientHandler {
    def forward_to_backend(request: HttpRequest, targetURL: String)(implicit
        system: ActorSystem,
        mat: Materializer,
        ec: ExecutionContext
    ): Future[HttpResponse] = {
        val clientIP = request.headers
            .find(_.name == "X-Forwarded-For")
            .map(_.value)
            .getOrElse(request.uri.authority.host.address())
        Logger.logServer(
          clientIP,
          request.method.value,
          request.uri.path.toString(),
          "200 (Forwarded)",
          s"We forward the request to the microservice: $targetURL"
        )

        val filteredHeaders = request.headers.filterNot(_.name == "Timeout-Access")

        val pool = ConnectionPoolSettings(system)
            .withIdleTimeout(3.seconds)

        Http()
            .singleRequest(
              HttpRequest(
                method = request.method,
                uri = targetURL,
                headers = request.headers,
                entity = request.entity
              ),
              settings = pool
            )
            .recover { case ex: Exception =>
                HttpResponse(
                  StatusCodes.InternalServerError,
                  entity = "Failed to send request to microservice"
                )
            }
    }
}
