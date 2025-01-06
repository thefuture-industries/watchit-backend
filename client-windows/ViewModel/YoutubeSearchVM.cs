using client.Internal.Core;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Windows.Input;

namespace client.ViewModel
{
    public class CategoryModel : INotifyPropertyChanged
    {
        private string text;
        public string Text
        {
            get { return text; }
            set
            {
                text = value;
                OnPropertyChanged();
            }
        }

        private bool isSelected;
        public bool IsSelected
        {
            get { return isSelected; }
            set
            {
                isSelected = value;
                OnPropertyChanged();
            }
        }

        public event PropertyChangedEventHandler PropertyChanged;
        public virtual void OnPropertyChanged(string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }
    }


    public class YoutubeSearchVM : INotifyPropertyChanged
    {
        public event PropertyChangedEventHandler PropertyChanged;
        public virtual void OnPropertyChanged(string propertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        }

        public ObservableCollection<CategoryModel> Categoryes { get; set; }

        private ICommand _categoryClickCommand;
        public ICommand CategoryClickCommand
        {
            get
            {
                return _categoryClickCommand ?? (_categoryClickCommand = new RelayCommand(category =>
                    {
                        if (category is CategoryModel categoryModel)
                        {
                            foreach (var c in Categoryes)
                            {
                                c.IsSelected = false;
                            }

                            categoryModel.IsSelected = true;
                        }   
                    }));
            }
        }

        public YoutubeSearchVM()
        {
            Categoryes = new ObservableCollection<CategoryModel>()
            {
                new CategoryModel() {Text = "Film", IsSelected = false},
                new CategoryModel() {Text = "Animation", IsSelected = false},
                new CategoryModel() {Text = "Autos", IsSelected = false},
                new CategoryModel() {Text = "Animals", IsSelected = false},
                new CategoryModel() {Text = "Sports", IsSelected = false},
                new CategoryModel() {Text = "Events", IsSelected = false},
                new CategoryModel() {Text = "Travel", IsSelected = false},
                new CategoryModel() {Text = "Gaming", IsSelected = false},
                new CategoryModel() {Text = "Blogs", IsSelected = false},
                new CategoryModel() {Text = "Howto", IsSelected = false},
                new CategoryModel() {Text = "Science", IsSelected = false},
                new CategoryModel() {Text = "Technology", IsSelected = false},
                new CategoryModel() {Text = "Education", IsSelected = false},
                new CategoryModel() {Text = "Movie and TV series trailers", IsSelected = false},
            };

            // CategoryClickCommand = new RelayCommand<object>(Click);
        }
    }
}
