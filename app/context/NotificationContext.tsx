import React, { createContext, useCallback, useContext, useState } from "react";

export type Notification = {
  id: string;
  title: string;
  message: string;
  type: "success" | "error" | "info" | "warning";
  timestamp: Date;
};

type NotificationContextType = {
  notifications: Notification[];
  addNotification: (
    title: string,
    message: string,
    type?: "success" | "error" | "info" | "warning"
  ) => void;
  removeNotification: (id: string) => void;
};

const NotificationContext = createContext<NotificationContextType | undefined>(
  undefined
);

export const NotificationProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [notifications, setNotifications] = useState<Notification[]>([]);

  const removeNotification = useCallback((id: string) => {
    setNotifications((prev) => prev.filter((notif) => notif.id !== id));
  }, []);

  const addNotification = useCallback(
    (
      title: string,
      message: string,
      type: "success" | "error" | "info" | "warning" = "info"
    ) => {
      const id = Date.now().toString();
      const newNotification: Notification = {
        id,
        title,
        message,
        type,
        timestamp: new Date(),
      };

      setNotifications((prev) => [newNotification, ...prev]);

      // Auto remove after 5 seconds
      setTimeout(() => {
        removeNotification(id);
      }, 5000);
    },
    [removeNotification]
  );

  return (
    <NotificationContext.Provider
      value={{ notifications, addNotification, removeNotification }}
    >
      {children}
    </NotificationContext.Provider>
  );
};

export const useNotification = () => {
  const context = useContext(NotificationContext);
  if (!context) {
    throw new Error("useNotification must be used inside NotificationProvider");
  }
  return context;
};
