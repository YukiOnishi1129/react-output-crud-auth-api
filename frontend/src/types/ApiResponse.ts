export interface ResponseType<T = undefined> {
  code: number;
  data?: T;
  message?: string;
}

export type IErrorResponse = {
  code: string;
  config: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  message: string;
  request: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  response: {
    config: any; // eslint-disable-line @typescript-eslint/no-explicit-any
    data: {
      error: string;
      message: string;
    };
    headers: any; // eslint-disable-line @typescript-eslint/no-explicit-any
    request: any; // eslint-disable-line @typescript-eslint/no-explicit-any
    status: number;
    statusText: string;
  };
};
