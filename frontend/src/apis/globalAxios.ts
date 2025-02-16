import axios, { AxiosError } from "axios";

const BASE_API_URL = import.meta.env.VITE_REACT_APP_API_URL;

const getToken = () =>
  localStorage.getItem("authentication")
    ? localStorage.getItem("authentication")
    : null;

const apiClient = axios.create({
  baseURL: BASE_API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

apiClient.interceptors.request.use(
  (config) => {
    const token = getToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const setAxiosAuthentication = (token: string) => {
  localStorage.setItem("authentication", token);
  apiClient.defaults.headers.Authorization = `Bearer ${token}`;
};

export default apiClient;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const isAxiosError = (error: any): error is AxiosError =>
  !!error.isAxiosError;
