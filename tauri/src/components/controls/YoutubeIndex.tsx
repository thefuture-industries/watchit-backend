import { Play } from "lucide-react";
import { useEffect, useState } from "react";
import { DefaultYoutube, YoutubeModel } from "~/types/youtube";

interface Props {
  video: YoutubeModel[] | [];
}

const YoutubeIndex = (prop: Props) => {
  const popular_data = ["🔥 Popular", "🆕 News", "🔝 Top", "✨ Famous"];
  const [popular, setPopular] = useState<string>("");
  const [video, setVideo] = useState<YoutubeModel[]>([]);

  useEffect(() => {
    if (prop.video.length == 0) {
      setVideo([DefaultYoutube]);
    } else {
      setVideo(prop.video);
    }

    setPopular(popular_data[Math.floor(Math.random() * 4)]);
  }, [prop.video]);

  return (
    <>
      <div>
        <div className="relative">
          <img
            src={`${video[0]?.snippet?.thumbnails.high.url}`}
            style={{ opacity: 0.6 }}
            className="object-cover w-full max-h-[350px] h-[50vh] rounded-xl mt-4"
            alt=""
          />

          <div
            style={{ background: "rgba(0, 0, 0, 0.6)" }}
            className="absolute top-5 left-5 cursor-pointer text-[14px] inline-block py-[0.5px] px-3 rounded-xl"
          >
            <p className="font-light tracking-wide px-[0.3rem] pr-[0.8rem] pb-[3px]">
              {popular}
            </p>
          </div>

          <div>
            <div className="absolute bottom-5 left-6">
              <p className="max-w-[50vw] text-[2rem] text-balance font-medium leading-[3rem] my-2">
                {video[0]?.snippet?.title}
              </p>
              <p className="max-w-[50vw] text-[#e6e6e6] text-balance tracking-wide text-[13px]">
                {video[0]?.snippet?.description}
              </p>

              <div className="cursor-pointer bg-[#fff] hover:bg-[#999] transition py-[5px] px-4 mt-3 rounded-2xl inline-block motion-preset-confetti">
                <div className="flex items-center">
                  <Play fill="#000" size={21} />
                  <p className="pr-2 text-[#000] text-[14px] ml-2 font-medium -mt-[1px]">
                    Watch
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default YoutubeIndex;
