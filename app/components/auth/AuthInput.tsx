import { View, TextInput, TouchableOpacity } from "react-native";
import { FontAwesome } from "@expo/vector-icons";
import { ComponentProps, useState } from "react";

type AuthInputProps = {
  icon: ComponentProps<typeof FontAwesome>["name"];
  placeholder?: string;
  secureTextEntry?: boolean;
  value?: string;
  onChangeText?: (text: string) => void;
};

export default function AuthInput({
  icon,
  placeholder,
  secureTextEntry = false,
  value,
  onChangeText,
}: AuthInputProps) {
  const [hidePassword, setHidePassword] = useState(secureTextEntry);
  const [isFocused, setIsFocused] = useState(false);

  return (
    <View
      className={`flex-row items-center bg-gray-100 rounded-xl px-4 py-3 mb-4 ${
        isFocused ? "border-2 border-blue-500" : "border-2 border-transparent"
      }`}
    >
      <FontAwesome
        name={icon}
        size={20}
        color={isFocused ? "#3b82f6" : "#9ca3af"}
      />
      <TextInput
        value={value}
        onChangeText={onChangeText}
        placeholder={placeholder}
        placeholderTextColor="#9ca3af"
        secureTextEntry={hidePassword}
        autoCapitalize="none"
        autoCorrect={false}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        className="flex-1 ml-3 text-base text-gray-800"
      />
      {secureTextEntry && (
        <TouchableOpacity
          onPress={() => setHidePassword(!hidePassword)}
          className="ml-2"
        >
          <FontAwesome
            name={hidePassword ? "eye-slash" : "eye"}
            size={20}
            color={isFocused ? "#3b82f6" : "#9ca3af"}
          />
        </TouchableOpacity>
      )}
    </View>
  );
}
