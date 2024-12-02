import { TriangleAlert } from "lucide-react";
import app from "~/core/app";

interface Props {
  errorMessage: string;
}

const Warning = (prop: Props) => {
  return (
    <>
      <div className="bg-[#0009] fixed top-0 left-0 h-screen w-screen z-[100]">
        <div className="flex fixed top-[44%] w-screen justify-center">
          <div className="bg-[#111] max-w-[40vw] p-3 border border-[#333] rounded ">
            <div className="flex items-start">
              <div className="bg-[#8f9000] rounded-3xl p-2 pt-[6px] mr-3 flex items-center justify-center">
                <TriangleAlert size={28} color="#dfff00" />
              </div>
              <p className="text-[0.9rem] tracking-wide">{prop.errorMessage}</p>
            </div>
            <div className="flex items-center justify-end mt-4">
              <div
                className="border border-[#8e8200] mr-3 cursor-pointer hover:bg-[#8e8200] transition inline-block px-7 rounded-2xl"
                onClick={() => window.location.reload()}
              >
                <p className="text-[0.9rem]">Ok</p>
              </div>
              <div
                className="border border-[#8e8200] cursor-pointer hover:bg-[#8e8200] transition inline-block px-7 rounded-2xl"
                onClick={() => app.exit()}
              >
                <p className="text-[0.9rem]">Exist</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Warning;
