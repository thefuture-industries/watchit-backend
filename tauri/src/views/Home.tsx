import Movie from "~/components/controls/Movie";
import YoutubeIndex from "~/components/controls/YoutubeIndex";
import Header from "~/components/Header";
import Navigation from "~/components/Navigation";
import { useEffect, useState } from "react";
import movieService from "~/services/movie-service";
import { ChevronsUpDown, MoveRight } from "lucide-react";
import youtubeService from "~/services/youtube-service";
import { YoutubeModel } from "~/types/youtube";
import { MovieModel } from "~/types/movie";
import Error from "~/components/Error";
import SearchResult from "./SearchResult";

const Home = () => {
  const [isButtonVisible, setIsButtonVisible] = useState<boolean>(false);
  const [isAtBottom, setIsAtBottom] = useState<boolean>(false);

  const [searchPage, setSearchPage] = useState<boolean>(false);

  const [errorMessage, setErrorMessage] = useState<string>("");
  const [isError, setIsError] = useState<boolean>(false);
  const [movies, setMovies] = useState<MovieModel[]>([]);
  const [video, setVideo] = useState<YoutubeModel[]>([]);

  // Скрол вниз/верх
  const scrollToTopOrBottom = () => {
    if (isAtBottom) {
      window.scrollTo({ top: 0, behavior: "smooth" });
    } else {
      window.scrollTo({ top: document.body.scrollHeight, behavior: "smooth" });
    }
  };

  const handleSearch = async (searchQuery: string) => {
    let movies_search = await movieService.search_movies(searchQuery);
    setMovies(movies_search);
    setSearchPage(true);
  };

  // Показ кнопки вниз/верх
  useEffect(() => {
    const threshold = 100;
    const handleScroll = () => {
      const currentScrollY = window.scrollY;
      const isScrolledDown = currentScrollY > 0;
      setIsButtonVisible(isScrolledDown);

      setIsAtBottom(
        window.innerHeight + window.scrollY >=
          document.body.offsetHeight - threshold
      );
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  // Получение популярных фильмов
  useEffect(() => {
    const movies_and_video = async () => {
      try {
        const [moviesPopular, videoData] = await Promise.all([
          movieService.get_popular_movies(),
          youtubeService.get_video(),
        ]);

        // Удаление повторяющихся фильмов
        setMovies(moviesPopular);
        setVideo(videoData);
      } catch (err: any) {
        setErrorMessage(err);
        setIsError(true);
      }
    };

    movies_and_video();
  }, []);

  return (
    <>
      {/* ERROR */}
      {isError && <Error errorMessage={errorMessage} />}
      {searchPage ? (
        <SearchResult movies={movies} videos={null} />
      ) : (
        <div className="container flex w-screen m-2">
          <div className="left">
            <Navigation />
          </div>
          <div className="right ml-[19rem] w-[67vw]">
            <Header onSearch={handleSearch} />

            {/* Популярное видео */}
            <YoutubeIndex video={video} />

            {/* Фильмы */}
            <div className="flex justify-center items-center flex-wrap mt-4">
              {movies.map((movie, index) => (
                <Movie movies={movie} key={index} />
              ))}
              <div
                className="flex items-center text-[#555] hover:text-[#fff] ml-3 cursor-pointer transition"
                onClick={async () => {
                  movieService.increment_page_popular();
                  setMovies(await movieService.get_popular_movies());
                }}
              >
                <p className="text-[1.1rem]">More</p>
                <MoveRight className="ml-2 pt-2" size={31} />
              </div>
            </div>

            {/* Кнопка вверх/вниз */}
            {isButtonVisible && (
              <div
                id="scroll-btn"
                className="fixed bottom-3 right-3 bg-[#fff] p-1 cursor-pointer inline-block rounded transform transition-transform hover:scale-110"
                onClick={scrollToTopOrBottom}
              >
                <ChevronsUpDown color="#000" size={15} />
              </div>
            )}
          </div>
        </div>
      )}
    </>
  );
};

export default Home;
