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
    user?: {
      name: string;
      email: string;
      username?: string;
      role?: {
        id?: number | bigint;
        name: string;
      };
    };
  };
};

export type RegisterRequest = {
  name: string;
  email: string;
  username: string;
  password: string;
  confirmpassword: string;
};

export type RegisterResponse = {
  status: boolean;
  message: string;
  data: any;
};
