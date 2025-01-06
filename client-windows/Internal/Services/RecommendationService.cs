using client.Internal.Interfaces;
using client.Models;
using client.Services;
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
        private readonly HttpClient _client;

        /// <summary>
        /// Подключение класса конфига
        /// </summary>
        private readonly Config _config;

        private readonly IUserService _userService;

        public RecommendationService()
        {
            this._exception = new UIException();
            this._client = new HttpClient();
            this._config = new Config();
            this._userService = new UserService();
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

                var response = await this._client.PostAsync($"{this._config.ReturnConfig().SERVER_URL}/recommendations", content);
                response.EnsureSuccessStatusCode();
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error("Network or HTTP Error", ex.Message);
                return;
            }
            catch (Exception ex)
            {
                this._exception.Error("Error server", ex.Message);
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

                var response = await this._client.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/recommendations/{uuid}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error("Network or HTTP Error", ex.Message);
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error("Error server", ex.Message);
                return null;
            }
        }
    }
}
