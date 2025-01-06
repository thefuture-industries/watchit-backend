using client.Services;
using System;
using System.IO;
using System.Net;
using System.Threading;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Media;
using System.Windows.Media.Imaging;

namespace client.UI.Components
{
    /// <summary>
    /// Логика взаимодействия для MovieCard.xaml
    /// </summary>
    public partial class MovieCard : UserControl
    {
        private readonly Config _config;

        public MovieCard()
        {
            InitializeComponent();
            this._config = new Config();
        }

        /// <summary>
        /// Проверка элемента в видимости
        /// </summary>
        private void LoadVisibleImages()
        {
        }

        /// <summary>
        /// Загрузка постеров
        /// </summary>
        /// <param name="path">путь изображения</param>
        private async void LoadImageAsync(string path)
        {
            MovieBrash.ImageSource = new BitmapImage(new Uri("pack://application:,,,/Public/images/image_preloader.png"));

            await Task.Run(async () =>
            {
                try
                {
                    // Загрузка постера
                    using (WebClient client = new WebClient())
                    {
                        using (var ctx = new CancellationTokenSource(TimeSpan.FromSeconds(10)))
                        {
                            try
                            {
                                byte[] imageBytes = await client.DownloadDataTaskAsync(new Uri($"{this._config.ReturnConfig().SERVER_URL}/image/w500{path}", UriKind.Absolute));
                                Dispatcher.Invoke(() =>
                                {
                                    using (MemoryStream memoryStream = new MemoryStream(imageBytes))
                                    {
                                        var bitmap = new BitmapImage();
                                        bitmap.BeginInit();
                                        bitmap.StreamSource = memoryStream;
                                        bitmap.CacheOption = BitmapCacheOption.OnLoad;
                                        bitmap.EndInit();
                                        MovieBrash.ImageSource = bitmap;
                                    }
                                });
                            }
                            catch (TaskCanceledException)
                            {
                                // Ошибка загрузки постера
                                Dispatcher.Invoke(() =>
                                {
                                    MovieBrash.ImageSource = new BitmapImage(new Uri("pack://application:,,,/Public/images/default_image.jpg"));
                                });
                            }
                        }
                    }
                }
                catch (Exception)
                {
                    // Ошибка загрузки постера
                    Dispatcher.Invoke(() =>
                    {
                        MovieBrash.ImageSource = new BitmapImage(new Uri("pack://application:,,,/Public/images/default_image.jpg"));
                    });
                }
            });
        }

        private static void OnPosterPathChanged(DependencyObject d, DependencyPropertyChangedEventArgs e)
        {
            var window = (MovieCard)d;
            if (e.NewValue != null)
            {
                window.LoadImageAsync(e.NewValue.ToString());
            }
            else
            {
                return;
            }
        }

        /// <summary>
        /// Title Dependency
        /// </summary>
        public static readonly DependencyProperty TitleProperty = DependencyProperty.Register("Title", typeof(string), typeof(MovieCard), new PropertyMetadata(string.Empty));
        public string Title
        {
            get { return (string)GetValue(TitleProperty); }
            set { SetValue(TitleProperty, value); }
        }

        /// <summary>
        /// PosterPath Dependency
        /// </summary>
        public static readonly DependencyProperty PosterPathProperty = DependencyProperty.Register("PosterPath", typeof(string), typeof(MovieCard), new PropertyMetadata(null, OnPosterPathChanged));
        public string PosterPath
        {
            get { return (string)GetValue(PosterPathProperty); }
            set { SetValue(PosterPathProperty, value); }
        }

        public ScaleTransform ScaleTransform => scaleTransform;
    }
}
