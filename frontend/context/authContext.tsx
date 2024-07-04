"use client";
import * as React from "react";
import { Midwife } from "@/lib/models";

export interface AuthContextProps {
  user: Midwife | null;
  login: (user: Midwife) => void;
  logout: () => void;
}

const AuthContext = React.createContext<AuthContextProps>({
  user: null,
  login: () => {},
  logout: () => {},
});

const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = React.useState<Midwife | null>(null);

  const login = (user: Midwife) => {
    setUser(user);
  };

  const logout = () => {
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
