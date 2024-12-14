import axios from "axios";
import { KeyRound } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import userService from "~/services/user-service";
import { UserAddPayload } from "~/types/user";

const LoadStart = () => {
  const navigate = useNavigate();
  const [isHoverBtn, setIsHoverBtn] = useState<boolean>(false);
  const [secretWord, setSecretWord] = useState<string>("");
  const [isSend, setIsSend] = useState<boolean>(false);

  return (
    <>
      <div>
        <div className="fixed">
          <video
            src="/src/assets/video-bg-sf09wjq.mp4"
            autoPlay
            muted
            loop
            className="w-screen h-screen object-cover"
          ></video>
        </div>
        <div className="z-[100] relative h-screen flex flex-col justify-center items-center px-[3rem]">
          <div
            className="p-5 py-6 rounded max-w-[37vw]"
            style={{ background: "rgba(0, 0, 0, 0.6)" }}
          >
            <div className="flex justify-center items-center mb-8">
              <img
                src="/src/assets/flicksfi_ico.png"
                className="max-w-[3rem] opacity-[0.3]"
                alt=""
              />
              <p className="opacity-[0.3] ml-3 tracking-wide text-[1.5rem]">
                FlicksFI
              </p>
            </div>
            <p className="text-[1.8rem] mb-5">The secret word</p>
            <p className="text-[#fff] font-light mb-2 tracking-wide">
              If you are not registered, enter your secret word to log in or
              register in the system. The secret word can contain any symbols
              and meaning, you have no restrictions, but remember it is your
              safety!
            </p>
            <div className="relative">
              <div>
                <input
                  type="password"
                  className="relative p-3 pl-[3rem] text-[1.1rem] border-2 border-[#333] my-2 w-full tracking-[0.4rem] font-semibold focus:border-[blue]"
                  style={{ background: "rgba(0, 0, 0, 0.4)" }}
                  value={secretWord}
                  onChange={(e) => setSecretWord(e.target.value)}
                />
              </div>
              <div className="absolute top-[1.5rem] left-4">
                <KeyRound size={20} />
              </div>
            </div>
            <div
              onMouseEnter={() => setIsHoverBtn(true)}
              onMouseLeave={() => setIsHoverBtn(false)}
              className="bg-[#2413ff] p-3 rounded-lg flex items-center justify-center mt-3 cursor-pointer font-light duration-150"
              onClick={async () => {
                const controller = new AbortController();
                const timeoutId = setTimeout(() => controller.abort(), 10000);

                try {
                  setIsSend(true);
                  const response = await axios.get("http://ip-api.com/json", {
                    signal: controller.signal,
                  });

                  const data: UserAddPayload = {
                    secret_word: secretWord,
                    ip_address: response.data.query || "255.255.255.255",
                    latitude: response.data.lat.toString() || "0",
                    longitude: response.data.lon.toString() || "0",
                    country: response.data.country || "NONE",
                    region_name: response.data.regionName || "NONE",
                    zip: response.data.zip || "000000",
                  };

                  const createdUser = await userService.add_user(data);

                  if (
                    typeof createdUser == "string" &&
                    createdUser.startsWith("ERROR")
                  ) {
                    setIsSend(false);
                  } else {
                    sessionStorage.setItem(
                      "_sess",
                      JSON.stringify(createdUser)
                    );
                    navigate("/");
                  }
                } catch (error) {
                  if (axios.isCancel(error)) {
                    const data: UserAddPayload = {
                      secret_word: secretWord,
                      ip_address: "255.255.255.255",
                      latitude: "0",
                      longitude: "0",
                      country: "NONE",
                      region_name: "NONE",
                      zip: "000000",
                    };

                    const createdUser = await userService.add_user(data);

                    if (
                      typeof createdUser == "string" &&
                      createdUser.startsWith("ERROR")
                    ) {
                      setIsSend(false);
                    } else {
                      sessionStorage.setItem(
                        "_sess",
                        JSON.stringify(createdUser)
                      );
                      navigate("/");
                    }
                  }
                } finally {
                  setIsSend(false);
                  clearTimeout(timeoutId);
                }
              }}
              style={{
                boxShadow: isHoverBtn
                  ? "rgba(0, 53, 255, 0.9) 0px 0px 50px 0px"
                  : "rgb(0 53 255 / 90%) 0px 0px 16px 0px",
              }}
            >
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
                <p>Save and verify</p>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default LoadStart;
