import { loginApi, registerApi } from "@/api/auth.api";
import { useAuth } from "@/context/AuthContext";
import { saveTokens } from "@/helpers/storage";
import { LoginRequest, RegisterRequest } from "@/types/auth";

export const useAuthActions = () => {
  const { login, logout } = useAuth();

  const logoutHandle = async () => {
    await logout();
  };

  const loginReq = async (payload: LoginRequest) => {
    const res = await loginApi(payload);

    if (!res.data?.access_token) {
      throw new Error("Login gagal");
    }

    await saveTokens(res.data.access_token, res.data.refresh_token);

    // Pass full user data including role
    const userData = res.data.user || {
      name: payload.identifier,
      email: payload.identifier.includes("@") ? payload.identifier : "",
    };

    await login(res.data.access_token, userData);
  };

  const registerReq = async (payload: RegisterRequest) => {
    const res = await registerApi(payload);

    if (res.status !== true) {
      throw new Error("Register gagal");
    }
  };

  return { loginReq, registerReq, logoutHandle };
};
