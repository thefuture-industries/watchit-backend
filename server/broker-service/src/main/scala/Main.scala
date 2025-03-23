import scala.collection.mutable
import java.net.{ ServerSocket, Socket }
import java.io.{ BufferedReader, InputStreamReader, PrintWriter }
import com.fasterxml.jackson.databind.{ JsonNode, ObjectMapper }

import Database.Database
import Packages.Logger
import Config.Config

object TCPServer {
    val PORT     = 8888
    val database = new Database()
    val config   = Config()

    val clients = mutable.Set[PrintWriter]()

    def main(args: Array[String]): Unit = {
        val serverSocket = new ServerSocket(PORT)
        database.createTableIfNotExists()

        println(
          "\n" + """██████╗ ██████╗  ██████╗ ██╗  ██╗███████╗██████╗     ███████╗ ██████╗ █████╗ ██╗      █████╗
██╔══██╗██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝██╔══██╗    ██╔════╝██╔════╝██╔══██╗██║     ██╔══██╗
██████╔╝██████╔╝██║   ██║█████╔╝ █████╗  ██████╔╝    ███████╗██║     ███████║██║     ███████║
██╔══██╗██╔══██╗██║   ██║██╔═██╗ ██╔══╝  ██╔══██╗    ╚════██║██║     ██╔══██║██║     ██╔══██║
██████╔╝██║  ██║╚██████╔╝██║  ██╗███████╗██║  ██║    ███████║╚██████╗██║  ██║███████╗██║  ██║
╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝    ╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
                                                                                             """
        )
        println(s"Server is listening on port $PORT")
        println(
          s"\nServer online at ${config.security}://${config.server_addr}:${config.server_port}/\nPress RETURN to stop..."
        )
        while (true) {
            val socket = serverSocket.accept()
            Logger.logSocket("", "", "", "New client connected")
            println("New client connected")
            new ClientHandler(socket, database).start()
        }
    }
}

class ClientHandler(socket: Socket, database: Database) extends Thread {
    val objectMapper = new ObjectMapper()

    override def run(): Unit = {
        val in  = new BufferedReader(new InputStreamReader(socket.getInputStream, "UTF-8"))
        val out =
            new PrintWriter(socket.getOutputStream, true, java.nio.charset.StandardCharsets.UTF_8)

        // Добавляем клиента в список
        TCPServer.clients.add(out)

        try {
            if (in.ready()) {
                val receivedData = in.readLine()
                if (receivedData != null && receivedData.nonEmpty) {
                    println("DATA PAYLOAD!!!")
                    database.insertData(receivedData)
                    notifyClients(s"""$receivedData""")
                    out.println("Data received and stored")
                }
            } else {
                sendBrokersData(out)
            }
        } catch {
            case e: Exception =>
                Logger.logError(e.getMessage)
                out.println(s"""{"status": "error", "message": "${e.getMessage}"}""")
        } finally {
            TCPServer.clients.remove(out)
            out.close()
            in.close()
            socket.close()
        }

        // if (!in.ready()) {
        //     sendBrokersData(out)
        // } else {
        //     val receivedData = in.readLine()
        //     database.insertData(receivedData)
        //     out.println("Data received and stored")
        // }

        socket.close()
    }

    def sendBrokersData(out: PrintWriter): Unit = {
        val brokersData = database.getBrockerData()

        val json_data = objectMapper.writeValueAsString(brokersData)
        out.println(json_data)
    }

    def notifyClients(message: String): Unit = {
        TCPServer.clients.foreach { client =>
            client.println(message)
        }
    }
}
