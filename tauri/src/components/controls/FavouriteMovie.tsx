import { Heart, Play } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import lazyService from "~/services/lazy-service";
import { FavouriteModel } from "~/types/favourites";
import Skeleton from "../Skeletons/Skeleton";
import { Link } from "react-router-dom";
import favouritesService from "~/services/favourites-service";

interface Props {
  movie: FavouriteModel;
}

const FavouriteMovie = (prop: Props) => {
  const [isSuccessFavourite, setIsSuccessFavourite] = useState<boolean>(false);
  const [poster, setPoster] = useState<string>("");
  const [loaded, setLoaded] = useState<boolean>(false);
  const [isHovered, setIsHovered] = useState(false);
  const imgRef = useRef(null);

  // Загрузка изображения
  useEffect(() => {
    const cleanup = lazyService.createImageObserver(
      imgRef,
      `${import.meta.env.VITE_SERVER_URL}/image/w500${prop.movie.moviePoster}`,
      setPoster,
      setLoaded
    );

    return cleanup;
  }, [prop.movie.moviePoster]);

  return (
    <div>
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
                />
                <div
                  className={`absolute top-3 right-3 p-2 rounded-[50%] flex items-center justify-center cursor-pointer hover:text-[red] transition ${
                    isSuccessFavourite ? "motion-preset-confetti" : ""
                  }`}
                  style={{
                    background: "rgba(0, 0, 0, 0.3)",
                    boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.3)",
                    border: "1px solid #fff",
                  }}
                  onClick={async () => {
                    setIsSuccessFavourite(true);
                    await favouritesService.delete(prop.movie.movieId);
                    location.reload();
                  }}
                >
                  <Heart
                    size={18}
                    color="red"
                    className="fill-[red] transition"
                  />
                </div>
                <div
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
                    to={`/movie/${prop.movie?.movieId}`}
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
  );
};

export default FavouriteMovie;
