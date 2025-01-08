using System.Windows.Controls;

namespace client
{
    /// <summary>
    /// Логика взаимодействия для MainControl.xaml
    /// </summary>
    public partial class MainControl : UserControl
    {
        public MainControl()
        {
            InitializeComponent();

            Loaded += MainControl_Loaded;
        }

        private void MainControl_Loaded(object sender, System.Windows.RoutedEventArgs e)
        {
            FrameContent.Source = new System.Uri("pack://application:,,,/View/Home.xaml", System.UriKind.Absolute);
        }
    }
}
