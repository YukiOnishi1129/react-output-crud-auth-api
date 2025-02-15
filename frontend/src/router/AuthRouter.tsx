import { NAVIGATION_LIST } from "../constants/navigation";
import { Routes, Route } from "react-router";
import { LoginPage } from "../pages";

export const AuthRouter = () => {
  return (
    <Routes>
      <Route index path={NAVIGATION_LIST.LOGIN} element={<LoginPage />} />
    </Routes>
  );
};
