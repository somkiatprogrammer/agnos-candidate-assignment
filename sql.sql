CREATE TYPE gender AS ENUM ('M', 'F');

CREATE TABLE patients (
	id smallserial,
	first_name_th TEXT,
	middle_name_th TEXT,
	last_name_th TEXT,
	first_name_en TEXT,
	middle_name_en TEXT,
	last_name_en TEXT,
	date_of_birth TIMESTAMP,
	patient_hn TEXT,
	national_id INT,
	passport_id TEXT,
	phone_number TEXT,
	email TEXT,
	gender gender,
	PRIMARY KEY (id)
);

INSERT INTO public.patients(
	first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender) 
	VALUES ('ดัมมี่1 ชื่อ', 'ดัมมี่1 ชื่อกลาง', 'ดัมมี่1 นามสกุล', 'Dummy1 firstname', 'Dummy1 middlename', 'Dummy1 lastname', '1980-12-11', 'dummy hospital', 1, 'DUMMY1', '000000000', 'dummy1@dummy.com', 'M');

INSERT INTO public.patients(
	first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender) 
	VALUES ('ดัมมี่2 ชื่อ', 'ดัมมี่2 ชื่อกลาง', 'ดัมมี่2 นามสกุล', 'Dummy2 firstname', 'Dummy2 middlename', 'Dummy2 lastname', '1992-12-11', 'dummy hospital', 2, 'DUMMY2', '000000000', 'dummy2@dummy.com', 'M');

INSERT INTO public.patients(
	first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender) 
	VALUES ('ดัมมี่3 ชื่อ', 'ดัมมี่3 ชื่อกลาง', 'ดัมมี่3 นามสกุล', 'Dummy3 firstname', 'Dummy3 middlename', 'Dummy3 lastname', '2000-07-20', 'dummy hospital', 3, 'DUMMY3', '000000000', 'dummy3@dummy.com', 'F');

INSERT INTO public.patients(
	first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender) 
	VALUES ('ดัมมี่4 ชื่อ', 'ดัมมี่4 ชื่อกลาง', 'ดัมมี่4 นามสกุล', 'Dummy4 firstname', 'Dummy4 middlename', 'Dummy4 lastname', '2003-02-12', 'dummy hospital 2', 4, 'DUMMY4', '000000000', 'dummy4@dummy.com', 'F');

INSERT INTO public.patients(
	first_name_th, middle_name_th, last_name_th, first_name_en, middle_name_en, last_name_en, date_of_birth, patient_hn, national_id, passport_id, phone_number, email, gender) 
	VALUES ('ดัมมี่5 ชื่อ', 'ดัมมี่5 ชื่อกลาง', 'ดัมมี่5 นามสกุล', 'Dummy5 firstname', 'Dummy5 middlename', 'Dummy5 lastname', '1993-01-01', 'dummy hospital 2', 5, 'DUMMY5', '000000000', 'dummy5@dummy.com', 'M');

CREATE TABLE staffs (
	id smallserial,
	username TEXT NOT NULL,
	password TEXT  NOT NULL,
	hospital TEXT NOT NULL,
	PRIMARY KEY (id), UNIQUE (username)
);

INSERT INTO public.staffs(
	username, password, hospital)
	VALUES ('staff1', '1234', 'dummy hospital');

INSERT INTO public.staffs(
	username, password, hospital)
	VALUES ('staff2', '1234', 'dummy hospital');

INSERT INTO public.staffs(
	username, password, hospital)
	VALUES ('staff3', '1234', 'dummy hospital 2');
