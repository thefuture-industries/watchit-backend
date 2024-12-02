import "~/App.css";
import { createBrowserRouter } from "react-router-dom";
import { lazy, Suspense } from "react";
import * as page from "~/pages/export.page.config";
import ProtectedRouter from "./Protected";
const Loader = lazy(() => import("./components/Loader"));
const MovieDetailsSkeleton = lazy(
  () => import("./components/Skeletons/MovieDetailsSkeleton")
);

const App = createBrowserRouter([
  {
    path: "/",
    children: [
      {
        path: "",
        element: (
          <Suspense fallback={<Loader />}>
            <ProtectedRouter>
              <page.home />
            </ProtectedRouter>
          </Suspense>
        ),
      },
      {
        path: "/load",
        element: <page.load />,
      },
      {
        path: "movie/filter",
        element: (
          <Suspense fallback={<Loader />}>
            <page.movie_filter />
          </Suspense>
        ),
      },
      {
        path: "youtube/filter",
        element: (
          <Suspense fallback={<Loader />}>
            <page.youtube_filter />
          </Suspense>
        ),
      },
      {
        path: "story",
        element: (
          <Suspense fallback={<Loader />}>
            <page.story />
          </Suspense>
        ),
      },
      {
        path: "search/result",
        element: (
          <Suspense fallback={<Loader />}>
            <page.search_result />
          </Suspense>
        ),
      },
      {
        path: "movie/:id",
        element: (
          <Suspense fallback={<MovieDetailsSkeleton />}>
            <page.movie_details />
          </Suspense>
        ),
      },
      {
        path: "*",
        element: (
          <Suspense fallback={<Loader />}>
            <page.not_found />
          </Suspense>
        ),
      },
    ],
  },
]);

export default App;
