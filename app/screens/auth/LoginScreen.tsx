import { View, Text, TouchableOpacity, Image, Alert } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import AuthInput from "@/components/auth/AuthInput";
import { AuthStackParamList } from "@/navigation/AuthNavigator";
import { useState } from "react";
import { useAuthActions } from "@/hooks/useAuthAction";

type Props = NativeStackScreenProps<AuthStackParamList, "Login">;

const LOGO = require("@/assets/images/logo-remove-bg.png");

export default function LoginScreen({ navigation }: Props) {
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
      // navigation.replace("Home");
    } catch (err: any) {
      Alert.alert("Login gagal", err?.response?.data?.message || err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-10">
        <Image source={LOGO} className="w-48 h-48  mb-4" resizeMode="contain" />
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
        onPress={() => navigation.navigate("Register")}
      >
        <Text className="text-center text-slate-600">
          Belum punya akun? <Text className="text-blue-500">Daftar</Text>
        </Text>
      </TouchableOpacity>
    </View>
  );
}
