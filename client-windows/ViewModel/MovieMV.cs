using client.API;
using client.Internal.Core;
using client.Models;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Runtime.CompilerServices;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;

namespace client.ViewModel
{
    public class MovieMV : INotifyPropertyChanged
    {
        /// <summary>
        /// Работа с UI элементами
        /// </summary>
        private readonly UIActions _action;

        /// <summary>
        /// Сервис для работы с фильмами
        /// </summary>
        private readonly MovieService _movieService;

        /// <summary>
        /// Хранение фильмов
        /// </summary>
        private ObservableCollection<MovieModel> _movies;
        public ObservableCollection<MovieModel> Movies
        {
            get { return this._movies; }
            set
            {
                this._movies = value;
                OnPropertyChanged(nameof(Movies));
            }
        }

        /// <summary>
        /// Конструктор
        /// </summary>
        public MovieMV()
        {
            this._action = new UIActions(Application.Current.MainWindow as MainWindow);
            this._movieService = new MovieService();

            Application.Current.Dispatcher.Invoke(() =>
            {
                this._action.LoaderVisibilityVisible();
            });

            Task.Run(async () =>
            {
                var movies = await this._movieService.Get();
                Movies = new ObservableCollection<MovieModel>();

                Application.Current.Dispatcher.Invoke(() =>
                {
                    if (movies != null)
                    {
                        this._action.LoaderVisibilityCollapsed();
                        foreach (var movie in movies)
                        {
                            Movies.Add(movie);
                        }
                    }
                });
            });
        }

        public event PropertyChangedEventHandler PropertyChanged;
        protected virtual void OnPropertyChanged([CallerMemberName] string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }
}
