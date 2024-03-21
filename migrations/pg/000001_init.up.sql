CREATE TABLE goods (
    id                  serial primary key,
    name			    varchar(255) UNIQUE not null
);


CREATE TABLE shelves (
    id           serial primary key,
    name         varchar(30) not null
);

CREATE TABLE goods_shelves (
    good_id             integer references goods(id) not null,
    shelf_id           integer references shelves(id) not null,
    main_or_additional       varchar(30) not null
);

CREATE TABLE orders (
     id           serial primary key,
     number         numeric unique not null
);

CREATE TABLE goods_orders (
   good_id           integer references goods(id) not null,
   order_id           integer references orders(id) not null,
   good_count             numeric not null
);

INSERT INTO goods (name) VALUES
    ('Ноутбук'),
    ('Монитор'),
    ('Телефон'),
    ('Системный блок'),
    ('Часы'),
    ('Микрофон');

INSERT INTO shelves (name) VALUES
    ('А'),
    ('Б'),
    ('В'),
    ('Ж'),
    ('З');


INSERT INTO goods_shelves (good_id, shelf_id, main_or_additional) VALUES
    (1, 1, 'главный'),
    (2, 1, 'главный'),
    (3, 2, 'главный'),
    (3, 3, 'дополнительный'),
    (3, 5, 'дополнительный'),
    (4, 4, 'главный'),
    (5, 4, 'главный'),
    (5, 1, 'дополнительный'),
    (6, 4, 'главный');

INSERT INTO orders (number) VALUES
    (10),
    (11),
    (14),
    (15);

INSERT INTO goods_orders (good_id, order_id, good_count) VALUES
    (1, 1, 2),
    (1, 3, 3),
    (3, 1, 1),
    (6, 1, 1),
    (2, 2, 3),
    (4, 3, 3),
    (5, 4, 1);