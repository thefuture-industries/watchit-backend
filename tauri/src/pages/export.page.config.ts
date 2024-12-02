import { lazy } from "react";
const home = lazy(() => import("./home"));
const story = lazy(() => import("./story"));
const movie_filter = lazy(() => import("./movie"));
const youtube_filter = lazy(() => import("./youtube"));
const search_result = lazy(() => import("./search_result"));
const movie_details = lazy(() => import("./movie_details"));
const load = lazy(() => import("./load"));
const not_found = lazy(() => import("./not_found"));

export {
  home,
  movie_filter,
  youtube_filter,
  story,
  search_result,
  movie_details,
  load,
  not_found,
};
