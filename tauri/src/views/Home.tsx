import Movie from "~/components/controls/Movie";
import YoutubeIndex from "~/components/controls/YoutubeIndex";
import Header from "~/components/Header";
import Navigation from "~/components/Navigation";
import React, { useCallback, useEffect, useState } from "react";
import movieService from "~/services/movie-service";
import { MoveRight } from "lucide-react";
import youtubeService from "~/services/youtube-service";
import { YoutubeModel } from "~/types/youtube";
import { MovieModel } from "~/types/movie";
import SearchResult from "./SearchResult";
import recommendationsService from "~/services/recommendations-service";
import { throttle } from "lodash";
import { useQuery } from "react-query";
import userService from "~/services/user-service";
import { invoke } from "@tauri-apps/api/core";
import StateRequest from "~/components/StateRequest";

// Компонент список фильмов
const MovieList = React.memo(
  ({
    movies,
    onLoadMore,
  }: {
    movies: MovieModel[];
    onLoadMore: () => void;
  }) => (
    <div className="flex justify-center items-center flex-wrap mt-4">
      {movies.map((movie, index) => (
        <Movie movies={movie} key={movie.id || index} />
      ))}
      <div
        className="flex items-center text-[#555] hover:text-[#fff] ml-3 cursor-pointer transition"
        onClick={onLoadMore}
      >
        <p className="text-[1.1rem]">More</p>
        <MoveRight className="ml-2 pt-2" size={31} />
      </div>
    </div>
  )
);

// Компонент кнопка вверх/вниз
const ScrollButton = React.memo(
  ({
    isVisible,
    onClick,
  }: {
    isVisible: boolean;
    isAtBottom: boolean;
    onClick: () => void;
  }) => {
    if (!isVisible) return null;

    return (
      <div
        id="scroll-btn"
        className="fixed bottom-3 right-3 px-2 py-[0.6rem] cursor-pointer inline-block rounded transform transition-transform hover:scale-110"
        style={{
          background: "rgba(255, 255, 255, 0.4)",
          boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.3)",
        }}
        onClick={onClick}
      >
        <p className="text-[1.2rem]">👇</p>
      </div>
    );
  }
);

const Home = () => {
  const [error, setError] = useState<string>("");
  const [isError, setIsError] = useState<boolean>(false);
  const [isButtonVisible, setIsButtonVisible] = useState<boolean>(false);
  const [isAtBottom, setIsAtBottom] = useState<boolean>(false);
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [movies, setMovies] = useState<MovieModel[]>([]);

  const { data: video } = useQuery("video", youtubeService.get_video, {
    initialData: [] as YoutubeModel[],
    refetchOnWindowFocus: false,
  });

  const { refetch } = useQuery("refetch", recommendationsService.get, {
    initialData: [] as MovieModel[],
    onSuccess: (data) => {
      setMovies(data);
      sessionStorage.setItem("sess_movies", JSON.stringify(data));
    },
    refetchOnWindowFocus: false,
  });

  // Скрол вниз/верх
  const scrollToTopOrBottom = useCallback(() => {
    if (isAtBottom) {
      window.scrollTo({ top: 0, behavior: "smooth" });
    } else {
      window.scrollTo({ top: document.body.scrollHeight, behavior: "smooth" });
    }
  }, [isAtBottom]);

  // Поиск фильмов
  const handleSearch = useCallback(async (searchQuery: string) => {
    let movies_search = await movieService.search(searchQuery);
    setMovies(movies_search);
    setSearchPage(true);
  }, []);

  // Загрузка большего количества фильмов
  const handleLoadMore = useCallback(async () => {
    try {
      const newMovies: MovieModel[] = await invoke("get_recommendations", {
        uuid: userService.get_uuid(),
      });

      setMovies((prevMovies) => {
        const updateMovie = [...prevMovies, ...(newMovies as MovieModel[])];
        sessionStorage.setItem("sess_movies", JSON.stringify(updateMovie));
        return updateMovie;
      });
    } catch (error) {
      setIsError(true);
      setError(error as string);
    }
  }, [refetch]);

  useEffect(() => {
    const threshold = 100;

    const handleScroll = throttle(() => {
      const currentScrollY = window.scrollY;
      const isScrolledDown = currentScrollY > 0;
      setIsButtonVisible(isScrolledDown);

      setIsAtBottom(
        window.innerHeight + window.scrollY >=
          document.body.offsetHeight - threshold
      );
    }, 200);

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
      handleScroll.cancel();
    };
  }, []);

  return (
    <>
      {/* ERROR */}
      {isError && (
        <StateRequest
          statusCode={400}
          message={error as string}
          state={isError}
          setState={setIsError}
        />
      )}
      {searchPage ? (
        <SearchResult movies={movies as MovieModel[]} videos={null} />
      ) : (
        <div className="container flex w-screen m-2">
          <div className="left">
            <Navigation />
          </div>
          <div className="right ml-[19rem] w-[67vw]">
            <Header onSearch={handleSearch} />

            {/* Популярное видео */}
            <YoutubeIndex video={video as YoutubeModel[]} />

            {/* Фильмы */}
            <MovieList movies={movies} onLoadMore={handleLoadMore} />

            {/* Кнопка вверх/вниз */}
            <ScrollButton
              isVisible={isButtonVisible}
              isAtBottom={isAtBottom}
              onClick={scrollToTopOrBottom}
            />
          </div>
        </div>
      )}
    </>
  );
};

export default React.memo(Home);
