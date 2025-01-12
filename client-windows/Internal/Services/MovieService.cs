using client.Internal.Interfaces;
using client.Internal.Services;
using client.Models;
using client.Services;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace client.API
{
    public class MovieService : IMovieService
    {
        /// <summary>
        /// Инициализация ошибок UI
        /// </summary>
        private readonly UIException _exception;

        /// <summary>
        /// Инициализация запросов к серверу с кешом
        /// </summary>
        private readonly CacheQuery _cacheQuery;

        /// <summary>
        /// Класс для отправки запросов
        /// </summary>
        private readonly HttpClient _httpClient;

        /// <summary>
        /// Класс конфиг приложения
        /// </summary>
        private readonly Config _config;

        /// <summary>
        /// Сервис для работы с пользователем
        /// </summary>
        private readonly IUserService _userService;

        public MovieService()
        {
            this._exception = new UIException();
            this._cacheQuery = new CacheQuery();
            this._config = new Config();
            this._userService = new UserService();
        }

        /// <summary>
        /// Получение фильмов по фильтрам
        /// </summary>
        public async Task<List<MovieModel>> GetByFilters(string search, string genre, string date)
        {
            try
            {
                var response = await this._httpClient.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/movies?s={search}&genre_id={genre}&date={date}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Server Error");
                return null;
            }
        }

        /// <summary>
        /// Поиск фильма по titla или overview (...)
        /// </summary>
        public async Task<List<MovieModel>> GetBySearch(string search)
        {
            try
            {
                var response = await this._httpClient.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/movies?s={Uri.EscapeDataString(search)}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Server Error");
                return null;
            }
        }

        /// <summary>
        /// Получние деталей фильма по ID
        /// </summary>
        public async Task<MovieModel> GetDetails(int id)
        {
            try
            {
                var response = await this._httpClient.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/movie/{id}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<MovieModel>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Server Error");
                return null;
            }
        }

        /// <summary>
        /// Получение похожих фильмов
        /// </summary>
        public async Task<List<MovieModel>> GetPlot(string text, string lege)
        {
            try
            {
                string json = JsonSerializer.Serialize(new
                {
                    uuid = this._userService.GetUUID(),
                    text = text,
                    lege = lege
                });
                var content = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await this._httpClient.PostAsync($"{this._config.ReturnConfig().SERVER_URL}/text/movies", content);
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Server Error");
                return null;
            }
        }

        /// <summary>
        /// Поиск фильмов по сюжету
        /// </summary>
        public async Task<List<MovieModel>> GetSimilar(int[] genre_ids, string title, string overview)
        {
            try
            {
                var response = await this._httpClient.GetAsync($"{this._config.ReturnConfig().SERVER_URL}/movies/similar?genre_id={string.Join(",", genre_ids.Select(x => x.ToString()))}&title={Uri.EscapeDataString(title)}&overview={Uri.EscapeDataString(overview)}");
                response.EnsureSuccessStatusCode();

                return JsonSerializer.Deserialize<List<MovieModel>>(await response.Content.ReadAsStringAsync());
            }
            catch (HttpRequestException ex)
            {
                this._exception.Error(ex.Message, "Network or HTTP Error");
                return null;
            }
            catch (Exception ex)
            {
                this._exception.Error(ex.Message, "Server Error");
                return null;
            }
        }
    }
}
