export type UserType = {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
};

export type AuthType = {
  token: string;
  user: UserType;
};
