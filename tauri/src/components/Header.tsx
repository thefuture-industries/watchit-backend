import { LogOut } from "lucide-react";
import { useState } from "react";
import app from "~/core/app";

const Header: React.FC<{ onSearch: (query: string) => Promise<void> }> = ({
  onSearch,
}) => {
  const [searchInput, setSearchInput] = useState<string>("");

  const handleKeyDown = async (event: any) => {
    if (event.key === "Enter") {
      await onSearch(searchInput);
    }
  };

  return (
    <>
      <div>
        <div className="flex items-center justify-between">
          <div className="left">
            <input
              className="w-[35vw]"
              placeholder="Search movies"
              onKeyDown={handleKeyDown}
              onChange={(e) => setSearchInput(e.target.value)}
              value={searchInput}
            />
          </div>
          <div className="right">
            <div
              className="flex items-center cursor-pointer bg-[#111] hover:bg-[#222] border border-[#222] py-2 px-3 rounded"
              onClick={() => app.exit()}
            >
              <LogOut size={18} strokeWidth={2.5} className="mt-[1.3px]" />
              <p className="tracking-wide ml-2">Exit</p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Header;
