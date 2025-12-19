import { View, Text, Button, Alert } from "react-native";
import { useAuthActions } from "@/hooks/useAuthAction";

export default function HomeScreen() {
  const { logoutHandle } = useAuthActions();

  const LogoutAlert = () => {
    Alert.alert("Peringatan", "Yakin keluar?", [
      {
        text: "Cancel",
        style: "cancel",
      },
      {
        text: "Keluar",
        onPress: () => handleLogout(),
        style: "destructive",
      },
    ]);
  };

  const handleLogout = async () => {
    await logoutHandle();

    Alert.alert("Berhasil", "berhasil logout!");
  };

  return (
    <View>
      <Text className="text-black dark:text-white">Hi there</Text>
      <Button onPress={LogoutAlert} title="logout" />
    </View>
  );
}
