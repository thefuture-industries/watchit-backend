using client.Internal.Interfaces;
using client.Models;
using client.Services;
using client.ViewModel;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace client.Internal.Services
{
    public class RecommendationService : IRecommendationService
    {
        /// <summary>
        /// Инициализация ошибок UI
        /// </summary>
        private readonly UIException _exception;

        /// <summary>
        /// Создание клиента
        /// </summary>
        private readonly HttpClient _httpClient;

        /// <summary>
        /// Подключение класса конфига
        /// </summary>
        private readonly Config _config;

        private readonly IUserService _userService;

        private readonly CacheQuery _cacheQuery;

        public RecommendationService()
        {
            this._exception = new UIException();
            this._httpClient = new HttpClient();
            this._config = new Config();
            this._userService = new UserService();
            this._cacheQuery = new CacheQuery();
        }

        /// <summary>
        /// Добовление рекомендаций
        /// </summary>
        public async void Add(RecommendationAddPayload recommendation)
        {
            try
            {
                string json = JsonSerializer.Serialize(recommendation);
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._httpClient.PostAsync($"{this._config.ReturnConfig().SERVER_URL}/recommendations", content);
                response.EnsureSuccessStatusCode();
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Error server");
                return;
            }
        }

        /// <summary>
        /// Получение рекомендаций
        /// </summary>
        public async Task<List<MovieModel>> Get()
        {
            try
            {
                string uuid = this._userService.GetUUID();

                var response = await this._cacheQuery.Get($"{this._config.ReturnConfig().SERVER_URL}/recommendations/{uuid}");
                if (response.IsError)
                {
                    this._exception.Error(response.Error, "Error Server");
                    return null;
                }

                var movieJSON = System.Text.Json.JsonSerializer.Serialize(response.Result);
                return System.Text.Json.JsonSerializer.Deserialize<List<MovieModel>>(movieJSON);

                /*var response = await this._httpClient.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/recommendations/{uuid}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());*/
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Error server");
                return null;
            }
        }
    }
}
