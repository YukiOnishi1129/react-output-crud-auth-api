import { BrowserRouter } from "react-router";
import { AuthRouter } from "./AuthRouter";
import { TodoRouter } from "./TodoRouter";

export const Router = () => {
  return (
    <BrowserRouter>
      <AuthRouter />
      <TodoRouter />
    </BrowserRouter>
  );
};
