using System.Text.Json.Serialization;

namespace client.Models
{
    public class UserAddPayload
    {
        [JsonPropertyName("secret_word")]
        public string SecretWord { get; set; }

        [JsonPropertyName("ip_address")]
        public string IpAddress { get; set; }

        [JsonPropertyName("country")]
        public string Country { get; set; }

        [JsonPropertyName("region_name")]
        public string RegionName { get; set; }

        [JsonPropertyName("zip")]
        public string Zip { get; set; }
    }
}
