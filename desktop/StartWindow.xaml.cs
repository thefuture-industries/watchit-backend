using System.IO;
using System.Windows;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Services;
using Path = System.IO.Path;

namespace flick_finder;

public partial class StartWindow : Window
{
    /// <summary>
    /// Метод для отправка запросов на сервер
    /// POST, GET, PUT, DELETE
    /// </summary>
    private readonly IHttpRequest _http;

    private readonly IYouTube _youtube;

    public StartWindow()
    {
        InitializeComponent();

        this._http = new HttpRequest();
        this._youtube = new YouTube();
    }

    /// <summary>
    /// Загрузка окна
    /// Создание папок Config и
    /// Регестрация пользователя в БД, если не существует 
    /// </summary>
    private async void StartWindow_Loaded(object sender, RoutedEventArgs e)
    {
        // Получение данных пользователя
        // -----------------------------

        // Directory
        string documentPath = Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments);

        // Версия устройства
        string osVersion = Environment.OSVersion.VersionString;

        // UserName
        string userName = Environment.UserName;

        // OS для Email
        string oc = null;

        string[] parts = osVersion.Split(" ");
        if (parts.Length >= 2 && parts[1] == "Windows")
        {
            oc = string.Join(" ", parts[1], parts[3]);
        }
        else
        {
            oc = osVersion;
        }

        // Email
        string email = $"{userName}_{oc}@gmail.com";

        // Отправка запрос на сервер
        // Создание нового потока
        await Task.Run(() =>
        {
            string response_user = this._http.SendRequest("signup", "POST", new
            {
                username = userName,
                email = email,
                directory = documentPath,
                oc = osVersion
            });

            this._youtube.SetVideos();
        }).ContinueWith(_ =>
        {
            Application.Current.Dispatcher.Invoke(() =>
            {
                this.Hide();
                new MainWindow().Show();
            });
        });
    }
}
