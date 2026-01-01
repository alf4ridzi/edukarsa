import { useRouter } from "expo-router";
import {
  Award,
  BookOpen,
  FolderOpen,
  Plus,
  TrendingUp,
  Users,
} from "lucide-react-native";
import { useEffect, useState } from "react";
import {
  ActivityIndicator,
  Alert,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";

import {
  ClassData,
  createClass,
  getClasses,
  joinClass,
  leaveClass,
} from "@/api/class.api";
import { useAuth } from "@/context/AuthContext";
import { useNotification } from "@/context/NotificationContext";

const classColors = ["#3b82f6", "#8b5cf6", "#10b981", "#f59e0b"];

export default function ClassesScreen() {
  const router = useRouter();
  const { user } = useAuth();
  const { addNotification } = useNotification();
  const isStudent = user?.role?.name?.toLowerCase() === "siswa";
  const isTeacher = user?.role?.name?.toLowerCase() === "teacher";
  const [classes, setClasses] = useState<ClassData[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchClasses();
  }, []);

  const fetchClasses = async () => {
    try {
      setLoading(true);
      const response = await getClasses();
      setClasses(response.data);
    } catch (error: any) {
      Alert.alert(
        "Error",
        error?.response?.data?.message || "Gagal mengambil data kelas"
      );
    } finally {
      setLoading(false);
    }
  };

  const handleCreateClass = async () => {
    Alert.prompt(
      "Buat Kelas Baru",
      "Masukkan nama kelas",
      [
        { text: "Batal", style: "cancel" },
        {
          text: "Buat",
          onPress: async (className?: string) => {
            if (!className?.trim()) {
              Alert.alert("Error", "Nama kelas tidak boleh kosong");
              return;
            }
            try {
              const response = await createClass({ name: className });
              addNotification(
                "Kelas Berhasil Dibuat! 🎉",
                `${className} - Kode: ${response.data.code}`,
                "success"
              );
              Alert.alert(
                "Kelas Berhasil Dibuat!",
                `${response.message}\n\nKode Kelas: ${response.data.code}\n\nBagikan kode ini kepada siswa untuk bergabung ke kelas.`,
                [{ text: "OK", onPress: () => fetchClasses() }]
              );
            } catch (error: any) {
              addNotification(
                "Gagal Membuat Kelas",
                error?.response?.data?.message || "Terjadi kesalahan",
                "error"
              );
              Alert.alert(
                "Error",
                error?.response?.data?.message || "Gagal buat kelas"
              );
            }
          },
        },
      ],
      "plain-text"
    );
  };

  const handleJoinClass = async () => {
    Alert.prompt(
      "Bergabung Kelas",
      "Masukkan kode kelas",
      [
        { text: "Batal", style: "cancel" },
        {
          text: "Bergabung",
          onPress: async (code?: string) => {
            if (!code?.trim()) {
              Alert.alert("Error", "Kode kelas tidak boleh kosong");
              return;
            }
            try {
              const response = await joinClass(code);
              addNotification(
                "Berhasil Bergabung! 🎓",
                response.message,
                "success"
              );
              Alert.alert("Sukses", response.message);
              fetchClasses();
            } catch (error: any) {
              addNotification(
                "Gagal Bergabung",
                error?.response?.data?.message || "Terjadi kesalahan",
                "error"
              );
              Alert.alert(
                "Error",
                error?.response?.data?.message || "Gagal bergabung dengan kelas"
              );
            }
          },
        },
      ],
      "plain-text"
    );
  };

  const handleLeaveClass = async (code: string, className: string) => {
    Alert.alert(
      "Keluar dari Kelas",
      `Apakah Anda yakin ingin keluar dari kelas ${className}?`,
      [
        { text: "Batal", style: "cancel" },
        {
          text: "Keluar",
          style: "destructive",
          onPress: async () => {
            try {
              const response = await leaveClass(code);
              Alert.alert("Sukses", response.message);
              fetchClasses();
            } catch (error: any) {
              Alert.alert(
                "Error",
                error?.response?.data?.message || "Gagal keluar dari kelas"
              );
            }
          },
        },
      ]
    );
  };

  const handleShowClassCode = (code: string, className: string) => {
    Alert.alert(
      "Kode Kelas",
      `Kode untuk kelas ${className}:\n\n${code}\n\nBagikan kode ini kepada siswa untuk bergabung.`,
      [{ text: "OK" }]
    );
  };

  const handleNavigateToClassDetail = (
    classId: string,
    className: string,
    classCode: string
  ) => {
    router.push({
      pathname: "/(tabs)/class-detail" as any,
      params: {
        classId: classId.toString(),
        className: className.toString(),
        classCode: classCode.toString(),
      },
    });
  };

  return (
    <View className="flex-1 bg-blue-100">
      <View className="p-5 pt-14 bg-white border-b border-gray-200 rounded-b-3xl">
        <View className="flex-row justify-between items-center mb-3">
          <View className="flex-1">
            <Text className="text-2xl font-bold text-slate-900">
              Kelas Saya
            </Text>
            <Text className="text-sm text-gray-600 mt-1">
              {classes.length} Kelas
            </Text>
            <Text className="text-xs text-gray-400 mt-1">
              Role: {user?.role?.name || "No role"} | Teacher:{" "}
              {isTeacher ? "Yes" : "No"} | Student: {isStudent ? "Yes" : "No"}
            </Text>
          </View>
          {(isTeacher || !user?.role) && (
            <TouchableOpacity
              className="w-10 h-10 rounded-full bg-blue-600 justify-center items-center"
              onPress={handleCreateClass}
            >
              <Plus size={20} color="#ffffff" />
            </TouchableOpacity>
          )}
        </View>
        {(isStudent || !user?.role) && (
          <TouchableOpacity
            className="bg-blue-50 p-3 rounded-3xl border border-blue-200"
            onPress={handleJoinClass}
          >
            <Text className="text-sm font-semibold text-blue-600 text-center">
              + Bergabung dengan Kelas
            </Text>
          </TouchableOpacity>
        )}
      </View>

      {loading ? (
        <View className="flex-1 justify-center items-center">
          <ActivityIndicator size="large" color="#3b82f6" />
          <Text className="mt-3 text-sm text-gray-600">
            Memuat data kelas...
          </Text>
        </View>
      ) : (
        <ScrollView className="flex-1">
          <View className="flex-row p-5 gap-3">
            <View className="flex-1 bg-white p-4 rounded-2xl items-center">
              <Users size={24} color="#3b82f6" />
              <Text className="text-xl font-bold text-slate-900 mt-2">
                {classes.length}
              </Text>
              <Text className="text-xs text-gray-600 mt-1 text-center">
                Total Kelas
              </Text>
            </View>
            <View className="flex-1 bg-white p-4 rounded-2xl items-center">
              <TrendingUp size={24} color="#10b981" />
              <Text className="text-xl font-bold text-slate-900 mt-2">85</Text>
              <Text className="text-xs text-gray-600 mt-1 text-center">
                Rata-rata Nilai
              </Text>
            </View>
            <View className="flex-1 bg-white p-4 rounded-2xl items-center">
              <FolderOpen size={24} color="#f59e0b" />
              <Text className="text-xl font-bold text-slate-900 mt-2">12</Text>
              <Text className="text-xs text-gray-600 mt-1 text-center">
                Tugas Aktif
              </Text>
            </View>
          </View>

          <View className="p-5">
            <Text className="text-lg font-bold text-slate-900 mb-4">
              Daftar Kelas
            </Text>
            {classes.length === 0 ? (
              <View className="bg-white p-6 rounded-2xl items-center">
                <Text className="text-gray-600 text-sm">
                  Belum ada kelas. Buat kelas baru dengan tombol (+)
                </Text>
              </View>
            ) : (
              classes.map((classItem, index) => (
                <TouchableOpacity
                  key={classItem.id}
                  className="bg-white rounded-2xl p-4 mb-3"
                >
                  <View
                    className="flex-row justify-between items-center pb-3 border-b border-gray-100 border-l-4 pl-3"
                    style={{
                      borderLeftColor: classColors[index % classColors.length],
                    }}
                  >
                    <View className="flex-1">
                      <Text className="text-lg font-bold text-slate-900">
                        {classItem.name}
                      </Text>
                      <Text className="text-xs text-gray-500 mt-0.5">
                        Kode: {classItem.code}
                      </Text>
                    </View>
                    <View
                      className="w-12 h-12 rounded-full justify-center items-center"
                      style={{
                        backgroundColor:
                          classColors[index % classColors.length] + "20",
                      }}
                    >
                      <Text
                        className="text-base font-bold"
                        style={{
                          color: classColors[index % classColors.length],
                        }}
                      >
                        {index + 1}
                      </Text>
                    </View>
                  </View>

                  <View className="flex-row gap-2 pt-3 border-t border-gray-100">
                    <TouchableOpacity
                      className="flex-1 flex-row items-center justify-center p-2 bg-gray-100 rounded-lg gap-1"
                      onPress={() =>
                        handleNavigateToClassDetail(
                          classItem.id,
                          classItem.name,
                          classItem.code
                        )
                      }
                    >
                      <BookOpen size={16} color="#3b82f6" />
                      <Text className="text-xs font-semibold text-blue-600">
                        Materi
                      </Text>
                    </TouchableOpacity>
                    <TouchableOpacity
                      className="flex-1 flex-row items-center justify-center p-2 bg-gray-100 rounded-lg gap-1"
                      onPress={() =>
                        handleNavigateToClassDetail(
                          classItem.id,
                          classItem.name,
                          classItem.code
                        )
                      }
                    >
                      <Users size={16} color="#3b82f6" />
                      <Text className="text-xs font-semibold text-blue-600">
                        Siswa
                      </Text>
                    </TouchableOpacity>
                    <TouchableOpacity
                      className="flex-1 flex-row items-center justify-center p-2 bg-gray-100 rounded-lg gap-1"
                      onPress={() =>
                        handleNavigateToClassDetail(
                          classItem.id,
                          classItem.name,
                          classItem.code
                        )
                      }
                    >
                      <TrendingUp size={16} color="#3b82f6" />
                      <Text className="text-xs font-semibold text-blue-600">
                        Analisis
                      </Text>
                    </TouchableOpacity>
                    {isTeacher && (
                      <TouchableOpacity
                        className="flex-1 flex-row items-center justify-center p-2 bg-green-100 rounded-lg gap-1"
                        onPress={() =>
                          handleShowClassCode(classItem.code, classItem.name)
                        }
                      >
                        <Text className="text-xs font-semibold text-green-600">
                          Kode
                        </Text>
                      </TouchableOpacity>
                    )}
                    {isStudent && (
                      <TouchableOpacity
                        className="flex-1 flex-row items-center justify-center p-2 bg-red-100 rounded-lg gap-1"
                        onPress={() =>
                          handleLeaveClass(classItem.code, classItem.name)
                        }
                      >
                        <Text className="text-xs font-semibold text-red-600">
                          Keluar
                        </Text>
                      </TouchableOpacity>
                    )}
                  </View>
                </TouchableOpacity>
              ))
            )}
          </View>

          <View className="p-5">
            <View className="flex-row justify-between items-center mb-4">
              <Text className="text-lg font-bold text-slate-900">
                Pencapaian Kelas
              </Text>
              <TouchableOpacity>
                <Text className="text-sm text-blue-600 font-semibold">
                  Lihat Semua
                </Text>
              </TouchableOpacity>
            </View>
            <View className="bg-white p-4 rounded-2xl flex-row gap-4 items-center">
              <View className="w-16 h-16 rounded-full bg-amber-100 justify-center items-center">
                <Award size={32} color="#f59e0b" />
              </View>
              <View className="flex-1">
                <Text className="text-sm font-bold text-slate-900">
                  Kelas Terbaik Bulan Ini
                </Text>
                <Text className="text-base font-bold text-amber-500 mt-1">
                  {classes.length > 0 ? classes[0].name : "N/A"}
                </Text>
                <Text className="text-xs text-gray-600 mt-1">
                  Rata-rata nilai tertinggi
                </Text>
              </View>
            </View>
          </View>
        </ScrollView>
      )}
    </View>
  );
}
