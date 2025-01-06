using client.Models;
using System.Threading.Tasks;

namespace client.Internal.Interfaces
{
    public interface IUserService
    {
        /// <summary>
        /// Добовление/Вход пользователя
        /// </summary>
        Task<string> Add(UserAddPayload user);

        /// <summary>
        /// Получение uuid пользователя
        /// </summary>
        string GetUUID();

        /// <summary>
        /// Обновление данных пользователя
        /// </summary>
        Task<string> Update(UserUpdatePayload user);
    }
}
