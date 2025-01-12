using client.Internal.Core;
using client.Internal.Interfaces;
using client.Internal.Services;
using client.Models;
using client.Services;
using System;
using System.ComponentModel;
using System.Net.Http;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Input;
using System.Windows.Threading;

namespace client.ViewModel
{
    public class StartedVM : INotifyPropertyChanged
    {
        public event PropertyChangedEventHandler PropertyChanged;
        public virtual void OnPropertyChanged(string PropertyName = null)
        {
            PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(PropertyName));
        }

        /// <summary>
        /// Клиент для отправки запросов
        /// </summary>
        private readonly HttpClient _client;

        /// <summary>
        /// Обработка ошибок в UI
        /// </summary>
        private readonly UIException _exception;

        /// <summary>
        /// Сервис работы с пользователем
        /// </summary>
        private readonly IUserService _userService;

        public StartedVM()
        {
            this._client = new HttpClient();
            this._exception = new UIException();
            this._userService = new UserService();
        }

        /// <summary>
        /// Секретное слово пользователя
        /// </summary>
        private string _secretWord;
        public string SecretWord
        {
            get { return _secretWord; }
            set 
            {
                _secretWord = value;
                OnPropertyChanged(nameof(SecretWord));
            }
        }

        /// <summary>
        /// Активность кнопки
        /// </summary>
        private bool _isBusy;
        public bool IsBusy
        {
            get { return _isBusy; }
            set
            {
                _isBusy = value;
                OnPropertyChanged(nameof(IsBusy));
            }
        }


        /// <summary>
        /// Отправка на сервер
        /// </summary>
        private ICommand _verifyCommand;
        public ICommand VerifyCommand
        {
            get
            {
                return _verifyCommand ?? (_verifyCommand = new RelayCommand(async (_) => {
                    this._isBusy = true;
                    using (var ctx = new CancellationTokenSource(TimeSpan.FromSeconds(10)))
                    {
                        try
                        {
                            var requestTask = this._client.GetAsync("http://ip-api.com/json", ctx.Token);
                            var delayTask = Task.Delay(Timeout.Infinite, ctx.Token);

                            var completedTask = await Task.WhenAny(requestTask, delayTask);

                            if (completedTask == requestTask)
                            {
                                var response = await requestTask;
                                if (response.IsSuccessStatusCode)
                                {
                                    GeoIPModel geo = JsonSerializer.Deserialize<GeoIPModel>(await response.Content.ReadAsStringAsync());
                                    this._userService.Add(new Models.UserAddPayload()
                                    {
                                        SecretWord = this._secretWord,
                                        IpAddress = geo.Query,
                                        Country = geo.Country,
                                        RegionName = geo.RegionName,
                                        Zip = geo.Zip,
                                    });
                                }
                            }
                            else
                            {
                                this._userService.Add(new Models.UserAddPayload()
                                {
                                    SecretWord = this._secretWord,
                                    IpAddress = "0.0.0.0",
                                    Country = "NONE",
                                    RegionName = "NONE",
                                    Zip = "000000",
                                });
                            }

                            // Работа с окнами
                            await Application.Current.Dispatcher.BeginInvoke(DispatcherPriority.Normal, new Action(() =>
                            {
                                UIActions actions = new UIActions(Application.Current.MainWindow as MainWindow);
                                actions.SetFrameIndex("pack://application:,,,/MainControl.xaml");
                            }));
                        }
                        catch (Exception ex)
                        {
                            this._exception.Error(ex.Message, "Error Server");
                            return;
                        }
                        finally
                        {
                            this._isBusy = false;
                        }
                    }
                }));
            }
        }
    }
}
