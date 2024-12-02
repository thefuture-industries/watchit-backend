using flick_finder.Domain.Models;

namespace flick_finder.Domain.Interfaces;

public interface IYouTube
{
    /// <summary>
    /// Запись в память текущии видео
    /// Данные зависят от модели YouTubeModel
    /// </summary>
    void SetVideos();
    
    /// <summary>
    /// Получение популярных видео с YouTube
    /// Возврощает только 3 популярных видео
    /// Данные зависят от модели YouTubeModel
    /// </summary>
    YouTubeModel[] PopularVideos();
    
    /// <summary>
    /// Получить текущии YouTube видео
    /// </summary>
    YouTubeModel[] ReturnVideos();
}