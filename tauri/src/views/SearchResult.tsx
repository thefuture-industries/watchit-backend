import React, {
  lazy,
  useCallback,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import Navigation from "~/components/Navigation";
import { MovieModel } from "~/types/movie";
import { YoutubeModel } from "~/types/youtube";
const Movie = lazy(() => import("~/components/controls/Movie"));
const Youtube = lazy(() => import("~/components/controls/Youtube"));

interface Props {
  movies: MovieModel[] | null;
  videos: YoutubeModel[] | null;
}

// Проверка типа
function isMovieModel(movie: any) {
  return (
    typeof movie === "object" &&
    movie != null &&
    "adult" in movie &&
    "overview" in movie &&
    "poster_path" in movie
  );
}

// Пустое состояние
const EmptyState = React.memo(() => (
  <div className="flex flex-col items-center justify-center h-screen">
    {/* <div className="text-[10vw]">🤷‍♂️</div> */}
    <img src="/src/assets/empty_img.webp" width="200px" alt="" />
    <p className="mt-[5vw] text-[1.5rem]">
      We didn't find what you were looking for
    </p>
  </div>
));

// Отображения массива
const RenderItem = React.memo(
  ({ item, index }: { item: MovieModel | YoutubeModel; index: number }) => {
    if (isMovieModel(item)) {
      return <Movie key={index} movies={item as MovieModel} />;
    }

    return <Youtube key={index} videos={item as YoutubeModel} />;
  }
);

// Компонент для отображения массива в container
const ResultsContainer = React.memo(
  ({
    containerRef,
    rendered,
  }: {
    containerRef: React.RefObject<HTMLDivElement>;
    rendered: (MovieModel | YoutubeModel)[];
  }) => (
    <div
      ref={containerRef}
      className="flex items-stretch flex-wrap justify-center"
      style={{
        height: "70vw",
        overflowY: "auto",
      }}
    >
      {rendered.map((item: any, index: number) => (
        <RenderItem key={index} item={item} index={index} />
      ))}
    </div>
  )
);

// IntersectionObserver
const createObserver = (
  containerRef: React.RefObject<HTMLDivElement>,
  callback: any,
  options: any
) => {
  const observer = new IntersectionObserver(callback, options);

  if (
    containerRef.current &&
    containerRef.current.lastChild instanceof Element
  ) {
    observer.observe(containerRef.current.lastChild);
  }

  return observer;
};

const SearchResult = (prop: Props) => {
  const [rendered, setRendered] = useState<MovieModel[] | YoutubeModel[]>([]);
  const containerRef = useRef<HTMLDivElement>(null);

  // Мемоизируем исходные данные
  const data = useMemo(
    () => prop.movies || prop.videos,
    [prop.movies, prop.videos]
  );

  // Инициализация данных
  useEffect(() => {
    setRendered(data ? data.slice(0, 10) : []);
  }, [data]);

  const handleIntersection = useCallback(
    (entries: IntersectionObserverEntry[]) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const data = prop.movies || prop.videos;

          if (data) {
            const newData = data.slice(0, rendered.length + 10);
            setRendered(newData);
          } else {
            console.warn("prop.movies and prop.videos null or undefined.");
          }
        }
      });
    },
    [prop.movies, prop.videos, rendered]
  );

  useEffect(() => {
    const observer = createObserver(containerRef, handleIntersection, {
      root: containerRef.current,
      rootMargin: "0px",
      threshold: 0.1,
    });

    return () => {
      if (observer) observer.disconnect();
    };
  }, [handleIntersection]);

  // Мемоизируем основной контент
  const content = useMemo(
    () =>
      rendered.length == 0 ? (
        <EmptyState />
      ) : (
        <ResultsContainer containerRef={containerRef} rendered={rendered} />
      ),
    [rendered]
  );

  return (
    <>
      <div className="container flex w-screen m-2">
        <div className="left">
          <Navigation />
        </div>
        <div className="right ml-[19rem] w-[67vw]">{content}</div>
      </div>
    </>
  );
};

export default SearchResult;
