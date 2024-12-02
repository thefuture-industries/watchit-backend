using flick_finder.Domain.Models;

namespace flick_finder.Domain.Interfaces;

public interface IUser
{
    /// <summary>
    /// Получение данных пользователя
    /// Получение зависит от модели UserModel
    /// </summary>
    UserModel GetUser();
}