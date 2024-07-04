"use server";

import {
  AuthMidwifeLoginInput,
  AuthMidwifeOutput,
  AuthMidwifeSignUpInput,
  CreatePatientInput,
  Midwife,
  Patient,
} from "@/lib/models";
import axios from "axios";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

export async function SignUp(
  signUpInput: AuthMidwifeSignUpInput
): Promise<AuthMidwifeOutput> {
  const url = `${process.env.BACKEND_URL}/api/auth/signup`;
  const res = await axios.post(url, signUpInput);
  if (res.status !== 201) {
    throw new Error(res.data.error);
  }
  cookies().set("token", res.data.token, {
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    httpOnly: true,
  });
  cookies().set("profileData", JSON.stringify(res.data.midwife as Midwife), {
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    httpOnly: true,
  });
  redirect("/dashboard");
  return res.data.midwife;
}

// TODO: Store token in server
export async function Login(
  loginInput: AuthMidwifeLoginInput
): Promise<string | undefined> {
  const url = `${process.env.BACKEND_URL}/api/auth/login`;
  const res = await axios.post(url, loginInput);
  if (res.status !== 200) {
    throw new Error(res.data.error);
  }
  cookies().set("token", res.data.token, {
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    httpOnly: true,
  });
  return res.data.token;
}

export async function GetProfile(token: string): Promise<Midwife> {
  const url = `${process.env.BACKEND_URL}/api/auth/profile`;
  const res = await axios.get(url, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (res.status !== 200) {
    throw new Error("Failed to fetch profile");
  }
  cookies().set("profileData", JSON.stringify(res.data.midwife as Midwife), {
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    httpOnly: true,
  });
  return res.data.midwife;
}

export async function GetPatients(token: string): Promise<Patient[]> {
  const url = `${process.env.BACKEND_URL}/api/patients`;
  const res = await axios.get(url, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (res.status !== 200) {
    throw new Error(res.data.error);
  }
  const patients: Patient[] = await res.data.patients;
  return patients;
}

export async function CreatePatient(patient: CreatePatientInput) {
  const midwifeData = cookies().get("profileData")?.value;
  if (!midwifeData) {
    throw new Error("Failed to get midwife data");
  }
  const midwife = JSON.parse(midwifeData);
  const token = cookies().get("token")?.value;
  if (!token) {
    throw new Error("Failed to get token");
  }
  patient.midwifeId = midwife.id;
  const url = `${process.env.BACKEND_URL}/api/patient`;
  const res = await axios.post(url, patient, {
    headers: { Authorization: `Bearer ${token}` },
  });
  if (res.status !== 201) {
    throw new Error(res.data.error);
  }
  redirect("/dashboard");
}

export async function UpdatePatient(patient: CreatePatientInput, id: number) {
  const midwifeData = cookies().get("profileData")?.value;
  if (!midwifeData) {
    throw new Error("Failed to get midwife data");
  }
  const midwife = JSON.parse(midwifeData);
  const token = cookies().get("token")?.value;
  if (!token) {
    throw new Error("Failed to get token");
  }
  const url = `${process.env.BACKEND_URL}/api/patient`;
  patient.midwifeId = midwife.id;
  const obj = { ...patient, id };
  console.log(obj);
  const res = await axios.patch(
    url,
    { ...obj },
    {
      headers: { Authorization: `Bearer ${token}` },
    }
  );
  if (res.status !== 201) {
    throw new Error(res.data.error);
  }
}
