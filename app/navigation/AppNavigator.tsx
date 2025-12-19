import HomeScreen from "@/screens/Home/HomeScreen";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

export type DashboardStackParamList = {
  Home: undefined;
};

const Stack = createNativeStackNavigator<DashboardStackParamList>();

export default function AppNavigator() {
  return (
    <Stack.Navigator>
      <Stack.Screen name="Home" component={HomeScreen} />
    </Stack.Navigator>
  );
}
