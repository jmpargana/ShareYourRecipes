create table recipe (
    item_id int primary_key,
    private int,
    time int,
    title text,
    ingridients text,
    method text
);

create table tag (
    tag_id int primary_key,
    title text
);


create table recipe_tag (
    recipe_id int foreign_key,
    tag_id int foreign_key
);



