create table if not exists ticks
(
    timestamp bigint      not null,
    symbol    varchar(20) not null,
    best_bid  double      not null,
    best_ask  double      not null
) engine InnoDB;