package Models

import java.sql.{Timestamp}
import com.fasterxml.jackson.databind.{JsonNode}

case class Broker(id: Int, sender: String, received: String, payload: JsonNode, status: String, created_at: Timestamp)
