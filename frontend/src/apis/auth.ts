import globalAxios, { isAxiosError } from "./globalAxios";

import { AuthType } from "../types/User";
import { IErrorResponse, ResponseType } from "../types/ApiResponse";

export const login = async (email: string, password: string) => {
  const res: ResponseType<AuthType> = {
    code: 500,
    message: "",
  };
  try {
    const response = await globalAxios.post<AuthType>("/auth/login", {
      email,
      password,
    });
    res.code = response.status;
    res.data = response.data;

    return res;
  } catch (error) {
    res.message = error as string;
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const register = async (
  name: string,
  email: string,
  password: string
) => {
  const res: ResponseType<AuthType> = {
    code: 500,
    message: "",
  };

  try {
    const response = await globalAxios.post<AuthType>("/auth/register", {
      name,
      email,
      password,
    });
    res.code = response.status;
    res.data = response.data;
    return res;
  } catch (error) {
    res.message = error as string;
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};

export const checkAuthentication = async () => {
  const res: ResponseType<AuthType> = {
    code: 500,
    message: "",
  };
  try {
    const response = await globalAxios.post<AuthType>("/auth/authentication");
    res.code = response.status;
    res.data = response.data;
    return res;
  } catch (error) {
    res.message = error as string;
    if (isAxiosError(error)) {
      const axiosError = error as IErrorResponse;
      res.code = axiosError.response.status;
      res.message = axiosError.response.data.message;
    }
    return res;
  }
};
