"use server";

import z from "zod";
import { SignUpSchema } from "@/lib/zodSchema";

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
