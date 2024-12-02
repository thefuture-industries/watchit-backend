using System.Windows;
using System.Windows.Input;
using flick_finder.Domain.Exceptions;
using flick_finder.Domain.Interfaces;
using flick_finder.Domain.Models;
using flick_finder.Domain.Services;
using UserControl = System.Windows.Controls.UserControl;

namespace flick_finder.View.Home;

public partial class TopView_Home : UserControl
{
    private readonly UIMessageException _uiexception;

    private readonly IMovies _movies;

    private readonly IUserControl _userControl;
    
    public TopView_Home()
    {
        InitializeComponent();

        _uiexception = new UIMessageException();
        _movies = new Movies();
        this._userControl = new Domain.Services.UserControl();
    }

    private void SearchBox_Home_KeyDown(object sender, KeyEventArgs e)
    {
        if (e.Key == Key.Enter)
        {
            this.SearchRequest();
        }
    }

    private void SearchRequest()
    {
        try
        {
            ResultsMovie[] movies = this._movies.SearchMovies(this.SearchBox_Home.Text);

            var home = new MainHome();
            this._userControl.DynamicMovie(home.MoviesBlock, movies);
        }
        catch (Exception ex)
        {
            this._uiexception.ShowError(ex.Message, "SERVER");
        }
    }
}