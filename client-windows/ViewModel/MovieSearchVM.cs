using GalaSoft.MvvmLight.Command;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Windows.Input;

namespace client.ViewModel
{
    public class GenreModel
    {
        public string Title { get; set; }
    }

    public class MovieSearchVM : INotifyPropertyChanged
    {
        public event PropertyChangedEventHandler PropertyChanged;
        public virtual void OnPropertyChanged(string propertyName)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }

        public ObservableCollection<GenreModel> Genres {  get; set; }
        public string _selectedGenre;

        private ICommand _genreClickCommand;
        public ICommand GenreClickCommand
        {
            get
            {
                return _genreClickCommand ?? (_genreClickCommand = new RelayCommand<string>(
                    g =>
                    {
                        this._selectedGenre = g;
                    }));
            }
        }

        public MovieSearchVM()
        {
            Genres = new ObservableCollection<GenreModel>
            {
                new GenreModel() {Title = "Action"},
                new GenreModel() {Title = "Adventure"},
                new GenreModel() {Title = "Animation"},
                new GenreModel() {Title = "Comedy"},
                new GenreModel() {Title = "Crime"},
                new GenreModel() {Title = "Documentary"},
                new GenreModel() {Title = "Drama"},
                new GenreModel() {Title = "Family"},
                new GenreModel() {Title = "Fantasy"},
                new GenreModel() {Title = "History"},
                new GenreModel() {Title = "Horror"},
                new GenreModel() {Title = "Music"},
                new GenreModel() {Title = "Mystery"},
                new GenreModel() {Title = "Romance"},
                new GenreModel() {Title = "Science Fiction"},
                new GenreModel() {Title = "TV Movie"},
                new GenreModel() {Title = "Thriller"},
                new GenreModel() {Title = "War"},
                new GenreModel() {Title = "Western"},
            };
        }
    }
}
