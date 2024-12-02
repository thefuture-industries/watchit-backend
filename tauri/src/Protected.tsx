import React from "react";
import { Navigate, useLocation } from "react-router-dom";

type Props = { children: React.ReactNode };

const ProtectedRouter = ({ children }: Props) => {
  const location = useLocation();
  const sess = sessionStorage.getItem("_sess");

  return sess != null ? (
    <> {children}</>
  ) : (
    <Navigate to="/load" state={{ from: location }} />
  );
};

export default ProtectedRouter;
