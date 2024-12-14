import {
  Clapperboard,
  LayoutTemplate,
  Server,
  Youtube,
  Text,
  Heart,
} from "lucide-react";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const itemNav = [
  {
    title: "Home",
    link: "/",
    icon: <LayoutTemplate size={22} />,
  },
  {
    title: "Favourites",
    link: "/favourites",
    icon: <Heart size={22} />,
  },
  {
    title: "API",
    link: "/api",
    icon: <Server size={22} />,
  },
  {
    title: "Youtube",
    link: "/youtube/filter",
    icon: <Youtube size={22} />,
  },
  {
    title: "Movies",
    link: "/movie/filter",
    icon: <Clapperboard size={22} />,
  },
  {
    title: "Story",
    link: "/story",
    icon: <Text size={22} />,
  },
];

const Navigation = () => {
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
      <div className="fixed">
        <div className="w-[18rem] h-screen bg-[#111] rounded-xl p-3 border border-[#222]">
          <Link to="/" className="flex items-center">
            <img src="/src/assets/gradient.webp" width={45} alt="" />
            <div className="ml-3">
              <p className="text-[#fff]">{user?.username}</p>
              <p className="font-normal text-[14px] text-[#888]">
                {user?.email ? (
                  <p>{user.email}</p>
                ) : (
                  <p>Specify the email address</p>
                )}
              </p>
            </div>
          </Link>

          <div className="mt-6">
            {itemNav.map((item, index) => (
              <Link
                to={item.link}
                key={index}
                className={`flex text-[#fff] hover:text-[#fff] items-center cursor-pointer mt-2 p-2 rounded bg-[${
                  location.pathname == item.link ? "#222" : "transparent"
                }] hover:bg-[#222] transition`}
              >
                <div className="mr-[5px]">{item.icon}</div>
                <p className="ml-1 tracking-wide">{item.title}</p>
              </Link>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default Navigation;
