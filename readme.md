# golang edu crud

## create users table

```
create table users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    last_name TEXT,
    second_name TEXT
)

insert into users (name, last_name, second_name) values ("egor", "warete", "")
```