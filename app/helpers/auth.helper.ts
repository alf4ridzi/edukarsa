import { loginApi, registerApi } from "@/api/auth.api";
import { saveTokens } from "./storage";
import { LoginRequest, RegisterRequest } from "@/types/auth";

export const login = async (payload: LoginRequest) => {
  const res = await loginApi(payload);
  if (!res?.data.access_token) {
    throw new Error("Login gagal: token tidak ditemukan");
  }

  await saveTokens(res.data.access_token, res.data.refresh_token);

  return true;
};

export const register = async (payload: RegisterRequest) => {
  const res = await registerApi(payload);

  if (res.status !== true) {
    throw new Error("register gagal");
  }

  return true;
};
