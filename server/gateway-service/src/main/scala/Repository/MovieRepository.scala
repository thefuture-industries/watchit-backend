package Repository

import scala.concurrent.{ Future, ExecutionContext }
import akka.http.scaladsl.server.Directives._
import scala.concurrent.ExecutionContextExecutor
import akka.http.scaladsl.server.Route
import akka.http.scaladsl.model._
import akka.actor.ActorSystem
import akka.stream.Materializer

import Interfaces.{ MovieInterface, ForwardInterface }
import Repository.{ ForwardRepository }
import Services.{ LoggerService, BannedWordService }
import Config.Config

class MovieRepository(implicit ec: ExecutionContext) extends MovieInterface {
    private val config                                = Config();
    private val _forward_repository: ForwardInterface = new ForwardRepository()

    def forward_movie(implicit
        system: ActorSystem,
        materializer: Materializer,
        executionContext: ExecutionContextExecutor
    ): Route = {
        pathPrefix("api" / "v1" / "movie") {
            extractRequest { request =>
                val newPath =
                    "/micro/movie" + request.uri.path.toString().substring("/api/v1/movie".length)
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
                            s"${config.security}://${config.moviemicro_addr}:${config.moviemicro_port}$newPath"
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
