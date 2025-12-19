import { View, Text, TouchableOpacity, Image, Alert } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import AuthInput from "../../components/auth/AuthInput";
import { AuthStackParamList } from "../../navigation/AuthNavigator";
import { useState } from "react";
import { useAuthActions } from "@/hooks/useAuthAction";

type Props = NativeStackScreenProps<AuthStackParamList, "Register">;

const LOGO = require("@/assets/images/logo-remove-bg.png");

export default function RegisterScreen({ navigation }: Props) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmpassword, setConfirmPassword] = useState("");
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

      if (password !== confirmpassword) {
        Alert.alert("terjadi kesalahan", "Password tidak sama");
        return;
      }

      await registerReq({
        name,
        username,
        email,
        password,
        confirmpassword,
      });

      Alert.alert("Berhasil", "Berhasil register, silahkan login");
      clearForm();
    } catch (err: any) {
      Alert.alert(
        "Register gagal",
        err?.response?.data?.message ?? "Terjadi kesalahan",
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-10">
        <Image source={LOGO} className="w-48 h-48 mb-4" resizeMode="contain" />
        <Text className="text-xl font-bold text-slate-800">
          Daftar ke Edukarsa
        </Text>
        <Text className="text-slate-500 mt-2">
          Memulai itu sulit, namun sangat penting!
        </Text>
      </View>

      <AuthInput
        icon="user"
        placeholder="Nama Lengkap"
        value={name}
        onChangeText={setName}
      />
      <AuthInput
        icon="user"
        placeholder="Username"
        value={username}
        onChangeText={setUsername}
      />
      <AuthInput
        icon="envelope"
        placeholder="Email"
        value={email}
        onChangeText={setEmail}
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
        value={confirmpassword}
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
          {loading ? "Loading..." : "Register"}
        </Text>
      </TouchableOpacity>

      <TouchableOpacity
        className="mt-6"
        onPress={() => navigation.navigate("Login")}
      >
        <Text className="text-center text-slate-600">
          Sudah punya akun? <Text className="text-blue-500">Login</Text>
        </Text>
      </TouchableOpacity>
    </View>
  );
}
