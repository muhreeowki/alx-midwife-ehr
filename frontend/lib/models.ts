export type Midwife = {
  id: string;
  firstName: string;
  lastName: string;
  token: string;
  email: string;
};

export type Patient = {
  id: number;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;

  // Patient's personal details
  firstName: string;
  lastName: string;
  birthDate: string | null;
  email: string | null;
  phone: string | null;
  address: string | null;
  partnerName: string | null;
  imageURL: string | null;

  // Patients' medical details
  lmp: string | null; // Last Menstrual Period
  conceptionDate: string | null; // Date of conception
  sonoDate: string | null; // Date of sonogram
  crl: number | null; // Crown Rump Length
  crlDate: string | null; // Date of CRL
  edd: string | null; // Estimated Due Date
  rhFactor: string | null; // Rh Factor

  // Delivery details
  delivered: boolean; // Has the patient delivered
  deliveryDate: string | null; // Date of delivery

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
