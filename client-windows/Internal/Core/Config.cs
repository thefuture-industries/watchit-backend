using Newtonsoft.Json;

namespace client.Services
{
    public class AppConfig
    {
        [JsonProperty("SERVER_URL")]
        public string SERVER_URL { get; }

        public AppConfig()
        {
            this.SERVER_URL = "https://flicksfi-production.up.railway.app/api/v1";
        }
    }

    public class Config
    {
        public AppConfig ReturnConfig()
        {
            return new AppConfig();
        }
    }
}
