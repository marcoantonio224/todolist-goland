CREATE TABLE `todos`
(
  id bigint auto_increment,
  todo varchar(225) NOT NULL,
  completed boolean DEFAULT False,
  PRIMARY KEY (`id`)
);

INSERT INTO `todos` (`todo`) VALUES ('Learn Go!');