package Services

import java.io._
import java.text.SimpleDateFormat
import java.util.Date

object LoggerService {
    private val logDir     = new File("logs")
    private val dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss")

    private var currentLogDir = getCurrentLogDir()

    private def getCurrentLogDir(): File = {
        val dateFormatFolder = new SimpleDateFormat("dd.MM")
        val folderName       = dateFormatFolder.format(new Date())
        val dir              = new File(logDir, folderName)
        if (!dir.exists()) dir.mkdirs()
        dir
    }

    private def updateLogDirIfNeeded(): Unit = {
        val newLogDir = getCurrentLogDir()
        if (newLogDir != currentLogDir) {
            currentLogDir = newLogDir
        }
    }

    def logServer(
        clientIP: String,
        method: String,
        url: String,
        status: String,
        duration: Long,
        payload: String,
        message: String = ""
    ): Unit = {
        updateLogDirIfNeeded()
        val logFile_SOCKET = new File(currentLogDir, "server.log")

        val timestamp  = dateFormat.format(new Date())
        val body       = if (payload.nonEmpty) {
            payload.map(b => s" | Body: ${b.toString().take(999)}...").mkString
        } else { "" }
        val logMessage =
            f"[$timestamp%-20s] Device: $clientIP%-15s | Method: $method%-5s | URL: $url%-30s | Status: $status%-5s | Response Time: ${duration}ms | Body: $body | Message: $message"

        val writer = new PrintWriter(new FileWriter(logFile_SOCKET, true))
        try {
            writer.println(logMessage)
        } finally {
            writer.close()
        }
    }

    def logError(error_message: String): Unit = {
        updateLogDirIfNeeded()
        val logFile_ERROR = new File(currentLogDir, "errors.log")

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
