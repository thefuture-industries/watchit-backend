using client.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace client.Internal.Interfaces
{
    public interface IMovieService
    {
        /// <summary>
        /// Получение фильмов по фильтрам
        /// </summary>
        Task<List<MovieModel>> GetByFilters(string search, string genre, string date);

        /// <summary>
        /// Поиск фильма по titla или overview (...)
        /// </summary>
        Task<List<MovieModel>> GetBySearch(string search);

        /// <summary>
        /// Получние деталей фильма по ID
        /// </summary>
        Task<MovieModel> GetDetails(int id);

        /// <summary>
        /// Получение похожих фильмов
        /// </summary>
        Task<List<MovieModel>> GetSimilar(int[] genre_ids, string title, string overview);

        /// <summary>
        /// Поиск фильмов по сюжету
        /// </summary>
        Task<List<MovieModel>> GetPlot(string text, string lege);
    }
}
