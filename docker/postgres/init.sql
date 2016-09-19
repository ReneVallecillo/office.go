--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.4
-- Dumped by pg_dump version 9.5.1

-- Started on 2016-09-06 16:25:12 CST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12361)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2150 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 187 (class 1259 OID 16427)
-- Name: user; Type: TABLE; Schema: public; Owner: OfficeAdmin
--

CREATE TABLE "user" (
    user_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    first_name text,
    last_name text,
    email text,
    address text,
    contact_number text,
    status_id integer,
    user_level integer,
    password text,
    gender_id integer,
    pic_url text
);


ALTER TABLE "user" OWNER TO "OfficeAdmin";

--
-- TOC entry 188 (class 1259 OID 16430)
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: OfficeAdmin
--

CREATE SEQUENCE user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_user_id_seq OWNER TO "OfficeAdmin";

--
-- TOC entry 2154 (class 0 OID 0)
-- Dependencies: 188
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: OfficeAdmin
--

ALTER SEQUENCE user_user_id_seq OWNED BY "user".user_id;


--
-- TOC entry 2018 (class 2604 OID 16432)
-- Name: user_id; Type: DEFAULT; Schema: public; Owner: OfficeAdmin
--

ALTER TABLE ONLY "user" ALTER COLUMN user_id SET DEFAULT nextval('user_user_id_seq'::regclass);


--
-- TOC entry 2028 (class 2606 OID 16440)
-- Name: pk_user; Type: CONSTRAINT; Schema: public; Owner: OfficeAdmin
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT pk_user PRIMARY KEY (user_id);


--
-- TOC entry 2149 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2016-09-06 16:25:13 CST

--
-- PostgreSQL database dump complete
--

