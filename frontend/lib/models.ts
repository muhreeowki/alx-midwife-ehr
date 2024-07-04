export type Midwife = {
  id: string;
  firstName: string;
  lastName: string;
  token: string;
  email: string;
};

export type Patient = {
  // Admin data
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;

  // Patient's personal details
  firstName: string;
  lastName: string;
  birthDate: string;
  email: string;
  phone: string;
  address: string;
  partnerName: string;
  imageURL: string;

  // Patients' medical details
  lmp: string; // Last Menstrual Period
  conceptionDate: string; // Date of conception
  sonoDate: string; // Date of sonogram
  crl: number; // Crown Rump Length
  crlDate: string; // Date of CRL
  edd: string; // Estimated Due Date
  rhFactor: string; // Rh Factor

  // Delivery details
  delivered: boolean; // Has the patient delivered
  deliveryDate: string; // Date of delivery

  // Midwife details
  midwifeId: number;
};

export type CreatePatientInput = {
  // Patient's personal details
  firstName: string;
  lastName: string;
  birthDate: string;
  email: string;
  phone: string;
  address: string;
  partnerName: string;
  imageURL: string;

  // Patients' medical details
  lmp: string; // Last Menstrual Period
  conceptionDate: string; // Date of conception
  sonoDate: string; // Date of sonogram
  crl: number; // Crown Rump Length
  crlDate: string; // Date of CRL
  edd: string; // Estimated Due Date
  rhFactor: string; // Rh Factor

  // Delivery details
  delivered: boolean; // Has the patient delivered
  deliveryDate: string; // Date of delivery

  // Midwife details
  midwifeId: number;
};

export type AuthMidwifeLoginInput = {
  email: string;
  password: string;
};

export type AuthMidwifeSignUpInput = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
};

export type AuthMidwifeOutput = {
  id: number;
  email: string;
  firstName: string;
  lastName: string;
};

export const EmptyPatient: Patient = {
  id: 0,
  createdAt: "2022-01-12T00:00:00.000Z",
  updatedAt: "2022-01-12T00:00:00.000Z",
  deletedAt: "2022-01-12T00:00:00.000Z",
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
};
