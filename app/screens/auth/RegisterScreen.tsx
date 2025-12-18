import { View, Text, TouchableOpacity, Image } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import AuthInput from "../../components/auth/AuthInput";
import { AuthStackParamList } from "../../navigation/AuthNavigator";

type Props = NativeStackScreenProps<AuthStackParamList, "Register">;

export default function RegisterScreen({ navigation }: Props) {
  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-10">
        <Image
          source={require("../../assets/images/icon.png")}
          className="w-24 h-24 mb-4"
          resizeMode="contain"
        />
        <Text className="text-3xl font-bold text-slate-800">Daftar Akun</Text>
        <Text className="text-slate-500 mt-2">Buat akun EduKarsa</Text>
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
