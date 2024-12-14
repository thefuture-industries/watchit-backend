import { useState } from "react";
import Navigation from "~/components/Navigation";
import movieService from "~/services/movie-service";
import { MovieModel } from "~/types/movie";
import SearchResult from "./SearchResult";

const Story = () => {
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [movies, setMovies] = useState<MovieModel[]>([]);
  const [plotInput, setPlotInput] = useState<string>("");

  return (
    <>
      {searchPage ? (
        <SearchResult movies={movies} videos={null} />
      ) : (
        <div className="container flex w-screen m-2">
          <div className="left">
            <Navigation />
          </div>
          <div className="right ml-[19rem] w-[67vw]">
            <div className="flex flex-col items-center">
              {/* Search */}
              <div className="flex justify-center">
                <textarea
                  className="w-[68vw] mb-4 bg-[transparent] border border-[#222] rounded h-[10rem] text-[1.3rem] p-3 resize-none outline-none"
                  placeholder="Search"
                  value={plotInput}
                  onChange={(e) => setPlotInput(e.target.value)}
                ></textarea>
              </div>
              <div
                className="bg-[#111] mb-2 w-[68vw] p-2 border border-[#333] rounded flex justify-center items-center cursor-pointer text-[#999] hover:bg-[#000] hover:text-[#fff] transition"
                onClick={async () => {
                  await movieService
                    .plot(plotInput, "simple")
                    .then((movies) => {
                      setMovies(movies);
                      setSearchPage(true);
                    });
                }}
              >
                <p className="text-[0.9rem]">Not an accurate search</p>
              </div>
              <div
                className="bg-[#111] w-[68vw] p-2 border border-[#333] rounded flex justify-center items-center cursor-pointer text-[#999] hover:bg-[#000] hover:text-[#fff] transition"
                onClick={async () => {
                  await movieService.plot(plotInput, "exact").then((movies) => {
                    setMovies(movies);
                    setSearchPage(true);
                  });
                }}
              >
                <p className="text-[0.9rem]">Accurate search</p>
              </div>

              <div className="text-center mt-[4vh]">
                <img src="/src/assets/circle_animation.gif" alt="" />
                <p className="text-[1.2rem]">
                  The story or a short retelling of the film
                </p>
              </div>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default Story;
