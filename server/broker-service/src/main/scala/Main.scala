import java.net.{ServerSocket, Socket}
import java.io.{BufferedReader, InputStreamReader, PrintWriter}
import com.fasterxml.jackson.databind.{JsonNode, ObjectMapper}

import Database.Database
import Packages.Logger

object TCPServer {
    val PORT = 8888
    val database = new Database()

    def main(args: Array[String]): Unit = {
        val serverSocket = new ServerSocket(PORT)
        database.createTableIfNotExists()

        println(s"Server is listening on port $PORT")
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
        val in = new BufferedReader(new InputStreamReader(socket.getInputStream))
        val out = new PrintWriter(socket.getOutputStream, true)

        try {
            if (in.ready()) {
                val receivedData = in.readLine()
                if (receivedData != null && receivedData.nonEmpty) {
                    database.insertData(receivedData)
                    out.println("Data received and stored")
                }
            }
            else {
                sendBrokersData(out)
            }
        } catch {
            case e: Exception =>
                Logger.logError(e.getMessage)
                out.println(s"""{"status": "error", "message": "${e.getMessage}"}""")
        } finally {
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
}
