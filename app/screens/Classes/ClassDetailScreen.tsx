import { useLocalSearchParams, useRouter } from "expo-router";
import {
  Award,
  BookOpen,
  Calendar,
  Copy,
  Plus,
  TrendingUp,
  Users,
} from "lucide-react-native";
import { useCallback, useEffect, useState } from "react";
import {
  ActivityIndicator,
  Alert,
  Clipboard,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";

import {
  Assessment,
  createAssessment,
  getAssessments,
} from "@/api/assessment.api";
import { Exam, getExams } from "@/api/exam.api";
import { useAuth } from "@/context/AuthContext";
import { useNotification } from "@/context/NotificationContext";

type TabType = "materials" | "students" | "analytics";

export default function ClassDetailScreen() {
  const router = useRouter();
  const { classId, className, classCode } = useLocalSearchParams<{
    classId: string;
    className: string;
    classCode: string;
  }>();

  const { user } = useAuth();
  const { addNotification } = useNotification();
  const isTeacher = user?.role?.name?.toLowerCase() === "teacher";

  const [activeTab, setActiveTab] = useState<TabType>("materials");
  const [assessments, setAssessments] = useState<Assessment[]>([]);
  const [exams, setExams] = useState<Exam[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchAssessments = useCallback(async () => {
    try {
      setLoading(true);
      if (classId) {
        const [assessmentRes, examRes] = await Promise.all([
          getAssessments(classId),
          getExams(classId),
        ]);
        setAssessments(assessmentRes.data.data);
        setExams(examRes.data.data);
      }
    } catch (error: any) {
      Alert.alert(
        "Error",
        error?.response?.data?.message || "Gagal mengambil data"
      );
    } finally {
      setLoading(false);
    }
  }, [classId]);

  useEffect(() => {
    fetchAssessments();
  }, [classId, fetchAssessments]);

  const handleCopyCode = async () => {
    if (classCode) {
      await Clipboard.setString(classCode);
      addNotification("Kode Tersalin!", `Kode kelas: ${classCode}`, "success");
    }
  };

  const handleCreateAssessment = () => {
    Alert.prompt("Buat Assessment", "Masukkan nama assessment", [
      { text: "Batal", style: "cancel" },
      {
        text: "Lanjut",
        onPress: (name?: string) => {
          if (!name?.trim()) {
            Alert.alert("Error", "Nama assessment tidak boleh kosong");
            return;
          }

          Alert.alert(
            "Tanggal Deadline",
            "Masukkan tanggal deadline (YYYY-MM-DD HH:MM:SS)",
            [
              { text: "Batal", style: "cancel" },
              {
                text: "Buat",
                onPress: async (deadline?: string) => {
                  if (!deadline?.trim()) {
                    Alert.alert("Error", "Deadline tidak boleh kosong");
                    return;
                  }

                  try {
                    if (!classId) return;
                    const response = await createAssessment(classId, {
                      name: name,
                      deadline_at: deadline,
                    });

                    addNotification(
                      "Assessment Berhasil Dibuat! 🎉",
                      `${name} - Deadline: ${deadline}`,
                      "success"
                    );

                    Alert.alert("Sukses", response.data.message);
                    fetchAssessments();
                  } catch (error: any) {
                    addNotification(
                      "Gagal Membuat Assessment",
                      error?.response?.data?.message || "Terjadi kesalahan",
                      "error"
                    );
                    Alert.alert(
                      "Error",
                      error?.response?.data?.message ||
                        "Gagal membuat assessment"
                    );
                  }
                },
              },
            ]
          );
        },
      },
    ]);
  };

  const renderHeader = () => (
    <View className="bg-white rounded-b-3xl border-gray-200">
      <View className="p-5 pt-14">
        <View className="flex-row justify-between items-center mb-4">
          <View className="flex-1">
            <Text className="text-2xl font-bold text-slate-900">
              {className}
            </Text>
            <Text className="text-sm text-gray-600 mt-1">
              Kode: {classCode}
            </Text>
          </View>
          <TouchableOpacity
            className="w-10 h-10 rounded-full bg-blue-100 justify-center items-center"
            onPress={handleCopyCode}
          >
            <Copy size={18} color="#3b82f6" />
          </TouchableOpacity>
        </View>

        <View className="flex-row gap-2">
          <TouchableOpacity
            className={`flex-1 py-4 px-2 rounded-2xl flex-row items-center justify-center gap-1 ${
              activeTab === "materials"
                ? "bg-blue-600"
                : "bg-gray-100 border border-gray-200"
            }`}
            onPress={() => setActiveTab("materials")}
          >
            <BookOpen
              size={14}
              color={activeTab === "materials" ? "#ffffff" : "#3b82f6"}
            />
            <Text
              className={`text-xs font-semibold ${
                activeTab === "materials" ? "text-white" : "text-gray-600"
              }`}
            >
              Materi
            </Text>
          </TouchableOpacity>

          <TouchableOpacity
            className={`flex-1 py-4 px-2 rounded-2xl flex-row items-center justify-center gap-1 ${
              activeTab === "students"
                ? "bg-blue-600"
                : "bg-gray-100 border border-gray-200"
            }`}
            onPress={() => setActiveTab("students")}
          >
            <Users
              size={14}
              color={activeTab === "students" ? "#ffffff" : "#3b82f6"}
            />
            <Text
              className={`text-xs font-semibold ${
                activeTab === "students" ? "text-white" : "text-gray-600"
              }`}
            >
              Siswa
            </Text>
          </TouchableOpacity>

          <TouchableOpacity
            className={`flex-1 py-4 px-2 rounded-2xl flex-row items-center justify-center gap-1 ${
              activeTab === "analytics"
                ? "bg-blue-600"
                : "bg-gray-100 border border-gray-200"
            }`}
            onPress={() => setActiveTab("analytics")}
          >
            <TrendingUp
              size={14}
              color={activeTab === "analytics" ? "#ffffff" : "#3b82f6"}
            />
            <Text
              className={`text-xs font-semibold ${
                activeTab === "analytics" ? "text-white" : "text-gray-600"
              }`}
            >
              Analisis
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );

  const renderMaterialsTab = () => (
    <View className="p-5">
      <View className="mb-6">
        <View className="flex-row justify-between items-center mb-4">
          <Text className="text-lg font-bold text-slate-900">Assessment</Text>
          {isTeacher && (
            <TouchableOpacity
              className="bg-blue-600 px-3 py-2 rounded-lg flex-row items-center gap-1"
              onPress={handleCreateAssessment}
            >
              <Plus size={14} color="#ffffff" />
              <Text className="text-xs font-semibold text-white">Tambah</Text>
            </TouchableOpacity>
          )}
        </View>

        {assessments.length === 0 ? (
          <View className="bg-gray-50 p-6 rounded-2xl items-center">
            <BookOpen size={32} color="#9ca3af" />
            <Text className="text-gray-600 text-sm mt-3">
              {isTeacher
                ? "Belum ada assessment. Tambahkan assessment baru dengan tombol (+)"
                : "Belum ada assessment untuk kelas ini"}
            </Text>
          </View>
        ) : (
          assessments.map((assessment) => (
            <View
              key={assessment.id}
              className="bg-white p-4 rounded-xl mb-3 border border-gray-200"
            >
              <Text className="text-base font-bold text-slate-900 mb-2">
                {assessment.name}
              </Text>
              <View className="flex-row items-center gap-2">
                <Calendar size={14} color="#f59e0b" />
                <Text className="text-xs text-gray-600">
                  Deadline:{" "}
                  {new Date(assessment.deadline_at).toLocaleDateString(
                    "id-ID",
                    {
                      year: "numeric",
                      month: "short",
                      day: "numeric",
                      hour: "2-digit",
                      minute: "2-digit",
                    }
                  )}
                </Text>
              </View>
            </View>
          ))
        )}
      </View>

      <View>
        <View className="flex-row justify-between items-center mb-4">
          <Text className="text-lg font-bold text-slate-900">Ujian</Text>
          {isTeacher && (
            <TouchableOpacity
              className="bg-green-600 px-3 py-2 rounded-lg flex-row items-center gap-1"
              onPress={() =>
                router.push({
                  pathname: "/(tabs)/create-exam" as any,
                  params: { classId, className },
                })
              }
            >
              <Plus size={14} color="#ffffff" />
              <Text className="text-xs font-semibold text-white">Buat</Text>
            </TouchableOpacity>
          )}
        </View>

        {exams.length === 0 ? (
          <View className="bg-gray-50 p-6 rounded-2xl items-center">
            <Award size={32} color="#9ca3af" />
            <Text className="text-gray-600 text-sm mt-3">
              {isTeacher
                ? "Belum ada ujian. Buat ujian baru dengan tombol (+)"
                : "Belum ada ujian untuk kelas ini"}
            </Text>
          </View>
        ) : (
          exams.map((exam) => (
            <View
              key={exam.id}
              className="bg-white p-4 rounded-xl mb-3 border border-gray-200"
            >
              <Text className="text-base font-bold text-slate-900 mb-2">
                {exam.name}
              </Text>
              <View className="flex-row items-center gap-2 mb-2">
                <Calendar size={14} color="#10b981" />
                <Text className="text-xs text-gray-600">
                  {new Date(exam.start_at).toLocaleDateString("id-ID", {
                    year: "numeric",
                    month: "short",
                    day: "numeric",
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </Text>
              </View>
              <View className="flex-row items-center gap-2">
                <Award size={14} color="#f59e0b" />
                <Text className="text-xs text-gray-600">
                  Durasi: {exam.duration} menit
                </Text>
              </View>
            </View>
          ))
        )}
      </View>
    </View>
  );

  const renderStudentsTab = () => (
    <View className="p-5">
      <Text className="text-lg font-bold text-slate-900 mb-4">
        Daftar Siswa
      </Text>
      <View className="bg-gray-50 p-6 rounded-2xl items-center">
        <Users size={32} color="#9ca3af" />
        <Text className="text-gray-600 text-sm mt-3">
          Fitur daftar siswa sedang dikembangkan
        </Text>
      </View>
    </View>
  );

  const renderAnalyticsTab = () => (
    <View className="p-5">
      <Text className="text-lg font-bold text-slate-900 mb-4">Analisis</Text>
      <View className="bg-gray-50 p-6 rounded-2xl items-center">
        <TrendingUp size={32} color="#9ca3af" />
        <Text className="text-gray-600 text-sm mt-3">
          Fitur analisis sedang dikembangkan
        </Text>
      </View>
    </View>
  );

  const renderContent = () => {
    switch (activeTab) {
      case "materials":
        return renderMaterialsTab();
      case "students":
        return renderStudentsTab();
      case "analytics":
        return renderAnalyticsTab();
      default:
        return null;
    }
  };

  return (
    <View className="flex-1 bg-blue-100">
      {renderHeader()}
      {loading ? (
        <View className="flex-1 justify-center items-center">
          <ActivityIndicator size="large" color="#3b82f6" />
        </View>
      ) : (
        <ScrollView className="flex-1">{renderContent()}</ScrollView>
      )}
    </View>
  );
}
