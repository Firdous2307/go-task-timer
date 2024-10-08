## Persistent Storage with SQLite

Go-Task-Timer now implements SQLite for persistent task storage, enhancing the application's functionality and user experience. This addition brings several benefits:

1. **Data Persistence**: Tasks are now stored in a SQLite database, allowing users to retain their task history even after closing the application.

2. **Improved Scalability**: SQLite provides efficient data management for a growing number of tasks, outperforming in-memory storage for larger datasets.

3. **Data Integrity**: SQLite ensures ACID (Atomicity, Consistency, Isolation, Durability) compliance, maintaining data integrity even in case of unexpected shutdowns.

4. **Lightweight and Self-contained**: SQLite is a serverless, zero-configuration database engine that's perfect for applications like Go-Task-Timer.

5. **Cross-platform Compatibility**: SQLite works seamlessly across different operating systems, maintaining Go-Task-Timer's portability.

6. **Query Capabilities**: With SQLite, we can implement more advanced features like searching, filtering, and reporting on task data.

The implementation uses the `database/sql` package along with the `github.com/mattn/go-sqlite3` driver, providing a robust and Go-idiomatic way to interact with the SQLite database.

[Learn about Go-Task-Timer's features](features.md)