using System.Net.Http;
using System.Windows;
using System.Windows.Controls;
using flick_finder.Domain.Core;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Services;

namespace flick_finder.Domain;

public class RoutesApp
{
    // Строка содержащия путь к файлу
    // MAIN_HOME: Путь к файлу MainGome.xaml 
    // Путь относительно внутри приложения
    public string MAIN_HOME { get; set; }
    
    // Строка содержащия путь к файлу
    // MAIN_TEXT_SEARCH: Путь к файлу MainTextSearch.xaml 
    // Путь относительно внутри приложения
    public string MAIN_TEXT_SEARCH { get; set; }

    // Строка содержащия путь к файлу
    // MAIN_MOVIE_POPUP: Путь к файлу Movie_PopUp.xaml 
    // Путь относительно внутри приложения
    public string MAIN_MOVIE_POPUP { get; set; }

    // Строка содержащия путь к файлу
    // MAIN_YOUTUBE_POPUP: Путь к файлу YouTube_PopUp.xaml 
    // Путь относительно внутри приложения
    public string MAIN_YOUTUBE_POPUP { get; set; }

    public RoutesApp()
    {
        this.MAIN_HOME = "pack://application:,,,/View/Home/MainHome.xaml";
        this.MAIN_TEXT_SEARCH = "pack://application:,,,/View/Text/MainTextSearch.xaml";
        this.MAIN_MOVIE_POPUP = "pack://application:,,,/View/PopUps/Movie_PopUp.xaml";
        this.MAIN_YOUTUBE_POPUP = "pack://application:,,,/View/PopUps/YouTube_PopUp.xaml";
    }
}

// Класс для работы с навигацией в приложении
// Два метода 1) Навигация 2) Получение пути
public class Routes
{
    /// <summary>
    /// Навигация в приложении за счет путей xaml
    /// </summary>
    public void NavigateToControl(string control)
    {
        var window = (MainWindow)Application.Current.Windows.OfType<MainWindow>().FirstOrDefault();
        
        window.DataWindow.Navigate(new Uri(control, UriKind.Absolute));
    }
    
    /// <summary>
    /// Получение путей(файлов xaml) для навигации в приложении
    /// </summary>
    public RoutesApp Get()
    {
        return new RoutesApp();
    }
}