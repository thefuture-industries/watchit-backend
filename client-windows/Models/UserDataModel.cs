using System.Text.Json.Serialization;

namespace client.Models
{
    public class UserDataModel
    {
        [JsonPropertyName("uuid")]
        public string UUID { get; set; }

        [JsonPropertyName("username")]
        public string UserName { get; set; }

        [JsonPropertyName("email")]
        public string Email { get; set; }
    }
}
