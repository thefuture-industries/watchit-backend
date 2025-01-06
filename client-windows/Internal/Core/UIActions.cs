using System;
using System.Windows;
using System.Windows.Input;

namespace client.Internal.Core
{
    public class UIActions
    {
        private readonly MainWindow _main_window;

        public UIActions(MainWindow mainWindow)
        {
            this._main_window = mainWindow;
        }

        public void SetFrameContent(string content_uri)
        {
            this._main_window.FrameContent.Navigate(new System.Uri(content_uri, UriKind.Absolute));
        }

        public void LoaderVisibilityVisible() => this._main_window.Loader.Visibility = Visibility.Visible;
        public void LoaderVisibilityCollapsed() => this._main_window.Loader.Visibility = Visibility.Collapsed;
    }

    public class RelayCommand : ICommand
    {
        private readonly Action<object> _execute;
        private readonly Func<object, bool> _canExecute;

        public RelayCommand(Action<object> execute, Func<object, bool> canExecute = null)
        {
            _execute = execute ?? throw new ArgumentNullException(nameof(execute));
            _canExecute = canExecute;
        }

        public bool CanExecute(object parameter) => _canExecute == null || _canExecute(parameter);

        public void Execute(object parameter) => _execute(parameter);

        public event EventHandler CanExecuteChanged
        {
            add => CommandManager.RequerySuggested += value;
            remove => CommandManager.RequerySuggested -= value;
        }
    }
}
