CREATE TABLE tb_user (
    user_id SERIAL PRIMARY KEY,
    user_email VARCHAR(200) NOT NULL,
    user_pass VARCHAR(300) NOT NULL
)

CREATE TABLE tb_company(
    com_user_id INTEGER PRIMARY KEY REFERENCES tb_user(user_id),
    com_name VARCHAR(300) NOT NULL,
    com_website VARCHAR(250) NOT NULL,
    com_address VARCHAR(250) NOT NULL,
    com_phone VARCHAR(20) NOT NULL,
    com_email VARCHAR(100) NOT NULL,
    com_logo VARCHAR(250)
)

CREATE TABLE tb_provider(
    prov_id SERIAL PRIMARY KEY,
    prov_name VARCHAR(300) NOT NULL,
    prov_address VARCHAR(250) NOT NULL,
    prov_email VARCHAR(100) NOT NULL,
    prov_phone VARCHAR(20) NOT NULL,
    com_user_id INTEGER REFERENCES tb_company(com_user_id)
)

CREATE TABLE tb_category (
        cat_id SERIAL PRIMARY KEY,
        cat_name VARCHAR(150) NOT NULL
)

CREATE TABLE tb_product (
        prod_id SERIAL PRIMARY KEY,
        prod_name VARCHAR(150) NOT NULL,
        prod_desc TEXT NOT NULL,
        prod_discount SMALLINT NOT NULL DEFAULT 0, -- 10%, 20%
        prod_price NUMERIC(10, 2) NOT NULL,
        prod_stk INT NOT NULL CHECK (prod_stk >= 0),
        cat_id INTEGER REFERENCES tb_category (cat_id)
)

CREATE TABLE tb_kardex (
    kar_id SERIAL PRIMARY KEY,
    kar_desc VARCHAR(250) NOT NULL,
    kar_type VARCHAR(60) NOT NULL CHECK (kar_type IN ('SALIDA', 'ENTRADA')),
    kar_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)

CREATE TABLE tb_provider_product(
    prov_prod_id SERIAL PRIMARY KEY,
    prov_id INTEGER REFERENCES tb_provider(prov_id),
    prod_id INTEGER REFERENCES tb_product(prod_id)
)

CREATE TABLE tb_product_kardex (
    pro_kar_id SERIAL PRIMARY KEY,
    pro_kar_amount INTEGER NOT NULL,
    prod_id INTEGER REFERENCES tb_product(prod_id),
    kar_id INTEGER REFERENCES tb_kardex(kar_id)
)


INSERT INTO tb_category ("cat_name")
VALUES
    ('TECNOLOGÍA'),
    ('ALIMENTOS Y BEBIDAS'),
    ('LACTEOS Y DERIVADOS'),
    ('LIMPIEZA Y CUIDADO DEL HOGAR'),
    ('HIGIENE Y CUIDADO PERSONAL'),
    ('BEBES Y NIÑOS'),
    ('ELECTRODOMÉSTICOS'),
    ('ROPA Y CALZADO'),
    ('MUEBLES Y DECORACIÓN'),
    ('MASCOTAS'),
    ('JUGUETERÍA'),
    ('OFICINA Y PAPELERÍA');


INSERT INTO tb_product("prod_name","prod_desc","prod_discount","prod_price","prod_stk","cat_id")
VALUES
 ('Yogur', 'Yogur de 1Lt Gloria sabor fresa', 0, 12.0, 120,3),
 ('Escoba', 'Juego de Escoba Y Recogedor 2 en 1', 2, 18.0, 90,4),
 ('Licuadora', 'Licuadora Oster® con control de textura BLST3B Niquelada + Accesorios', 0, 899.0, 10,7),
 ('Escritorio', 'Organizador Escritorio Oficina Papelería con Cajones SJ-159', 20, 19.90, 20,12);



CREATE OR REPLACE FUNCTION kardex_update_prodcut_stock()
RETURNS TRIGGER AS $$

DECLARE kar_type VARCHAR(60);
BEGIN

    SELECT k.kar_type INTO kar_type
    FROM tb_kardex AS k
    WHERE k.kar_id = NEW.kar_id;

    IF kar_type = 'ENTRADA' THEN
        UPDATE tb_product
        SET prod_stk = prod_stk + NEW.pro_kar_amount 
        WHERE prod_id = NEW.prod_id;

    ELSEIF kar_type = 'SALIDA' THEN
        UPDATE tb_product
        SET prod_stk = prod_stk - NEW.pro_kar_amount
        WHERE prod_id = NEW.prod_id;
    END IF;
    RETURN NEW;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_kardex_update_stock
AFTER INSERT ON tb_product_kardex
FOR EACH ROW
EXECUTE FUNCTION kardex_update_prodcut_stock();


SELECT 
	k.kar_id, k.kar_desc, k.kar_type, k.kar_created_at,
	pk.pro_kar_amount, p.prod_id, p.prod_name
	FROM tb_kardex k
	JOIN tb_product_kardex pk ON k.kar_id = pk.kar_id
	JOIN tb_product p ON pk.prod_id = p.prod_id
	ORDER BY k.kar_created_at DESC;