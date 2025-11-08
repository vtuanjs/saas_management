## When atlas conflict happening, you need to:

Check your teammate migration file first

1. If not conflict (Most case) or conflict but easily to resolve
- Rebase your migration
`make migrate-rebase <your_migration_file_name>`
- Update your migration:
Add some condition: CREATE IF NOT EXIST, update something if not exist...
- Update alias hash
`make migration-hash`
- Apply migration
`make migrate-apply`

Optional: Clean your local migration
- Manual query DB
SELECT * FROM public.atlas_schema_revisions
ORDER BY version DESC 
- Delete old <your_migration_file_name>

2. If hard to resolve (Rare case)
- Down your local migration
`make migrate-down`
- Re-generate your local migration
`make migrate-remove <your_migration_file_name>`
`make migrate-generate <your_migration_name>`
- Apply migration
`make migrate-apply`

2. If very hard (Bad case)
- Drop your local db
`make schema-clean`
- Apply migration
`make migrate-apply`
- Seed db

## Tips to avoid conflict migration
- Code first with unit test
- After finish code -> merge code in dev first -> Generate migration -> Apply migration -> Integration test
- Apply trunk based development
