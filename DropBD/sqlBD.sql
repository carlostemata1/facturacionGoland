CREATE TABLE `promocion` (
  `id` int PRIMARY KEY,
  `descipcion` varchar(100),
  `porcentaje` double,
  `fecha_inicio` date,
  `fecha_fin` date
);

CREATE TABLE `Factura` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `fecha_crear` date,
  `pago_total` double
);

CREATE TABLE `Medicamento` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nombre` varchar(50),
  `precio` double,
  `ubicacion` varchar(50)
);

ALTER TABLE `Medicamento` ADD FOREIGN KEY (`id`) REFERENCES `Factura` (`id`);
