import { Clapperboard, MoveLeft, Youtube, Text, Monitor } from "lucide-react";
import Skeleton from "./Skeleton";

const MovieDetailsSkeleton = () => {
  return (
    <>
      <div className="my-5 ml-[5rem]">
        <div className="flex items-center justify-between">
          <MoveLeft
            size={27}
            className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
          />
          <div className="flex items-center gap-[1rem] mr-[4rem]">
            <Youtube
              size={22}
              className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
            />
            <Clapperboard
              size={22}
              className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
            />
            <Text
              size={22}
              className="cursor-pointer text-[#4b4b4b] hover:text-[#fff] transition"
            />
          </div>
        </div>
        <div className="flex mt-7 items-center">
          {/* POSTER */}
          <div>
            <Skeleton width={27} height={35} />
          </div>
          <div className="ml-[4rem]">
            {/* TITLE */}
            <p className="h-3 w-full bg-[#444] rounded animate-pulse"></p>

            {/* MINI DEATILS */}
            <div className="flex items-center mt-2">
              <p className="w-[3rem] h-2 bg-[#444] rounded animate-pulse"></p>
              <span className="mx-2 text-[#888]">|</span>
              <p className="w-[3rem] h-2 bg-[#444] rounded animate-pulse"></p>
            </div>

            {/* OVERVIEW */}
            <p className="mt-5 max-w-[30vw] gap-x-[0.5rem] gap-y-[0.5rem] flex flex-wrap items-center">
              <div className="w-[10rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[4rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[2rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[5rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[2rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[9rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[10rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[7rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[4rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[14rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[8rem] h-2 bg-[#444] rounded animate-pulse"></div>
              <div className="w-[7rem] h-2 bg-[#444] rounded animate-pulse"></div>
            </p>

            {/* RANGE */}
            <div className="mt-5">
              <p className="mt-2 flex items-center">
                <span className="text-[#777]">Vote Average</span>
                <div className="ml-[3.5rem] h-2 w-[2rem] bg-[#444] rounded animate-pulse"></div>
              </p>
              <p className="mt-2 flex items-center">
                <span className="text-[#777]">Genre</span>
                <div className="ml-[6rem] h-2 w-[5rem] bg-[#444] rounded animate-pulse"></div>
              </p>
              <p className="mt-2 flex items-center">
                <span className="text-[#777]">Language</span>
                <div className="ml-[4.8rem] h-2 w-[4rem] bg-[#444] rounded animate-pulse"></div>
              </p>
            </div>

            <div className="flex items-center mt-10 gap-[2rem]">
              <div className="cursor-pointer bg-[#b7b7b7] hover:bg-[#fff] transition py-2 px-9 rounded inline-block font-semibold">
                <p className="uppercase text-[#000]">trailer</p>
              </div>
              <div className="text-[#888] hover:text-[#fff] transition flex items-center gap-[0.8rem] cursor-pointer">
                <Monitor strokeWidth={1.5} />
                <p className="uppercase">watch</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default MovieDetailsSkeleton;
