import { useEffect, useState } from "react";
import Navigation from "~/components/Navigation";

const User = () => {
  const [username, setUsername] = useState<string>("");

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
                <input type="text" className="w-full" />
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
            </div>
            <div className="mt-[10px]">
              <p>Profile picture</p>
              <img
                src="/src/assets/gradient.webp"
                alt=""
                width="230rem"
                className="rounded-[52%]"
              />
              <div className="bg-[#111] border border-[#222] text-center p-2 rounded mt-[1rem] cursor-pointer">
                <p>Сhange</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default User;
