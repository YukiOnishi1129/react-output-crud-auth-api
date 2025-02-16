import globalAxios, { isAxiosError } from "./globalAxios";

import { AuthType } from "../types/User";
import { IErrorResponse, ResponseType } from "../types/ApiResponse";

export const login = async (email: string, password: string) => {
  try {
    const response = await globalAxios.post<AuthType>("/auth/login", {
      email,
      password,
    });
    const res: ResponseType<AuthType> = {
      code: response.status,
      data: response.data,
    };

    return res;
  } catch (error) {
    const res: ResponseType = {
      code: 500,
      message: "",
    };
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
  try {
    const response = await globalAxios.post<AuthType>("/auth/register", {
      name,
      email,
      password,
    });
    return response.data;
  } catch (error) {
    console.error(error);
  }
};

export const checkAuthentication = async () => {
  try {
    const response = await globalAxios.post<AuthType>("/auth/authentication");
    return response.data;
  } catch (error) {
    console.error(error);
  }
};
