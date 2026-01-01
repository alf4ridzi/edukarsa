import Constants from "expo-constants";
import client from "./client";

const EXPO_PUBLIC_API_URL = Constants.expoConfig?.extra?.EXPO_PUBLIC_API_URL;

export interface Assessment {
  id: string;
  name: string;
  deadline_at: string;
}

export interface CreateAssessmentRequest {
  name: string;
  deadline_at: string;
}

export interface AssessmentResponse {
  data: Assessment;
  message: string;
  status: boolean;
}

export interface GetAssessmentsResponse {
  data: Assessment[];
  message: string;
  status: boolean;
}

export const getAssessments = async (classId: string) => {
  const response = await client.get<GetAssessmentsResponse>(
    `${EXPO_PUBLIC_API_URL}/classes/${classId}/assessment`
  );
  return response;
};

export const createAssessment = async (
  classId: string,
  payload: CreateAssessmentRequest
) => {
  const response = await client.post<AssessmentResponse>(
    `${EXPO_PUBLIC_API_URL}/classes/${classId}/assessment`,
    payload
  );
  return response;
};
