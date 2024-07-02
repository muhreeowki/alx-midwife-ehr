"use client";

import Link from "next/link";
import * as React from "react";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { SignUpSchema, LoginSchema } from "@/lib/zodSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import SignUpForm from "./SignUpForm";
import LoginForm from "./LoginForm";

const AuthForm = () => {
  const [authType, setAuthType] = React.useState<"signUp" | "login">("signUp");

  type SignUpInputs = z.infer<typeof SignUpSchema>;
  type LoginInputs = z.infer<typeof LoginSchema>;

  const signUpForm = useForm<SignUpInputs>({
    resolver: zodResolver(SignUpSchema),
    defaultValues: {
      first_name: "",
      last_name: "",
      email: "",
      password: "",
    },
  });

  const loginForm = useForm<LoginInputs>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  return (
    <main className="flex items-center justify-center h-screen">
      {authType === "signUp" ? (
        <SignUpForm form={signUpForm} setAuthType={setAuthType} />
      ) : (
        <LoginForm form={loginForm} setAuthType={setAuthType} />
      )}
    </main>
  );
};

export default AuthForm;
