package Config

case class Config(
    security: String = "http",
    server_addr: String = "127.0.0.1",
    server_port: String = "8080"
)

object Config {
    def apply(): Config = new Config()
}
