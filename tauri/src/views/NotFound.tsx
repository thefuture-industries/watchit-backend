import { ChevronLeft } from "lucide-react";
import { useNavigate } from "react-router-dom";
import "~/../sass/404.sass";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <>
      <div className="flex h-screen justify-center items-center">
        <div className="text-center">
          <p className="text-center tracking-wide text-[#888] cursor-pointer">
            Sorry, the page you requested was not found in this app.
          </p>
          <p className="text-[24rem] font-semibold mb-[12rem] mt-[10rem] cursor-pointer">
            404
          </p>
          <div
            className="flex items-center justify-center cursor-pointer hover:text-[#ff0000] transition"
            onClick={() => {
              navigate(-1);
            }}
          >
            <ChevronLeft size={20} strokeWidth={2.5} className="pt-1" />
            <p className="tracking-wide font-medium">Go back to home page</p>
          </div>
        </div>
      </div>
    </>
  );
};

export default NotFound;
