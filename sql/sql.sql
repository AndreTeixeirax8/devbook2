CREATE DATABASE IF NOT EXISTS devbook DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE devbook;


DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuario;

CREATE TABLE usuario (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    nick VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    senha VARCHAR(200) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,


    seguidor_id INT NOT NULL
    FOREIGN KEY(seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id,seguidor_id)
)

CREATE TABLE publicacoes(
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuario(id)
    ON DELETE CASCADE,
    curtidas INT DEFAULT 0,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;