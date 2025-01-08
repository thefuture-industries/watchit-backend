using client.Internal.Interfaces;
using client.Models;
using client.Services;
using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;

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

        private UserDataModel _userData;

        public UserService()
        {
            this._exception = new UIException();
            this._client = new HttpClient();
            this._config = new Config();
        }

        /// <summary>
        /// Добовление/Вход пользователя
        /// </summary>
        public async void Add(UserAddPayload user)
        {
            try
            {
                string json = JsonSerializer.Serialize(user);
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._client.PostAsync($"{this._config.ReturnConfig().SERVER_URL}/user/add", content);
                response.EnsureSuccessStatusCode();

                UserDataModel user_data = JsonSerializer.Deserialize<UserDataModel>(await response.Content.ReadAsStringAsync());
                this._userData = user_data;
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

        public UserDataModel GetUserData()
        {
            return this._userData;
        }

        /// <summary>
        /// Получение uuid пользователя
        /// </summary>
        public string GetUUID()
        {
            return this._userData.UUID;
        }

        /// <summary>
        /// Обновление данных пользователя
        /// </summary>
        public async void Update(UserUpdatePayload user)
        {
            try
            {
                string json = JsonSerializer.Serialize(user);
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._client.PutAsync($"{this._config.ReturnConfig().SERVER_URL}/user/update", content);
                response.EnsureSuccessStatusCode();

                UserDataModel user_data = JsonSerializer.Deserialize<UserDataModel>(await response.Content.ReadAsStringAsync());
                this._userData = user_data;
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
    }
}
