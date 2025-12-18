export type Role = {
  id: BigUint64Array;
  name: string;
};

export type User = {
  id: BigUint64Array;
  name: string;
  email: string;
  username: string;
  created_at: string;
  updated_at: string;
  role: Role;
};
