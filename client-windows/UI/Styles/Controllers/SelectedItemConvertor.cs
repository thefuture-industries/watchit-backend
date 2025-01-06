using System;
using System.Globalization;
using System.Windows.Data;
using System.Windows.Media;

namespace client.UI.Styles.Controllers
{
    public class SelectedItemConvertor : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, CultureInfo culture)
        {
            return (bool)value ? new SolidColorBrush(Colors.White) : new SolidColorBrush(Color.FromRgb(17, 17, 17));
        }

        public object ConvertBack(object value, Type targetType, object parameter, CultureInfo culture)
        {
            throw new NotImplementedException();
        }
    }
}
