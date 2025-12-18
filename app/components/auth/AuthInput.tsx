import { View, TextInput } from "react-native";
import { FontAwesome } from "@expo/vector-icons";
import { ComponentProps } from "react";

type AuthInputProps = {
  icon: ComponentProps<typeof FontAwesome>["name"];
  placeholder?: string;
  secureTextEntry?: boolean;
};

export default function AuthInput({
  icon,
  placeholder,
  secureTextEntry,
}: AuthInputProps) {
  return (
    <View className="flex-row items-center border-b border-blue-500 mb-6">
      <FontAwesome name={icon} size={18} color="#2563eb" />
      <TextInput
        placeholder={placeholder}
        secureTextEntry={secureTextEntry}
        className="flex-1 ml-3 py-2 text-base"
      />
    </View>
  );
}
