using client.Services;
using System;
using System.Net.Http;
using System.Text.Json;
using System.Runtime.Caching;
using System.Threading.Tasks;

namespace client.Internal.Services
{
    public class CacheQueryModel
    {
        public object Result { get; set; }
        public bool IsLoading { get; set; }
        public bool IsError { get; set; }
        public string Error { get; set; }
    }

    public class CacheQuery
    {
        /// <summary>
        /// Класс кеширование
        /// </summary>
        private MemoryCache _cache = MemoryCache.Default;

        /// <summary>
        /// Создание клиента
        /// </summary>
        private readonly HttpClient _client;

        /// <summary>
        /// Подключение класса конфига
        /// </summary>
        private readonly Config _config;

        public CacheQuery()
        {
            this._client = new HttpClient();
            this._config = new Config();
        }

        /// <summary>
        /// Отправка GET запросов
        /// </summary>
        public async Task<CacheQueryModel> Get(string route)
        {
            // Получение кеша и проверка на его существование
            var cache_data = this._cache.Get($"{this._config.ReturnConfig().SERVER_URL}{route}");
            if (cache_data != null)
            {
                var result = JsonSerializer.Deserialize<dynamic>(cache_data as string);

                return new CacheQueryModel()
                {
                    Result = result,
                    IsLoading = false,
                    IsError = false,
                    Error = null
                };
            }

            try
            {
                var response = await this._client.GetAsync($"{this._config.ReturnConfig().SERVER_URL}{route}");
                response.EnsureSuccessStatusCode();

                var result = System.Text.Json.JsonSerializer.Deserialize<dynamic>(await response.Content.ReadAsStringAsync());

                this._cache.Set($"{this._config.ReturnConfig().SERVER_URL}{route}", await response.Content.ReadAsStringAsync(), DateTimeOffset.UtcNow.AddHours(1));
                return new CacheQueryModel()
                {
                    Result = result,
                    IsLoading = false,
                    IsError = false,
                    Error = null,
                };
            }
            catch (Exception ex)
            {
                return new CacheQueryModel()
                {
                    Result = null,
                    IsLoading = false,
                    IsError = true,
                    Error = ex.Message,
                };
            }
        }
    }
}
