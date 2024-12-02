import axios from "axios";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import userService from "~/services/user-service";
import { UserModel } from "~/types/user";

const LoadStart = () => {
  const navigate = useNavigate();
  useEffect(() => {
    const getUserData = async () => {
      try {
        const response = await axios.get("http://ip-api.com/json");

        const data: UserModel = {
          id: null,
          username: null,
          username_upper: null,
          email: null,
          email_upper: null,
          ip_address: response.data.query,
          latitude: response.data.lat.toString(),
          longitude: response.data.lon.toString(),
          country: response.data.country,
          region_name: response.data.regionName,
          zip: response.data.zip,
          created_at: null,
        };

        const createdUser = await userService.add_user(data);
        if (createdUser.startsWith("Error Error HTTP: 400 Bad Request")) {
          console.log("You are already registered");
        } else {
          sessionStorage.setItem("_sess", createdUser);
          navigate("/");
        }
      } catch (err) {
        return;
      }
    };

    getUserData();
  }, []);

  return (
    <>
      <div>
        <div className="fixed">
          <img
            src="/src/assets/bg.jpeg"
            className="w-screen h-screen object-cover opacity-[0.5]"
            alt=""
          />
        </div>
        <div className="z-[100] relative h-screen flex flex-col justify-center px-[3rem]">
          <div>
            <p className="text-[9vw] font-medium">Flick Finder</p>
            <p className="text-[2.6vw] text-[#999] mt-[5vw] ml-1">
              Find it quickly and start watching
            </p>
          </div>
        </div>
      </div>
    </>
  );
};

export default LoadStart;
