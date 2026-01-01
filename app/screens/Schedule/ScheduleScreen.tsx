import {
  BookOpen,
  Calendar,
  CheckCircle2,
  FileText,
  Video,
} from "lucide-react-native";
import { useState } from "react";
import { ScrollView, Text, TouchableOpacity, View } from "react-native";

export default function ScheduleScreen() {
  const [selectedSemester, setSelectedSemester] = useState(1);

  const semesters = [
    { id: 1, label: "Semester 1" },
    { id: 2, label: "Semester 2" },
  ];

  const schedule = {
    1: [
      {
        week: 1,
        title: "Pengenalan Trigonometri",
        topics: ["Sudut dan Satuan", "Perbandingan Trigonometri"],
        date: "1-5 Januari 2025",
        materials: 3,
        status: "completed",
      },
      {
        week: 2,
        title: "Identitas Trigonometri",
        topics: ["Identitas Dasar", "Pembuktian Identitas"],
        date: "8-12 Januari 2025",
        materials: 4,
        status: "completed",
      },
      {
        week: 3,
        title: "Grafik Fungsi Trigonometri",
        topics: ["Sinus dan Cosinus", "Tangen dan Cotangen"],
        date: "15-19 Januari 2025",
        materials: 5,
        status: "active",
      },

      {
        week: 4,
        title: "Persamaan Trigonometri",
        topics: ["Persamaan Sederhana", "Persamaan Kompleks"],
        date: "22-26 Januari 2025",
        materials: 4,
        status: "upcoming",
      },

      {
        week: 5,
        title: "Aturan Sinus dan Cosinus",
        topics: ["Aturan Sinus", "Aturan Cosinus", "Penerapan"],
        date: "29 Jan - 2 Feb 2025",
        materials: 5,
        status: "upcoming",
      },
    ],
    2: [
      {
        week: 1,
        title: "Matriks dan Operasinya",
        topics: ["Definisi Matriks", "Operasi Matriks"],
        date: "1-5 Juli 2025",
        materials: 4,
        status: "upcoming",
      },

      {
        week: 2,
        title: "Determinan dan Invers Matriks",
        topics: ["Determinan", "Invers Matriks"],
        date: "8-12 Juli 2025",
        materials: 5,
        status: "upcoming",
      },

      {
        week: 3,
        title: "Sistem Persamaan Linear",
        topics: ["Metode Eliminasi", "Metode Substitusi", "Metode Matriks"],
        date: "15-19 Juli 2025",
        materials: 6,
        status: "upcoming",
      },
    ],
  };

  const currentSchedule =
    schedule[selectedSemester as keyof typeof schedule] || [];

  const getStatusColor = (status: string) => {
    switch (status) {
      case "completed":
        return "#10b981";
      case "active":
        return "#3b82f6";
      case "upcoming":
        return "#94a3b8";
      default:
        return "#64748b";
    }
  };

  const getStatusLabel = (status: string) => {
    switch (status) {
      case "completed":
        return "Selesai";
      case "active":
        return "Berlangsung";
      case "upcoming":
        return "Akan Datang";
      default:
        return status;
    }
  };

  return (
    <View className="flex-1 bg-blue-100">
      <View className="p-5 pt-14 bg-white border-b border-gray-200 rounded-b-3xl">
        <Text className="text-2xl font-bold text-slate-900">
          Jadwal & Materi
        </Text>
        <Text className="text-sm text-gray-600 mt-1">
          Tahun Ajaran 2024/2025
        </Text>
      </View>

      <ScrollView className="flex-1 p-5">
        <View className="flex-row bg-white p-4 gap-3 rounded-3xl mb-5">
          {semesters.map((semester) => (
            <TouchableOpacity
              key={semester.id}
              className={`flex-1 py-3 rounded-2xl items-center ${
                selectedSemester === semester.id ? "bg-blue-600" : "bg-gray-100"
              }`}
              onPress={() => setSelectedSemester(semester.id)}
            >
              <Text
                className={`text-sm font-semibold ${
                  selectedSemester === semester.id
                    ? "text-white"
                    : "text-gray-600"
                }`}
              >
                {semester.label}
              </Text>
            </TouchableOpacity>
          ))}
        </View>
        <View className="bg-white p-5 rounded-2xl mb-5">
          <View className="flex-row justify-between items-center mb-3">
            <Text className="text-base font-semibold text-slate-900">
              Progress Semester {selectedSemester}
            </Text>
            <Text className="text-xl font-bold text-blue-600">40%</Text>
          </View>
          <View className="h-2 bg-gray-200 rounded-full overflow-hidden mb-2">
            <View className="h-full bg-blue-600" style={{ width: "40%" }} />
          </View>
          <Text className="text-xs text-gray-600">2 dari 5 minggu selesai</Text>
        </View>

        {currentSchedule.map((item, index) => (
          <TouchableOpacity
            key={index}
            className="bg-white p-4 rounded-2xl mb-4"
          >
            <View className="flex-row justify-between items-center mb-3">
              <View className="bg-gray-100 px-3 py-1.5 rounded-xl">
                <Text className="text-xs font-semibold text-gray-600">
                  Minggu {item.week}
                </Text>
              </View>
              <View
                className="flex-row items-center gap-1 px-2.5 py-1.5 rounded-xl"
                style={{ backgroundColor: getStatusColor(item.status) + "20" }}
              >
                {item.status === "completed" && (
                  <CheckCircle2 size={14} color={getStatusColor(item.status)} />
                )}
                <Text
                  className="text-xs font-semibold"
                  style={{ color: getStatusColor(item.status) }}
                >
                  {getStatusLabel(item.status)}
                </Text>
              </View>
            </View>

            <Text className="text-base font-bold text-slate-900 mb-2">
              {item.title}
            </Text>

            <View className="flex-row items-center gap-1.5 mb-3">
              <Calendar size={14} color="#64748b" />
              <Text className="text-xs text-gray-600">{item.date}</Text>
            </View>

            <View className="mb-3">
              <Text className="text-xs font-semibold text-gray-600 mb-2">
                Topik Pembelajaran:
              </Text>
              {item.topics.map((topic, idx) => (
                <View key={idx} className="flex-row items-center gap-2 mb-1.5">
                  <View className="w-1 h-1 rounded-full bg-blue-600" />
                  <Text className="text-sm text-slate-900">{topic}</Text>
                </View>
              ))}
            </View>

            <View className="flex-row justify-between items-center pt-3 border-t border-gray-100">
              <View className="flex-row items-center gap-1.5">
                <FileText size={16} color="#3b82f6" />
                <Text className="text-xs text-blue-600 font-semibold">
                  {item.materials} Materi
                </Text>
              </View>
              <TouchableOpacity className="px-4 py-2 bg-blue-600 rounded-xl">
                <Text className="text-xs font-semibold text-white">
                  Lihat Detail
                </Text>
              </TouchableOpacity>
            </View>
          </TouchableOpacity>
        ))}

        <View className="mt-5 pb-5">
          <Text className="text-lg font-bold text-slate-900 mb-4">
            Sumber Belajar
          </Text>

          <TouchableOpacity className="bg-white p-4 rounded-2xl mb-3 flex-row items-center gap-4 ">
            <View className="w-12 h-12 rounded-full bg-gray-100 justify-center items-center">
              <BookOpen size={24} color="#3b82f6" />
            </View>
            <View className="flex-1">
              <Text className="text-sm font-semibold text-slate-900">
                Bank Soal Trigonometri
              </Text>
              <Text className="text-xs text-gray-600 mt-0.5">
                150+ soal latihan dengan pembahasan
              </Text>
            </View>
          </TouchableOpacity>

          <TouchableOpacity className="bg-white p-4 rounded-2xl mb-3 flex-row items-center gap-4 ">
            <View className="w-12 h-12 rounded-full bg-gray-100 justify-center items-center">
              <Video size={24} color="#8b5cf6" />
            </View>
            <View className="flex-1">
              <Text className="text-sm font-semibold text-slate-900">
                Video Pembelajaran
              </Text>
              <Text className="text-xs text-gray-600 mt-0.5">
                25 video materi interaktif
              </Text>
            </View>
          </TouchableOpacity>

          <TouchableOpacity className="bg-white p-4 rounded-2xl mb-3 flex-row items-center gap-4 ">
            <View className="w-12 h-12 rounded-full bg-gray-100 justify-center items-center">
              <FileText size={24} color="#10b981" />
            </View>
            <View className="flex-1">
              <Text className="text-sm font-semibold text-slate-900">
                Rangkuman Materi
              </Text>
              <Text className="text-xs text-gray-600 mt-0.5">
                Ringkasan lengkap semua bab
              </Text>
            </View>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </View>
  );
}
