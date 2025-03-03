package Config

case class Config(
    security: String = "http",
    server_addr: String = "127.0.0.1",
    server_port: String = "8080",

    usermicro_addr: String = "127.0.0.1",
    usermicro_port: String = "8001",
)

object Config {
  def apply(): Config = new Config()
}
