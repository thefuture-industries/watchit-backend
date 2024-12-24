import { X } from "lucide-react";
import app from "~/core/app";

interface Props {
  errorMessage: string;
}

const Error = (prop: Props) => {
  return (
    <>
      <div className="bg-[#0009] fixed top-0 left-0 h-screen w-screen z-[100]">
        <div className="flex fixed top-[44%] w-screen justify-center">
          <div className="bg-[#111] max-w-[40vw] p-3 border border-[#333] rounded ">
            <div className="flex items-start">
              <div className="bg-[#900000] rounded-3xl p-2 mr-3">
                <X size={28} />
              </div>
              <p className="text-[0.9rem] tracking-wide">{prop.errorMessage}</p>
            </div>
            <div className="flex items-center justify-end mt-4">
              <div
                className="border border-[#900000] mr-3 cursor-pointer hover:bg-[#900000] transition inline-block px-7 rounded-2xl"
                onClick={() => app.restart()}
              >
                <p className="text-[0.9rem]">Restart</p>
              </div>
              <div
                className="border border-[#900000] cursor-pointer hover:bg-[#900000] transition inline-block px-7 rounded-2xl"
                onClick={() => app.exit()}
              >
                <p className="text-[0.9rem]">Exist</p>
              </div>
              <div className="border border-[#900000] cursor-pointer hover:bg-[#900000] transition inline-block px-7 rounded-2xl ml-4">
                <p className="text-[0.9rem]">OK</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Error;
