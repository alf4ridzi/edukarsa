import {
  Award,
  Bell,
  ChevronRight,
  FileText,
  LogOut,
  Moon,
  Settings,
  Shield,
  User,
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

import { getUserProfile, UserProfile } from "@/api/user.api";
import { useAuthActions } from "@/hooks/useAuthAction";
import { router } from "expo-router";

export default function ProfileScreen() {
  const { logoutHandle } = useAuthActions();
  const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchUserProfile();
  }, []);

  const fetchUserProfile = async () => {
    try {
      setLoading(true);
      const response = await getUserProfile();
      setUserProfile(response.data.user);
    } catch (error: any) {
      Alert.alert(
        "Error",
        error?.response?.data?.message || "Gagal mengambil data profil"
      );
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    Alert.alert("Konfirmasi", "Yakin ingin keluar?", [
      {
        text: "Batal",
        style: "cancel",
      },
      {
        text: "Keluar",
        onPress: async () => {
          await logoutHandle();
          router.replace("/auth/login");
        },
        style: "destructive",
      },
    ]);
  };

  const menuItems = [
    {
      section: "Akun",
      items: [
        {
          icon: User,
          label: "Edit Profil",
          color: "#3b82f6",
          action: "edit-profile",
        },
        {
          icon: Shield,
          label: "Keamanan & Privasi",
          color: "#8b5cf6",
          action: "security",
        },
        {
          icon: Bell,
          label: "Notifikasi",
          color: "#f59e0b",
          action: "notifications",
        },
      ],
    },
    {
      section: "Pembelajaran",
      items: [
        {
          icon: Award,
          label: "Pencapaian & Badge",
          color: "#10b981",
          action: "achievements",
        },
        {
          icon: FileText,
          label: "Riwayat Aktivitas",
          color: "#06b6d4",
          action: "history",
        },
      ],
    },
    {
      section: "Pengaturan",
      items: [
        {
          icon: Moon,
          label: "Mode Gelap",
          color: "#64748b",
          action: "dark-mode",
        },
      ],
    },
  ];

  const stats = [
    { label: "Kelas", value: "4", color: "#3b82f6" },
    { label: "Siswa", value: "115", color: "#10b981" },
    { label: "Quiz", value: "48", color: "#f59e0b" },
  ];

  return (
    <View className="flex-1 bg-blue-100">
      <View className="p-5 pt-14 bg-white border-b border-gray-200 rounded-b-3xl">
        <Text className="text-2xl font-bold text-slate-900">Profil</Text>
      </View>

      {loading ? (
        <View className="flex-1 justify-center items-center p-5">
          <ActivityIndicator size="large" color="#3b82f6" />
          <Text className="mt-3 text-sm text-gray-600">
            Memuat data profil...
          </Text>
        </View>
      ) : (
        <ScrollView className="flex-1">
          <View className="bg-white m-5 p-6 rounded-2xl items-center">
            <View className="relative mb-4">
              <View className="w-24 h-24 rounded-full bg-blue-600 justify-center items-center">
                <User size={48} color="#ffffff" />
              </View>
              <TouchableOpacity className="absolute bottom-0 right-0 w-8 h-8 rounded-full bg-green-500 justify-center items-center border-[3px] border-white">
                <Settings size={16} color="#ffffff" />
              </TouchableOpacity>
            </View>

            <Text className="text-xl font-bold text-slate-900 mb-1">
              {userProfile?.name || "Nama Pengguna"}
            </Text>
            <Text className="text-sm text-gray-600 mb-1">
              {userProfile?.role.name === "student"
                ? "Siswa"
                : userProfile?.role.name === "teacher"
                  ? "Guru"
                  : "Admin"}
            </Text>
            <Text className="text-sm text-gray-600 mb-1">
              {userProfile?.email || "email@example.com"}
            </Text>
            <Text className="text-sm text-blue-600 mb-4">
              @{userProfile?.username || "username"}
            </Text>

            <View className="flex-row gap-8 pt-5 border-t border-gray-100">
              {stats.map((stat, index) => (
                <View key={index} className="items-center">
                  <Text
                    className="text-2xl font-bold"
                    style={{ color: stat.color }}
                  >
                    {stat.value}
                  </Text>
                  <Text className="text-xs text-gray-600 mt-1">
                    {stat.label}
                  </Text>
                </View>
              ))}
            </View>
          </View>

          {menuItems.map((section, sectionIndex) => (
            <View key={sectionIndex} className="px-5 mb-5">
              <Text className="text-base font-bold text-slate-900 mb-3">
                {section.section}
              </Text>
              <View className="bg-white rounded-2xl overflow-hidden shadow-2xl">
                {section.items.map((item, itemIndex) => (
                  <TouchableOpacity
                    key={itemIndex}
                    className={`flex-row justify-between items-center p-4 ${
                      itemIndex !== section.items.length - 1
                        ? "border-b border-gray-100"
                        : ""
                    }`}
                  >
                    <View className="flex-row items-center gap-3">
                      <View
                        className="w-10 h-10 rounded-full justify-center items-center"
                        style={{ backgroundColor: item.color + "20" }}
                      >
                        <item.icon size={20} color={item.color} />
                      </View>
                      <Text className="text-sm font-semibold text-slate-900">
                        {item.label}
                      </Text>
                    </View>
                    <ChevronRight size={20} color="#94a3b8" />
                  </TouchableOpacity>
                ))}
              </View>
            </View>
          ))}
          <View className="px-5 mb-5">
            <View className="rounded-xl p-4 items-center">
              <Text className="text-sm font-semibold text-slate-900 mb-1">
                Edukarsa v1.0.0
              </Text>
              <Text className="text-xs text-gray-600 mb-3">
                Platform Pendidikan Terpadu
              </Text>
            </View>
          </View>

          <TouchableOpacity
            className="flex-row items-center justify-center gap-2 bg-white mx-5 p-4 rounded-xl border border-red-100 mb-5"
            onPress={handleLogout}
          >
            <LogOut size={20} color="#ef4444" />
            <Text className="text-sm font-semibold text-red-500">Keluar</Text>
          </TouchableOpacity>

          <View className="p-5 items-center">
            <Text className="text-xs text-gray-400">
              © 2025 Edukarsa. All rights reserved.
            </Text>
          </View>
        </ScrollView>
      )}
    </View>
  );
}
