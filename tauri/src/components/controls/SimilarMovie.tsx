import { useEffect, useRef, useState } from "react";
import { MovieModel } from "~/types/movie";
import Skeleton from "../Skeletons/Skeleton";
import { Link } from "react-router-dom";
import { Play } from "lucide-react";
import lazyService from "~/services/lazy-service";

interface Props {
  movie: MovieModel;
}

const SimilarMovie = (prop: Props) => {
  const [poster, setPoster] = useState<string>("");
  const [loaded, setLoaded] = useState<boolean>(false);
  const [isHovered, setIsHovered] = useState(false);
  const imgRef = useRef(null);

  // Загрузка изображений
  useEffect(() => {
    const cleanup = lazyService.createImageObserver(
      imgRef,
      `${import.meta.env.VITE_SERVER_URL}/image/w500${prop.movie.poster_path}`,
      setPoster,
      setLoaded
    );

    return cleanup;
  }, [prop.movie.poster_path]);

  return (
    <>
      <div>
        <div
          className="relative"
          ref={imgRef}
          onMouseEnter={() => setIsHovered(true)}
          onMouseLeave={() => setIsHovered(false)}
        >
          {loaded ? (
            <>
              <img
                src={poster}
                className={`${
                  isHovered ? "opacity-[0.3]" : "opacity-[1]"
                } duration-150 rounded object-cover max-w-[10rem] min-h-[15rem]`}
                onError={() => {
                  setPoster("/src/assets/default_image.jpg");
                }}
              />
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
                  to={`/movie/${prop.movie.id}`}
                  onClick={location.reload}
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
            <Skeleton width={10} height={15} />
          )}
        </div>

        <div className="mt-2">
          <p className="tracking-wide max-w-[10rem]">{prop.movie.title}</p>
          <p className="tracking-wide text-[0.8rem] text-[#999]">
            {prop.movie.vote_average}
          </p>
        </div>
      </div>
    </>
  );
};

export default SimilarMovie;
