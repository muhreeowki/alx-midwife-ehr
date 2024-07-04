import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { GetPatients } from "@/app/serverActions";
import { Patient } from "@/lib/models";
import DashboardClient from "./DashboardClient";

export default async function Dashboard() {
  const data = cookies().get("token");
  if (!data) {
    redirect("/auth");
  }
  const token = data.value;

  const patients: Patient[] = await GetPatients(token);
  console.log(patients);

  return <DashboardClient patients={patients} />;
}
