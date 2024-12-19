import { Heart, Play } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";
import { MovieModel } from "~/types/movie";
import Skeleton from "../Skeletons/Skeleton";
import lazyService from "~/services/lazy-service";
import favouritesService from "~/services/favourites-service";
import recommendationsService from "~/services/recommendations-service";

interface Props {
  movies: MovieModel;
}

const Movie = (prop: Props) => {
  const [isSuccessFavourite, setIsSuccessFavourite] = useState<boolean>(false);
  const [isHoverFavourites, setIsHoverFavourites] = useState<boolean>(false);
  const [poster, setPoster] = useState<string>("");
  const [loaded, setLoaded] = useState<boolean>(false);
  const [isHovered, setIsHovered] = useState(false);
  const imgRef = useRef(null);

  // Загрузка изображения
  useEffect(() => {
    const cleanup = lazyService.createImageObserver(
      imgRef,
      `https://image.tmdb.org/t/p/w500${prop.movies.poster_path}`,
      setPoster,
      setLoaded
    );

    return cleanup;
  }, [prop.movies.poster_path]);

  return (
    <>
      <div className="">
        <div className="flex items-center flex-wrap">
          <div
            className="m-2 relative"
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
          >
            <div ref={imgRef}>
              {loaded ? (
                <>
                  <img
                    src={poster}
                    className={`${
                      isHovered ? "opacity-[0.3]" : "opacity-[1]"
                    } duration-150 rounded-xl object-cover w-[13.2rem] h-[18rem] pre-load`}
                    // onError={() => {
                    //   setPoster("/src/assets/default_image.jpg");
                    // }}
                  />
                  <div
                    onMouseEnter={() => setIsHoverFavourites(true)}
                    onMouseLeave={() => setIsHoverFavourites(false)}
                    className={`absolute top-3 right-3 p-2 rounded-[50%] flex items-center justify-center cursor-pointer ${
                      isSuccessFavourite ? "motion-preset-confetti" : ""
                    }`}
                    style={{
                      background: "rgba(0, 0, 0, 0.3)",
                      boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.3)",
                      border: "1px solid #fff",
                    }}
                    onClick={async () => {
                      setIsSuccessFavourite(true);
                      await Promise.all([
                        favouritesService.add({
                          movie_id: prop.movies.id,
                          movie_poster: prop.movies.poster_path,
                        }),
                        recommendationsService.add({
                          title: prop.movies.title,
                          genre: prop.movies.genre_ids.join(", "),
                        }),
                      ]);
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
                  {/* <div
                    className={`absolute duration-200 ${
                      isHovered ? "opacity-100 visible" : "opacity-0 invisible"
                    }`}
                    style={{
                      left: "50%",
                      top: "50%",
                      transform: "translate(-50%, -50%)",
                    }}
                  >
                    <Link
                      to={`/movie/${prop.movies.id}`}
                      className="flex items-center"
                    >
                      <div className="bg-[transparent] hover:bg-[#fff] duration-300 p-2 border-[2px] rounded-3xl">
                        <div
                          className="cursor-pointer bg-[#fff] rounded-[50%] p-3 transform transition-transform hover:scale-125"
                          style={{
                            boxShadow: "0 0 10px 0 rgba(255, 255, 255, 0.3)",
                          }}
                        >
                          <Play color="#000" fill="#000" size={19} />
                        </div>
                      </div>
                    </Link>
                  </div> */}
                  <div
                    className={`absolute flex flex-col items-center duration-200 ${
                      isHovered ? "opacity-100 visible" : "opacity-0 invisible"
                    }`}
                    style={{
                      left: "50%",
                      bottom: "10%",
                      transform: "translateX(-50%)",
                    }}
                  >
                    <p className="text-center mb-[0.7rem] tracking-wide text-[1.1rem]">
                      {prop.movies.title}
                    </p>
                    <Link
                      to={`/movie/${prop.movies.id}`}
                      className="flex items-center"
                    >
                      <div className="bg-[transparent] hover:bg-[#fff] duration-300 p-2 border-[2px] rounded-3xl">
                        <div
                          className="cursor-pointer bg-[#fff] rounded-[50%] p-3 transform transition-transform hover:scale-125"
                          style={{
                            boxShadow: "0 0 10px 0 rgba(255, 255, 255, 0.3)",
                          }}
                        >
                          <Play color="#000" fill="#000" size={19} />
                        </div>
                      </div>
                    </Link>
                  </div>
                </>
              ) : (
                <Skeleton width={13.2} height={18} />
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Movie;
