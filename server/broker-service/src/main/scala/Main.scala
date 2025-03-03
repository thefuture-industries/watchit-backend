import java.net.{ServerSocket, Socket}
import java.io.{BufferedReader, InputStreamReader, PrintWriter}
import com.fasterxml.jackson.databind.{JsonNode, ObjectMapper}

object TCPServer {
    val PORT = 8888
    val database = new Database()

    def main(args: Array[String]): Unit = {
        val serverSocket = new ServerSocket(PORT)
        database.createTableIfNotExists()

        println(s"Server is listening on port $PORT")
        while (true) {
            val socket = serverSocket.accept()
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

        if (!in.ready()) {
            sendBrokersData(out)
        } else {
            val receivedData = in.readLine()
            database.insertData(receivedData)
            out.println("Data received and stored")
        }

        socket.close()
    }

    def sendBrokersData(out: PrintWriter): Unit = {
        val brokersData = database.getBrockerData()
        brokersData.foreach(out.println)
    }
}
