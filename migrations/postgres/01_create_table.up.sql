CREATE TABLE "customers" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR , 
    "phone" VARCHAR 
);

CREATE TABLE "users" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR , 
    "phone_number" VARCHAR 
);

CREATE TABLE "couriers" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR , 
    "phone_number" VARCHAR 
);

CREATE TABLE "categories" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR 
);

CREATE TABLE "products" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR , 
    "price" NUMERIC ,
    "category_id" UUID REFERENCES categories("id")
);

CREATE TABLE "orders" (
    "id" UUID PRIMARY KEY ,
    "name" VARCHAR ,
    "price" NUMERIC ,
    "phone_number" VARCHAR ,
    "latitude" NUMERIC,
    "longtitude" NUMERIC, 
    "user_id" UUID,
    "customer_id" UUID,
    "courier_id" UUID,
    "product_id" UUID,
    "quantity" NUMERIC
);