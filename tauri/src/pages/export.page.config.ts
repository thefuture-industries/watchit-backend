import { lazy } from "react";
const home = lazy(() => import("./home"));
const story = lazy(() => import("./story"));
const movie_filter = lazy(() => import("./movie"));
const youtube_filter = lazy(() => import("./youtube"));
const movie_details = lazy(() => import("./movie_details"));
const favourites = lazy(() => import("./favourites"));
const load = lazy(() => import("./load"));
const not_found = lazy(() => import("./not_found"));
const user = lazy(() => import("./user"));

export {
  home,
  movie_filter,
  youtube_filter,
  story,
  movie_details,
  favourites,
  load,
  not_found,
  user,
};
