using client.Internal.Interfaces;
using client.Internal.Services;
using client.Models;
using client.Services;
using System.Collections.Generic;
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

        public MovieService()
        {
            this._exception = new UIException();
            this._cacheQuery = new CacheQuery();
        }

        public async Task<List<MovieModel>> Get()
        {
            var movies = await _cacheQuery.Get("/movies/popular");
            if (movies.IsError)
            {
                this._exception.Error("Server Error", movies.Error);
                return null;
            }

            var movieJSON = System.Text.Json.JsonSerializer.Serialize(movies.Result);

            return System.Text.Json.JsonSerializer.Deserialize<List<MovieModel>>(movieJSON);
        }

        /// <summary>
        /// Получение фильмов по фильтрам
        /// </summary>
        public Task<List<MovieModel>> GetByFilters(string search, string genre, string date)
        {
            throw new System.NotImplementedException();
        }

        public Task<List<MovieModel>> GetBySearch(string search)
        {
            throw new System.NotImplementedException();
        }

        public Task<MovieModel> GetDetails(int id)
        {
            throw new System.NotImplementedException();
        }

        public Task<List<MovieModel>> GetPlot(string text, string lege)
        {
            throw new System.NotImplementedException();
        }

        public Task<List<MovieModel>> GetSimilar(int[] genre_ids, string title, string overview)
        {
            throw new System.NotImplementedException();
        }
    }
}
