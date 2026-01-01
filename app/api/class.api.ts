import apiClient from "./client";

export interface ClassCreator {
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
}

export interface ClassData {
  id: string;
  code: string;
  name: string;
  created_by?: ClassCreator;
  created_at: string;
  updated_at: string;
}

export interface GetClassesResponse {
  data: ClassData[];
  message: string;
  status: boolean;
}

export interface CreateClassRequest {
  name: string;
}

export interface CreateClassResponse {
  data: ClassData;
  message: string;
  status: boolean;
}

export interface JoinLeaveResponse {
  data: null;
  message: string;
  status: boolean;
}

export const getClasses = async (): Promise<GetClassesResponse> => {
  const response = await apiClient.get<GetClassesResponse>("/classes");
  return response.data;
};

export const createClass = async (
  classData: CreateClassRequest
): Promise<CreateClassResponse> => {
  const response = await apiClient.post<CreateClassResponse>(
    "/classes",
    classData
  );
  return response.data;
};

export const joinClass = async (code: string): Promise<JoinLeaveResponse> => {
  const response = await apiClient.post<JoinLeaveResponse>(
    `/classes/code/${code}/join`
  );
  return response.data;
};

export const leaveClass = async (code: string): Promise<JoinLeaveResponse> => {
  const response = await apiClient.delete<JoinLeaveResponse>(
    `/classes/code/${code}/leave`
  );
  return response.data;
};
