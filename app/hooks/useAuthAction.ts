import { useAuth } from "@/context/AuthContext";
import { loginApi, registerApi } from "@/api/auth.api";
import { saveTokens } from "@/helpers/storage";
import { LoginRequest, RegisterRequest } from "@/types/auth";

export const useAuthActions = () => {
  const { login } = useAuth();

  const loginReq = async (payload: LoginRequest) => {
    const res = await loginApi(payload);

    if (!res.data?.access_token) {
      throw new Error("Login gagal");
    }

    await saveTokens(res.data.access_token, res.data.refresh_token);
    await login(res.data.access_token);
  };

  const registerReq = async (payload: RegisterRequest) => {
    const res = await registerApi(payload);

    if (res.status !== true) {
      throw new Error("Register gagal");
    }
  };

  return { loginReq, registerReq };
};
