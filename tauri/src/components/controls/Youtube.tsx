import { useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";
import lazyService from "~/services/lazy-service";
import { YoutubeModel } from "~/types/youtube";
import Skeleton from "../Skeletons/Skeleton";

interface Props {
  videos: YoutubeModel;
}

function timeSince(dateString: string) {
  const date = new Date(dateString);
  const now = new Date();
  const seconds = Math.floor((now.getTime() - date.getTime()) / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);
  const weeks = Math.floor(days / 7);
  const months = Math.floor(days / 30.44); // Приблизительное значение для средней длины месяца
  const years = Math.floor(days / 365.25); // Приблизительное значение для средней длины года

  if (years > 0) {
    return `${years} year${years > 1 ? "s" : ""} ago`;
  } else if (months > 0) {
    return `${months} month${months > 1 ? "s" : ""} ago`;
  } else if (weeks > 0) {
    return `${weeks} week${weeks > 1 ? "s" : ""} ago`;
  } else if (days > 0) {
    return `${days} day${days > 1 ? "s" : ""} ago`;
  } else if (hours > 0) {
    return `${hours} hour${hours > 1 ? "s" : ""} ago`;
  } else if (minutes > 0) {
    return `${minutes} minute${minutes > 1 ? "s" : ""} ago`;
  } else {
    return `just now`;
  }
}

const Youtube = (prop: Props) => {
  const [isHover, setIsHover] = useState<boolean>(false);
  const [poster, setPoster] = useState<string>("");
  const [loaded, setLoaded] = useState<boolean>(false);
  const imgRef = useRef(null);

  useEffect(() => {
    const cleanup = lazyService.createImageObserver(
      imgRef,
      prop.videos.snippet.thumbnails.medium.url,
      setPoster,
      setLoaded
    );

    return cleanup;
  }, [prop.videos.snippet.thumbnails.medium.url]);

  return (
    <>
      <div className="max-w-[20rem] my-3 mx-2">
        <Link
          to="/"
          className="rounded-xl"
          onMouseEnter={() => setIsHover(true)}
          onMouseLeave={() => setIsHover(false)}
        >
          <div ref={imgRef} className="min-w-[320px] min-h-[180px]">
            {loaded ? (
              <img
                src={poster}
                className={`${
                  isHover ? "scale-[0.9]" : ""
                } duration-100 rounded max-w-[20rem] max-h-[13rem] min-w-[320px] min-h-[180px] object-cover`}
              />
            ) : (
              <Skeleton width={20} height={11.4} />
            )}
          </div>
          <p className="text-[1.2rem] text-[#fff] mt-[1rem] mb-[0.2rem]">
            {prop.videos.snippet.title}
          </p>
          <p className="text-[#999] text-[1rem] font-normal tracking-wide">
            {prop.videos.snippet.channelTitle}{" "}
            <span className="ml-[0.5rem]">
              {timeSince(prop.videos.snippet.publishedAt)}
            </span>
          </p>
        </Link>
      </div>
    </>
  );
};

export default Youtube;
