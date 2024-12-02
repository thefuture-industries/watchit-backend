namespace flick_finder.Domain.Interfaces;

public interface IHttpRequest
{
    /// <summary>
    /// Функция отправки данных на сервер
    /// Используется любой body с любим типом
    /// </summary>
    string SendRequest<T>(string route, string method, T body = default);

    /// <summary>
    /// Функция отправки данных на сервер GET
    /// Не используется body для отпарвки на сервер
    /// Подходит для пустых GET запросов
    /// </summary>
    string SendGETRequest(string route, string method);
}