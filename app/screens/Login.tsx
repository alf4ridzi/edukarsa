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

const LoginScreen = ({ navigation }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);

  const handleLogin = () => {
    console.log("Login dengan:", email, password);
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
        <View className="flex-1 px-6 justify-center">
          <View className="items-center mb-12">
            <Image
              source={require("../assets/logo.png")}
              className="w-32 h-32 mb-6"
              resizeMode="contain"
            />
            <Text className="text-3xl font-bold text-gray-800 mb-2">
              EduKarsa
            </Text>
            <Text className="text-lg text-gray-600">Selamat Datang</Text>
          </View>

          <View className="mb-6">
            {/* Email Input */}
            <View className="mb-6">
              <Text className="text-gray-700 mb-2 text-base">
                Username/NIS/Email
              </Text>
              <View className="flex-row items-center border-b-2 border-blue-500 pb-2">
                <FontAwesome name="user" size={20} color="#6B7280" />
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

            <View className="mb-8">
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

            <TouchableOpacity
              className="bg-blue-500 rounded-full py-4 items-center mb-4"
              onPress={handleLogin}
              activeOpacity={0.8}
            >
              <Text className="text-white text-lg font-semibold">Login</Text>
            </TouchableOpacity>

            <TouchableOpacity className="items-center mb-4">
              <Text className="text-blue-500 text-sm">Lupa Password?</Text>
            </TouchableOpacity>

            <View className="flex-row justify-center items-center">
              <Text className="text-gray-600 text-sm">Belum punya akun? </Text>
              <TouchableOpacity onPress={() => navigation.navigate("Register")}>
                <Text className="text-blue-500 text-sm font-semibold">
                  Daftar Sekarang
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </ScrollView>
    </KeyboardAvoidingView>
  );
};

export default LoginScreen;
