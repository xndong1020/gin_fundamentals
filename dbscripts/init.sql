CREATE TABLE IF NOT EXISTS public.albums (
	id serial NOT NULL,
	title VARCHAR(255),
	artist VARCHAR(255),
	price DECIMAL,
	CONSTRAINT "PK_tbl_albums" PRIMARY KEY (id)
);

INSERT INTO public.albums ("title", "artist", "price") VALUES
('Blue Train','John Coltrane',56.99)
,('Jeru','Gerry Mulligan',17.99)
,('Sarah Vaughan and Clifford Brown','Sarah Vaughan',39.99)
;