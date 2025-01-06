using client.Models;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace client.Internal.Interfaces
{
    public interface IRecommendationService
    {
        /// <summary>
        /// Получение рекомендаций
        /// </summary>
        Task<List<MovieModel>> Get();

        /// <summary>
        /// Добовление рекомендаций
        /// </summary>
        void Add(RecommendationAddPayload recommendation);
    }
}
