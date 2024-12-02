using System.Windows;
using System.Windows.Controls;
using System.Windows.Media.Imaging;
using flick_finder.Domain.Exceptions;
using WpfAnimatedGif;

namespace flick_finder.View.Text;

public partial class MainTextSearch : UserControl
{
    public MainTextSearch()
    {
        InitializeComponent();

        Loader.Visibility = Visibility.Visible;
        
        // Задержка на подгрузку img
        Task.Run(() =>
        {
            Task.Delay(1000).ContinueWith(_ =>
            {
                Dispatcher.BeginInvoke(new Action(() =>
                {
                    var img = (Image)FindName("ImageGifAnimation");
                    // gif:ImageBehavior.AnimatedSource="Public/Gif/circle_animation.gif"
                    ImageBehavior.SetAnimatedSource(img,
                        new BitmapImage(new Uri(
                            "pack://application:,,,/Public/Gif/circle_animation.gif"))); // pack://application:,,,/Public/Gif/circle_animation.gif
            
                    Loader.Visibility = Visibility.Collapsed;
                }));
            });
        });
    }
}