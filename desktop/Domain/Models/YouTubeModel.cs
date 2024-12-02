using System.Text.Json.Serialization;

namespace flick_finder.Domain.Models;

// YouTubeModel: Результат поиска содержит информацию о видео на YouTube,
// канале или плейлисте, которые соответствуют параметрам поиска, указанным в запросе API
//. Хотя результат поиска указывает на однозначно идентифицируемый ресурс,
// такой как видео, у него нет собственных постоянных данных.
public class YouTubeModel
{
    // Etag: Etag этого ресурса.
    [JsonPropertyName("etag")]
    public string Etag { get; set; }
    
    // Id: Объект id содержит информацию, которая может быть использована для уникальной идентификации
    // ресурса, соответствующего поисковому запросу.
    [JsonPropertyName("id")]
    public ResourceId Id { get; set; }
    
    // Kind: Указывает, что это за ресурс. Значение: фиксированная строка
    // "youtube#SearchResult".
    [JsonPropertyName("kind")]
    public string Kind { get; set; }
    
    // Фрагмент: Объект snippet содержит основные сведения о результате поиска,
    // такие как его заголовок или описание. Например, если результатом поиска является
    // видео, то заголовком будет название видео, а описанием -
    // описание видео.
    [JsonPropertyName("snippet")]
    public SearchResultSnippet Snippet { get; set; }
}

// ResourceId: Идентификатор ресурса - это общая ссылка, которая указывает на другой ресурс
// YouTube.
public class ResourceId
{
    // Kind: Тип ресурса API.
    [JsonPropertyName("kind")]
    public string Kind { get; set; }

    // VideoID: Идентификатор, который YouTube использует для уникальной идентификации ресурса, на который ссылается
    //, если этот ресурс является видео. Это свойство присутствует только в том случае, если значение
    // resourceId.kind равно youtube#video.
    [JsonPropertyName("videoId")]
    public string VideoId { get; set; }
}

// SearchResultSnippet: Основные сведения о результатах поиска, включая заголовок,
// описание и миниатюры элемента, на который ссылается результат поиска.
public class SearchResultSnippet
{
    // channelId: значение, которое YouTube использует для уникальной идентификации канала,
    // опубликовавшего ресурс, указанный в результатах поиска.
    [JsonPropertyName("channelId")]
    public string ChannelId { get; set; }
    
    // ChannelTitle: название канала, опубликовавшего ресурс, который
    // идентифицируется в результате поиска.
    [JsonPropertyName("channelTitle")]
    public string ChannelTitle { get; set; }
    
    // Описание: Описание результата поиска.
    [JsonPropertyName("description")]
    public string Description { get; set; }
    
    // LiveBroadcastContent: Указывает, есть ли на ресурсе (видео или канале)
    // предстоящий/активный контент для прямой трансляции. Или "нет", если
    // предстоящих/активных прямых трансляций нет.
    //
    // Possible values:
    //   "none"
    //   "upcoming" - The live broadcast is upcoming.
    //   "live" - The live broadcast is active.
    //   "completed" - The live broadcast has been completed.
    [JsonPropertyName("liveBroadcastContent")]
    public string LiveBroadcastContent { get; set; }
    
    // publishedAt: дата и время создания ресурса, который идентифицируется в результате поиска
    [JsonPropertyName("publishedAt")]
    public string PublishedAt { get; set; }
    
    // Миниатюры: карта уменьшенных изображений, связанных с результатом поиска. Для
    // каждого объекта на карте ключом является название уменьшенного изображения, а значением
    // является объект, содержащий другую информацию об уменьшенном изображении.
    [JsonPropertyName("thumbnails")]
    public ThumbnailDetails Thumbnails { get; set; }
    
    // Title: The title of the search result.
    [JsonPropertyName("title")]
    public string Title { get; set; }
}

// Миниатюра: Миниатюра - это изображение, представляющее ресурс YouTube.
public class Thumbnail
{
    // Высота: (необязательно) Высота уменьшенного изображения.
    [JsonPropertyName("height")]
    public int Height { get; set; }
    
    // Url: URL-адрес уменьшенного изображения.
    [JsonPropertyName("url")]
    public string Url { get; set; }
    
    // Ширина: (необязательно) Ширина уменьшенного изображения.
    [JsonPropertyName("width")]
    public int Width { get; set; }
}

// ThumbnailDetails: внутреннее представление миниатюр для ресурса YouTube
public class ThumbnailDetails
{
    // По умолчанию: изображение по умолчанию для данного ресурса.
    [JsonPropertyName("default")]
    public Thumbnail Default { get; set; }
    
    // High: Изображение высокого качества для данного ресурса.
    [JsonPropertyName("high")]
    public Thumbnail High { get; set; }
    
    // Maxres: изображение с максимальным разрешением для данного ресурса.
    [JsonPropertyName("maxres")]
    public Thumbnail Maxres { get; set; }
    
    // Средний: изображение среднего качества для данного ресурса.
    [JsonPropertyName("medium")]
    public Thumbnail Medium { get; set; }
    
    // Стандарт: изображение стандартного качества для данного ресурса.
    [JsonPropertyName("standard")]
    public Thumbnail Standard { get; set; }
}