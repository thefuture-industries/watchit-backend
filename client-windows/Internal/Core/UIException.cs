using System.Windows;

namespace client.Services
{
    public class UIException
    {
        public void Error(string overview, string title)
        {
            MessageBox.Show(overview.ToString(), title.ToString(), MessageBoxButton.OK, MessageBoxImage.Error);
        }

        public void Warning(string overview, string title)
        {
            MessageBox.Show(overview.ToString(), title.ToString(), MessageBoxButton.OK, MessageBoxImage.Warning);
        }
    }
}
