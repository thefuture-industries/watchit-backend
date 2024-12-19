import { ExternalLink, LogOut, Search } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import app from "~/core/app";

const Header: React.FC<{ onSearch: (query: string) => Promise<void> }> = ({
  onSearch,
}) => {
  const [searchInput, setSearchInput] = useState<string>("");
  const [isOpenHistory, setIsOpenHistory] = useState<boolean>(false);
  const [historyElements, setHistoryElements] = useState<string[]>([]);
  const historyRef = useRef(null);

  const handleKeyDown = async (event: any) => {
    if (event.key === "Enter") {
      const history = JSON.parse(
        sessionStorage.getItem("sess_history__search") || "[]"
      );

      if (!history.includes(searchInput)) {
        history.push(searchInput);
      }

      sessionStorage.setItem("sess_history__search", JSON.stringify(history));
      await onSearch(searchInput);
    }
  };

  useEffect(() => {
    const history = sessionStorage.getItem("sess_history__search");

    if (!history) {
      setIsOpenHistory(false);
    }
  }, [isOpenHistory]);

  useEffect(() => {
    setHistoryElements(
      JSON.parse(sessionStorage.getItem("sess_history__search") || "[]")
    );
  }, []);

  return (
    <>
      <div>
        <div className="flex items-center justify-between">
          <div className="left">
            <div className="relative">
              <input
                className="w-[35vw] h-[2.3rem] pl-9"
                placeholder="Search movies"
                onKeyDown={handleKeyDown}
                onChange={(e) => setSearchInput(e.target.value)}
                value={searchInput}
                onFocus={() => setIsOpenHistory(true)}
                onBlur={() => setIsOpenHistory(false)}
              />
              <Search
                className="absolute top-[0.55rem] left-[0.6rem] text-[#999]"
                size={19}
              />
            </div>
          </div>
          <div className="right">
            <div
              className="flex items-center cursor-pointer bg-[#111] hover:bg-[#222] border border-[#222] py-2 px-3 h-[2.3rem] rounded"
              onClick={() => app.exit()}
            >
              <LogOut size={18} strokeWidth={2.5} className="mt-[1.3px]" />
              <p className="tracking-wide ml-2">Exit</p>
            </div>
          </div>
        </div>

        {/* HISTORY */}
        <div
          ref={historyRef}
          className={`p-3 ${
            isOpenHistory ? "" : "hidden"
          } absolute z-[100] w-[35vw] rounded-lg border mt-2 border-[#333]`}
          style={{ background: "rgba(0, 0, 0, 0.9)" }}
        >
          <div>
            {historyElements.map((el, index) => (
              <div
                key={index}
                className="bg-[transparent] hover:bg-[#111] transition rounded py-[0.4rem] px-[0.6rem] flex items-center cursor-pointer"
                onClick={async () => {
                  await onSearch(el);
                  console.log("CLICK" + el);
                }}
              >
                <ExternalLink size={18} color="#666" strokeWidth={1.2} />
                <p className="text-[#999] ml-4 text-[0.95rem] font-light tracking-wide">
                  {el}
                </p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default Header;
