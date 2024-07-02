import Link from "next/link";
import * as React from "react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

import {
  Form,
  FormControl,
  FormDescription,
  FormItem,
  FormLabel,
  FormMessage,
  FormField,
} from "@/components/ui/form";

import { UseFormReturn } from "react-hook-form";
import z from "zod";
import { SignUpSchema } from "@/lib/zodSchema";
import { onSignUpSubmit } from "@/app/actions";

const SignUpForm = ({
  form,
  setSignUpOrLogin,
}: {
  form: UseFormReturn<
    {
      first_name: string;
      last_name: string;
      email: string;
      password: string;
    },
    any,
    undefined
  >;
  setSignUpOrLogin: React.Dispatch<React.SetStateAction<"signUp" | "login">>;
}) => {
  return (
    <Form {...form}>
      <form
        action="#"
        onSubmit={form.handleSubmit((data: z.infer<typeof SignUpSchema>) =>
          onSignUpSubmit({
            email: data.email,
            password: data.password,
            first_name: data.first_name,
            last_name: data.last_name,
          })
        )}
        className="space-y-8"
      >
        <Card className="mx-auto max-w-sm">
          <CardHeader>
            <CardTitle className="text-xl">Sign Up</CardTitle>
            <CardDescription>
              Enter your information to create an account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="grid gap-2">
                  <FormField
                    control={form.control}
                    name="first_name"
                    render={({ field, fieldState, formState }) => (
                      <FormItem>
                        <FormLabel>First name</FormLabel>
                        <FormControl>
                          <Input placeholder="Mary" required {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <div className="grid gap-2">
                  <FormField
                    control={form.control}
                    name="last_name"
                    render={({ field, fieldState, formState }) => (
                      <FormItem>
                        <FormLabel>Last Name</FormLabel>
                        <FormControl>
                          <Input placeholder="Robinson" required {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
              <div className="grid gap-2">
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field, fieldState, formState }) => (
                    <FormItem>
                      <FormLabel>Email</FormLabel>
                      <FormControl>
                        <Input
                          placeholder="name@example.com"
                          type="email"
                          required
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="grid gap-2">
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field, fieldState, formState }) => (
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input type="password" required {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <Button type="submit" className="w-full">
                Create an account
              </Button>
              <Button variant="outline" className="w-full">
                Sign up with GitHub
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Already have an account?{" "}
              <Link
                href="#"
                className="underline"
                onClick={() => setSignUpOrLogin("login")}
              >
                Login
              </Link>
            </div>
          </CardContent>
        </Card>
      </form>
    </Form>
  );
};

export default SignUpForm;
