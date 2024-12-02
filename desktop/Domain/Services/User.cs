using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Models;

namespace flick_finder.Domain.Services;

public class User : IUser
{
    /// <summary>
    /// Инициализация модели пользователя
    /// Позволяет хранить информацию в отдельном классе
    /// </summary>
    private UserModel userGlobal = new UserModel()
    {
        UserName = Environment.UserName,
        Directory = Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments),
        OS = Environment.OSVersion.VersionString,
        Email = $"{Environment.UserName}_{Environment.OSVersion.VersionString}@gmail.com",
    };
    
    /// <summary>
    /// Получение данных пользователя
    /// Получение зависит от модели UserModel
    /// </summary>
    public UserModel GetUser()
    {
        return this.userGlobal;
    }
}