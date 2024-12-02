using System.Text.Json;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Models;

namespace flick_finder.Domain.Services;

public class Movies : IMovies
{
    /// <summary>
    /// Текущие фильмы
    /// </summary>
    private ResultsMovie[] _movies = null;

    /// <summary>
    /// Инициализация запросов на сервер
    /// </summary>
    private readonly IHttpRequest _http;

    /// <summary>
    /// Текущая страницы популярных фильмов
    /// </summary>
    public int Page { get; set; } = 0;
    public int PagePopular { get; set; } = 1;

    public Movies()
    {
        this._http = new HttpRequest();
    }

    /// <summary>
    /// Запись в память текущии фильмы
    /// </summary>
    public void AddMovies(ResultsMovie[] movies)
    {
        // Создаем новый массив с увеличенной длиной
        int length = this._movies == null ? movies.Length : this._movies.Length + movies.Length;
        
        // Определение размера массива
        ResultsMovie[] newMovies = new ResultsMovie[length];
        
        // Копируем существующие данные
        if (this._movies != null)
        {
            Array.Copy(this._movies, newMovies, this._movies.Length);
        }
        
        // Копируем новые данные
        Array.Copy(movies, 0, newMovies, this._movies == null ? 0 : this._movies.Length, movies.Length);
        
        // Обновляем _movies
        this._movies = newMovies;
    }

    /// <summary>
    /// Получение популярных фильмов
    /// </summary>
    public void PopularMovies()
    {
        try
        {
            // Отправка запросов на сервер
            string jsonMovies = this._http.SendGETRequest($"movies/popular?page={this.PagePopular}", "GET");

            // Проверка что ответ не пустой
            if (jsonMovies == null)
            {
                throw new Exception("Movies not found");
            }

            // Конвертация из JSON в model
            ResultsMovie[] movies = JsonSerializer.Deserialize<ResultsMovie[]>(jsonMovies);

            // Добовляем фильмы в выделенный массив
            this.AddMovies(movies);
        }
        catch (Exception ex)
        {
            throw new Exception(ex.Message);
        }
    }

    public ResultsMovie[] SearchMovies(string search)
    {
        try
        {
            string jsonMovies = this._http.SendGETRequest($"movies?s={search}&page={this.Page}", "GET");

            if (jsonMovies == null)
            {
                throw new Exception("Movies not found");
            }

            ResultsMovie[] movies = JsonSerializer.Deserialize<ResultsMovie[]>(jsonMovies);

            return movies;
        }
        catch (Exception ex)
        {
            throw new Exception(ex.Message);
        }
    }

    /// <summary>
    /// Получить текущие фильмы
    /// </summary>
    public ResultsMovie[] ReturnMovies()
    { 
        return this._movies;
    }
}