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

  return (
    <View className="flex-row items-center border-b border-blue-500 mb-6">
      <FontAwesome name={icon} size={18} color="#2563eb" />

      <TextInput
        value={value}
        onChangeText={onChangeText}
        placeholder={placeholder}
        secureTextEntry={hidePassword}
        autoCapitalize="none"
        autoCorrect={false}
        className="flex-1 ml-3 py-2 text-base"
      />

      {secureTextEntry && (
        <TouchableOpacity onPress={() => setHidePassword(!hidePassword)}>
          <FontAwesome
            name={hidePassword ? "eye-slash" : "eye"}
            size={18}
            color="#2563eb"
          />
        </TouchableOpacity>
      )}
    </View>
  );
}
