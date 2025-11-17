# Repository usage for the application

- For command-based applications:
Use the root repository (for example, `user.go`) and follow a lightweight DDD ("Lite") approach. This approach helps keep the code easy to maintain and prevents breakage if the database schema changes.

- For query-based applications:
You can use `sqlc` directly; you do not need to follow DDD, which reduces code complexity.