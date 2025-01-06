using client.Internal.Interfaces;
using client.Models;
using client.Services;
using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace client.Internal.Services
{
    public class UserService : IUserService
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

        private string _uuid = null;

        public UserService()
        {
            this._exception = new UIException();
            this._client = new HttpClient();
            this._config = new Config();
        }

        /// <summary>
        /// Добовление/Вход пользователя
        /// </summary>
        public async Task<string> Add(UserAddPayload user)
        {
            try
            {
                string json = JsonSerializer.Serialize(user);
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._client.PostAsync($"{this._config.ReturnConfig().SERVER_URL}/user/add", content);
                response.EnsureSuccessStatusCode();

                this._uuid = await response.Content.ReadAsStringAsync();
                return await response.Content.ReadAsStringAsync();
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

        /// <summary>
        /// Получение uuid пользователя
        /// </summary>
        public string GetUUID()
        {
            return this._uuid;
        }

        /// <summary>
        /// Обновление данных пользователя
        /// </summary>
        public async Task<string> Update(UserUpdatePayload user)
        {
            try
            {
                string json = JsonSerializer.Serialize(user);
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._client.PutAsync($"{this._config.ReturnConfig().SERVER_URL}/user/update", content);
                response.EnsureSuccessStatusCode();

                this._uuid = await response.Content.ReadAsStringAsync();
                return await response.Content.ReadAsStringAsync();
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
