import { router, useLocalSearchParams } from "expo-router";
import { ArrowLeft, Plus, Trash2 } from "lucide-react-native";
import { useState } from "react";
import {
  ActivityIndicator,
  Alert,
  ScrollView,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";

import { createExamQuestions } from "@/api/exam.api";
import { useNotification } from "@/context/NotificationContext";

interface Question {
  id: string;
  question: string;
  options: string[];
  correct_index: number;
  explanation: string;
}

export default function CreateQuestionScreen() {
  const params = useLocalSearchParams<{
    examId: string;
    examName: string;
    className: string;
  }>();

  const { addNotification } = useNotification();
  const [loading, setLoading] = useState(false);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [currentQuestion, setCurrentQuestion] = useState<Question>({
    id: "0",
    question: "",
    options: ["", "", "", ""],
    correct_index: 0,
    explanation: "",
  });

  const addQuestion = () => {
    const trimmedQuestion = currentQuestion.question.trim();
    const hasEmptyOption = currentQuestion.options.some((o) => !o.trim());

    if (!trimmedQuestion) {
      Alert.alert("Error", "Mohon isi pertanyaan terlebih dahulu");
      return;
    }

    if (hasEmptyOption) {
      Alert.alert("Error", "Mohon isi semua pilihan jawaban (A, B, C, D)");
      return;
    }

    setQuestions([
      ...questions,
      { ...currentQuestion, id: Date.now().toString() },
    ]);
    setCurrentQuestion({
      id: "",
      question: "",
      options: ["", "", "", ""],
      correct_index: 0,
      explanation: "",
    });

    addNotification(
      "Pertanyaan Ditambahkan",
      `Pertanyaan #${questions.length + 1} berhasil ditambahkan`,
      "success"
    );
  };

  const removeQuestion = (index: number) => {
    setQuestions(questions.filter((_, i) => i !== index));
    addNotification(
      "Pertanyaan Dihapus",
      `Pertanyaan #${index + 1} berhasil dihapus`,
      "success"
    );
  };

  const handleSaveQuestions = async () => {
    if (questions.length === 0) {
      Alert.alert("Error", "Tambahkan minimal 1 pertanyaan");
      return;
    }

    try {
      setLoading(true);

      await createExamQuestions(
        params.examId || "",
        questions.map((q) => ({
          question: q.question,
          options: q.options,
          correct_index: q.correct_index,
          explanation: q.explanation,
        }))
      );

      addNotification(
        "Ujian Berhasil Dibuat! 🎉",
        `${params.examName} dengan ${questions.length} pertanyaan`,
        "success"
      );

      Alert.alert(
        "Sukses",
        `Ujian "${params.examName}" berhasil dibuat dengan ${questions.length} pertanyaan`,
        [
          {
            text: "OK",
            onPress: () => {
              router.back();
              router.back();
            },
          },
        ]
      );
    } catch (error: any) {
      const errorMessage =
        error?.response?.data?.message ||
        error?.message ||
        "Terjadi kesalahan tidak diketahui";

      addNotification("Gagal Menyimpan Pertanyaan", errorMessage, "error");
      Alert.alert("Error", errorMessage);
    } finally {
      setLoading(false);
    }
  };

  if (!params.examId) {
    return (
      <View className="flex-1 bg-blue-50 items-center justify-center">
        <Text className="text-slate-900">Loading...</Text>
      </View>
    );
  }

  return (
    <View className="flex-1 bg-blue-50">
      <View className="bg-white px-5 py-4 pt-14 border-b border-gray-200">
        <View className="flex-row items-center justify-between">
          <TouchableOpacity
            onPress={() => router.back()}
            className="w-10 h-10 bg-gray-100 rounded-full items-center justify-center"
          >
            <ArrowLeft size={20} color="#1e293b" />
          </TouchableOpacity>
          <View className="flex-1 mx-3">
            <Text className="text-lg font-bold text-slate-900">
              Buat Pertanyaan
            </Text>
            <Text className="text-xs text-gray-600">{params.examName}</Text>
          </View>
          <View className="px-3 py-1 rounded-full bg-green-100">
            <Text className="text-xs font-semibold text-green-600">
              {questions.length} Soal
            </Text>
          </View>
        </View>
      </View>

      <ScrollView className="flex-1 p-4">
        <Text className="text-base font-bold text-slate-900 mb-4">
          ❓ Pertanyaan Ujian
        </Text>

        {questions.map((q, idx) => (
          <View
            key={q.id}
            className="bg-white p-4 rounded-xl mb-3 border border-gray-200"
          >
            <View className="flex-row justify-between items-start mb-2">
              <Text className="text-sm font-bold text-slate-900 flex-1">
                {idx + 1}. {q.question}
              </Text>
              <TouchableOpacity onPress={() => removeQuestion(idx)}>
                <Trash2 size={18} color="#ef4444" />
              </TouchableOpacity>
            </View>
            <Text className="text-xs text-gray-600">
              Jawaban: {["A", "B", "C", "D"][q.correct_index]}
            </Text>
          </View>
        ))}

        <View className="bg-white p-4 rounded-xl border-2 border-blue-200 mb-4">
          <Text className="text-sm font-bold text-slate-900 mb-3">
            ➕ Tambah Pertanyaan Baru
          </Text>

          <TextInput
            className="bg-gray-50 border border-gray-300 rounded-lg px-3 py-2 text-slate-900 mb-3"
            placeholder="Masukkan pertanyaan..."
            multiline
            value={currentQuestion.question}
            onChangeText={(text) =>
              setCurrentQuestion({ ...currentQuestion, question: text })
            }
          />

          {["A", "B", "C", "D"].map((label, idx) => (
            <View key={idx} className="mb-2">
              <View className="flex-row items-center">
                <TouchableOpacity
                  onPress={() =>
                    setCurrentQuestion({
                      ...currentQuestion,
                      correct_index: idx,
                    })
                  }
                  className={`w-6 h-6 rounded-full border-2 mr-2 items-center justify-center ${
                    currentQuestion.correct_index === idx
                      ? "bg-green-500 border-green-500"
                      : "border-gray-400"
                  }`}
                >
                  {currentQuestion.correct_index === idx && (
                    <View className="w-3 h-3 bg-white rounded-full" />
                  )}
                </TouchableOpacity>
                <Text className="text-sm font-semibold text-slate-700 w-8">
                  {label}.
                </Text>
                <TextInput
                  className="flex-1 bg-gray-50 border border-gray-300 rounded-lg px-3 py-2 text-slate-900"
                  placeholder={`Pilihan ${label}`}
                  value={currentQuestion.options[idx]}
                  onChangeText={(text) => {
                    const newOptions = [...currentQuestion.options];
                    newOptions[idx] = text;
                    setCurrentQuestion({
                      ...currentQuestion,
                      options: newOptions,
                    });
                  }}
                />
              </View>
            </View>
          ))}

          <TextInput
            className="bg-gray-50 border border-gray-300 rounded-lg px-3 py-2 text-slate-900 mt-2"
            placeholder="Penjelasan (opsional)"
            multiline
            value={currentQuestion.explanation}
            onChangeText={(text) =>
              setCurrentQuestion({ ...currentQuestion, explanation: text })
            }
          />

          <TouchableOpacity
            onPress={addQuestion}
            className="bg-blue-600 rounded-lg py-3 items-center mt-4"
          >
            <View className="flex-row items-center">
              <Plus size={18} color="#fff" />
              <Text className="text-white font-bold text-sm ml-2">
                Tambah Pertanyaan
              </Text>
            </View>
          </TouchableOpacity>
        </View>

        {questions.length > 0 && (
          <TouchableOpacity
            onPress={handleSaveQuestions}
            disabled={loading}
            className={`rounded-xl py-4 items-center mb-6 ${
              loading ? "bg-green-300" : "bg-green-600"
            }`}
          >
            {loading ? (
              <ActivityIndicator color="#fff" />
            ) : (
              <Text className="text-white font-bold">Simpan & Selesai ✓</Text>
            )}
          </TouchableOpacity>
        )}
      </ScrollView>
    </View>
  );
}
