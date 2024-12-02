using System.Text.Json.Serialization;

namespace flick_finder.Domain.Models;

// UserModel: Модель пользователя описывает его поля
// UserName, Directory, OS, Email
public class UserModel
{
    // Это свойство представляет имя пользователя, которое используется для идентификации пользователя в системе.
    // Оно должно быть уникальным для каждого пользователя
    [JsonPropertyName("username")]
    public string UserName { get; set; }
    
    // Это свойство представляет каталог пользователя, который хранит файлы и настройки пользователя.
    // Default значение директория MyDocuments
    [JsonPropertyName("directory")]
    public string Directory { get; set; }

    // Это свойство представляет операционную систему, используемую пользователем.
    // Пример для Window: Microsoft Windows NT 10.0.22631.0
    [JsonPropertyName("os")]
    public string OS { get; set; }
    
    // Электронная почта пользователя.
    // Это свойство представляет адрес электронной почты пользователя.
    [JsonPropertyName("email")]
    public string Email { get; set; }
}