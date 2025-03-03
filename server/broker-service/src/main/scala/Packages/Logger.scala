package Packages

import java.io._
import java.text.SimpleDateFormat
import java.util.Date

object Logger {
    private val logFile = new File("logs/app.log")
    private val dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss")

    def logToFile(sender: String, receiver: String, payload: String): Unit = {
        val timestamp = dateFormat.format(new Date())
        val logMessage = f"[$timestamp%-20s] Sender: $sender Receiver: $receiver | Payload: $payload"

        val writer = new PrintWriter(new FileWriter(logFile, true))
        try {
            writer.println(logMessage)
        } finally {
            writer.close()
        }
    }
}