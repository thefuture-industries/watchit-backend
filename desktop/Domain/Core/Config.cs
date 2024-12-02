using System.Text.Json.Serialization;

namespace flick_finder.Domain.Core;

/// <summary>
/// Класс модель, для хранения переменных
/// </summary>
public class AppConfig
{
    // Это свойство содержит адрес сервера API,
    // к которому отправляются HTTP-запросы.
    [JsonPropertyName("SERVER_URL")]
    public string SERVER_URL { get; set; }

    // Это свойство содержит IP-адрес прокси-сервера, который используется
    // для маршрутизации HTTP-запросов.
    [JsonPropertyName("PROXY_SERVER_IP")]
    public string PROXY_SERVER_IP { get; set; }
    
    // Это свойство содержит дополнительный параметр, который передается
    // прокси-серверу.
    [JsonPropertyName("PROXY_SERVER_PARAM")]
    public string PROXY_SERVER_PARAM { get; set; }

    /// <summary>
    /// Инициализация переменных
    /// </summary>
    public AppConfig()
    {
        this.SERVER_URL = "http://localhost:8080/api/v1/";
        this.PROXY_SERVER_IP = "51.159.195.58";
        this.PROXY_SERVER_PARAM = "__cpo=aHR0cHM6Ly9pbWFnZS50bWRiLm9yZw";
    }
}

/// <summary>
/// Вывод переменных из конфига
/// </summary>
public class Config
{
    public AppConfig ReturnConfig()
    {
        return new AppConfig();
    }
}