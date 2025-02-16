import { NAVIGATION_LIST } from "../constants/navigation";
import { Routes, Route } from "react-router";
import {
  TodoCreatePage,
  TodoDetailPage,
  TodoEditPage,
  TodoListPage,
} from "../pages";

export const TodoRouter = () => {
  return (
    <Routes>
      <Route path={NAVIGATION_LIST.TOP} element={<TodoListPage />} />
      <Route path={NAVIGATION_LIST.DETAIL} element={<TodoDetailPage />} />
      <Route path={NAVIGATION_LIST.CREATE} element={<TodoCreatePage />} />
      <Route path={NAVIGATION_LIST.EDIT} element={<TodoEditPage />} />
    </Routes>
  );
};
