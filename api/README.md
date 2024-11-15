# Archpass API

Create .env file and add the values for the ff:
```
DATABASE_URL=
ENV=local
```

Running the app. Add --build flag if this is your first time running the app.
```sh 
docker compose up
```

Creating migrations
```sh
make create_migration name="migration_name"
```

Apply migrations
```sh
make apply_migrations 
```

Handle database migration conflicts
```sh
make migrate_hash
```

Code gen. You need to run ent code gen everytime you update ent/schema files
```sh
make ent_generate
```
