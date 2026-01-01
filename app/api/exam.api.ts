import Constants from "expo-constants";
import client from "./client";

const EXPO_PUBLIC_API_URL = Constants.expoConfig?.extra?.EXPO_PUBLIC_API_URL;

export interface ExamOption {
  id: number;
  text: string;
}

export interface ExamQuestion {
  id: number;
  question: string;
  explanation: string;
  options: ExamOption[];
  correct_option_id: number;
  created_at: string;
}

export interface Exam {
  id: string;
  name: string;
  start_at: string;
  duration: number;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface CreateExamRequest {
  name: string;
  class_id: string;
  duration: number;
  start_at: string;
  end_at: string;
}

export interface UpdateExamRequest {
  name?: string;
  duration?: number;
  start_at?: string;
  end_at?: string;
  status?: string;
}

export interface CreateQuestionRequest {
  question: string;
  options: string[];
  correct_index: number;
  explanation: string;
}

export interface CreateExamResponse {
  data: Exam;
  message: string;
  status: boolean;
}

export interface GetExamsResponse {
  data: Exam[];
  message: string;
  status: boolean;
}

export interface GetQuestionsResponse {
  data: ExamQuestion[];
  message: string;
  status: boolean;
}

export interface CreateQuestionsResponse {
  data: ExamQuestion[];
  message: string;
  status: boolean;
}

export interface UpdateExamResponse {
  data: Exam;
  message: string;
  status: boolean;
}

export interface SubmitAnswerRequest {
  option_id: number;
}

export interface SubmitAnswerResponse {
  data: any;
  message: string;
  status: boolean;
}

export const getExams = async (classId: string) => {
  const response = await client.get<GetExamsResponse>(
    `${EXPO_PUBLIC_API_URL}/classes/${classId}/exams`
  );
  return response;
};

export const createExam = async (payload: CreateExamRequest) => {
  const url = `${EXPO_PUBLIC_API_URL}/classes/${payload.class_id}/exams`;
  const response = await client.post<CreateExamResponse>(url, payload);
  return response;
};

export const getExamQuestions = async (examId: string) => {
  const response = await client.get<GetQuestionsResponse>(
    `${EXPO_PUBLIC_API_URL}/exams/${examId}/questions`
  );
  return response;
};

export const createExamQuestions = async (
  examId: string,
  questions: CreateQuestionRequest[]
) => {
  const response = await client.post<CreateQuestionsResponse>(
    `${EXPO_PUBLIC_API_URL}/exams/${examId}/questions`,
    questions
  );
  return response;
};

export const updateExam = async (
  examId: string,
  payload: UpdateExamRequest
) => {
  const response = await client.patch<UpdateExamResponse>(
    `${EXPO_PUBLIC_API_URL}/exams/${examId}`,
    payload
  );
  return response;
};

export const getStudentClassExams = async (classId: string) => {
  const response = await client.get<GetExamsResponse>(
    `${EXPO_PUBLIC_API_URL}/student/classes/${classId}/exams`
  );
  return response;
};

export const getStudentExamQuestions = async (examId: string) => {
  const response = await client.get<GetQuestionsResponse>(
    `${EXPO_PUBLIC_API_URL}/student/exams/${examId}/questions`
  );
  return response;
};

export const submitStudentAnswer = async (
  examId: string,
  questionId: number,
  payload: SubmitAnswerRequest
) => {
  const response = await client.post<SubmitAnswerResponse>(
    `${EXPO_PUBLIC_API_URL}/student/exams/${examId}/questions/${questionId}/answer`,
    payload
  );
  return response;
};
