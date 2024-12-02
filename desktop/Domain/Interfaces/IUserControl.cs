using System.Windows.Controls;
using flick_finder.Domain.Models;

namespace flick_finder.Domain.Interfaces;

public interface IUserControl
{
    /// <summary>
    /// Генерация массива фильмов
    /// </summary>
    void DynamicMovie(WrapPanel control, ResultsMovie[] array);
}