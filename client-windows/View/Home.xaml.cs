using System.Windows;
using System.Windows.Controls;

namespace client.View
{
    /// <summary>
    /// Логика взаимодействия для Home.xaml
    /// </summary>
    public partial class Home : UserControl
    {
        public Home()
        {
            InitializeComponent();
        }

        private void SearchTextBox_TextChanged(object sender, TextChangedEventArgs e)
        {
            if (SearchTextBox.Text.Length > 0)
            {
                PlaceholderSearch.Visibility = Visibility.Collapsed;
            }
            else
            {
                PlaceholderSearch.Visibility= Visibility.Visible;
            }
        }
    }
}
