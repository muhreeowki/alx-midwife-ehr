import { z } from "zod";

export const SignUpSchema = z.object({
  firstName: z.string({ required_error: "First name is required" }).min(2),
  lastName: z.string({ required_error: "First name is required" }).min(2),
  email: z.string().email({ message: "Please provide a valid email" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters" }),
});

export const LoginSchema = z.object({
  email: z.string({ required_error: "Please provide an email" }).email(),
  password: z.string(),
});
