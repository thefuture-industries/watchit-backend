import { useEffect, useState } from "react";
import Navigation from "~/components/Navigation";
import userService from "~/services/user-service";

const User = () => {
  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [secret_word, setSecretWord] = useState<string>("");

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
                  type="text"
                  className="w-full"
                  value={email || user?.email || ""}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <p className="mb-[9px] mt-[12px]">Nickname</p>
                <input
                  type="text"
                  className="w-full"
                  value={username || user?.username}
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
                className="bg-[#111] w-full hover:bg-[#2413ff] transition duration-150 border border-[#222] text-center p-2 rounded mt-[1rem] cursor-pointer"
                onClick={async () => {
                  const response = await userService.update_user({
                    username: username,
                    email: email,
                    secret_word: secret_word,
                  });

                  sessionStorage.setItem("_sess", JSON.stringify(response));
                }}
              >
                Сhange
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default User;
