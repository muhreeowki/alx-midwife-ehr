import { Patient } from "@/lib/models";
import React from "react";
import Image from "next/image";
import Link from "next/link";
import {
  ChevronLeft,
  Home,
  LineChart,
  Package,
  Package2,
  PanelLeft,
  PlusCircle,
  Search,
  Settings,
  ShoppingCart,
  Upload,
  Users2,
} from "lucide-react";

import { Badge } from "@/components/ui/badge";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Textarea } from "@/components/ui/textarea";
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Slider } from "@/components/ui/slider";

import { CreatePatientInput } from "@/lib/models";
import { UseFormReturn } from "react-hook-form";
import { UpdatePatient } from "@/app/serverActions";
import { EmptyPatient } from "@/lib/models";
import { revalidatePath } from "next/cache";

const PatientView = ({
  patient,
  form,
  setSelectedPatient,
}: {
  patient: Patient;
  form: UseFormReturn<CreatePatientInput, any, undefined>;
  setSelectedPatient: React.Dispatch<React.SetStateAction<Patient>>;
}) => {
  const handleSubmit = async (data: CreatePatientInput) => {
    try {
      console.log(patient);
      await UpdatePatient(data, patient.id);
      console.log("Patient updated");
      setSelectedPatient(EmptyPatient);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex min-h-screen w-full flex-col bg-muted/40">
      <div className="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
        <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
          <Form {...form}>
            <form
              action="#"
              onSubmit={form.handleSubmit(handleSubmit)}
              className="space-y-8"
            >
              <div className="mx-auto grid max-w-[59rem] flex-1 auto-rows-max gap-4">
                <div className="flex items-center gap-4">
                  <Button
                    variant="outline"
                    size="icon"
                    className="h-7 w-7"
                    type="reset"
                    onClick={() => {
                      form.reset();
                      setSelectedPatient(EmptyPatient);
                    }}
                  >
                    <ChevronLeft className="h-4 w-4" />
                    <span className="sr-only">Back</span>
                  </Button>
                  <h1 className="flex-1 shrink-0 whitespace-nowrap text-3xl font-semibold tracking-tight sm:grow-0">
                    {patient.firstName} {patient.lastName}
                  </h1>
                  <div className="hidden items-center gap-2 md:ml-auto md:flex">
                    <Button size="sm" type="submit">
                      Save Patient
                    </Button>
                  </div>
                </div>
                <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
                  <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
                    <Card x-chunk="dashboard-07-chunk-0">
                      <CardHeader>
                        <CardTitle>Personal Info</CardTitle>
                        <CardDescription>
                          Enter the patient's personal information
                        </CardDescription>
                      </CardHeader>
                      <CardContent>
                        <div className="grid gap-6">
                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.firstName}
                              name="firstName"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>First Name</FormLabel>
                                  <FormControl>
                                    <Input placeholder="Jane" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={patient.lastName}
                              name="lastName"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Last Name</FormLabel>
                                  <FormControl>
                                    <Input placeholder="Doe" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>

                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.phone}
                              name="phone"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Phone</FormLabel>
                                  <FormControl>
                                    <Input
                                      type="tel"
                                      placeholder="0712345678"
                                      {...field}
                                    />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={patient.email}
                              name="email"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Email</FormLabel>
                                  <FormControl>
                                    <Input
                                      placeholder="patient@gmail.com"
                                      {...field}
                                    />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.birthDate.split("T")[0]}
                              name="birthDate"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Birth Date</FormLabel>
                                  <FormControl>
                                    <Input
                                      className="w-full"
                                      type="date"
                                      {...field}
                                    />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={patient.address}
                              name="address"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Address</FormLabel>
                                  <FormControl>
                                    <Input className="w-full" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                    <Card x-chunk="dashboard-07-chunk-1">
                      <CardHeader>
                        <CardTitle>Pregnancy Details</CardTitle>
                        <CardDescription>
                          Enter the patient's current pregnancy details
                        </CardDescription>
                      </CardHeader>
                      <CardContent>
                        <div className="grid gap-6">
                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.lmp.split("T")[0]}
                              name="lmp"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>LMP</FormLabel>
                                  <FormControl>
                                    <Input type="date" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={
                                patient.conceptionDate.split("T")[0]
                              }
                              name="conceptionDate"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Conception</FormLabel>
                                  <FormControl>
                                    <Input type="date" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.edd.split("T")[0]}
                              name="edd"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>EDD</FormLabel>
                                  <FormControl>
                                    <Input type="date" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={patient.sonoDate.split("T")[0]}
                              name="sonoDate"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Sono Date</FormLabel>
                                  <FormControl>
                                    <Input type="date" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                          <div className="grid grid-cols-4 gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.crl}
                              name="crl"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>CRL</FormLabel>
                                  <FormControl>
                                    <div className="grid grid-cols-12 gap-3">
                                      <Slider
                                        className="col-span-9"
                                        defaultValue={[patient.crl]}
                                        max={95}
                                        min={15}
                                        step={1}
                                        onValueChange={(vals) => {
                                          field.onChange(vals[0]);
                                        }}
                                      />
                                      <h4 className="text-sm font-medium tracking-tight">
                                        {patient.crl}mm
                                      </h4>
                                    </div>
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                            <FormField
                              control={form.control}
                              defaultValue={patient.crlDate.split("T")[0]}
                              name="crlDate"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>CRL Date</FormLabel>
                                  <FormControl>
                                    <Input type="date" {...field} />
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                  </div>
                  <div className="grid auto-rows-max items-start gap-4 lg:gap-8">
                    <Card x-chunk="dashboard-07-chunk-3">
                      <CardHeader>
                        <CardTitle>Pregnancy Status</CardTitle>
                        <CardDescription>
                          {" "}
                          Has the patient delivered?
                        </CardDescription>{" "}
                      </CardHeader>
                      <CardContent>
                        <div className="grid gap-6">
                          <div className="grid gap-3">
                            <FormField
                              control={form.control}
                              defaultValue={patient.delivered}
                              name="delivered"
                              render={({ field }) => (
                                <FormItem className="col-span-2">
                                  <FormLabel>Status</FormLabel>
                                  <FormControl>
                                    <Select
                                      defaultValue={
                                        patient.delivered ? "true" : "false"
                                      }
                                      onValueChange={field.onChange}
                                    >
                                      <SelectTrigger
                                        id="status"
                                        aria-label="Select status"
                                      >
                                        <SelectValue placeholder="Select status" />
                                      </SelectTrigger>
                                      <SelectContent>
                                        <SelectItem value={"true"}>
                                          Delivered
                                        </SelectItem>
                                        <SelectItem value={"false"}>
                                          Active
                                        </SelectItem>
                                      </SelectContent>
                                    </Select>
                                  </FormControl>
                                  <FormMessage />
                                </FormItem>
                              )}
                            />
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                    <Card
                      className="overflow-hidden"
                      x-chunk="dashboard-07-chunk-4"
                    >
                      <CardHeader>
                        <CardTitle>Image</CardTitle>
                        <CardDescription>
                          Upload an image of the patient (optional)
                        </CardDescription>
                      </CardHeader>
                      <CardContent>
                        <div className="grid gap-2">
                          <Image
                            alt="Product image"
                            className="aspect-square w-full rounded-md object-cover"
                            height="300"
                            src="/placeholder.svg"
                            width="300"
                          />
                          <div className="grid grid-cols-3 gap-2">
                            <button>
                              <Image
                                alt="Product image"
                                className="aspect-square w-full rounded-md object-cover"
                                height="84"
                                src="/placeholder.svg"
                                width="84"
                              />
                            </button>
                            <button>
                              <Image
                                alt="Product image"
                                className="aspect-square w-full rounded-md object-cover"
                                height="84"
                                src="/placeholder.svg"
                                width="84"
                              />
                            </button>
                            <button className="flex aspect-square w-full items-center justify-center rounded-md border border-dashed">
                              <Upload className="h-4 w-4 text-muted-foreground" />
                              <span className="sr-only">Upload</span>
                            </button>
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                  </div>
                </div>
                <div className="flex items-center justify-end gap-2 md:hidden">
                  <Button variant="outline" size="sm">
                    Discard
                  </Button>
                  <Button size="sm" type="submit">
                    Add Patient
                  </Button>
                </div>
              </div>
            </form>
          </Form>
        </main>
      </div>
    </div>
  );
};

export default PatientView;
