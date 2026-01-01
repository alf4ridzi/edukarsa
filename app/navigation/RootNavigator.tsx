import { useAuth } from "@/context/AuthContext";
import { Stack } from "expo-router";
import AuthNavigator from "./AuthNavigator";

export default function RootNavigator() {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) return null;

  if (!isAuthenticated) {
    return <AuthNavigator />;
  }

  return (
    <Stack screenOptions={{ headerShown: false }}>
      <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
      <Stack.Screen
        name="TeacherDashboard"
        options={{
          headerShown: true,
          headerTitle: "Dashboard Guru",
          headerTintColor: "#0066FF",
          headerTitleStyle: {
            fontWeight: "bold",
          },
        }}
      />
    </Stack>
  );
}
