import { View, Text, TouchableOpacity, Image } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import AuthInput from "../../components/auth/AuthInput";
import { AuthStackParamList } from "@/navigation/AuthNavigator";

type Props = NativeStackScreenProps<AuthStackParamList, "Login">;

export default function LoginScreen({ navigation }: Props) {
  return (
    <View className="flex-1 bg-white px-8 justify-center">
      <View className="items-center mb-10">
        <Image
          source={require("../../assets/images/icon.png")}
          className="w-24 h-24 mb-4"
          resizeMode="contain"
        />
        <Text className="text-3xl font-bold text-slate-800">EduKarsa</Text>
        <Text className="text-slate-500 mt-2">Selamat Datang</Text>
      </View>

      <AuthInput icon="user" placeholder="Username / NIS / Email" />
      <AuthInput icon="lock" placeholder="Password" secureTextEntry />

      <TouchableOpacity className="bg-blue-500 py-4 rounded-full mt-4">
        <Text className="text-center text-white text-lg font-semibold">
          Login
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
