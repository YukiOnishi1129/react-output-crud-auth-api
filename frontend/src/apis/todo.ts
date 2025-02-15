import globalAxios from "./globalAxios";
import {
  TodoListType,
  TodoType,
  GetTodoRequest,
  CreateTodoRequest,
  UpdateTodoRequest,
  DeleteTodoRequest,
} from "../types/Todo";

export const getTodos = async () => {
  try {
    const response = await globalAxios.get<TodoListType>("/todos");
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const getTodo = async (request: GetTodoRequest) => {
  try {
    const response = await globalAxios.get<TodoType>(`/todos/${request.id}`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const createTodo = async (request: CreateTodoRequest) => {
  try {
    const response = await globalAxios.post<TodoType>("/todos", request);
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const updateTodo = async (request: UpdateTodoRequest) => {
  try {
    const response = await globalAxios.put<TodoType>(`/todos/${request.id}`, {
      title: request.title,
      content: request.content,
    });
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const deleteTodo = async (request: DeleteTodoRequest) => {
  try {
    await globalAxios.delete(`/todos/${request.id}`);
  } catch (error) {
    console.error(error);
  }
};
