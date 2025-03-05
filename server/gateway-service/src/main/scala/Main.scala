import akka.http.scaladsl.Http
import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives._
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer

import scala.io.StdIn
import scala.concurrent.ExecutionContextExecutor

import Repository.{ UserRepository, MovieRepository }

import Services.{ LoggerService, BannedWordService }
import Config.Config

object Main extends App {
    implicit val system: ActorSystem                        = ActorSystem("gateway-service")
    implicit val materializer: ActorMaterializer            = ActorMaterializer()
    implicit val executionContext: ExecutionContextExecutor = system.dispatcher

    val config = Config()

    val routes = concat(
      new UserRepository().forward_user,
      new MovieRepository().forward_movie
    )

    val bindingFuture = Http().bindAndHandle(routes, config.server_addr, config.server_port.toInt)

    println("""   ______      __                              _____                 _
  / ____/___ _/ /____ _      ______ ___  __   / ___/___  ______   __(_)_______
 / / __/ __ `/ __/ _ \ | /| / / __ `/ / / /   \__ \/ _ \/ ___/ | / / / ___/ _ \
/ /_/ / /_/ / /_/  __/ |/ |/ / /_/ / /_/ /   ___/ /  __/ /   | |/ / / /__/  __/
\____/\__,_/\__/\___/|__/|__/\__,_/\__, /   /____/\___/_/    |___/_/\___/\___/
                                  /____/                                       """)
    println(
      s"\nServer online at ${config.security}://${config.server_addr}:${config.server_port}/\nPress RETURN to stop..."
    )
    StdIn.readLine()
    bindingFuture
        .flatMap(_.unbind())
        .onComplete(_ => system.terminate())
}
