export type LoginRequest = {
  identifier: string;
  password: string;
};

export type LoginResponse = {
  status: boolean;
  message: string;
  data: {
    access_token: string;
    refresh_token: string;
  };
};
