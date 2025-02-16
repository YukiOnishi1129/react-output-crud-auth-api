import globalAxios, { isAxiosError } from "./globalAxios";
import {
  TodoListType,
  TodoType,
  GetTodoRequest,
  CreateTodoRequest,
  UpdateTodoRequest,
  DeleteTodoRequest,
} from "../types/Todo";
import { IErrorResponse, ResponseType } from "../types/ApiResponse";

export const getTodos = async () => {
  try {
    const response = await globalAxios.get<TodoListType>("/todos");
    const res: ResponseType<TodoListType> = {
      code: response.status,
      data: response.data,
    };
    return res;
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: error as string,
    };
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const getTodo = async (request: GetTodoRequest) => {
  try {
    const response = await globalAxios.get<TodoType>(`/todos/${request.id}`);
    const res: ResponseType<TodoType> = {
      code: response.status,
      data: response.data,
    };
    return res;
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: error as string,
    };
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const createTodo = async (request: CreateTodoRequest) => {
  try {
    const response = await globalAxios.post<TodoType>("/todos", request);
    const res: ResponseType<TodoType> = {
      code: response.status,
      data: response.data,
    };
    return res;
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: error as string,
    };
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const updateTodo = async (request: UpdateTodoRequest) => {
  try {
    const response = await globalAxios.put<TodoType>(`/todos/${request.id}`, {
      title: request.title,
      content: request.content,
    });
    const res: ResponseType<TodoType> = {
      code: response.status,
      data: response.data,
    };
    return res;
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: error as string,
    };
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const deleteTodo = async (request: DeleteTodoRequest) => {
  try {
    await globalAxios.delete(`/todos/${request.id}`);
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: error as string,
    };
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};
