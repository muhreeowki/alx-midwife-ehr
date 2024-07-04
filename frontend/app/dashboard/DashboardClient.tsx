"use client";

import React from "react";
import { EmptyPatient, Patient } from "@/lib/models";
import PatientView from "./PatientView";
import AllPatientsVeiw from "./AllPatientsVeiw";

import z from "zod";
import { CreatePatientSchema } from "@/lib/zodSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

const DashboardClient = ({ patients }: { patients: Patient[] }) => {
  const [selectedPatient, setSelectedPatient] =
    React.useState<Patient>(EmptyPatient);
  const form = useForm<z.infer<typeof CreatePatientSchema>>({
    resolver: zodResolver(CreatePatientSchema),
    defaultValues: {
      midwifeId: 0,
      deliveryDate: EmptyPatient.deliveryDate,
      rhFactor: EmptyPatient.rhFactor,
      partnerName: EmptyPatient.partnerName,
      imageURL: EmptyPatient.imageURL,
    },
  });
  return (
    <>
      {selectedPatient.id !== 0 ? (
        <PatientView
          setSelectedPatient={setSelectedPatient}
          patient={selectedPatient}
          form={form}
        />
      ) : (
        <AllPatientsVeiw
          setSelectedPatient={setSelectedPatient}
          patients={patients}
        />
      )}
    </>
  );
};

export default DashboardClient;
