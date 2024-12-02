use serde::{Deserialize,Serialize};

// YouTubeModel: Результат поиска содержит информацию о видео на YouTube,
// канале или плейлисте, которые соответствуют параметрам поиска, указанным в запросе API
//. Хотя результат поиска указывает на однозначно идентифицируемый ресурс,
// такой как видео, у него нет собственных постоянных данных.
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct YoutubeModel {
  // Etag: Etag этого ресурса.
  pub etag: String,
  // Id: Объект id содержит информацию, которая может быть использована для уникальной идентификации
  // ресурса, соответствующего поисковому запросу.
  pub id: ResourceId,

  // Kind: Указывает, что это за ресурс. Значение: фиксированная строка
  // "youtube#SearchResult".
  pub kind: String,

  // Фрагмент: Объект snippet содержит основные сведения о результате поиска,
  // такие как его заголовок или описание. Например, если результатом поиска является
  // видео, то заголовком будет название видео, а описанием -
  // описание видео.
  pub snippet: SearchResultSnippet,
}

// ResourceId: Идентификатор ресурса - это общая ссылка, которая указывает на другой ресурс
// YouTube.
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ResourceId
{
  // Kind: Тип ресурса API.
  pub kind: String,

  // VideoID: Идентификатор, который YouTube использует для уникальной идентификации ресурса, на который ссылается
  //, если этот ресурс является видео. Это свойство присутствует только в том случае, если значение
  // resourceId.kind равно youtube#video.
  #[serde(rename="videoId")]
  pub video_id: String,
}

// SearchResultSnippet: Основные сведения о результатах поиска, включая заголовок,
// описание и миниатюры элемента, на который ссылается результат поиска.
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SearchResultSnippet
{
  // channelId: значение, которое YouTube использует для уникальной идентификации канала,
  // опубликовавшего ресурс, указанный в результатах поиска.
  #[serde(rename="channelId")]
  pub channel_id: String,

  // ChannelTitle: название канала, опубликовавшего ресурс, который
  // идентифицируется в результате поиска.
  #[serde(rename="channelTitle")]
  pub channel_title: String,

  // Описание: Описание результата поиска.
  #[serde(rename = "description", default)]
  pub description: Option<String>,

  // LiveBroadcastContent: Указывает, есть ли на ресурсе (видео или канале)
  // предстоящий/активный контент для прямой трансляции. Или "нет", если
  // предстоящих/активных прямых трансляций нет.
  //
  // Possible values:
  //   "none"
  //   "upcoming" - The live broadcast is upcoming.
  //   "live" - The live broadcast is active.
  //   "completed" - The live broadcast has been completed.
  #[serde(rename="liveBroadcastContent")]
  pub live_broadcast_content: String,

  // publishedAt: дата и время создания ресурса, который идентифицируется в результате поиска
  #[serde(rename="publishedAt")]
  pub published_at: String,

  // Миниатюры: карта уменьшенных изображений, связанных с результатом поиска. Для
  // каждого объекта на карте ключом является название уменьшенного изображения, а значением
  // является объект, содержащий другую информацию об уменьшенном изображении.
  pub thumbnails: ThumbnailDetails,

  // Title: The title of the search result.
  pub title: String,
}

// Миниатюра: Миниатюра - это изображение, представляющее ресурс YouTube.
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Thumbnail
{
  // Высота: (необязательно) Высота уменьшенного изображения.
  pub height: i32,

  // Url: URL-адрес уменьшенного изображения.
  pub url: String,

  // Ширина: (необязательно) Ширина уменьшенного изображения.
  pub width: i32,
}

// ThumbnailDetails: внутреннее представление миниатюр для ресурса YouTube
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ThumbnailDetails
{
  // По умолчанию: изображение по умолчанию для данного ресурса.
  pub default: Thumbnail,

  // High: Изображение высокого качества для данного ресурса.
  pub high: Thumbnail,

  // Средний: изображение среднего качества для данного ресурса.
  pub medium: Thumbnail,
}
