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
