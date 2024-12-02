import { useState } from "react";
import Navigation from "~/components/Navigation";
import { YoutubeModel } from "~/types/youtube";
import SearchResult from "./SearchResult";
import youtubeService from "~/services/youtube-service";

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
  const [searchInput, setSearchInput] = useState<string>("");
  const [category, setCategory] = useState<string>("");
  const [top, setTop] = useState<string>("");
  const [searchPage, setSearchPage] = useState<boolean>(false);
  const [videos, setVideos] = useState<YoutubeModel[]>([]);

  return (
    <>
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
                  await youtubeService
                    .get_youtube_videos(searchInput, category, top)
                    .then((videos) => {
                      setVideos(videos);
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

export default YoutubeFilter;
