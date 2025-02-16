import { BrowserRouter } from "react-router";
import { AuthRouter } from "./AuthRouter";
import { TodoRouter } from "./TodoRouter";
import { AuthProvider } from "../contexts/AuthContext";

export const Router = () => {
  return (
    <BrowserRouter>
      <AuthProvider>
        <AuthRouter />
        <TodoRouter />
      </AuthProvider>
    </BrowserRouter>
  );
};
