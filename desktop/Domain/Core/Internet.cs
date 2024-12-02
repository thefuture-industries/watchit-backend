using System.Net;
using System.Net.NetworkInformation;

namespace flick_finder.Domain.Core;

public class Internet
{
    public static bool OK()
    {
        // try
        // {
        //     Dns.GetHostEntry("dotnet.beget.tech");
        //     return true;
        // }
        // catch
        // {
        //     return false;
        // }

        try
        {
            // Пытаемся получить список доступных сетевых адаптеров
            var networkInterfaces = NetworkInterface.GetAllNetworkInterfaces();

            // Проверяем каждый адаптер
            foreach (var networkInterface in networkInterfaces)
            {
                // Если адаптер включен и имеет IP-адрес
                if (networkInterface.OperationalStatus == OperationalStatus.Up &&
                    networkInterface.GetIPProperties().UnicastAddresses.Count > 0)
                {
                    return true;
                }
            }
        }
        catch
        {
            // Возвращаем false в случае ошибки
            return false;
        }

        // Интернет недоступен
        return false;
    }
}