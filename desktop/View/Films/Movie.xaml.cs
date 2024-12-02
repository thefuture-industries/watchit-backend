using System.Globalization;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Media.Imaging;

namespace flick_finder.View.Films;

public partial class Movie : UserControl
{
    public Movie()
    {
        InitializeComponent();
    }

    private int _id;
    private string _image;
    private string _title;
    private string _overview;
    private float _vote_average;

    public int Id
    {
        get { return _id;}
        set { _id = value; }
    }

    public string PosterPath
    {
        get { return _image; }
        set { 
            _image = value;
            this.PosterPathMovie.ImageSource = new BitmapImage(new Uri(value));
        }
    }

    public string Title
    {
        get { return _title; }
        set { 
            _title = value;
            this.TitleMovie.Text = value;
        }
    }

    public float VoteAverage
    {
        get { return _vote_average; }
        set
        {
            _vote_average = value; 
            this.VoteAverageMovie.Content = value; 
        }
    }

    /*public string Overview
    {
        get { return _overview; }
        set { _overview = value;
            OverviewMovie.Text = value;
        }
    }*/
}