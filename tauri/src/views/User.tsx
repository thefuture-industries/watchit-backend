import { useEffect, useState } from "react";
import Error from "~/components/Error";
import Navigation from "~/components/Navigation";
import userService from "~/services/user-service";

const User = () => {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [secret_word, setSecretWord] = useState<string>("");
  const [isSend, setIsSend] = useState<boolean>(false);
  const [isSuccess, setIsSuccess] = useState<boolean>(false);

  const [user, setUser] = useState<{
    uuid: string;
    username: string;
    email: string | null;
  }>();

  useEffect(() => {
    const _sess = sessionStorage.getItem("_sess");

    if (!_sess) {
      return;
    }

    setUser(JSON.parse(_sess));
    setUsername(JSON.parse(_sess).username);
    setEmail(JSON.parse(_sess).email);
  }, []);

  return (
    <>
      <div className="container flex w-screen m-2">
        <div className="left">
          <Navigation />
        </div>
        <div className="right ml-[19rem] w-[67vw]">
          <div className="flex justify-between">
            <div className="max-w-[25rem]">
              <p className="text-[1.5rem] pb-[1rem] pt-[10px]">
                User Information
              </p>
              <p className="text-[#999] max-w-[30rem] mb-[1rem]">
                Here you can edit public information about yourself. The changes
                will be displayed for other users within 5 minutes.
              </p>
              <div>
                <p className="mb-[9px]">Email address</p>
                <input
                  type="email"
                  className="w-full"
                  value={email || ""}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <p className="mb-[9px] mt-[12px]">Nickname</p>
                <input
                  type="text"
                  className="w-full"
                  value={username || ""}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </div>
              <div>
                <p className="mt-[1rem] mb-[0.7rem]">Secret word</p>
                <div className="flex gap-[0.7rem]">
                  <input
                    type="password"
                    className="w-full tracking-wide"
                    placeholder="The current secret word"
                  />
                  <input
                    type="password"
                    className="w-full tracking-wide"
                    placeholder="Your new secret word"
                    value={secret_word}
                    onChange={(e) => setSecretWord(e.target.value)}
                  />
                </div>
              </div>
            </div>
            <div className="mt-[10px]">
              <p>Profile picture</p>
              <img
                src="/src/assets/gradient.webp"
                alt=""
                width="230rem"
                className="rounded-[52%]"
              />
              <button
                className={`bg-[#111] flex justify-center items-center w-full hover:bg-[#2413ff] transition duration-150 border border-[#222] text-center p-2 rounded mt-[1rem] cursor-pointer ${
                  isSuccess ? "motion-preset-confetti" : ""
                }`}
                onClick={async () => {
                  setIsSend(true);
                  try {
                    const response = await userService.update_user({
                      username: username,
                      email: email,
                      secret_word: secret_word,
                    });
                    sessionStorage.setItem("_sess", JSON.stringify(response));
                    setIsSuccess(true);
                    setTimeout(() => {
                      setIsSuccess(false);
                    }, 4000);
                  } catch {
                  } finally {
                    setIsSend(false);
                  }
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
                  "Сhange"
                )}
              </button>
              {isSuccess && (
                <p className="flex justify-center text-[#10721a] mt-2 font-normal tracking-wide">
                  The data has been updated
                </p>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default User;
