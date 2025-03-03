import java.io._
import java.net.Socket
import com.fasterxml.jackson.databind.node.ObjectNode
import com.fasterxml.jackson.databind.ObjectMapper

object TCPClient {
    val SERVER_ADDR = "localhost"
    val SERVER_PORT = 8888
    val objectMapper = new ObjectMapper()

    def main(args: Array[String]): Unit = {
        val socket = new Socket(SERVER_ADDR, SERVER_PORT)
        val in = new BufferedReader(new InputStreamReader(socket.getInputStream, "UTF-8"))
        val out = new PrintWriter(socket.getOutputStream, true, java.nio.charset.StandardCharsets.UTF_8)

        try {
            println("Connected to the server.")

            // 1️⃣ Получаем данные от сервера
            println("Receiving data from server...")
            val receivedData = in.readLine()
            if (receivedData != null && receivedData.nonEmpty) {
                println(s"📥 Данные от сервера: $receivedData")
            } else {
                println("⚠️ Нет данных от сервера")
            }

            // 2️⃣ Создаём JSON-объект для отправки
            val jsonNode = objectMapper.createObjectNode()
            jsonNode.put("sender", "ClientA")
            jsonNode.put("receiver", "Server")
            jsonNode.put("payload", """{"message": "Hello, Server!"}""")

            val jsonString = objectMapper.writeValueAsString(jsonNode)

            // 3️⃣ Отправляем данные
            println(s"📤 Отправляем данные: $jsonString")
            out.println(jsonString)
            out.flush()

            // 4️⃣ Получаем ответ от сервера
            val serverResponse = in.readLine()
            println(s"🔄 Ответ от сервера: $serverResponse")
        } catch {
            case e: Exception =>
                println("Error: " + e)
        } finally {
            in.close()
            out.close()
            socket.close()
        }
    }
}
