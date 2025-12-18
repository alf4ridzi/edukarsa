import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  Image,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
} from "react-native";
import { FontAwesome } from "@expo/vector-icons";

const RegisterScreen = ({ navigation }) => {
  const [fullName, setFullName] = useState("");
  const [email, setEmail] = useState("");
  const [nis, setNis] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const handleRegister = () => {
    if (password !== confirmPassword) {
      alert("Password tidak cocok!");
      return;
    }
    console.log("Register dengan:", { fullName, email, nis, password });
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      className="flex-1 bg-white"
    >
      <ScrollView
        contentContainerStyle={{ flexGrow: 1 }}
        keyboardShouldPersistTaps="handled"
      >
        <View className="flex-1 px-6 pt-12 pb-6">
          <View className="items-center mb-8">
            <Image
              source={require("../assets/logo.png")}
              className="w-24 h-24 mb-4"
              resizeMode="contain"
            />
            <Text className="text-3xl font-bold text-gray-800 mb-2">
              EduKarsa
            </Text>
            <Text className="text-lg text-gray-600">Daftar Akun Baru</Text>
          </View>

          <View className="mb-6">
            <View className="mb-5">
              <Text className="text-gray-700 mb-2 text-base">Nama Lengkap</Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="user" size={20} color="#6B7280" />
                <TextInput
                  className="flex-1 ml-3 text-base text-gray-800"
                  placeholder="Masukkan nama lengkap"
                  value={fullName}
                  onChangeText={setFullName}
                />
              </View>
            </View>

            <View className="mb-5">
              <Text className="text-gray-700 mb-2 text-base">Email</Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="envelope" size={18} color="#6B7280" />
                <TextInput
                  className="flex-1 ml-3 text-base text-gray-800"
                  placeholder="email@gmail.com"
                  value={email}
                  onChangeText={setEmail}
                  keyboardType="email-address"
                  autoCapitalize="none"
                />
              </View>
            </View>

            <View className="mb-5">
              <Text className="text-gray-700 mb-2 text-base">NIS</Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="id-card" size={18} color="#6B7280" />
                <TextInput
                  className="flex-1 ml-3 text-base text-gray-800"
                  placeholder="Nomor Induk Siswa"
                  value={nis}
                  onChangeText={setNis}
                  keyboardType="numeric"
                />
              </View>
            </View>

            <View className="mb-5">
              <Text className="text-gray-700 mb-2 text-base">Password</Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="lock" size={20} color="#6B7280" />
                <TextInput
                  className="flex-1 ml-3 text-base text-gray-800"
                  placeholder="••••••••"
                  value={password}
                  onChangeText={setPassword}
                  secureTextEntry={!showPassword}
                />
                <TouchableOpacity
                  onPress={() => setShowPassword(!showPassword)}
                >
                  <FontAwesome
                    name={showPassword ? "eye" : "eye-slash"}
                    size={20}
                    color="#6B7280"
                  />
                </TouchableOpacity>
              </View>
            </View>

            <View className="mb-8">
              <Text className="text-gray-700 mb-2 text-base">
                Konfirmasi Password
              </Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="lock" size={20} color="#6B7280" />
                <TextInput
                  className="flex-1 ml-3 text-base text-gray-800"
                  placeholder="••••••••"
                  value={confirmPassword}
                  onChangeText={setConfirmPassword}
                  secureTextEntry={!showConfirmPassword}
                />
                <TouchableOpacity
                  onPress={() => setShowConfirmPassword(!showConfirmPassword)}
                >
                  <FontAwesome
                    name={showConfirmPassword ? "eye" : "eye-slash"}
                    size={20}
                    color="#6B7280"
                  />
                </TouchableOpacity>
              </View>
            </View>

            <TouchableOpacity
              className="bg-blue-500 rounded-full py-4 items-center mb-4"
              onPress={handleRegister}
              activeOpacity={0.8}
            >
              <Text className="text-white text-lg font-semibold">Daftar</Text>
            </TouchableOpacity>

            <View className="flex-row justify-center items-center">
              <Text className="text-gray-600 text-sm">Sudah punya akun? </Text>
              <TouchableOpacity onPress={() => navigation.navigate("Login")}>
                <Text className="text-blue-500 text-sm font-semibold">
                  Login Sekarang
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
};

export default RegisterScreen;
