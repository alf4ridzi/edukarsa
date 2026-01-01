import axios from "axios";
import * as SecureStore from "expo-secure-store";
import React, { createContext, useContext, useEffect, useState } from "react";

type User = {
  name: string;
  email: string;
  role?: {
    id?: number | bigint;
    name: string;
  };
};

type AuthContextType = {
  isAuthenticated: boolean;
  isLoading: boolean;
  user: User | null;
  login: (token: string, userData: User) => Promise<void>;
  logout: () => Promise<void>;
  refreshUser: () => Promise<void>;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [user, setUser] = useState<User | null>(null);

  const fetchUserData = async (token: string) => {
    try {
      const response = await axios.get(
        `${process.env.EXPO_PUBLIC_API_URL}/users/me`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      console.log("=== API Response /users/me ===");
      console.log("Full response:", JSON.stringify(response.data, null, 2));
      const userData = response.data.data.user;
      console.log("User data extracted:", JSON.stringify(userData, null, 2));
      console.log("User role:", userData?.role);
      console.log("User role name:", userData?.role?.name);

      await SecureStore.setItemAsync("user_data", JSON.stringify(userData));
      setUser(userData);
      return userData;
    } catch (error) {
      console.error("Error fetching user data:", error);
      throw error;
    }
  };

  useEffect(() => {
    const bootstrap = async () => {
      try {
        const isAvailable = await SecureStore.isAvailableAsync();
        if (isAvailable) {
          const token = await SecureStore.getItemAsync("access_token");
          if (token) {
            await fetchUserData(token);
            setIsAuthenticated(true);
          }
        }
      } catch (error) {
        console.error("Error checking authentication:", error);
        await SecureStore.deleteItemAsync("access_token");
        await SecureStore.deleteItemAsync("user_data");
        setIsAuthenticated(false);
      } finally {
        setIsLoading(false);
      }
    };
    bootstrap();
  }, []);

  const login = async (token: string, userData: User) => {
    await SecureStore.setItemAsync("access_token", token);
    try {
      await fetchUserData(token);
    } catch (error) {
      await SecureStore.setItemAsync("user_data", JSON.stringify(userData));
      setUser(userData);
    }
    setIsAuthenticated(true);
  };

  const refreshUser = async () => {
    const token = await SecureStore.getItemAsync("access_token");
    if (token) {
      await fetchUserData(token);
    }
  };

  const logout = async () => {
    await SecureStore.deleteItemAsync("access_token");
    await SecureStore.deleteItemAsync("user_data");
    setUser(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider
      value={{ isAuthenticated, isLoading, user, login, logout, refreshUser }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used inside AuthProvider");
  }
  return context;
};
