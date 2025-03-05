package Packages

import java.io._
import java.text.SimpleDateFormat
import java.util.Date

object Logger {
    private val logDir         = new File("logs")
    private val logFile_ERROR  = new File(logDir, "errors.log")
    private val logFile_SOCKET = new File(logDir, "socket.log")
    private val dateFormat     = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss")

    if (!logDir.exists()) logDir.mkdirs()

    def logSocket(sender: String, receiver: String, payload: String, message: String = ""): Unit = {
        val timestamp  = dateFormat.format(new Date())
        val logMessage =
            f"[$timestamp%-20s] Sender: $sender Receiver: $receiver | Payload: $payload | Message: $message"

        val writer = new PrintWriter(new FileWriter(logFile_SOCKET, true))
        try {
            writer.println(logMessage)
        } finally {
            writer.close()
        }
    }

    def logError(error_message: String): Unit = {
        val timestamp  = dateFormat.format(new Date())
        val logMessage = f"[$timestamp%-20s] ERROR: $error_message"

        val writer = new PrintWriter(new FileWriter(logFile_ERROR, true))
        try {
            writer.println(logMessage)
        } finally {
            writer.close()
        }
    }
}
