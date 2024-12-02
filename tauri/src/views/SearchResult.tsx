import { lazy, useCallback, useEffect, useRef, useState } from "react";
import Navigation from "~/components/Navigation";
import { MovieModel } from "~/types/movie";
import { YoutubeModel } from "~/types/youtube";
const Movie = lazy(() => import("~/components/controls/Movie"));
const Youtube = lazy(() => import("~/components/controls/Youtube"));

interface Props {
  movies: MovieModel[] | null;
  videos: YoutubeModel[] | null;
}

function checkMovieModel(movie: any) {
  return (
    typeof movie === "object" &&
    movie != null &&
    "adult" in movie &&
    "overview" in movie &&
    "poster_path" in movie
  );
}

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

  useEffect(() => {
    const data = prop.movies || prop.videos;
    setRendered(data ? data.slice(0, 10) : []);
  }, []);

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

  return (
    <>
      <div className="container flex w-screen m-2">
        <div className="left">
          <Navigation />
        </div>
        <div className="right ml-[19rem] w-[67vw]">
          {rendered.length == 0 ? (
            <div className="flex flex-col items-center justify-center h-screen">
              <div className="text-[10vw]">🤷‍♂️</div>
              <p className="mt-[6.4vw] text-[1.5rem]">
                We didn't find what you were looking for
              </p>
            </div>
          ) : (
            <div
              ref={containerRef}
              className="flex items-stretch flex-wrap justify-center"
              style={{
                height: "70vw",
                overflowY: "auto",
              }}
            >
              {rendered.map((item, index) => {
                if (checkMovieModel(item)) {
                  return <Movie key={index} movies={item as MovieModel} />;
                } else {
                  return <Youtube key={index} videos={item as YoutubeModel} />;
                }
              })}
              {/* <div
              className="flex items-center text-[#555] hover:text-[#fff] ml-3 cursor-pointer transition"
              onClick={async () => {
                let moviesPage =
                  await movieService.increment_page_movie_search();
                setMovies([...movies, ...moviesPage]);
              }}
            >
              <p className="text-[1.1rem]">More</p>
              <MoveRight className="ml-2 pt-2" size={31} />
            </div> */}
            </div>
          )}
        </div>
      </div>
    </>
  );
};

export default SearchResult;
