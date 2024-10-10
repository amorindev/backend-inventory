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

INSERT INTO tb_user (user_email, user_pass) 
VALUES
('calidad@gmail.com', 'Calidad2024')

-- First create user account
INSERT INTO tb_company (com_user_id, com_name, com_website, com_address, com_phone, com_email, com_logo)
VALUES 
(1, 'Limdes - Chimbote', 'https://www.limdes.com', 'Av. Avenida Industrial', '923456789', 'limbes@gmail.com', 'logo.png');


INSERT INTO tb_provider (prov_name, prov_address, prov_email, prov_phone, com_user_id) VALUES
('Distribuciones Alimenticias S.A.', 'Calle Principal 123, Ciudad Centro', 'contacto@distribuciones.com', '555-0123', 1),
('Suministros de Oficina Ltda.', 'Avenida de la Libertad 456, Barrio Nuevo', 'info@suministros.com', '555-4567', 1),
('Soluciones Informáticas S.R.L.', 'Calle Secundaria 789, Sector Industrial', 'soporte@solucionesinf.com', '555-7890', 1),
('Materiales de Construcción El Roble', 'Avenida 10 de Agosto 321, Zona Industrial', 'ventas@elroble.com', '555-3456', 1),
('Productos Químicos del Norte', 'Calle del Comercio 654, Polígono Norte', 'ventas@productosquimicos.com', '555-2345', 1),
('Tecnología Avanzada S.A.C.', 'Calle San Martín 987, Parque Tecnológico', 'info@tecnologiaavanzada.com', '555-6789', 1),
('Transporte y Logística Global', 'Carretera Nacional 202, Ciudad Empresarial', 'contacto@transporte.com', '555-8901', 1),
('Servicios de Limpieza Eficaz', 'Calle 5 de Junio 543, Urbanización Jardines', 'info@limpiezaeficaz.com', '555-4321', 1),
('Electrodomésticos y Más', 'Avenida La Paz 159, Plaza Comercial', 'ventas@electrodomesticos.com', '555-6780', 1),
('Ropa y Accesorios Elegantes', 'Calle de la Moda 111, Centro Comercial', 'atencion@modaelegante.com', '555-9876', 1);

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
 ('Yogurt Bebible GLORIA', 'Yogur de 1Lt Gloria sabor fresa', 0, 12.0, 210,3),
 ('Escoba HUDE', 'Juego de Escoba Y Recogedor 2 en 1', 2, 18.0, 350,4),
 ('Licuadora OSTER BLST3AR2G053 Xpert ', 'Licuadora Oster® con control de textura BLST3B Niquelada + Accesorios', 0, 899.0, 10,7),
 ('Escritorio VIVA HOME Nilo', 'Organizador Escritorio Oficina Papelería con Cajones SJ-159', 20, 19.90, 400,12);
 ('Blusa Hypnotic ', 'Mujer Manga Larga Delfin', 0, 110.90, 240,8);
 ('Carrito Estante Organizador', ' Organizador de Oficina Papelería Almacenamiento Multiusos FH4', 2, 199.90, 310,12);


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