drop database if exists `news`;
create database if not exists `news`;
use `news`;

drop table if exists `roles`;
drop table if exists `comments`;
drop table if exists `news`;
drop table if exists `categories`;
drop table if exists `users`;
drop table if exists `tags`;
drop table if exists `news_tags`;
drop table if exists `censor`;

create table if not exists `roles`(
	`id` 	int(11) auto_increment,
    `name` 	nvarchar(255) ,
    primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `users` (
	`id` 			int(11) auto_increment,
    `username` 		varchar(50) not null unique,
    `password` 		varchar(255) not null,
    `role_id` 		int(11) not null,
    `name` 			nvarchar(50) not null,
    `phone_number` 	varchar(12) not null,
    `sex` 			bool not null,	-- 1_male, 0_female;
    `email` 		varchar(60) unique,
	 primary key(`id`),
	constraint `users_ibfk1` foreign key (`role_id`) references `roles` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `categories`(
	`id` 					int(11) auto_increment,
    `name`	 				nvarchar(255) not null,
    `parent_category_id` 	int(11),
    primary key (`id`),
    constraint `category_ibfk1` foreign key(`parent_category_id`) references `categories`(`id`)	
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `news`(
	`id` 			int(11) auto_increment,
    `title` 		nvarchar(255),
    `avatar` 		nvarchar(255),
    `summary`		text,
    `content` 		longtext,
    `user_id` 		int not null,
    `date_post` 	datetime,
    `category_id` 	int not null,
    `views` 		int not null default 0,
    `is_premium`		bool not null default false,
    primary key (`id`),
    constraint `news_ibfk1` foreign key (`user_id`) references `users` (`id`),
    constraint `news_ibfk2` foreign key (`category_id`) references `categories` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `censor`(
	`news_id` 		int(11),
    `user_id`		int(11),
    `is_public`		bool not null default false,
    `date_censor`	timestamp not null,
    `date_public`	timestamp not null,
    primary key(`news_id`, `user_id`),
    constraint `censor_ibfk1` foreign key (`news_id`) references `news` (`id`),
    constraint `censor_ibfk2` foreign key (`user_id`) references `users` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `comments`(
	`id` 		int(11) auto_increment,
    `news_id` 	int(11) not null,
    `message` 	text,
    `user_id` 	int(11) not null,
    `date_post` datetime,
    primary key (`id`),
    constraint `comment_ibfk1` foreign key (`user_id`) references `users` (`id`),
    constraint `comment_ibfk2` foreign key (`news_id`) references `news` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `tags`(
	`id`	int(11) auto_increment,
    `name`	varchar(50) not null,
     primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table if not exists `news_tags`(
	`news_id`	int(11),
    `tag_id`	int(11),
	primary key (`news_id`, `tag_id`),
    constraint `news_tags_ibfk1` foreign key (`news_id`) references `news` (`id`),
    constraint `news_tags_ibfk2` foreign key (`tag_id`) references `tags` (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

insert into `categories` (`name`,`parent_category_id`) value
('Xã hội', null),
('Thời sự', 1),
('Giao thông', 1),
('Môi trường - Khí hậu', 1),
('Pháp luật', 1),
('Thế giới', null),
('Văn hóa', null),
('Nghệ thuật', 7),
('Ẩm thực', 7),
('Du lịch', 7),
('Kinh tế',null),
('Kinh doanh', 11),
('Lao động - Việc làm', 11),
('Chứng khoán', 11),
('Tài chính', 11),
('Giáo dục', null),
('Học bổng - Du học', 16),
('Đào tạo - Thi cử', 16),
('Thể thao', null),
('Bóng đá quốc tế', 19),
('Bóng đá việt nam', 19);

insert into `roles` (`name`) value
('độc giả'),
('premium'),
('writer'),
('editor'),
('administrator');
