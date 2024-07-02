"use server";

import axios from "axios";
import { redirect } from "next/navigation";

export async function SignUp(data: {
  first_name: string;
  last_name: string;
  email: string;
  password: string;
}) {
  const url = `${process.env.BACKEND_URL}/api/auth/signup`;
  const res = await axios.post(url, data);
  if (res.status !== 201) {
    return undefined;
  } else {
    if (res) {
      redirect("/");
    }
  }
}

export async function Login(data: { email: string; password: string }) {
  const url = `${process.env.BACKEND_URL}/api/auth/login`;
  const res = await axios.post(url, data);
  if (res.status === 200) {
    return res.data;
  }
  return undefined;
}
