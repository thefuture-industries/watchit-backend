import akka.http.scaladsl.Http
import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives._
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer

import scala.io.StdIn
import scala.concurrent.ExecutionContextExecutor

import Packages.{ Logger }
import Security.{ BannedWord }
import Networking.{ ClientHandler }
import Config.Config

object Main extends App {
    implicit val system: ActorSystem                        = ActorSystem("gateway-service")
    implicit val materializer: ActorMaterializer            = ActorMaterializer()
    implicit val executionContext: ExecutionContextExecutor = system.dispatcher

    val config = Config()

    val route = pathPrefix("api" / "v1" / "user") {
        extractRequest { request =>
            val newPath =
                "/micro/user" + request.uri.path.toString().substring("/api/v1/user".length)
            val body    = request.entity.dataBytes.runFold("")(_ + _.utf8String)

            onSuccess(body) { bodyContent =>
                if (
                  BannedWord.is_request_suspicious(bodyContent) || BannedWord
                      .is_request_suspicious(request.uri.toString())
                ) {
                    complete(
                      HttpResponse(
                        StatusCodes.Forbidden,
                        entity = "Suspicious content detected in request"
                      )
                    )
                } else {
                    val targetURL =
                        s"${config.security}://${config.usermicro_addr}:${config.usermicro_port}$newPath"
                    onComplete(ClientHandler.forward_to_backend(request, targetURL)) {
                        case scala.util.Success(response) => complete(response)
                        case scala.util.Failure(ex)       =>
                            complete(
                              HttpResponse(
                                StatusCodes.InternalServerError,
                                entity = "Failed to forward request"
                              )
                            )
                    }
                }
            }
        }
    }

    val bindingFuture = Http().bindAndHandle(route, config.server_addr, config.server_port.toInt)

    println(
      s"Server online at http://${config.server_addr}:${config.server_port}/\nPress RETURN to stop..."
    )
    StdIn.readLine()
    bindingFuture
        .flatMap(_.unbind())
        .onComplete(_ => system.terminate())
}

// Ограничени времени ответа от микросервиса
