using client.Internal.Core;
using client.Models;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Windows;
using System.Windows.Input;

namespace client.ViewModel
{
    public class PanelVM : INotifyPropertyChanged
    {
        public event PropertyChangedEventHandler PropertyChanged;
        public virtual void OnPropertyChanged(string propertyName)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }

        /// <summary>
        /// Название меню в панеле
        /// </summary>
        public ObservableCollection<MenuItem> MenuItems { get; set; }

        /// <summary>
        /// Команды
        /// </summary>
        private ICommand _menuItemClickCommand;
        public ICommand MenuItemClickCommand
        {
            get
            {
                return _menuItemClickCommand ?? (_menuItemClickCommand = new RelayCommand(path =>
                    {
                        UIActions _action = new UIActions(Application.Current.MainWindow as MainWindow);
                        _action.SetFrameContent(path as string);
                    }));
            }
        }

        public PanelVM()
        {
            MenuItems = new ObservableCollection<MenuItem>
            {
                new MenuItem { Image="pack://application:,,,/Public/ico/layout-dashboard.png", PathView="pack://application:,,,/View/Home.xaml", Title="Home" },
                new MenuItem { Image="pack://application:,,,/Public/ico/heart.png", PathView="pack://application:,,,/View/Favourite.xaml", Title="Favourites" },
                new MenuItem { Image="pack://application:,,,/Public/ico/youtube.png", PathView="pack://application:,,,/View/YoutubeSearch.xaml", Title="Youtube" },
                new MenuItem { Image="pack://application:,,,/Public/ico/clapperboard.png", PathView="pack://application:,,,/View/MovieSearch.xaml", Title="Movies" },
                new MenuItem { Image="pack://application:,,,/Public/ico/text.png", PathView="pack://application:,,,/View/Story.xaml", Title="Story" },
            };
        }
    }
}
