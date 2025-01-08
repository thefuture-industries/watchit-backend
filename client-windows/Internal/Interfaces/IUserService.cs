using client.Models;
using System.Threading.Tasks;

namespace client.Internal.Interfaces
{
    public interface IUserService
    {
        /// <summary>
        /// Добовление/Вход пользователя
        /// </summary>
        void Add(UserAddPayload user);

        /// <summary>
        /// Получение uuid пользователя
        /// </summary>
        string GetUUID();

        /// <summary>
        /// Получение данных пользователя
        /// </summary>
        UserDataModel GetUserData();

        /// <summary>
        /// Обновление данных пользователя
        /// </summary>
        void Update(UserUpdatePayload user);
    }
}
