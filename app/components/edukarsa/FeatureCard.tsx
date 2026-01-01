import React from "react";
import { Text, View } from "react-native";

interface FeatureCardProps {
  icon: string;
  title: string;
  description: string;
}

export default function FeatureCard({
  icon,
  title,
  description,
}: FeatureCardProps) {
  return (
    <View className="bg-blue-50 rounded-2xl p-5 mb-4 border-l-4 border-blue-600">
      <View className="flex-row items-center mb-3">
        <Text className="text-3xl mr-3">{icon}</Text>
        <Text className="text-lg font-bold text-blue-900 flex-1">{title}</Text>
      </View>
      <Text className="text-sm text-gray-600 leading-5">{description}</Text>
    </View>
  );
}
