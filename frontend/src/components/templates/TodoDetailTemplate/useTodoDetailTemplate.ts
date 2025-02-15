import { useParams } from "react-router";
import { useCallback, useState, useEffect } from "react";
import { getTodo } from "../../../apis/todo";
import { TodoType } from "../../../types/Todo";

export const useTodoDetailTemplate = () => {
  const { id } = useParams();
  const [todo, setTodo] = useState<TodoType | null>(null);

  const fetchTodo = useCallback(async () => {
    if (!id) return;
    const response = await getTodo({ id });
    if (!response) return;

    setTodo(response);
  }, [id]);

  useEffect(() => {
    fetchTodo();
  }, [fetchTodo]);

  return { todo };
};
