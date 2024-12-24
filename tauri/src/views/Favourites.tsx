import React, { useEffect, useState } from "react";
import { useQuery } from "react-query";
import FavouriteMovie from "~/components/controls/FavouriteMovie";
import Loader from "~/components/Loader";
import Navigation from "~/components/Navigation";
import StateRequest from "~/components/StateRequest";
import favouritesService from "~/services/favourites-service";
import { FavouriteModel } from "~/types/favourites";

// Мемоизированный список фильмов
const FavouritesList = React.memo(
  ({ movies }: { movies: FavouriteModel[] }) => (
    <div className="flex justify-center items-center flex-wrap mt-4">
      {movies.map((movie) => (
        <FavouriteMovie movie={movie} key={movie.id} />
      ))}
    </div>
  )
);

// Пустое состояние
const FavouritesEmpty = React.memo(() => (
  <div className="flex flex-col items-center justify-center h-screen">
    <img src="/src/assets/empty_img.webp" width="200px" alt="" />
    <p className="mt-[5vw] text-[1.5rem]">
      You haven't added any movies to your favorites yet
    </p>
  </div>
));

const Favourites = () => {
  const [isButtonVisible, setIsButtonVisible] = useState<boolean>(false);
  const [isAtBottom, setIsAtBottom] = useState<boolean>(false);
  const {
    data: movies,
    isLoading,
    isError,
    error,
  } = useQuery("movies", favouritesService.get, {
    initialData: [] as FavouriteModel[],
    refetchOnWindowFocus: false,
  });

  // Скрол вниз/верх
  const scrollToTopOrBottom = () => {
    if (isAtBottom) {
      window.scrollTo({ top: 0, behavior: "smooth" });
    } else {
      window.scrollTo({ top: document.body.scrollHeight, behavior: "smooth" });
    }
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

  return (
    <>
      {isLoading && <Loader />}
      {isError && (
        <StateRequest
          message={error as string}
          statusCode={500}
          state={isError}
          setState={() => void 0}
        />
      )}

      <div className="container flex w-screen m-2">
        <div className="left">
          <Navigation />
        </div>
        <div className="right ml-[19rem] w-[67vw]">
          {movies?.length == 0 && <FavouritesEmpty />}

          {/* Фильмы */}
          <FavouritesList movies={movies as FavouriteModel[]} />

          {/* Кнопка вверх/вниз */}
          {isButtonVisible && (
            <div
              id="scroll-btn"
              className="fixed bottom-3 right-3 px-2 py-[0.6rem] cursor-pointer inline-block rounded transform transition-transform hover:scale-110"
              style={{
                background: "rgba(255, 255, 255, 0.4)",
                boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.3)",
              }}
              onClick={scrollToTopOrBottom}
            >
              <p className="text-[1.2rem]">👇</p>
            </div>
          )}
        </div>
      </div>
    </>
  );
};

export default Favourites;
