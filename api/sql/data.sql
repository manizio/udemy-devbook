insert into users (name, nick, email, password)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$OOKLVCjDGOWq/YNm7gb7/.a9ZnYEvaxVGUkkMQ9j1DQgf0ia7T21C"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$OOKLVCjDGOWq/YNm7gb7/.a9ZnYEvaxVGUkkMQ9j1DQgf0ia7T21C"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$OOKLVCjDGOWq/YNm7gb7/.a9ZnYEvaxVGUkkMQ9j1DQgf0ia7T21C");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);
