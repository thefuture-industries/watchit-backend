using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Models;
using System.Text.Json;

namespace flick_finder.Domain.Services;

public class YouTube : IYouTube
{
    private YouTubeModel[] youtube_videos = null;
    
    /// <summary>
    /// Инициализация запросов на сервер
    /// </summary>
    private readonly IHttpRequest _http;

    public YouTube()
    {
        this._http = new HttpRequest();
    }

    /// <summary>
    /// Запись в память текущии видео
    /// Данные зависят от модели YouTubeModel
    /// </summary>
    private void AddVideos(YouTubeModel[] videos)
    {
        // Создаем новый массив с увеличенной длиной
        // Либо с длиной массива videos
        int length = this.youtube_videos == null ? videos.Length : this.youtube_videos.Length + videos.Length;

        // Создание нового массива с типом моделе YouTubeModel
        // С длинной length
        YouTubeModel[] new_videos = new YouTubeModel[length];

        if (this.youtube_videos != null)
        {
            Array.Copy(this.youtube_videos, new_videos, this.youtube_videos.Length);
        }
        
        Array.Copy(videos, 0, new_videos, this.youtube_videos == null ? 0 : this.youtube_videos.Length, videos.Length);

        this.youtube_videos = new_videos;
    }

    /// <summary>
    /// Добовление в массив данные 3 популярных видео с ютуба
    /// </summary>
    public void SetVideos()
    {
        try
        {
            string jsonVideos = this._http.SendGETRequest("youtube/video/popular", "GET");

            // Проверка что ответ не пустой
            if (jsonVideos == null)
            {
                throw new Exception("Movies not found");
            }

            // Конвертация из JSON в model
            YouTubeModel[] videos = JsonSerializer.Deserialize<YouTubeModel[]>(jsonVideos);

            // Добовляем фильмы в выделенный массив
            AddVideos(videos);
        }
        catch(Exception ex)
        {
            throw new Exception(ex.Message);
        }
    }

    /// <summary>
    /// Получение популярных видео с YouTube
    /// Возврощает только 3 популярных видео
    /// Данные зависят от модели YouTubeModel
    /// </summary>
    public YouTubeModel[] PopularVideos()
    {
        throw new NotImplementedException();
    }

    /// <summary>
    /// Получить текущии YouTube видео
    /// </summary>
    public YouTubeModel[] ReturnVideos()
    {
        throw new NotImplementedException();
    }
}