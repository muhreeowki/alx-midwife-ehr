"use server";

import { User } from "@/context/authContext";
import axios from "axios";
import exp from "constants";
import { cookies } from "next/headers";
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

// TODO: Store token in server
export async function Login(data: { email: string; password: string }) {
  const url = `${process.env.BACKEND_URL}/api/auth/login`;
  const res = await axios.post(url, data);
  if (res.status === 200) {
    cookies().set("token", res.data.token, {
      expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
      httpOnly: true,
    });
    return res.data.token;
  }
  return undefined;
}

// TODO: Store profile in server
export async function GetProfile(token: string) {
  const url = `${process.env.BACKEND_URL}/api/auth/profile`;
  const res = await axios.get(url, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (res.status === 200) {
    // TODO: Encrypt profile data
    cookies().set("profileData", JSON.stringify(res.data.midwife as User), {
      expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
      httpOnly: true,
    });
    return res.data.midwife;
  }
  return undefined;
}

export async function getPatients(token: string) {
  try {
    const url = `${process.env.BACKEND_URL}/api/patients`;
    const res = await axios.get(url, {
      headers: { Authorization: `Bearer ${token}` },
    });
    if (res.status === 200) {
      const patients = await res.data.patients;
      return patients;
    }
    return undefined;
  } catch (error) {
    console.error(error);
  }
}
