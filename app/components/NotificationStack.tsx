import { Notification, useNotification } from "@/context/NotificationContext";
import { AlertCircle, CheckCircle, Info, XCircle } from "lucide-react-native";
import { useEffect, useRef } from "react";
import { Animated, Text, TouchableOpacity, View } from "react-native";

export default function NotificationStack() {
  const { notifications, removeNotification } = useNotification();

  return (
    <View className="absolute top-0 left-0 right-0 z-50 pointer-events-none">
      {notifications.map((notif, index) => (
        <NotificationItem
          key={notif.id}
          notification={notif}
          onClose={() => removeNotification(notif.id)}
          index={index}
        />
      ))}
    </View>
  );
}

function NotificationItem({
  notification,
  onClose,
  index,
}: {
  notification: Notification;
  onClose: () => void;
  index: number;
}) {
  const slideAnim = useRef(new Animated.Value(-100)).current;

  useEffect(() => {
    Animated.timing(slideAnim, {
      toValue: 0,
      duration: 300,
      useNativeDriver: true,
    }).start();
  }, [slideAnim]);

  const getIcon = () => {
    switch (notification.type) {
      case "success":
        return <CheckCircle size={20} color="#10b981" />;
      case "error":
        return <XCircle size={20} color="#ef4444" />;
      case "warning":
        return <AlertCircle size={20} color="#f59e0b" />;
      case "info":
      default:
        return <Info size={20} color="#3b82f6" />;
    }
  };

  const getBackgroundColor = () => {
    switch (notification.type) {
      case "success":
        return "bg-green-50 border-l-4 border-green-500";
      case "error":
        return "bg-red-50 border-l-4 border-red-500";
      case "warning":
        return "bg-amber-50 border-l-4 border-amber-500";
      case "info":
      default:
        return "bg-blue-50 border-l-4 border-blue-500";
    }
  };

  const getTextColor = () => {
    switch (notification.type) {
      case "success":
        return "text-green-900";
      case "error":
        return "text-red-900";
      case "warning":
        return "text-amber-900";
      case "info":
      default:
        return "text-blue-900";
    }
  };

  return (
    <Animated.View
      style={{
        transform: [{ translateY: slideAnim }],
        marginTop: index === 0 ? 12 : 8,
      }}
    >
      <TouchableOpacity
        activeOpacity={0.7}
        onPress={onClose}
        className={`mx-3 p-4 rounded-lg flex-row items-start gap-3 ${getBackgroundColor()}`}
      >
        <View className="pt-0.5">{getIcon()}</View>
        <View className="flex-1">
          <Text className={`text-sm font-semibold ${getTextColor()}`}>
            {notification.title}
          </Text>
          <Text className={`text-xs mt-1 ${getTextColor()} opacity-80`}>
            {notification.message}
          </Text>
        </View>
        <TouchableOpacity onPress={onClose} className="pt-0.5">
          <Text className={`text-lg font-bold ${getTextColor()}`}>×</Text>
        </TouchableOpacity>
      </TouchableOpacity>
    </Animated.View>
  );
}
