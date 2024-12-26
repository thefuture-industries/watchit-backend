import {
  Clapperboard,
  Youtube,
  Text,
  Monitor,
  LayoutDashboard,
  MoveLeft,
  Heart,
} from "lucide-react";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import SimilarMovie from "~/components/controls/SimilarMovie";
import Skeleton from "~/components/Skeletons/Skeleton";
import StateRequest from "~/components/StateRequest";
import favouritesService from "~/services/favourites-service";
import movieService from "~/services/movie-service";
import { MovieModel } from "~/types/movie";

const MovieDetails = () => {
  const [isSuccessFavourite, setIsSuccessFavourite] = useState<boolean>(false);
  const [isHoverFavourites, setIsHoverFavourites] = useState<boolean>(false);
  const genres: Record<number, string> = {
    28: "Action",
    12: "Adventure",
    16: "Animation",
    35: "Comedy",
    80: "Crime",
    99: "Documentary",
    18: "Drama",
    10751: "Family",
    14: "Fantasy",
    36: "History",
    27: "Horror",
    10402: "Music",
    9648: "Mystery",
    10749: "Romance",
    878: "Science Fiction",
    10770: "TV Movie",
    53: "Thriller",
    10752: "War",
    37: "Western",
  };

  const [poster, setPoster] = useState<string>("");
  const [movieDetails, setMovieDetails] = useState<MovieModel>();
  const [similarMovies, setSimilarMovies] = useState<MovieModel[]>([]);
  const [loaded, setLoaded] = useState<boolean>(false);
  const [isError, setIsError] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  const [totalPage, setTotalPage] = useState<number>(5);

  const navigate = useNavigate();
  let { id } = useParams();

  // Получение данных о фильме
  useEffect(() => {
    // Получение деталий фильма по ID
    const movieDetails = async () => {
      try {
        // Получение деталий фильма по ID
        let movieDetails: MovieModel = await movieService.details(Number(id));
        setMovieDetails(movieDetails);

        // Получение похожих фильмов
        try {
          let similarMovies = await movieService.similar(
            movieDetails.genre_ids,
            movieDetails.title,
            movieDetails.overview
          );

          // Удаление повторяющихся фильмов
          const uniqueIDs = new Set();
          const filteredMovies = similarMovies.filter((movie) => {
            const id = movie.id;
            if (!uniqueIDs.has(id)) {
              uniqueIDs.add(id);
              return true;
            }

            return false;
          });
          setSimilarMovies(filteredMovies);
        } catch (err) {
          setSimilarMovies([]);
        }

        // Загрузка и отображение изображение фильма
        const img = new Image();
        img.src = `https://image.tmdb.org/t/p/w500${movieDetails?.poster_path}`;
        img.onload = () => {
          setPoster(
            `https://image.tmdb.org/t/p/w500${movieDetails?.poster_path}`
          );
          setLoaded(true);
        };
        img.onerror = () => {
          setPoster("/src/assets/default_image.jpg");
          setLoaded(true);
        };
      } catch (error) {
        setIsError(true);
        setError("error movie details");
      }
    };

    movieDetails();
  }, [movieDetails?.poster_path]);

  return (
    <>
      {/* ERROR */}
      {isError && (
        <StateRequest
          message={error}
          statusCode={500}
          state={isError}
          setState={setIsError}
        />
      )}
      
      <div className="my-5 ml-[5rem]">
        <div className="flex items-center justify-between">
          <MoveLeft
            size={27}
            className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
            onClick={() => {
              navigate(-1);
            }}
          />
          <div className="flex items-center gap-[1rem] mr-[4rem]">
            <Link to="/youtube/filter">
              <Youtube
                size={22}
                className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
              />
            </Link>
            <Link to="/movie/filter">
              <Clapperboard
                size={22}
                className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
              />
            </Link>
            <Link to="/story">
              <Text
                size={22}
                className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
              />
            </Link>
          </div>
        </div>
        <div className="flex mt-7 items-center">
          {/* POSTER */}
          <div>
            {loaded ? (
              <img
                src={poster}
                style={{ opacity: 0.8 }}
                className="rounded-xl object-cover max-w-[40vw]"
                onError={() => {
                  setPoster("/src/assets/default_image.jpg");
                }}
              />
            ) : (
              <Skeleton width={27} height={35} />
            )}
          </div>
          <div className="ml-[4rem]">
            {/* TITLE */}
            <p className="uppercase text-[2rem] font-semibold leading-10">
              {movieDetails?.title}
            </p>

            {/* MINI DEATILS */}
            <div className="flex items-center mt-2">
              <p className="text-[#888]">
                {movieDetails?.release_date.slice(0, 4)}
              </p>
              <span className="mx-2 text-[#888]">|</span>
              <p className="text-[#888]">18+</p>
            </div>

            {/* OVERVIEW */}
            <p className="mt-5 text-[#e6e6e6] font-light tracking-wide max-w-[40vw]">
              {movieDetails?.overview}
            </p>

            {/* RANGE */}
            <div className="mt-5">
              <p className="mt-2">
                <span className="text-[#777]">Vote Average</span>
                <span className="ml-[3.3rem]">
                  {movieDetails?.vote_average}
                </span>
              </p>
              <p className="mt-2">
                <span className="text-[#777]">Genre</span>
                <span className="ml-[6rem]">
                  {movieDetails?.genre_ids.map((genre, index: number) => (
                    <span key={index}>{genres[genre]}, </span>
                  ))}
                </span>
              </p>
              <p className="mt-2">
                <span className="text-[#777]">Language</span>
                <span className="ml-[4.8rem] uppercase">
                  {movieDetails?.original_language}
                </span>
              </p>
            </div>

            <div className="flex items-center mt-10 gap-[2rem]">
              <Link
                to={`https://www.justwatch.com/us/movie/${movieDetails?.title
                  .toLocaleLowerCase()
                  .replace(/ /g, "-")}`}
                className="flex items-center cursor-pointer bg-[#b7b7b7] hover:bg-[#fff] transition py-2 px-9 rounded inline-block font-semibold"
              >
                <Monitor strokeWidth={1.5} color="#000" className="mr-3" />
                <p className="uppercase text-[#000]">trailer</p>
              </Link>
              {/* <div className="text-[#888] hover:text-[#fff] transition flex items-center gap-[0.8rem] cursor-pointer">
                <Monitor strokeWidth={1.5} />
                <p className="uppercase">watch</p>
              </div> */}
              <div
                onMouseEnter={() => setIsHoverFavourites(true)}
                onMouseLeave={() => setIsHoverFavourites(false)}
                className={`p-2 rounded-[50%] flex items-center justify-center cursor-pointer ${
                  isSuccessFavourite ? "motion-preset-confetti" : ""
                }`}
                style={{
                  background: "rgba(0, 0, 0, 0.3)",
                  boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.3)",
                  border: "1px solid #fff",
                }}
                onClick={async () => {
                  setIsSuccessFavourite(true);
                  await favouritesService.add({
                    movie_id: movieDetails?.id as number,
                    movie_poster: movieDetails?.poster_path as string,
                  });
                }}
              >
                <Heart
                  size={18}
                  color={`${isHoverFavourites ? "red" : "white"}`}
                  className={`fill-[${
                    isHoverFavourites ? "red" : ""
                  }] transition`}
                />
              </div>
            </div>
          </div>
        </div>

        {/* SIMILAR */}
        <div className="mt-5 mr-[4rem]">
          <div className="w-full h-[1px] bg-[#333] mb-5" />
          <div className="flex items-center">
            <LayoutDashboard size={20} />
            <p className="ml-3">Similar Movies</p>
          </div>
          <div className="w-full h-[1px] bg-[#333] mt-5" />

          <div className="mt-5 flex gap-y-[1.5rem] justify-between flex-wrap">
            {similarMovies.slice(0, totalPage).map((movie, index) => (
              <SimilarMovie movie={movie} key={index} />
            ))}
          </div>
          {totalPage < similarMovies.length && (
            <div
              className="bg-[#111] p-3 rounded-lg flex justify-center cursor-pointer mt-3"
              onClick={() => setTotalPage(totalPage + 10)}
            >
              <p className="text-[#555] tracking-wide text-[0.9rem]">
                Show More
              </p>
            </div>
          )}
        </div>
      </div>
    </>
  );
};

export default MovieDetails;
