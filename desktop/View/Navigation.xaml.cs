using System.Windows;
using System.Windows.Input;
using flick_finder.Domain;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Services;
using UserControl = System.Windows.Controls.UserControl;

namespace flick_finder.View;

public partial class Navigation : UserControl
{
    /// <summary>
    /// Работа с данными пользователя
    /// </summary>
    private readonly IUser _user;
    
    /// <summary>
    /// Пути приложения XAML
    /// </summary>
    private readonly Routes _router;
    
    public Navigation()
    {
        InitializeComponent();

        this._router = new Routes();
        this._user = new User();
    }

    private void Navigation_Loaded(object sender, RoutedEventArgs e)
    {
        this.UserName_Label.Content = this._user.GetUser().UserName;
        this.Email_Label.Content = this._user.GetUser().Email;
    }

    /// <summary>
    /// Навигация на Home
    /// </summary>
    private void HomePage_Click(object sender, MouseButtonEventArgs e)
    {
        this._router.NavigateToControl(this._router.Get().MAIN_HOME);
    }

    /// <summary>
    /// Навигация выход из приложения
    /// </summary>
    private void ExitBtn_Click(object sender, MouseButtonEventArgs e)
    {
        Application.Current.Shutdown();
    }

    /// <summary>
    /// Навигация на поиск по Фильмам
    /// </summary>
    /// <param name="sender"></param>
    /// <param name="e"></param>
    private void Movies_Click(object sender, MouseButtonEventArgs e)
    {
        this._router.NavigateToControl(this._router.Get().MAIN_MOVIE_POPUP);
    }

    /// <summary>
    /// Навигация на поиск по YouTube
    /// </summary>
    private void YouTube_Click(object sender, MouseButtonEventArgs e)
    {
        this._router.NavigateToControl(this._router.Get().MAIN_YOUTUBE_POPUP);
    }
    
    /// <summary>
    /// Кнопка показа окна
    /// </summary>
    private void TextBtnSearch(object sender, MouseButtonEventArgs e)
    { 
        this._router.NavigateToControl(this._router.Get().MAIN_TEXT_SEARCH);
    }
}