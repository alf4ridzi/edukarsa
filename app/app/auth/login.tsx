import AuthInput from "@/components/auth/AuthInput";
import { useAuthActions } from "@/hooks/useAuthAction";
import { router } from "expo-router";
import { useState } from "react";
import { Alert, Image, Text, TouchableOpacity, View } from "react-native";

const LOGO = require("@/assets/images/logo-remove-bg.png");

export default function LoginScreen() {
  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const { loginReq } = useAuthActions();

  function clearForm() {
    setIdentifier("");
    setPassword("");
  }

  const handleLogin = async () => {
    try {
      setLoading(true);

      await loginReq({
        identifier,
        password,
      });

      Alert.alert("Berhasil", "berhasil login!");
      clearForm();
      router.replace("/(tabs)");
    } catch (err: any) {
      Alert.alert("Login gagal", err?.response?.data?.message || err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-10">
        <Image source={LOGO} className="w-48 h-48 mb-4" resizeMode="contain" />
        <Text className="text-xl font-bold text-slate-800">
          Login ke Edukarsa
        </Text>
        <Text className="text-slate-500 mt-2">
          Yuk buat dunia lebih baik bersama-sama!
        </Text>
      </View>

      <AuthInput
        icon="user"
        placeholder="Username / Email"
        value={identifier}
        onChangeText={setIdentifier}
      />

      <AuthInput
        icon="lock"
        placeholder="Password"
        secureTextEntry
        value={password}
        onChangeText={setPassword}
      />

      <TouchableOpacity
        onPress={handleLogin}
        disabled={loading}
        className={`bg-blue-500 py-4 rounded-full mt-4 ${
          loading ? "opacity-60" : ""
        }`}
      >
        <Text className="text-center text-white text-lg font-semibold">
          {loading ? "Loading..." : "Login"}
        </Text>
      </TouchableOpacity>

      <TouchableOpacity
        className="mt-6"
        onPress={() => router.push("/auth/register")}
      >
        <Text className="text-center text-slate-600">
          Belum punya akun? <Text className="text-blue-500">Daftar</Text>
        </Text>
      </TouchableOpacity>
    </View>
  );
}
