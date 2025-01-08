using client.Services;
using System;
using System.Globalization;
using System.Windows.Data;

namespace client.UI.Styles.Controllers
{
    /// <summary>
    /// Изображение фильмов с сервера
    /// </summary>
    public class ImageUriConvertor : IValueConverter
    {
        private readonly Config _config;

        public ImageUriConvertor()
        {
            this._config = new Config();
        }

        public object Convert(object value, Type targetType, object parameter, CultureInfo culture)
        {
            if (value is string posterPath)
            {
                return new Uri($"{this._config.ReturnConfig().SERVER_URL}/image/w500/{posterPath}", UriKind.Absolute);
                /*return new ImageBrush(new BitmapImage(uri));*/
            }

            return null;
        }

        public object ConvertBack(object value, Type targetType, object parameter, CultureInfo culture)
        {
            throw new NotImplementedException();
        }
    }
}
