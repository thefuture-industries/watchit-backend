using client.UI.Components;
using System;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Input;

namespace client.Internal.Core
{
    public class UIActions
    {
        private readonly MainWindow _main;

        public UIActions(MainWindow main)
        {
            this._main = main;
        }

        public void SetFrameContent(string content_uri)
        {
            if (this._main.FrameIndex.Content is MainControl mainControl)
            {
                mainControl.FrameContent.Navigate(new System.Uri(content_uri, UriKind.Absolute));
            }
        }

        public void SetFrameIndex(string content_uri)
        {
            this._main.FrameIndex.Navigate(new System.Uri(content_uri, UriKind.Absolute));
        }

        public void LoaderVisibilityVisible()
        {
            if (this._main.FrameIndex.Content is MainControl mainControl)
            {
                if (mainControl.FrameContent.Content is UserControl userControl)
                {
                    if (userControl.FindName("Loader") is SkeletonLoader skeletonLoader)
                    {
                        skeletonLoader.Visibility = Visibility.Visible;
                    }
                }
            }
        }

        public void LoaderVisibilityCollapsed()
        {
            if (this._main.FrameIndex.Content is MainControl mainControl)
            {
                if (mainControl.FrameContent.Content is UserControl userControl)
                {
                    if (userControl.FindName("Loader") is SkeletonLoader skeletonLoader)
                    {
                        skeletonLoader.Visibility = Visibility.Collapsed;
                    }
                }
            }
        }
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
