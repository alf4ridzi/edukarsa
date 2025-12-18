import { loginApi } from "@/api/auth.api";
import { saveTokens } from "./storage";
import { LoginRequest } from "@/types/auth";

export const login = async (payload: LoginRequest) => {
  const res = await loginApi(payload);
  if (!res?.data.access_token) {
    throw new Error("Login gagal: token tidak ditemukan");
  }

  await saveTokens(res.data.access_token, res.data.refresh_token);

  return true;
};
