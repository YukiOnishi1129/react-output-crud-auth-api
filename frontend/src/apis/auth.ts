import globalAxios, { isAxiosError } from "./globalAxios";

import { AuthType } from "../types/User";

export const login = async (email: string, password: string) => {
  try {
    const response = await globalAxios.post<AuthType>("/auth/login", {
      email,
      password,
    });
    return response.data;
  } catch (error) {
    if (isAxiosError(error)) {
      console.log(error);
    }
    console.log(error);
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
