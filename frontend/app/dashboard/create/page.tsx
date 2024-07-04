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
      birthDate: "",
      email: "",
      phone: "",
      address: "",
      partnerName: "",
      imageURL: "",
      lmp: "",
      conceptionDate: "",
      sonoDate: "",
      crl: 0,
      crlDate: "",
      edd: "",
      rhFactor: "",
      delivered: false,
      deliveryDate: "",
      midwifeId: 0,
    },
  });

  
  return <CreatePatientForm form={form} />;
};

export default CreatePatientPage;
