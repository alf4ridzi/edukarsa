import apiClient from "./client";
import { LoginRequest, LoginResponse } from "@/types/auth";

export const loginApi = async (payload: LoginRequest) => {
  const res = await apiClient.post<LoginResponse>("/auth/login", payload);
  return res.data;
};

export const refreshTokenApi = async (refreshToken: string) => {
  const res = await apiClient.post("/auth/refresh", {
    refresh_token: refreshToken,
  });
  return res.data;
};
