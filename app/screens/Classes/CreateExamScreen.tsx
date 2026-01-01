import { router, useLocalSearchParams } from "expo-router";
import { ArrowLeft, Calendar, Clock, Plus } from "lucide-react-native";
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

import { createExam } from "@/api/exam.api";
import { useNotification } from "@/context/NotificationContext";

export default function CreateExamScreen() {
  const params = useLocalSearchParams<{
    classId: string;
    className: string;
  }>();

  const classId = params.classId;
  const className = params.className;

  const { addNotification } = useNotification();

  const [loading, setLoading] = useState(false);
  const [examData, setExamData] = useState({
    name: "",
    duration: "",
    start_date: "",
    start_time: "",
    end_date: "",
    end_time: "",
  });

  if (!classId) {
    return (
      <View className="flex-1 bg-blue-100 items-center justify-center">
        <Text className="text-slate-900">Loading...</Text>
      </View>
    );
  }

  const formatToISO = (date: string, time: string) => {
    return `${date}T${time}:00+07:00`;
  };

  const handleCreateExam = async () => {
    if (
      !examData.name.trim() ||
      !examData.duration.trim() ||
      !examData.start_date.trim() ||
      !examData.start_time.trim() ||
      !examData.end_date.trim() ||
      !examData.end_time.trim()
    ) {
      addNotification(
        "Data Tidak Lengkap",
        "Mohon isi semua informasi ujian",
        "error"
      );
      Alert.alert("Error", "Semua field informasi ujian harus diisi");
      return;
    }

    try {
      setLoading(true);

      const start_at = formatToISO(examData.start_date, examData.start_time);
      const end_at = formatToISO(examData.end_date, examData.end_time);

      console.log("Creating Exam");
      console.log("Exam Data:", {
        name: examData.name.trim(),
        class_id: classId,
        duration: parseInt(examData.duration),
        start_at,
        end_at,
      });
      console.log("Class ID being used:", classId);

      const examRes = await createExam({
        name: examData.name.trim(),
        class_id: classId || "",
        duration: parseInt(examData.duration),
        start_at,
        end_at,
      });

      console.log("Exam created, response:", examRes.data);

      const examId = examRes.data.data.id;
      console.log("Exam ID:", examId);

      addNotification(
        "Ujian Berhasil Dibuat! 🎉",
        `${examData.name} - Lanjut membuat pertanyaan`,
        "success"
      );

      router.push({
        pathname: "/(tabs)/create-question",
        params: {
          examId: examId,
          examName: examData.name.trim(),
          className: className,
        },
      });
    } catch (error: any) {
      console.error("=== Error Creating Exam ===");
      console.error("Error object:", error);
      console.error("Error response:", error?.response);
      console.error("Error data:", error?.response?.data);
      console.error("Error message:", error?.message);

      const errorMessage =
        error?.response?.data?.message ||
        error?.message ||
        "Terjadi kesalahan tidak diketahui";

      addNotification("Gagal Membuat Ujian", errorMessage, "error");
      Alert.alert("Error", errorMessage);
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="flex-1 bg-blue-100">
      <View className="bg-white px-5 py-4 pt-14 rounded-b-3xl">
        <View className="flex-row items-center justify-between mb-2">
          <TouchableOpacity
            onPress={() => router.back()}
            className="w-10 h-10 bg-gray-100 rounded-full items-center justify-center"
          >
            <ArrowLeft size={20} color="#1e293b" />
          </TouchableOpacity>
          <View className="flex-1 mx-3">
            <Text className="text-xl font-bold text-slate-900">Buat Ujian</Text>
            <Text className="text-xs text-gray-600">{className}</Text>
          </View>
          <View className="px-3 py-1.5 rounded-full bg-blue-100">
            <Text className="text-xs font-semibold text-blue-600">
              Langkah 1/2
            </Text>
          </View>
        </View>
      </View>

      <ScrollView className="flex-1 p-5" showsVerticalScrollIndicator={false}>
        <View className="bg-blue-600 p-4 rounded-2xl mb-5">
          <Text className="text-white font-bold text-base mb-1">
            Informasi Ujian
          </Text>
          <Text className="text-blue-100 text-xs">
            Isi detail ujian sebelum membuat pertanyaan
          </Text>
        </View>

        <View className="bg-white p-4 rounded-2xl mb-4 shadow-sm">
          <View className="flex-row items-center mb-3">
            <View className="w-8 h-8 bg-blue-100 rounded-lg items-center justify-center mr-2">
              <Plus size={16} color="#2563eb" />
            </View>
            <Text className="text-sm font-bold text-slate-900">Nama Ujian</Text>
          </View>
          <TextInput
            className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-base"
            placeholder="Contoh: Ujian Tengah Semester Matematika"
            placeholderTextColor="#94a3b8"
            value={examData.name}
            onChangeText={(text) => setExamData({ ...examData, name: text })}
          />
        </View>

        <View className="bg-white p-4 rounded-2xl mb-4 shadow-sm">
          <View className="flex-row items-center mb-3">
            <View className="w-8 h-8 bg-amber-100 rounded-lg items-center justify-center mr-2">
              <Clock size={16} color="#f59e0b" />
            </View>
            <Text className="text-sm font-bold text-slate-900">
              Durasi (Menit)
            </Text>
          </View>
          <TextInput
            className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-base"
            placeholder="Contoh: 90"
            placeholderTextColor="#94a3b8"
            keyboardType="numeric"
            value={examData.duration}
            onChangeText={(text) =>
              setExamData({ ...examData, duration: text })
            }
          />
        </View>

        <View className="bg-white p-4 rounded-2xl mb-4 shadow-sm">
          <View className="flex-row items-center mb-3">
            <View className="w-8 h-8 bg-green-100 rounded-lg items-center justify-center mr-2">
              <Calendar size={16} color="#10b981" />
            </View>
            <Text className="text-sm font-bold text-slate-900">
              Waktu Mulai
            </Text>
          </View>
          <View className="flex-row gap-2">
            <View className="flex-1">
              <Text className="text-xs text-gray-500 mb-2">Tanggal</Text>
              <TextInput
                className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-sm"
                placeholder="2026-01-15"
                placeholderTextColor="#94a3b8"
                value={examData.start_date}
                onChangeText={(text) =>
                  setExamData({ ...examData, start_date: text })
                }
              />
            </View>
            <View className="flex-1">
              <Text className="text-xs text-gray-500 mb-2">Jam</Text>
              <TextInput
                className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-sm"
                placeholder="08:00"
                placeholderTextColor="#94a3b8"
                value={examData.start_time}
                onChangeText={(text) =>
                  setExamData({ ...examData, start_time: text })
                }
              />
            </View>
          </View>
        </View>

        <View className="bg-white p-4 rounded-2xl mb-6 shadow-sm">
          <View className="flex-row items-center mb-3">
            <View className="w-8 h-8 bg-red-100 rounded-lg items-center justify-center mr-2">
              <Calendar size={16} color="#ef4444" />
            </View>
            <Text className="text-sm font-bold text-slate-900">
              Waktu Selesai
            </Text>
          </View>
          <View className="flex-row gap-2">
            <View className="flex-1">
              <Text className="text-xs text-gray-500 mb-2">Tanggal</Text>
              <TextInput
                className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-sm"
                placeholder="2026-01-16"
                placeholderTextColor="#94a3b8"
                value={examData.end_date}
                onChangeText={(text) =>
                  setExamData({ ...examData, end_date: text })
                }
              />
            </View>
            <View className="flex-1">
              <Text className="text-xs text-gray-500 mb-2">Jam</Text>
              <TextInput
                className="bg-gray-50 border border-gray-200 rounded-xl px-4 py-3 text-slate-900 text-sm"
                placeholder="23:59"
                placeholderTextColor="#94a3b8"
                value={examData.end_time}
                onChangeText={(text) =>
                  setExamData({ ...examData, end_time: text })
                }
              />
            </View>
          </View>
        </View>

        <TouchableOpacity
          className="bg-blue-600 py-4 rounded-2xl items-center flex-row justify-center shadow-sm active:bg-blue-700 mb-10"
          onPress={handleCreateExam}
          disabled={loading}
        >
          {loading ? (
            <ActivityIndicator color="#ffffff" />
          ) : (
            <>
              <Text className="text-white font-bold text-base mr-2">
                Lanjut Buat Pertanyaan
              </Text>
              <ArrowLeft
                size={16}
                color="#ffffff"
                style={{ transform: [{ rotate: "180deg" }] }}
              />
            </>
          )}
        </TouchableOpacity>
      </ScrollView>
    </View>
  );
}
