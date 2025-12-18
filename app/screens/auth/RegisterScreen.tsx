import { View, Text, TouchableOpacity, Image } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import AuthInput from "../../components/auth/AuthInput";
import { AuthStackParamList } from "../../navigation/AuthNavigator";

type Props = NativeStackScreenProps<AuthStackParamList, "Register">;

const LOGO = require("@/assets/images/logo-remove-bg.png");

export default function RegisterScreen({ navigation }: Props) {
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

      <AuthInput icon="user" placeholder="Nama Lengkap" />
      <AuthInput icon="envelope" placeholder="Email" />
      <AuthInput icon="lock" placeholder="Password" secureTextEntry />
      <AuthInput
        icon="lock"
        placeholder="Konfirmasi Password"
        secureTextEntry
      />

      <TouchableOpacity className="bg-blue-500 py-4 rounded-full mt-4">
        <Text className="text-center text-white text-lg font-semibold">
          Register
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
