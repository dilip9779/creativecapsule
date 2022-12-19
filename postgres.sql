CREATE TABLE IF NOT EXISTS public.scratch_card
(
    id bigserial PRIMARY KEY NOT NULL,
    discount_amount double precision NOT NULL,
    expiry_date date NOT NULL,
    is_scratched boolean NOT NULL DEFAULT false,
    is_active boolean NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS public."user"
(
    id bigserial PRIMARY KEY NOT NULL,
    user_email character varying NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    UNIQUE (user_email)
);

CREATE TABLE IF NOT EXISTS public.transaction
(
    id bigserial PRIMARY KEY NOT NULL,
    transaction_amount double precision NOT NULL,
    user_id integer NOT NULL,
    scratch_card_id integer NOT NULL,
    date_of_transaction timestamp with time zone NOT NULL DEFAULT now(),
    UNIQUE (user_id, scratch_card_id),
    CONSTRAINT transaction_scratch_card_id_fkey FOREIGN KEY (scratch_card_id)
        REFERENCES public.scratch_card (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT transaction_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);