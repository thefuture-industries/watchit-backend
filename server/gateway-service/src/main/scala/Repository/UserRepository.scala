package Repository

import scala.concurrent.{ Future, ExecutionContext }
import akka.http.scaladsl.server.Directives._
import scala.concurrent.ExecutionContextExecutor
import akka.http.scaladsl.server.Route
import akka.http.scaladsl.model._
import akka.actor.ActorSystem
import akka.stream.Materializer

import Interfaces.{ UserInterface, ForwardInterface }
import Repository.{ ForwardRepository }
import Config.Config
import Services.{ LoggerService, BannedWordService }

class UserRepository(implicit ec: ExecutionContext) extends UserInterface {
    private val config                                = Config();
    private val _forward_repository: ForwardInterface = new ForwardRepository()

    def forward_user(implicit
        system: ActorSystem,
        materializer: Materializer,
        executionContext: ExecutionContextExecutor
    ): Route = {
        pathPrefix("api" / "v1" / "user") {
            extractRequest { request =>
                val newPath =
                    "/micro/user" + request.uri.path.toString().substring("/api/v1/user".length)
                val body    = request.entity.dataBytes.runFold("")(_ + _.utf8String)

                onSuccess(body) { bodyContent =>
                    if (
                      BannedWordService.is_request_suspicious(bodyContent) || BannedWordService
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
                        onComplete(
                          this._forward_repository.forward_to_backend(request, targetURL)
                        ) {
                            case scala.util.Success(response: HttpResponse) => complete(response)
                            case scala.util.Failure(ex)                     =>
                                LoggerService.logError(ex.getMessage)
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
    }
}
