PGDMP  	    *                |            mega_xxx    16.3    16.3 (               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16478    mega_xxx    DATABASE     |   CREATE DATABASE mega_xxx WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE mega_xxx;
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false                       0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    4            �            1259    16479    users    TABLE     $  CREATE TABLE public.users (
    id integer NOT NULL,
    phone text NOT NULL,
    password text NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    id_addr integer,
    role character varying(64) DEFAULT 'user'::character varying NOT NULL,
    ip_addres character varying(64)
);
    DROP TABLE public.users;
       public         heap    postgres    false    4            �            1259    16485    User_id_seq    SEQUENCE     �   ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public."User_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    215            �            1259    16486    addres_user    TABLE     �   CREATE TABLE public.addres_user (
    id integer NOT NULL,
    city text,
    street text,
    house text,
    flat text,
    entrance text
);
    DROP TABLE public.addres_user;
       public         heap    postgres    false    4            �            1259    16491    addres_user_id_seq    SEQUENCE     �   ALTER TABLE public.addres_user ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.addres_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    217            �            1259    16492    category    TABLE     v   CREATE TABLE public.category (
    id integer NOT NULL,
    name text NOT NULL,
    main_category integer NOT NULL
);
    DROP TABLE public.category;
       public         heap    postgres    false    4            �            1259    16497    category_id_seq    SEQUENCE     �   ALTER TABLE public.category ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    219    4            �            1259    16498    main_category    TABLE     W   CREATE TABLE public.main_category (
    id integer NOT NULL,
    name text NOT NULL
);
 !   DROP TABLE public.main_category;
       public         heap    postgres    false    4            �            1259    16503    main_category_id_seq    SEQUENCE     �   ALTER TABLE public.main_category ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.main_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    221            �            1259    16504 
   parameters    TABLE     �   CREATE TABLE public.parameters (
    id integer NOT NULL,
    name text NOT NULL,
    value text NOT NULL,
    id_product integer NOT NULL
);
    DROP TABLE public.parameters;
       public         heap    postgres    false    4            �            1259    16509    parameters_id_seq    SEQUENCE     �   ALTER TABLE public.parameters ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.parameters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    4    223            �            1259    16510    product    TABLE     �   CREATE TABLE public.product (
    id integer NOT NULL,
    name text NOT NULL,
    price integer NOT NULL,
    description text,
    category integer NOT NULL
);
    DROP TABLE public.product;
       public         heap    postgres    false    4            �            1259    16515    product_id_seq    SEQUENCE     �   ALTER TABLE public.product ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    225    4                      0    16486    addres_user 
   TABLE DATA           N   COPY public.addres_user (id, city, street, house, flat, entrance) FROM stdin;
    public          postgres    false    217   e+                 0    16492    category 
   TABLE DATA           ;   COPY public.category (id, name, main_category) FROM stdin;
    public          postgres    false    219   �+                 0    16498    main_category 
   TABLE DATA           1   COPY public.main_category (id, name) FROM stdin;
    public          postgres    false    221   N,                 0    16504 
   parameters 
   TABLE DATA           A   COPY public.parameters (id, name, value, id_product) FROM stdin;
    public          postgres    false    223   �,                 0    16510    product 
   TABLE DATA           I   COPY public.product (id, name, price, description, category) FROM stdin;
    public          postgres    false    225   -       	          0    16479    users 
   TABLE DATA           ]   COPY public.users (id, phone, password, name, surname, id_addr, role, ip_addres) FROM stdin;
    public          postgres    false    215   �-                  0    0    User_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public."User_id_seq"', 5, true);
          public          postgres    false    216                       0    0    addres_user_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.addres_user_id_seq', 2, true);
          public          postgres    false    218                       0    0    category_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.category_id_seq', 7, true);
          public          postgres    false    220                       0    0    main_category_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.main_category_id_seq', 5, true);
          public          postgres    false    222                        0    0    parameters_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.parameters_id_seq', 3, true);
          public          postgres    false    224            !           0    0    product_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.product_id_seq', 7, true);
          public          postgres    false    226            s           2606    16517    parameters PK_id 
   CONSTRAINT     P   ALTER TABLE ONLY public.parameters
    ADD CONSTRAINT "PK_id" PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.parameters DROP CONSTRAINT "PK_id";
       public            postgres    false    223            k           2606    16519    users User_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY public.users
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);
 ;   ALTER TABLE ONLY public.users DROP CONSTRAINT "User_pkey";
       public            postgres    false    215            m           2606    16521    addres_user addres_user_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.addres_user
    ADD CONSTRAINT addres_user_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.addres_user DROP CONSTRAINT addres_user_pkey;
       public            postgres    false    217            o           2606    16523    category category_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.category DROP CONSTRAINT category_pkey;
       public            postgres    false    219            q           2606    16525     main_category main_category_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.main_category
    ADD CONSTRAINT main_category_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.main_category DROP CONSTRAINT main_category_pkey;
       public            postgres    false    221            u           2606    16527    product product_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.product DROP CONSTRAINT product_pkey;
       public            postgres    false    225            x           2606    16528    parameters FK_id_product    FK CONSTRAINT     ~   ALTER TABLE ONLY public.parameters
    ADD CONSTRAINT "FK_id_product" FOREIGN KEY (id_product) REFERENCES public.product(id);
 D   ALTER TABLE ONLY public.parameters DROP CONSTRAINT "FK_id_product";
       public          postgres    false    4725    223    225            v           2606    16533    users User_fk5    FK CONSTRAINT     u   ALTER TABLE ONLY public.users
    ADD CONSTRAINT "User_fk5" FOREIGN KEY (id_addr) REFERENCES public.addres_user(id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT "User_fk5";
       public          postgres    false    4717    217    215            w           2606    16538    category category_fk2    FK CONSTRAINT     �   ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_fk2 FOREIGN KEY (main_category) REFERENCES public.main_category(id);
 ?   ALTER TABLE ONLY public.category DROP CONSTRAINT category_fk2;
       public          postgres    false    221    219    4721            y           2606    16543    product product_fk5    FK CONSTRAINT     v   ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_fk5 FOREIGN KEY (category) REFERENCES public.category(id);
 =   ALTER TABLE ONLY public.product DROP CONSTRAINT product_fk5;
       public          postgres    false    219    4719    225               E   x�3估�b�ņ�/l������[/콰�7ps�pra��w��¾�8��8�9��b���� G_%�         �   x�m��	�@��w�H�ͤ��<��!�,g~Z��ȹɋ/��|3s$���?*�FaH�h'����o�/l�\��7堭�Y���}ʛ�A����zq�a:&����4�r��XJ���~�nr�         T   x�=ʱ	�0�:��@P�q�J�*Q�`\��F���i�jV��8�mpl��N��OEn
�FGf����J6uf�=4j         D   x�3估����/l���¾[.vsp�pq^�va�� ��5�0�d��]vp����qqq ��W         �   x�M��	�@��3UL��n�/^z��DP,��dQ�_o:rvA����Ǜ폧s)xj�I/�b��9G����C�I�6�
�k�7������3{�܉^�l�R\sAx�0�ZЊ���%h����&I������Ê���e�      	   �   x�u�Mo�0 ���wx.�����F�P`�b� �+�������ɛ��<:��'�X�|��
]0��V���j�
�R��96#�=I���!�8�K2J-Nw�-���O qe#8
��W�&ϛ������|V�9w�t�>-�*��nN^����I,�I���G� �����Z���&����0
�6
����[��Y�-w����e�L�q��o���w��Y&�ȀHנN�#T�fEX�     