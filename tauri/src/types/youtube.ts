// YouTubeModel: Результат поиска содержит информацию о видео на YouTube,
// канале или плейлисте, которые соответствуют параметрам поиска, указанным в запросе API
//. Хотя результат поиска указывает на однозначно идентифицируемый ресурс,
// такой как видео, у него нет собственных постоянных данных.
export type YoutubeModel = {
  // Etag: Etag этого ресурса.
  etag: string;
  // Id: Объект id содержит информацию, которая может быть использована для уникальной идентификации
  // ресурса, соответствующего поисковому запросу.
  id: ResourceId;

  // Kind: Указывает, что это за ресурс. Значение: фиксированная строка
  // "youtube#SearchResult".
  kind: string;

  // Фрагмент: Объект snippet содержит основные сведения о результате поиска,
  // такие как его заголовок или описание. Например, если результатом поиска является
  // видео, то заголовком будет название видео, а описанием -
  // описание видео.
  snippet: SearchResultSnippet;
};

// ResourceId: Идентификатор ресурса - это общая ссылка, которая указывает на другой ресурс
// YouTube.
type ResourceId = {
  // Kind: Тип ресурса API.
  kind: string;

  // VideoID: Идентификатор, который YouTube использует для уникальной идентификации ресурса, на который ссылается
  //, если этот ресурс является видео. Это свойство присутствует только в том случае, если значение
  // resourceId.kind равно youtube#video.
  videoId: string;
};

// SearchResultSnippet: Основные сведения о результатах поиска, включая заголовок,
// описание и миниатюры элемента, на который ссылается результат поиска.
type SearchResultSnippet = {
  // channelId: значение, которое YouTube использует для уникальной идентификации канала,
  // опубликовавшего ресурс, указанный в результатах поиска.
  channelId: string;

  // ChannelTitle: название канала, опубликовавшего ресурс, который
  // идентифицируется в результате поиска.
  channelTitle: string;

  // Описание: Описание результата поиска.
  description: string;

  // LiveBroadcastContent: Указывает, есть ли на ресурсе (видео или канале)
  // предстоящий/активный контент для прямой трансляции. Или "нет", если
  // предстоящих/активных прямых трансляций нет.
  //
  // Possible values:
  //   "none"
  //   "upcoming" - The live broadcast is upcoming.
  //   "live" - The live broadcast is active.
  //   "completed" - The live broadcast has been completed.
  liveBroadcastContent: string;

  // publishedAt: дата и время создания ресурса, который идентифицируется в результате поиска
  publishedAt: string;

  // Миниатюры: карта уменьшенных изображений, связанных с результатом поиска. Для
  // каждого объекта на карте ключом является название уменьшенного изображения, а значением
  // является объект, содержащий другую информацию об уменьшенном изображении.
  thumbnails: ThumbnailDetails;

  // Title: The title of the search result.
  title: string;
};

// Миниатюра: Миниатюра - это изображение, представляющее ресурс YouTube.

type Thumbnail = {
  // Высота: (необязательно) Высота уменьшенного изображения.
  height: number;

  // Url: URL-адрес уменьшенного изображения.
  url: string;

  // Ширина: (необязательно) Ширина уменьшенного изображения.
  width: number;
};

// ThumbnailDetails: внутреннее представление миниатюр для ресурса YouTube
type ThumbnailDetails = {
  // По умолчанию: изображение по умолчанию для данного ресурса.
  default: Thumbnail;

  // High: Изображение высокого качества для данного ресурса.
  high: Thumbnail;

  // Средний: изображение среднего качества для данного ресурса.
  medium: Thumbnail;
};

export const DefaultYoutube = {
  etag: "Ov4RWA4PbayLAidueaiMvqgMFn8",
  id: {
    kind: "youtube#video",
    videoId: "fyfFRPhCkGM",
  },
  kind: "youtube#searchResult",
  snippet: {
    channelId: "UCwVg9btOceLQuNCdoQk9CXg",
    channelTitle: "Ben Azelart",
    description:
      "Spider-Man: Into the Spider-Verse is a 2018 American animated superhero film featuring the Marvel Comics character Miles Morales / Spider-Man, produced by Columbia Pictures and Sony Pictures Animation in association with Marvel Entertainment, and distributed by Sony Pictures Releasing. It is the first animated film in the Spider-Man franchise and the first film in the Spider-Verse franchise.",
    liveBroadcastContent: "none",
    publishedAt: "2024-11-16T15:52:41Z",
    thumbnails: {
      default: {
        height: 90,
        url: "/src/assets/default_poster.webp",
        width: 120,
      },
      high: {
        height: 360,
        url: "/src/assets/default_poster.webp",
        width: 480,
      },
      medium: {
        height: 180,
        url: "/src/assets/default_poster.webp",
        width: 320,
      },
    },
    title: "Spider-Man: Into the Spider-Verse",
  },
};
