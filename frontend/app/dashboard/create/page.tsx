"use client";

import React from "react";

import z from "zod";
import { CreatePatientSchema } from "@/lib/zodSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import CreatePatientForm from "./CreatePatientForm";

const CreatePatientPage = () => {
  const form = useForm<z.infer<typeof CreatePatientSchema>>({
    resolver: zodResolver(CreatePatientSchema),
    defaultValues: {
      firstName: "",
      lastName: "",
      birthDate: "2022-01-12T00:00:00.000Z",
      email: "",
      phone: "",
      address: "",
      partnerName: "",
      imageURL: "",
      lmp: "2022-01-12T00:00:00.000Z",
      conceptionDate: "2022-01-12T00:00:00.000Z",
      sonoDate: "2022-01-12T00:00:00.000Z",
      crl: 15,
      crlDate: "2022-01-12T00:00:00.000Z",
      edd: "2022-01-12T00:00:00.000Z",
      rhFactor: "2022-01-12T00:00:00.000Z",
      delivered: false,
      deliveryDate: "2022-01-12T00:00:00.000Z",
      midwifeId: 0,
    },
  });

  return <CreatePatientForm form={form} />;
};

export default CreatePatientPage;
