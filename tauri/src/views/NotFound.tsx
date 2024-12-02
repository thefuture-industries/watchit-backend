import { useNavigate } from "react-router-dom";
import "~/../sass/404.sass";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <>
      <div className="flex h-screen justify-center items-center">
        <div className="text-center">
          <h1 className="text-[6vw] uppercase" id="text-404">
            404 Not Found
          </h1>
          <p className="mt-[3.7vw] cursor-pointer hover:scale-[1.3] transition" onClick={() => navigate(-1)}>
            Back
          </p>
        </div>
      </div>
    </>
  );
};

export default NotFound;
