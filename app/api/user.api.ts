import apiClient from "./client";

export type UserProfile = {
  id: number;
  role: {
    id: number;
    name: string;
    created_at: string;
    updated_at: string;
  };
  name: string;
  email: string;
  username: string;
  birthday: string | null;
  created_at: string;
  updated_at: string;
};

export type UserProfileResponse = {
  data: {
    user: UserProfile;
  };
  message: string;
  status: boolean;
};

export const getUserProfile = async () => {
  const res = await apiClient.get<UserProfileResponse>("/users/me");
  return res.data;
};
