import { BrowserRouter } from "react-router";
import { TodoRouter } from "./TodoRouter";

export const Router = () => {
  return (
    <BrowserRouter>
      <TodoRouter />
    </BrowserRouter>
  );
};
