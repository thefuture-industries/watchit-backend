using System.Windows;

namespace flick_finder.Domain.Exceptions;

public class UIMessageException
{
    /// <summary>
    /// Дефолтный текст в окне ошибки
    /// Текст будет использоваться если не передан type
    /// </summary>
    private string _typeException = "Error about the warning";
    
    /// <summary>
    /// Метод для показа окна с типом ошибка
    /// </summary>
    public void ShowError(string error, string type = null)
    {
        if (type == "INTERNET") this._typeException = "Internet connection error";
        if (type == "SERVER") this._typeException = "Error on the server side";
        if (type == "CLIENT") this._typeException = "Error, please try again later";
        
        MessageBox.Show(
            error, 
            this._typeException, 
            MessageBoxButton.OK, 
            MessageBoxImage.Error);

        Application.Current.Dispatcher.Invoke(() =>
        {
            System.Windows.Application.Current.Shutdown();
        });
    }

    /// <summary>
    /// Метод для показа окна с типом предупреждение
    /// </summary>
    public void ShowWarning(string warning, string type = null)
    {
        if (type == "INTERNET") this._typeException = "Internet connection error";
        if (type == "SERVER") this._typeException = "Warning on the server side";
        if (type == "CLIENT") this._typeException = "Warning, think about it";

        MessageBox.Show(
            warning, 
            this._typeException,
            MessageBoxButton.OK,
            MessageBoxImage.Warning);
    }
}