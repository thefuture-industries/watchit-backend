using System.Windows;
using System.Windows.Controls;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using flick_finder.Domain.Exceptions;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Models;
using flick_finder.View.Films;

namespace flick_finder.Domain.Services;

public class UserControl : IUserControl
{
    private WrapPanel control;
    
    /// <summary>
    /// Работа с фильмами
    /// </summary>
    private readonly IMovies _movies;

    /// <summary>
    /// Обработчик ошибок
    /// </summary>
    private readonly UIMessageException _uiexception;

    public UserControl()
    {
        this._movies = new Movies();
        this._uiexception = new UIMessageException();
    }
    
    /// <summary>
    /// Генерация массива фильмов
    /// </summary>
    public void DynamicMovie(WrapPanel control, ResultsMovie[] array)
    {
        this.control = control;
        this.control.Children.Clear();
        
        Movie[] movies = new Movie[array.Length];

        for (int i = 0; i < array.Length; i++)
        {
            movies[i] = new Movie();
            
            movies[i].Id = array[i].Id;
            movies[i].PosterPath = $"https://avatars.mds.yandex.net/get-kinopoisk-image/1629390/f7eb0ee5-fbab-4398-b06d-df3db4536ce0/1920x"; // https://image.tmdb.org/t/p/w500{array[i].PosterPath}
            movies[i].Title = array[i].Title;
            movies[i].VoteAverage = array[i].VoteAverage;
            /*movies[i].Overview = TextLimiter(array[i].Overview, 9);*/

            movies[i].MouseLeftButtonDown += (sender, e) =>
            {
                Console.WriteLine("Clicked");
            };

            this.control.Children.Add(movies[i]);
        }

        this.control.Children.Add(ShowMore());
    }

    /// <summary>
    /// Создание кнопки показать еще
    /// </summary>
    private Border ShowMore()
    {
        Border border = new Border();
        
        border.Background = Brushes.White;
        border.CornerRadius = new CornerRadius(6);
        border.Padding = new Thickness(10);
        border.MinHeight = 40;
        border.MaxHeight = 40;
        border.Cursor = Cursors.Hand;

        Image image = new Image();

        BitmapImage bitImage = new BitmapImage(new Uri("pack://application:,,,/Public/Images/chevron-right.png"));
        image.Source = bitImage;

        image.Width = 20;
        image.Height= 20;

        border.Child = image;
        border.Name = "ShowMoreBtn";
        border.MouseLeftButtonDown += async (sender, e) =>
        {
            MainWindow window = new MainWindow();

            try
            {
                Application.Current.Dispatcher.Invoke(() => window.GetLoader().Visibility = Visibility.Visible);

                await Task.Run(() =>
                {
                    this._movies.PagePopular++;
                    this._movies.PopularMovies();
                });

                Application.Current.Dispatcher.Invoke(() =>
                {
                    DynamicMovie(this.control, this._movies.ReturnMovies());
                });
            }
            catch (Exception ex)
            {
                Application.Current.Dispatcher.Invoke(() => this._uiexception.ShowError(ex.Message));
            }
            finally
            {
                Application.Current.Dispatcher.Invoke(() => window.GetLoader().Visibility = Visibility.Collapsed);
            }
        };

        return border;
    }

    /// <summary>
    /// Сокращение Overview + ...
    /// </summary>
    private static string TextLimiter(string title, int maxLength)
    {
        string[] words = title.Split(" ");
        if (words.Length <= maxLength)
        {
            return title;
        }
        
        return string.Join(" ", words.Take(maxLength)) + "...";
    }
}