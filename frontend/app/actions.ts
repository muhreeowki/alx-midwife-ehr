"use server";

import z from "zod";
import { SignUpSchema } from "@/lib/zodSchema";

import axios from "axios";

export const onSignUpSubmit = async (data: {
  first_name: string;
  last_name: string;
  email: string;
  password: string;
}) => {
  const url = `${process.env.BACKEND_URL}/api/auth/signup`;
  const res = await axios.post(url, data);
  console.log(res);
};
