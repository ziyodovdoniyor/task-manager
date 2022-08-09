create table users(
    user_id uuid,
    user_name varchar(255) not null,
    password varchar(255) not null,
    primary key(user_id)
);

create table task(
    task_id uuid,
    user_id uuid,
    name varchar(255) not null,
    description text not null,
    creating_time timestamp,
    closing_time timestamp,
    primary key(task_id),
    constraint fk_user
        foreign key(user_id)
            references users(user_id)
);


