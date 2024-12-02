using flick_finder.Domain.Models;

namespace flick_finder.Domain.Interfaces;

public interface IMovies
{
    public int Page { get; set; }
    public int PagePopular { get; set; }
    
    /// <summary>
    /// Запись в память текущии фильмы
    /// </summary>
    void AddMovies(ResultsMovie[] movies);

    /// <summary>
    /// Получение популярных фильмов
    /// </summary>
    void PopularMovies();

    /// <summary>
    /// Поиск фильма по строке (Title, Description)
    /// </summary>
    ResultsMovie[] SearchMovies(string search);
    
    /// <summary>
    /// Получить текущие фильмы
    /// </summary>
    ResultsMovie[] ReturnMovies();
}