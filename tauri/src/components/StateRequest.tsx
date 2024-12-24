import { CircleCheck, CircleX, Info, X } from "lucide-react";
import { useEffect } from "react";

interface Props {
  statusCode: number;
  message: string;
  state: any;
  setState: any;
}

const StateRequest = (prop: Props) => {
  const statusCodeString = prop.statusCode.toString();
  console.log(statusCodeString);

  useEffect(() => {
    setTimeout(() => {
      prop.setState(false);
    }, 3000);
  }, [prop.state]);

  return (
    <>
      <div
        className="fixed z-[100] cursor-pointer"
        style={{ left: "50%", top: "1rem", transform: "translateX(-50%)" }}
      >
        <div
          className={`px-[2rem] py-[0.5rem] rounded-xl`}
          style={{
            background: `${
              statusCodeString.startsWith("2")
                ? "#006e1ce0"
                : statusCodeString.startsWith("3")
                ? "#6e4d00e0"
                : "#6e0000e0"
            }`,
          }}
        >
          <div className="flex items-center">
            {statusCodeString.startsWith("2") ? (
              <CircleCheck size={18} className="mr-2" color="#52ff47" />
            ) : statusCodeString.startsWith("3") ? (
              <Info size={18} className="mr-2" color="#fbc23ee0" />
            ) : (
              <CircleX size={18} className="mr-2" color="#ff4242" />
            )}
            <p
              className="text-[#ff4242] tracking-wide font-normal text-[0.9rem]"
              style={{
                fontFamily: "Inter",
                color: `${
                  statusCodeString.startsWith("2")
                    ? "#52ff47"
                    : statusCodeString.startsWith("3")
                    ? "#fbc23ee0"
                    : "#ff4242"
                }`,
              }}
            >
              {prop.message}
            </p>
          </div>
        </div>
      </div>
    </>
  );
};

export default StateRequest;
