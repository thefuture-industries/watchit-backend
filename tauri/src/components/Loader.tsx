import "~/../sass/Loader.sass";

const itemNav = [
  {
    title: "Home",
  },
  {
    title: "Favourites",
  },
  {
    title: "API",
  },
  {
    title: "Youtube",
  },
  {
    title: "Movies",
  },
  {
    title: "Story",
  },
];

const Loader = () => {
  return (
    <>
      <div className="w-screen h-screen m-2">
        <div className="fixed">
          <div className="w-[18rem] h-screen bg-[#111] rounded-xl p-3 border border-[#222]">
            <div className="flex items-center">
              <img
                src="/src/assets/gradient.png"
                className="animate-pulse"
                width={45}
                alt=""
              />
              <div className="ml-3">
                <div className="w-[120px] h-2 bg-[#555] rounded mb-3 animate-pulse"></div>
                <div className="w-[12rem] h-2 bg-[#555] rounded animate-pulse"></div>
              </div>
            </div>

            <div className="mt-6">
              {itemNav.map((_, index: number) => (
                <div
                  key={index}
                  className="flex text-[#fff] items-center mt-3 p-2 rounded bg-[#222]"
                >
                  <div className="mr-[5px] bg-[#555] w-[30px] h-6 rounded animate-pulse"></div>
                  <p className="ml-1 tracking-wide w-full h-2 bg-[#555] rounded animate-pulse"></p>
                </div>
              ))}
            </div>
          </div>
        </div>
        <div id="loader">
          <div>
            <div id="loading"></div>
            <p className="wait-text">Wait... we're uploading</p>
          </div>
        </div>
      </div>
    </>
  );
};

export default Loader;
