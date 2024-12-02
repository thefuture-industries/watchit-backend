import { useState } from "react";
import Navigation from "~/components/Navigation";
import movieService from "~/services/movie-service";
import SearchResult from "./SearchResult";
import { MovieModel } from "~/types/movie";

const categorys = [
  "Action",
  "Adventure",
  "Animation",
  "Comedy",
  "Crime",
  "Documentary",
  "Drama",
  "Family",
  "Fantasy",
  "History",
  "Horror",
  "Music",
  "Mystery",
  "Romance",
  "Science Fiction",
  "TV Movie",
  "Thriller",
  "War",
  "Western",
];

const MovieFilter = () => {
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [movies, setMovies] = useState<MovieModel[]>([]);
  const [genre, setGenre] = useState<string>("");
  const [search, setSearch] = useState<string>("");
  const [date, setDate] = useState<string>("");

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
            <div className="bg-[#111] border border-[#222] p-5 rounded-lg">
              <p className="text-[1.5rem] tracking-wide">Movie Filter</p>
              <div className="w-full bg-[#222] h-[1px] my-3 mt-5"></div>

              {/* Search */}
              <textarea
                className="w-full mb-4 bg-[transparent] border border-[#222] rounded h-[7rem] text-[1.3rem] p-3 resize-none outline-none"
                placeholder="Search"
                onChange={(e) => setSearch(e.target.value)}
                value={search}
              ></textarea>

              {/* Category */}
              <p className="text-[1rem] tracking-wide">Category</p>
              <div className="w-full bg-[#222] h-[1px] my-3"></div>
              <div>
                {categorys.map((item, index) => (
                  <div
                    key={index}
                    className={`bg-[${genre == item ? "#fff" : ""}] text-[${
                      genre == item ? "#000" : "#fff"
                    }] inline-block border border-[#555] py-[1px] pb-1 px-3 rounded-3xl m-1`}
                  >
                    <div
                      className="flex items-center cursor-pointer"
                      onClick={() => setGenre(item)}
                    >
                      <div
                        className={`w-[15px] h-[15px] rounded-3xl border-2 border-[${
                          genre == item ? "#000" : "#fff"
                        }] mr-[8px]`}
                      ></div>
                      <p>{item}</p>
                    </div>
                  </div>
                ))}
              </div>

              {/* TOPS */}
              <p className="text-[1rem] tracking-wide mt-3">
                The year of the film
              </p>
              <div className="w-full bg-[#222] h-[1px] my-3"></div>
              <div>
                <input
                  type="text"
                  className="w-full"
                  placeholder="Year"
                  value={date}
                  onChange={(e) => setDate(e.target.value)}
                />
              </div>

              <div
                className="mt-5 bg-[#ff2400] hover:bg-[#b21900] transition flex justify-center items-center min-h-[47px] rounded-lg cursor-pointer"
                style={{
                  boxShadow: "inset 0px -7px 0px 0px rgba(0, 0, 0, 0.4)",
                }}
                onClick={async () => {
                  await movieService
                    .get_movies(search, genre, date)
                    .then((movies) => {
                      setMovies(movies);
                      setSearchPage(true);
                    });
                }}
              >
                <p className="uppercase text-[17px] -mt-1">search</p>
              </div>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default MovieFilter;
