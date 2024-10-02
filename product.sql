CREATE TABLE category (
        cat_id SERIAL PRIMARY KEY,
        cat_name VARCHAR(150) NOT NULL
);

-- CREATE - TABLES WITH FK
CREATE TABLE product (
        prod_id SERIAL PRIMARY KEY,
        prod_name VARCHAR(150) NOT NULL,
        prod_desc TEXT NOT NULL,
        prod_discount SMALLINT DEFAULT 0, -- 10%, 20%
        prod_price NUMERIC(10, 2) NOT NULL,
        prod_stk INT NOT NULL CHECK (prod_stk >= 0),
        cat_id INTEGER REFERENCES category (cat_id)
);

CREATE TABLE kardex (
    kar_id SERIAL PRIMARY KEY,
    kar_desc VARCHAR(250) NOT NULL,
    kar_tipo VARCHAR(60) NOT  NULL,                           --entrada o salida
    kar_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE product_kardex (
    pro_kar_id SERIAL PRIMARY KEY,
    pro_kar_amount INTEGER NOT NULL,
    prod_id INTEGER REFERENCES product(prod_id),
    kar_id INTEGER REFERENCES kardex(kar_id)
);
ALTER TABLE kardex
ADD CONSTRAINT check_kardex_tipo CHECK (kar_tipo IN ('SALIDA', 'ENTRADA'));

-- INSERT - TABLES WITHOUT FK
INSERT INTO
    category ("cat_name")
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



INSERT INTO product("prod_name","prod_desc","prod_discount","prod_price","prod_stk","cat_id")
VALUES
 ('Yogur', 'Yogur de 1Lt Gloria sabor fresa', 0, 12.0, 120,3),
 ('Escoba', 'Juego de Escoba Y Recogedor 2 en 1', 2, 18.0, 90,4),
 ('Licuadora', 'Licuadora Oster® con control de textura BLST3B Niquelada + Accesorios', 0, 899.0, 10,7),
 ('Escritorio', 'Organizador Escritorio Oficina Papelería con Cajones SJ-159', 20, 19.90, 20,12);




CREATE OR REPLACE FUNCTION kardex_update_prodcut_stock()
RETURNS TRIGGER AS $$

DECLARE kar_tipo VARCHAR(60);
BEGIN

    SELECT k.kar_tipo INTO kar_tipo
    FROM  kardex AS k
    WHERE k.kar_id = NEW.kar_id;

    IF kar_tipo = 'ENTRADA' THEN
        UPDATE product
        SET prod_stk = prod_stk + NEW.pro_kar_amount 
        WHERE prod_id = NEW.prod_id;

    ELSEIF kar_tipo = 'SALIDA' THEN
        UPDATE product
        SET prod_stk = prod_stk - NEW.pro_kar_amount
        WHERE prod_id = NEW.prod_id;
    END IF;
    RETURN NEW;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_kardex_update_stock
AFTER INSERT ON product_kardex
FOR EACH ROW
EXECUTE FUNCTION kardex_update_prodcut_stock();


SELECT 
	k.kar_id, k.kar_desc, k.kar_tipo, k.kar_created_at,
	pk.pro_kar_amount, p.prod_id, p.prod_name
	FROM kardex k
	JOIN product_kardex pk ON k.kar_id = pk.kar_id
	JOIN product p ON pk.prod_id = p.prod_id
	ORDER BY k.kar_created_at DESC;



