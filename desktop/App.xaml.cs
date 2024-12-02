using System.Configuration;
using System.Data;
using System.Windows;
using System.Windows.Controls;
using flick_finder.Domain.Core;
using flick_finder.Domain.Exceptions;
using flick_finder.Domain.Services;

namespace flick_finder;

/// <summary>
/// Interaction logic for App.xaml
/// </summary>
public partial class App : Application
{
    private readonly UIMessageException _uiexception;

    public App()
    {
        this._uiexception = new UIMessageException();
    }
    
    protected override void OnStartup(StartupEventArgs e)
    {
        base.OnStartup(e);

        if (!Internet.OK())
        {
            this._uiexception.ShowError("Internet connection error, check the connection and try again later.", "INTERNET");
            
            Shutdown();
            return;
        }
    }
}