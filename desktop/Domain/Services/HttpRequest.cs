using System;
using System.Net.Http;
using System.Text;
using flick_finder.Domain.Core;
using flick_finder.Domain.Exceptions;
using flick_finder.Domain.Interfaces;
using Newtonsoft.Json;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace flick_finder.Domain.Services;

public class HttpRequest : IHttpRequest
{
    /// <summary>
    /// Создание клиента, для отправки запросов
    /// Этот клиент используется для отправки запросов на сервер и получения ответов.
    /// </summary>
    private readonly HttpClient client = new HttpClient();
    
    /// <summary>
    /// Инициализация конфига
    /// Содержит настройки, такие как URL сервера, ключи API и другие параметры,
    /// необходимые для работы с сервером.
    /// </summary>
    private readonly Config _config;
    
    /// <summary>
    /// Обработчик ошибок
    /// Используется для отображения сообщений об ошибках, возникающих во время работы с сервером.
    /// </summary>
    private readonly UIMessageException _uiexception;
    
    /// <summary>
    /// Констуктор класса
    /// </summary>
    public HttpRequest()
    {
        this._config = new Config();
        this._uiexception = new UIMessageException();
    }
    
    /// <summary>
    /// Функция отправки данных на сервер
    /// Используется любой body с любим типом
    /// </summary>
    public string SendRequest<T>(string route, string method, T body = default)
    {
        // Вылидация данных
        // Проверка что необходиммые данные не пустые
        if (string.IsNullOrEmpty(route) || string.IsNullOrEmpty(method))
        {
            throw new ArgumentNullException("Invalid Route or Method");
        }
        
        // Создание HttpRequestMessage
        var request = new HttpRequestMessage(new HttpMethod(method), this._config.ReturnConfig().SERVER_URL + route);
        
        // Серелизация данных
        string json = JsonSerializer.Serialize(body);
            
        // Установка контента
        var content = new StringContent(JsonConvert.SerializeObject(body), Encoding.UTF8, "application/json");
        request.Content = content;

        try
        {
            // Отправка запроса на сервер
            var response = this.client.SendAsync(request, HttpCompletionOption.ResponseContentRead).Result;

            // Обработка ошибок
            if (!response.IsSuccessStatusCode)
            {
                // Ошибка
                throw new Exception(response.Content.ReadAsStringAsync().Result);
            }
            else
            {
                // Все ОК
                return response.Content.ReadAsStringAsync().Result;
            }
        }
        catch (Exception ex)
        {
            this._uiexception.ShowError(ex.Message, "SERVER");
            return "";
        }
    }
    
    /// <summary>
    /// Функция отправки данных на сервер GET
    /// Не используется body для отпарвки на сервер
    /// Подходит для пустых GET запросов
    /// </summary>
    public string SendGETRequest(string route, string method)
    {
        // Вылидация данных
        if (string.IsNullOrEmpty(route) || string.IsNullOrEmpty(method))
        {
            throw new ArgumentNullException("Invalid Route or Method");
        }
        
        // Создание HttpRequestMessage
        var request = new HttpRequestMessage(new HttpMethod(method), this._config.ReturnConfig().SERVER_URL + route);

        // Отправка запроса на сервер
        try
        {
            var response = this.client.SendAsync(request, HttpCompletionOption.ResponseContentRead).Result;

            // Обработка ошибок
            if (!response.IsSuccessStatusCode)
            {
                // Ошибка
                return response.Content.ReadAsStringAsync().Result;
            }
            else
            {
                // Все ОК
                return response.Content.ReadAsStringAsync().Result;
            }
        }
        catch (Exception ex)
        {
            throw new ApplicationException(ex.Message);
        }
    }
}