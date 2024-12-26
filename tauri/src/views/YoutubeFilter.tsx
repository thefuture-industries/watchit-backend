import { useState } from "react";
import Navigation from "~/components/Navigation";
import { YoutubeModel } from "~/types/youtube";
import SearchResult from "./SearchResult";
import youtubeService from "~/services/youtube-service";
import StateRequest from "~/components/StateRequest";

const categorys = [
  "Film",
  "Animation",
  "Autos",
  "Animals",
  "Sports",
  "Events",
  "Travel",
  "Gaming",
  "Blogs",
  "Howto",
  "Science",
  "Technology",
  "Education",
  "Movie and TV series trailers",
];

const tops = [
  "Programming training",
  "Foreign languages",
  "Finance and investment",
  "News world",
  "Delay search",
  "Auto and stophams",
  "Meal",
];

const YoutubeFilter = () => {
  const [isError, setIsError] = useState<boolean>(false);
  const [error, setError] = useState<string>("");
  const [isSend, setIsSend] = useState<boolean>(false);
  const [searchInput, setSearchInput] = useState<string>("");
  const [category, setCategory] = useState<string>("");
  const [top, setTop] = useState<string>("");
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [videos, setVideos] = useState<YoutubeModel[]>([]);

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

      {searchPage ? (
        <SearchResult movies={null} videos={videos} />
      ) : (
        <div className="container flex w-screen m-2">
          <div className="left">
            <Navigation />
          </div>
          <div className="right ml-[19rem] w-[67vw]">
            <div className="bg-[#111] border border-[#222] p-5 rounded-lg">
              <p className="text-[1.5rem] tracking-wide">YouTube Filter</p>
              <div className="w-full bg-[#222] h-[1px] my-3 mt-5"></div>

              {/* Search */}
              <textarea
                className="w-full mb-4 bg-[transparent] border border-[#222] rounded h-[7rem] text-[1.3rem] p-3 resize-none outline-none"
                placeholder="Search"
                onChange={(e) => setSearchInput(e.target.value)}
                value={searchInput}
              ></textarea>

              {/* Category */}
              <p className="text-[1rem] tracking-wide">Category</p>
              <div className="w-full bg-[#222] h-[1px] my-3"></div>
              <div>
                {categorys.map((item, index) => (
                  <div
                    key={index}
                    className={`bg-[${category == item ? "#fff" : ""}] text-[${
                      category == item ? "#000" : "#fff"
                    }] inline-block border border-[#555] py-[1px] pb-1 px-3 rounded-3xl m-1`}
                  >
                    <div
                      className="flex items-center cursor-pointer"
                      onClick={() => setCategory(item)}
                    >
                      <div
                        className={`w-[15px] h-[15px] rounded-3xl border-2 border-[${
                          category == item ? "#000" : "#fff"
                        }] mr-[8px]`}
                      ></div>
                      <p>{item}</p>
                    </div>
                  </div>
                ))}
              </div>

              {/* TOPS */}
              <p className="text-[1rem] tracking-wide mt-3">TOPs</p>
              <div className="w-full bg-[#222] h-[1px] my-3"></div>
              <div>
                {tops.map((item, index) => (
                  <div
                    key={index}
                    className={`bg-[${top == item ? "#fff" : ""}] text-[${
                      top == item ? "#000" : "#fff"
                    }] inline-block border border-[#555] py-[1px] pb-1 px-3 rounded-3xl m-1`}
                  >
                    <div
                      className="flex items-center cursor-pointer"
                      onClick={() => setTop(item)}
                    >
                      <div
                        className={`w-[15px] h-[15px] rounded-3xl border-2 border-[${
                          top == item ? "#000" : "#fff"
                        }] mr-[8px]`}
                      ></div>
                      <p>{item}</p>
                    </div>
                  </div>
                ))}
              </div>

              <div
                className="mt-6 bg-[#ff2400] hover:bg-[#b21900] transition flex justify-center items-center min-h-[47px] rounded-lg cursor-pointer"
                style={{
                  boxShadow: "inset 0px -7px 0px 0px rgba(0, 0, 0, 0.4)",
                }}
                onClick={async () => {
                  setIsSend(true);
                  await youtubeService
                    .get_youtube_videos(searchInput, category, top)
                    .then((videos) => {
                      setVideos(videos);
                      setSearchPage(true);
                    })
                    .catch((error) => {
                      setIsError(true);
                      setError(error.error as string);
                    })
                    .finally(() => {
                      setIsSend(false);
                    });
                }}
              >
                <p className="uppercase text-[17px] -mt-1">
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
                    "search"
                  )}
                </p>
              </div>
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default YoutubeFilter;
