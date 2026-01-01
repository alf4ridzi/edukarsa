import { useAuth } from "@/context/AuthContext";
import { useNotification } from "@/context/NotificationContext";
import {
  Award,
  Bell,
  CheckCircle,
  Clock,
  FileText,
  TrendingUp,
  Users,
} from "lucide-react-native";
import { ScrollView, Text, TouchableOpacity, View } from "react-native";

export default function HomeScreen() {
  const { user } = useAuth();
  const { notifications } = useNotification();
  const userName = user?.name || "Pengguna";

  const stats = [
    { label: "Total Siswa", value: "156", icon: Users, color: "#3b82f6" },
    { label: "Quiz Aktif", value: "12", icon: FileText, color: "#8b5cf6" },
    { label: "Tugas Masuk", value: "43", icon: CheckCircle, color: "#10b981" },
    {
      label: "Rata-rata Kelas",
      value: "85",
      icon: TrendingUp,
      color: "#f59e0b",
    },
  ];

  const recentActivities = [
    {
      title: "Quiz Matematika Bab 5",
      class: "XI IPA 1",
      time: "2 jam lalu",
      status: "completed",
    },
    {
      title: "Tugas Fisika - Gerak Parabola",
      class: "XI IPA 2",
      time: "3 jam lalu",
      status: "pending",
    },
    {
      title: "Ujian Tengah Semester",
      class: "XI IPA 1",
      time: "1 hari lalu",
      status: "completed",
    },
    {
      title: "Diskusi: Teori Relativitas",
      class: "XI IPA 3",
      time: "2 hari lalu",
      status: "active",
    },
  ];

  const topStudents = [
    { name: "Ahmad Rizki", class: "XI IPA 1", score: 95, badge: "gold" },
    { name: "Siti Nurhaliza", class: "XI IPA 2", score: 93, badge: "silver" },
    { name: "Budi Santoso", class: "XI IPA 1", score: 91, badge: "bronze" },
  ];

  return (
    <ScrollView className="flex-1 bg-blue-100">
      <View className="flex-row justify-between items-center p-5 pt-14 bg-white rounded-b-3xl">
        <View>
          <Text className="text-sm text-gray-600">Selamat Datang</Text>
          <Text className="text-xl font-bold text-slate-900 mt-1">
            {userName}
          </Text>
        </View>
        <TouchableOpacity className="w-11 h-11 rounded-full bg-gray-100 justify-center items-center">
          <Bell size={20} color="#64748b" />
          {notifications.length > 0 && (
            <View className="absolute top-2 right-2 bg-red-500 w-4.5 h-4.5 rounded-full justify-center items-center">
              <Text className="text-white text-[10px] font-bold">
                {notifications.length > 9 ? "9+" : notifications.length}
              </Text>
            </View>
          )}
        </TouchableOpacity>
      </View>

      <View className="flex-row flex-wrap p-4 gap-3">
        {stats.map((stat, index) => (
          <View
            key={index}
            className="flex-1 min-w-[45%] bg-gray-50 p-4 rounded-2xl"
          >
            <View
              className="w-12 h-12 rounded-xl justify-center items-center mb-3"
              style={{ backgroundColor: stat.color + "20" }}
            >
              <stat.icon size={24} color={stat.color} />
            </View>
            <Text className="text-2xl font-bold text-slate-900">
              {stat.value}
            </Text>
            <Text className="text-xs text-gray-600 mt-1">{stat.label}</Text>
          </View>
        ))}
      </View>

      <View className="p-5">
        <View className="flex-row justify-between items-center mb-4">
          <Text className="text-lg font-bold text-slate-900">
            Aktivitas Terbaru
          </Text>
          <TouchableOpacity>
            <Text className="text-sm text-blue-600 font-semibold">
              Lihat Semua
            </Text>
          </TouchableOpacity>
        </View>
        {recentActivities.map((activity, index) => (
          <View
            key={index}
            className="bg-white p-4 rounded-2xl  mb-3 flex-row justify-between items-center"
          >
            <View className="flex-row items-center flex-1">
              <View
                className="w-2 h-2 rounded-full mr-3"
                style={{
                  backgroundColor:
                    activity.status === "completed"
                      ? "#10b981"
                      : activity.status === "pending"
                        ? "#f59e0b"
                        : "#3b82f6",
                }}
              />
              <View className="flex-1">
                <Text className="text-sm font-semibold text-slate-900">
                  {activity.title}
                </Text>
                <Text className="text-xs text-gray-600 mt-0.5">
                  {activity.class}
                </Text>
              </View>
            </View>
            <View className="flex-row items-center gap-1">
              <Clock size={14} color="#94a3b8" />
              <Text className="text-xs text-gray-400">{activity.time}</Text>
            </View>
          </View>
        ))}
      </View>

      <View className="p-5">
        <View className="flex-row justify-between items-center mb-4">
          <Text className="text-lg font-bold text-slate-900">
            Siswa Terbaik Bulan Ini
          </Text>
          <TouchableOpacity>
            <Text className="text-sm text-blue-600 font-semibold">
              Leaderboard
            </Text>
          </TouchableOpacity>
        </View>
        {topStudents.map((student, index) => (
          <View
            key={index}
            className="bg-white p-4 rounded-2xl mb-3 flex-row justify-between items-center"
          >
            <View className="flex-row items-center gap-3">
              <View
                className="w-8 h-8 rounded-full justify-center items-center"
                style={{
                  backgroundColor:
                    student.badge === "gold"
                      ? "#fbbf24"
                      : student.badge === "silver"
                        ? "#94a3b8"
                        : "#fb923c",
                }}
              >
                <Text className="text-white text-sm font-bold">
                  {index + 1}
                </Text>
              </View>
              <View>
                <Text className="text-sm font-semibold text-slate-900">
                  {student.name}
                </Text>
                <Text className="text-xs text-gray-600 mt-0.5">
                  {student.class}
                </Text>
              </View>
            </View>
            <View className="flex-row items-center gap-1">
              <Award size={16} color="#f59e0b" />
              <Text className="text-base font-bold text-slate-900">
                {student.score}
              </Text>
            </View>
          </View>
        ))}
      </View>

      <View className="p-5 pb-10">
        <Text className="text-lg font-bold text-slate-900">Aksi Cepat</Text>
        <View className="flex-row flex-wrap gap-3 mt-4">
          <TouchableOpacity className="flex-1 min-w-[45%] bg-white p-5 rounded-2xl items-center">
            <FileText size={24} color="#3b82f6" />
            <Text className="text-xs font-semibold text-slate-900 mt-2">
              Buat Quiz
            </Text>
          </TouchableOpacity>
          <TouchableOpacity className="flex-1 min-w-[45%] bg-white p-5 rounded-2xl items-center">
            <CheckCircle size={24} color="#10b981" />
            <Text className="text-xs font-semibold text-slate-900 mt-2">
              Buat Tugas
            </Text>
          </TouchableOpacity>
          <TouchableOpacity className="flex-1 min-w-[45%] bg-white p-5 rounded-2xl items-center">
            <Users size={24} color="#8b5cf6" />
            <Text className="text-xs font-semibold text-slate-900 mt-2">
              Kelola Kelas
            </Text>
          </TouchableOpacity>
          <TouchableOpacity className="flex-1 min-w-[45%] bg-white p-5 rounded-2xl items-center">
            <TrendingUp size={24} color="#f59e0b" />
            <Text className="text-xs font-semibold text-slate-900 mt-2">
              Analisis
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </ScrollView>
  );
}
