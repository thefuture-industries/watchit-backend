package Database

import java.sql.{Connection, DriverManager, PreparedStatement, ResultSet, Statement, Timestamp}
import com.fasterxml.jackson.databind.{JsonNode, ObjectMapper}
import scala.collection.mutable.ListBuffer

case class Broker(id: Int, sender: String, received: String, payload: JsonNode, status: String, created_at: Timestamp)

class Database {
    val DB_URL = "jdbc:postgresql://localhost:5432/flicksfi"
    val DB_USER = "postgres"
    val DB_PASSWORD = "password"

    private val objectMapper = new ObjectMapper()
    private val conn: Connection = connectDB()

    private def connectDB(): Connection = {
        DriverManager.getConnection(DB_URL, DB_USER, DB_PASSWORD)
    }

    def createTableIfNotExists(): Unit = {
        val sql = """CREATE TABLE IF NOT EXISTS brokers (
                    |id SERIAL PRIMARY KEY,
                    |sender VARCHAR(100),
                    |receiver VARCHAR(100),
                    |payload JSONB,
                    |status VARCHAR(20) NOT NULL DEFAULT 'sent',
                    |created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP)""".stripMargin
        val stmt = conn.createStatement()
        stmt.executeUpdate(sql)
        stmt.close()
    }

    def getBrockerData(): List[Broker] = {
        val sql = "SELECT * FROM brokers"
        val stmt: Statement = conn.createStatement()
        val rs: ResultSet = stmt.executeQuery(sql)

        val brokers = ListBuffer[Broker]()
        while (rs.next()) {
            val id = rs.getInt("id")
            val sender = rs.getString("sender")
            val receiver = rs.getString("receiver")
            val payload = objectMapper.readTree(rs.getString("payload"))
            val status = rs.getString("status")
            val created_at = rs.getTimestamp("created_at")

            brokers += Broker(id, sender, receiver, payload, status, created_at)
        }

        println(brokers.toList)

        rs.close()
        stmt.close()
        brokers.toList
    }

    def insertData(jsonString: String): Unit = {
        try {
            val jsonNode: JsonNode = objectMapper.readTree(jsonString)
            val sender = jsonNode.get("payload").asText()
            val receiver = jsonNode.get("payload").asText()
            val payload = jsonNode.get("payload").asText()

            val sql = "INSERT INTO brokers (sender, receiver, payload) VALUES (?, ?, ?::jsonb)"
            val pstmt: PreparedStatement = conn.prepareStatement(sql)
            pstmt.setString(1, sender)
            pstmt.setString(2, receiver)
            pstmt.setString(3, payload)
            pstmt.executeUpdate()
            pstmt.close()
        } catch {
            case e: Exception => println(s"Error parsing JSON: ${e.getMessage}")
        }
    }
}
