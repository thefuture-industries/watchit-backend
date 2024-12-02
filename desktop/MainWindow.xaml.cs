using System.Windows;
using flick_finder.Domain;
using flick_finder.View;

namespace flick_finder;

/// <summary>
/// Interaction logic for MainWindow.xaml
/// </summary>
public partial class MainWindow : Window
{
    /// <summary>
    /// Метод для получений путей приложения XAML
    /// </summary>
    private readonly Routes _routes;
    
    public MainWindow()
    {
        InitializeComponent();

        this._routes = new Routes();
    }

    /// <summary>
    /// Загрузка стартового экрана
    /// Идет подгрузка и соединение с сервером
    /// </summary>
    private void Window_Loaded(object sender, RoutedEventArgs e)
    {
        this._routes.NavigateToControl(this._routes.Get().MAIN_HOME);
    }

    public Loader GetLoader()
    {
        return Loader;
    }
}