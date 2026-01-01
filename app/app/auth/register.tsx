import AuthInput from "@/components/auth/AuthInput";
import { useAuthActions } from "@/hooks/useAuthAction";
import { router } from "expo-router";
import { useState } from "react";
import { Alert, Image, Text, TouchableOpacity, View } from "react-native";

const LOGO = require("@/assets/images/logo-remove-bg.png");

export default function RegisterScreen() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const { registerReq } = useAuthActions();

  function clearForm() {
    setName("");
    setEmail("");
    setUsername("");
    setPassword("");
    setConfirmPassword("");
  }

  const handleRegister = async () => {
    try {
      setLoading(true);

      await registerReq({
        name,
        email,
        username,
        password,
        confirmpassword: confirmPassword,
      });

      Alert.alert("Berhasil", "Registrasi berhasil! Silakan login.", [
        { text: "OK", onPress: () => router.replace("/auth/login") },
      ]);
      clearForm();
    } catch (err: any) {
      Alert.alert(
        "Registrasi gagal",
        err?.response?.data?.message || err.message
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-8">
        <Image source={LOGO} className="w-32 h-32 mb-3" resizeMode="contain" />
        <Text className="text-xl font-bold text-slate-800">
          Daftar ke Edukarsa
        </Text>
        <Text className="text-slate-500 mt-2">
          Mari bergabung dengan komunitas kami!
        </Text>
      </View>

      <AuthInput
        icon="user"
        placeholder="Nama Lengkap"
        value={name}
        onChangeText={setName}
      />

      <AuthInput
        icon="mail"
        placeholder="Email"
        value={email}
        onChangeText={setEmail}
        keyboardType="email-address"
      />

      <AuthInput
        icon="user"
        placeholder="Username"
        value={username}
        onChangeText={setUsername}
      />

      <AuthInput
        icon="lock"
        placeholder="Password"
        secureTextEntry
        value={password}
        onChangeText={setPassword}
      />

      <AuthInput
        icon="lock"
        placeholder="Konfirmasi Password"
        secureTextEntry
        value={confirmPassword}
        onChangeText={setConfirmPassword}
      />

      <TouchableOpacity
        onPress={handleRegister}
        disabled={loading}
        className={`bg-blue-500 py-4 rounded-full mt-4 ${
          loading ? "opacity-60" : ""
        }`}
      >
        <Text className="text-center text-white text-lg font-semibold">
          {loading ? "Loading..." : "Daftar"}
        </Text>
      </TouchableOpacity>

      <TouchableOpacity className="mt-6" onPress={() => router.back()}>
        <Text className="text-center text-slate-600">
          Sudah punya akun? <Text className="text-blue-500">Login</Text>
        </Text>
      </TouchableOpacity>
    </View>
  );
}
