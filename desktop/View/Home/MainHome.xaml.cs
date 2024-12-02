using System.Windows;
using System.Windows.Input;
using flick_finder.Domain;
using flick_finder.Domain.Exceptions;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Services;
using UserControl = System.Windows.Controls.UserControl;

namespace flick_finder.View.Home;

public partial class MainHome : UserControl
{
    /// <summary>
    /// Обработчик ошибок
    /// </summary>
    private readonly UIMessageException _uiexception;

    /// <summary>
    /// Работа с фильмами
    /// </summary>
    private readonly IMovies _movies;

    /// <summary>
    /// Работа с UserControl
    /// </summary>
    private readonly IUserControl _userControl;
    
    public MainHome()
    {
        InitializeComponent();
        
        this._uiexception = new UIMessageException();
        this._movies = new Movies();
        this._userControl = new Domain.Services.UserControl();
    }

    /// <summary>
    /// Загрузка экрана
    /// Получение массив популярных фильмов
    /// Показ их в UI
    /// </summary>
    private async void MainHome_Loaded(object sender, RoutedEventArgs e)
    {
        try
        {
            // Показ окна загрузки
            Loader.Visibility = Visibility.Visible;
            
            // Отправка запроса на сервер
            // В отдельном потоке
            await Task.Run(() => this._movies.PopularMovies());
        
            // Вставка UI элементы
            _userControl.DynamicMovie(MoviesBlock, this._movies.ReturnMovies());
        }
        catch (Exception ex)
        {
            this._uiexception.ShowError(ex.Message, "SERVER");
        }
        finally
        {
            Loader.Visibility = Visibility.Collapsed;
        }
    }
}