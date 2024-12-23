import { useState } from "react";
import Navigation from "~/components/Navigation";
import movieService from "~/services/movie-service";
import { MovieModel } from "~/types/movie";
import SearchResult from "./SearchResult";

const Story = () => {
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [isSend, setIsSend] = useState<boolean>(false);
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
                  setIsSend(true);
                  try {
                    await movieService
                      .plot(plotInput, "simple")
                      .then((movies) => {
                        setMovies(movies);
                        setSearchPage(true);
                      });
                  } catch {
                    setMovies([]);
                    setSearchPage(true);
                  } finally {
                    setIsSend(false);
                  }
                }}
              >
                {isSend ? (
                  <svg
                    className="animate-spin border-indigo-600"
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 64 64"
                    fill="none"
                  >
                    <g id="Group 1000003699">
                      <circle
                        id="Ellipse 715"
                        cx="31.9989"
                        cy="31.8809"
                        r="24"
                        stroke="#888"
                        stroke-width="7"
                      />
                      <path
                        id="Ellipse 716"
                        d="M42.111 53.6434C44.9694 52.3156 47.5383 50.4378 49.6709 48.1172C51.8036 45.7967 53.4583 43.0787 54.5406 40.1187C55.6229 37.1586 56.1115 34.0143 55.9787 30.8654C55.8458 27.7165 55.094 24.6246 53.7662 21.7662C52.4384 18.9078 50.5606 16.339 48.24 14.2063C45.9194 12.0736 43.2015 10.4189 40.2414 9.33662C37.2814 8.25434 34.1371 7.76569 30.9882 7.89856C27.8393 8.03143 24.7473 8.78323 21.889 10.111"
                        stroke="#fff"
                        stroke-width="7"
                        stroke-linecap="round"
                      />
                    </g>
                  </svg>
                ) : (
                  <p className="text-[0.9rem]">Not an accurate search</p>
                )}
                {/* <p className="text-[0.9rem]">Not an accurate search</p> */}
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
