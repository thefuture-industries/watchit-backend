import org.scalatest.funsuite.AnyFunSuite

object BannedWord {
    def is_request_suspicious(request: String): Int = {
        val forbiddenPatterns = Seq("<script>")
        if (forbiddenPatterns.exists(request.contains)) {
            1
        } else {
            0
        }
    }
}

class BannedWordSpec extends AnyFunSuite {
    val testCases = Seq(
      ("Suspicious URL with script in path", "http://localhost:8080/api/v1/user/sync/<script>", 1),
      (
        "Suspicious query parameter with script",
        "http://localhost:8080/api/v1/user/sync?param=<script>",
        1
      ),
      (
        "Suspicious body with script tag",
        "POST /api/v1/user/create HTTP/1.1\r\n\r\n<script>alert('XSS');</script>",
        1
      ),
      ("Clean request", "http://localhost:8080/api/v1/user/sync", 0),
      ("Clean POST request", "POST /api/v1/user/create HTTP/1.1\r\n\r\nvalid data here", 0)
    )

    testCases.foreach { case (name, request, expectedResult) =>
        test(name) {
            assert(BannedWord.is_request_suspicious(request) == expectedResult)
        }
    }
}
