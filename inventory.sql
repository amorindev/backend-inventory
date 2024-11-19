CREATE TABLE tb_user (
    user_id SERIAL PRIMARY KEY,
    user_email VARCHAR(200) NOT NULL,
    user_pass VARCHAR(300) NOT NULL
)

CREATE TABLE tb_role (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL
)

CREATE TABLE tb_user_role (
    userrole_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES tb_user(user_id),
    role_id INTEGER REFERENCES tb_role(role_id)
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



INSERT INTO tb_user (user_email, user_pass) 
VALUES
('calidad@gmail.com', 'Calidad2024')

-- First create user account,get id
INSERT INTO tb_company (com_user_id, com_name, com_website, com_address, com_phone, com_email, com_logo)
VALUES 
(1, 'Limdes - Chimbote', 'https://www.limdes.com', 'Av. Avenida Industrial', '923456789', 'limbes@gmail.com', 'logo.png');

-- first inset  company, get id
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
    ('LIMPIEZA DEL HOGAR'),
    ('CUIDADO DE LA ROPA'),
    ('DESINFECCIÓN'),
    ('AROMATIZACIÓN Y AMBIENTACIÓN'),
    ('CUIDADO DEL AUTOMÓVIL'),
    ('PRODUCTOS ECOLÓGICOS'),
    ('HERRAMIENTAS DE LIMPIEZA'),
    ('CUIDADO DE MASCOTAS'),
    ('PISOS Y SUPERFICIES'),
    ('JARDÍN Y EXTERIORES');



INSERT INTO tb_product("prod_name", "prod_desc", "prod_discount", "prod_price", "prod_stk", "cat_id")
VALUES
    ('Limpiador Multiusos MR. CLEAN', 'Limpiador multiusos para todo tipo de superficies, botella de 1Lt', 0, 10.0, 200, 1),
    ('Detergente ARIEL', 'Detergente en polvo de 1.5Kg para lavar ropa blanca y de color', 5, 15.0, 300, 2),
    ('Desinfectante LYSOL', 'Desinfectante en aerosol para eliminación de gérmenes y bacterias', 10, 12.0, 150, 3),
    ('Aromatizante Glade', 'Aromatizante en aerosol de fragancia lavanda para el hogar', 0, 8.5, 220, 4),
    ('Limpiavidrios WINDEX', 'Limpiador de ventanas y superficies de cristal, 500ml', 0, 7.0, 180, 5),
    ('Jabón ecológico ECOS', 'Jabón biodegradable para limpieza de superficies, 500ml', 3, 6.5, 100, 6),
    ('Mopa Mágica', 'Mopa de microfibra con sistema de escurrido automático', 0, 25.0, 80, 7),
    ('Shampoo para Mascotas PETSAFE', 'Shampoo desinfectante y aromatizante para mascotas, 500ml', 0, 18.0, 120, 8),
    ('Desengrasante para Cocina EASY-OFF', 'Desengrasante potente para hornos y cocinas, 1Lt', 0, 14.5, 140, 9),
    ('Limpiador de Pisos CIF', 'Limpiador especializado para pisos y superficies duras, 1Lt', 5, 9.5, 160, 10);

INSERT INTO tb_kardex(kar_desc, kar_type)
VALUES
('Rebibir productos de Amacén en Arequipa', 'ENTRADA')
-- insert kardex, get id
INSERT INTO tb_product_kardex(pro_kar_amount,prod_id, kar_id)
VALUES
(20,2,1),
(43,5,1);
